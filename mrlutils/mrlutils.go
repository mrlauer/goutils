package mrlutils

import(
	"os"
	"exec"
)

type byteWriter struct {
	data *[]byte
}

func (w *byteWriter) Write (b []byte) (n int, err os.Error) {
	*w.data = append(*w.data, b...)
	return len(b), nil
}

func Execute(cmd string, args... string) (stdout []byte, stderr []byte, err os.Error) {
	command := exec.Command(cmd, args...)
	command.Stdout = &byteWriter{&stdout}
	command.Stderr = &byteWriter{&stderr}
	err = command.Run()
	return
}

