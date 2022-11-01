package pigpio

type ScriptStatus int

const (
    Unknown    ScriptStatus = -1
    Initiating ScriptStatus = 0
    Halted     ScriptStatus = 1
    Running    ScriptStatus = 2
    Waiting    ScriptStatus = 3
    Failed     ScriptStatus = 4
)

type Script struct {
    pi     *Pi
    handle int
    code   string
}

func (s *Script) Handle() int  { return s.handle }
func (s *Script) Code() string { return s.code }

func (s *Script) Run(params ...int) error {
    r, e := s.pi.socket.SendCommand(cmdScriptRun, s.handle, 0, convertToBytes(params[:]...))
    if e != nil || r < 0 {
        return NewPiError(r, e, "Run(handle: %d, code: %s, params: %v)", s.handle, s.code, params[:])
    }

    return nil
}

func (s *Script) Update(params ...int) error {
    r, e := s.pi.socket.SendCommand(cmdScriptUpdate, s.handle, 0, convertToBytes(params[:]...))
    if e != nil || r < 0 {
        return NewPiError(r, e, "Update(handle: %d, code: %s, params: %v)", s.handle, s.code, params[:])
    }

    return nil
}

func (s *Script) Status() (ScriptStatus, []int, error) {
    r, e := s.pi.socket.SendCommand(cmdScriptStatus, s.handle, 0, nil)
    if e != nil || r < 0 {
        return Unknown, nil, NewPiError(r, e, "Status(handle: %d)", s.handle)
    }

    if r > 0 {
        data, e := s.pi.socket.Read(r)
        dataInts := convertToInt32Array(data)
        if e != nil || dataInts[0] < 0 {
            return Unknown, nil, NewPiError(int(dataInts[0]), e, "Status(handle: %d)", s.handle)
        }

        return ScriptStatus(dataInts[0]), dataInts, nil
    }

    return ScriptStatus(r), []int{}, nil
}

func (s *Script) Stop() error {
    r, e := s.pi.socket.SendCommand(cmdScriptStop, s.handle, 0, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "Stop(handle: %d)", s.handle)
    }

    return nil
}

func (s *Script) Delete() error {
    r, e := s.pi.socket.SendCommand(cmdScriptDelete, s.handle, 0, nil)
    if e != nil || r < 0 {
        return NewPiError(r, e, "Delete(handle: %d)", s.handle)
    }

    return nil
}
