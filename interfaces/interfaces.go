package interfaces

import (
	"os"
	"io"
	"fmt"
)

func Copy(in ReadSeeker, out Writer) error {
	w := io.MultiWriter(out, os.Stdout)

	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()
	return nil
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadSeeker interface {
	Reader
	Seeker
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
