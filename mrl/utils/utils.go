package utils

import (
	"os"
	"exec"
	"bytes"
)

// A writer that appends to an internal byte array
// (Isn't there a standard package for that?)
type ByteWriter struct {
	data []byte
}

// Write by appending to the internal array
func (w *ByteWriter) Write(b []byte) (n int, err os.Error) {
	w.data = append(w.data, b...)
	return len(b), nil
}

// Return the ByteWriter's data
func (w *ByteWriter) Data() []byte {
	return w.data
}

// Create a ByteWriter
func NewByteWriter() *ByteWriter {
	return &ByteWriter{}
}

// Read the next line from a byte slice. Return the line and the unread portion,
// omitting the line delimiter. If there is no line delimiter present, or if
// the array ends with a delimiter, return the
// entire slice as the line and an empty unread portion.
// \n and \r\n are accepted as delimiters
func ReadLine(s []byte) (line, rest []byte) {
	idx := bytes.IndexByte(s, '\n')
	if idx < 0 {
		return s, nil
	}
	l := idx
	if l > 0 && s[l-1] == byte('\r') {
		l -= 1
	}
	return s[:l], s[idx+1:]
}

// Split a byte array into lines
func SplitLines(s []byte) (results [][]byte) {
	rest := s
	for len(rest) > 0 {
		line, r := ReadLine(rest)
		results = append(results, line)
		rest = r
	}
	return
}

// Execute a command, returning stdout, stdin, and the result as an error status
func Execute(cmd string, args ...string) (stdout []byte, stderr []byte, err os.Error) {
	command := exec.Command(cmd, args...)
	wout := NewByteWriter()
	werr := NewByteWriter()
	command.Stdout = wout
	command.Stderr = werr
	err = command.Run()
	return wout.Data(), werr.Data(), err
}
