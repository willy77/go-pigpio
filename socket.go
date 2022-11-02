package pigpio

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Socket struct {
	Connection net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
	address    *net.TCPAddr
	host       string
	port       int
	rwMutex    sync.Mutex
}

func (s *Socket) Address() *net.TCPAddr { return s.address }
func (s *Socket) Host() string          { return s.host }
func (s *Socket) Port() int             { return s.port }

// NewSocket
//
// Create a new socket connection to host:port
func NewSocket(host string, port int) (*Socket, error) {
	var e error
	var tcpAddr *net.TCPAddr
	var con net.Conn

	hostAddr := fmt.Sprintf("%s:%d", host, port)
	tcpAddr, e = net.ResolveTCPAddr("tcp", hostAddr)
	if e != nil {
		return nil, newPiError(1, e, "Error resolve address")
	}

	con, e = net.DialTCP("tcp", nil, tcpAddr)
	if e != nil {
		return nil, newPiError(2, e, "Error connect to host")
	}

	return &Socket{
		host:       host,
		port:       port,
		address:    tcpAddr,
		writer:     bufio.NewWriter(con),
		reader:     bufio.NewReader(con),
		Connection: con}, nil
}

func (s *Socket) Close() error {
	if s.Connection != nil {
		s.rwMutex.Lock()
		defer s.rwMutex.Unlock()
		return s.Connection.Close()
	}

	return nil
}

func (s *Socket) Read(len int) ([]byte, error) {
	buf := make([]byte, len)
	s.rwMutex.Lock()
	r, e := s.reader.Read(buf)
	s.rwMutex.Unlock()

	if e != nil || r != len {
		return nil, newPiError(4, e, "cmdGR(%d)", len)
	}

	return buf, nil
}

func (s *Socket) SendCommand(command Command, p1 int, p2 int, extData []byte) (int, error) {
	bufSize := 16
	if extData != nil {
		bufSize += len(extData)
	}

	buf := make([]byte, 0)
	buf = append(buf, convertToBytes(int(command), p1, p2)[:]...)
	if extData != nil && len(extData) > 0 {
		buf = append(buf, convertToBytes(len(extData))[:]...)
		buf = append(buf, extData[:]...)
	} else {
		buf = append(buf, convertToBytes(0)[:]...)
	}

	s.rwMutex.Lock()
	w, e := s.writer.Write(buf)
	fe := s.writer.Flush()
	s.rwMutex.Unlock()

	if e != nil || w != len(buf) {
		return -1, newPiError(3, e, "SendCommand(%d, %d, %d, %v)", int(command), p1, p2, extData)
	}

	if fe != nil {
		return -1, newPiError(3, fe, "SendCommand.Flush(%d, %d, %d, %v)", int(command), p1, p2, extData)
	}

	response := make([]byte, 16)
	_, e = s.reader.Read(response)
	if e != nil {
		return -1, newPiError(5, e, "SendCommand.Read(%d, %d, %d, %v)", int(command), p1, p2, extData)
	}

	return convertToInt32(response[12:]), nil
}
