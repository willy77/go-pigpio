package pigpio

type File struct {
	pi       *Pi
	handle   int
	mode     FileMode
	filename string
}

type FileSeekFrom int

const (
	FromStart   FileSeekFrom = 0
	FromCurrent FileSeekFrom = 1
	FromEnd     FileSeekFrom = 1
)

func (f *File) Handle() int      { return f.handle }
func (f *File) Mode() FileMode   { return f.mode }
func (f *File) Filename() string { return f.filename }

func (f *File) Close() error {
	r, e := f.pi.socket.SendCommand(cmdFileClose, f.handle, 0, nil)
	if e != nil || r < 0 {
		return newPiError(r, e, "Close(handle: %d)", f.handle)
	}

	return nil
}

func (f *File) Read(count int) ([]byte, error) {
	r, e := f.pi.socket.SendCommand(cmdFileRead, f.handle, count, nil)
	if e != nil || r < 0 {
		return nil, newPiError(r, e, "Read(handle: %d, count: %d)", f.handle, count)
	}

	if r >= 0 {
		data, e := f.pi.socket.Read(r)
		if e != nil || r < 0 {
			return nil, newPiError(r, e, "Read(handle: %d, count: %d)", f.handle, count)
		}

		return data, nil
	}

	return []byte{}, nil
}

func (f *File) Write(data []byte) error {
	r, e := f.pi.socket.SendCommand(cmdFileWrite, f.handle, 0, data)
	if e != nil || r < 0 {
		return newPiError(r, e, "Write(handle: %d, data: %v)", f.handle, data)
	}

	return nil
}

func (f *File) Seek(offset int, from FileSeekFrom) (int, error) {
	r, e := f.pi.socket.SendCommand(cmdFileSeek, f.handle, offset, convertToBytes(int(from)))
	if e != nil || r < 0 {
		return -1, newPiError(r, e, "Seek(handle: %d, offset: %v, from: %d)", f.handle, offset, from)
	}

	return r, nil
}
