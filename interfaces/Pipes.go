package interfaces

import (
	"os"
	"io"
)


func PipeExample() error {
	r, w := io.Pipe()

	go writeToPipe(*w, "test\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}

func writeToPipe(w io.PipeWriter, message string) {
	w.Write([]byte(message))
	w.Close()
}
