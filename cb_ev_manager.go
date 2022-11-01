package pigpio

import (
    "sync"
)

type callbackManager struct {
    pi          *Pi
    socket      *Socket
    running     bool
    handle      int
    lastLevel   int
    monitor     int
    handleCount int
    mu          sync.Mutex
}

const (
    notifyWatchDog = 1 << 5
    notifyEvent    = 1 << 7
    notifyGpio     = 31
)

func initializeCallbackManager(p *Pi, s *Socket) (*callbackManager, error) {
    var e error
    var ll, h int
    ll, e = p.ReadBank(Bank1)
    if e != nil {
        return nil, NewPiError(7, e, "Initialize")
    }

    h, e = s.SendCommand(cmdNoIB, 0, 0, nil)
    if e != nil {
        return nil, NewPiError(7, e, "Initialize")
    }

    cb := &callbackManager{
        pi:        p,
        socket:    s,
        running:   true,
        handle:    h,
        lastLevel: ll}

    go cb.Run()
    return cb, nil
}

func (cbm *callbackManager) nextHandle() int {
    cbm.handleCount++
    return cbm.handleCount
}

func (cbm *callbackManager) Stop() {
    cbm.running = false
}

func (cbm *callbackManager) Run() {
    var data []byte
    var e error

    for cbm.running {
        data, e = cbm.socket.Read(12)
        if e != nil {
            panic(e)
        }

        flags := convertToInt16(data[2:4])
        tick := convertToUint32(data[4:8])
        level := convertToInt32(data[8:12])

        cbm.mu.Lock()
        if flags == 0 {
            changed := level ^ cbm.lastLevel
            cbm.lastLevel = level
            for pin := 0; pin < 32; pin++ {
                bit := 1 << pin
                if bit&changed == bit {
                    l := 0
                    if bit&level == bit {
                        l = 1
                    }
                    cbm.pi.Gpio(BCM(pin)).invokeCallbacks(l, tick)
                }
            }
        } else {
            if flags&notifyWatchDog == notifyWatchDog {
                pin := flags & notifyGpio
                cbm.pi.Gpio(BCM(pin)).invokeCallbacks(2, tick)
            } else if flags&notifyEvent == notifyEvent {
                // TODO Events
            }
        }
        cbm.mu.Unlock()
    }

    _ = cbm.socket.Close()
}

func (cbm *callbackManager) append(bit int) error {
    cbm.mu.Lock()
    defer cbm.mu.Unlock()

    cbm.monitor |= bit
    r, e := cbm.socket.SendCommand(cmdNotifyBegin, cbm.handle, cbm.monitor, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "append(%d)", bit)
    }

    return nil
}

func (cbm *callbackManager) remove(bit int) error {
    cbm.mu.Lock()
    defer cbm.mu.Unlock()

    cbm.monitor &= ^bit
    r, e := cbm.socket.SendCommand(cmdNotifyBegin, cbm.handle, cbm.monitor, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "remove(%d)", bit)
    }

    return nil
}
