package pigpio

import "strings"

type FileMode int

const (
	FileRead      FileMode = 1
	FileWrite     FileMode = 2
	FileReadWrite FileMode = 3
	FileAppend    FileMode = 4
	FileCreate    FileMode = 8
	FileTrunc     FileMode = 16
)

func (p *Pi) OpenFile(filename string, mode FileMode) (*File, error) {
	r, e := p.socket.SendCommand(cmdFileOpen, int(mode), 0, []byte(filename))
	if e != nil || r < 0 {
		return nil, newPiError(r, e, "OpenFile(filename: %s, mode: %d)", filename, mode)
	}

	return &File{pi: p, handle: r, mode: mode, filename: filename}, nil
}

func (p *Pi) ListFiles(pattern string) ([]string, error) {
	r, e := p.socket.SendCommand(cmdFileList, 60000, 0, []byte(pattern))
	if e != nil || r < 0 {
		return nil, newPiError(r, e, "ListFiles(pattern: %s)", pattern)
	}

	if r > 0 {
		data, e := p.socket.Read(r)
		if e != nil {
			return nil, newPiError(9, e, "ListFiles(pattern: %s)", pattern)
		}

		list := string(data)
		return strings.Split(list, "\n"), nil
	}

	return []string{}, nil
}
