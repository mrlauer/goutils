// Package utils implements some generally useful things that I could not immediately find in
// the standard libraries
package utils

import (
	"io"
	"os"
	"exec"
	"bytes"
)

// ByteWriter writes into a byte array.
// (Isn't there a standard package for that?)
type ByteWriter struct {
	data []byte
}

// Write appends to the internal array.
func (w *ByteWriter) Write(b []byte) (n int, err os.Error) {
	w.data = append(w.data, b...)
	return len(b), nil
}

// Data returns the ByteWriter's byte slice.
func (w *ByteWriter) Data() []byte {
	return w.data
}

// NewByteWriter creates a ByteWriter initialized to an empty array.
func NewByteWriter() *ByteWriter {
	return &ByteWriter{}
}

// ReadLine reads the next line from a byte slice, 
// return the line and the unread portion,
// omitting the line delimiter. 
// If there is no line delimiter present the entire slice
// is returned as the line and nil as the unread portion.
// \n and \r\n are accepted as delimiters.
// Both return values are slices into the original slice, not copies.
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

// SplitLines splits a byte slice into lines. The returned values
// are not copies.
func SplitLines(s []byte) (results [][]byte) {
	rest := s
	for len(rest) > 0 {
		line, r := ReadLine(rest)
		results = append(results, line)
		rest = r
	}
	return
}

// Execute runs a command, returning stdout, stdin, and the result as an error status.
func Execute(cmd string, args ...string) (stdout []byte, stderr []byte, err os.Error) {
	return ExecuteWithStdin(nil, cmd, args...)
}

// ExecuteWithStdin runs a command, returning stdout, stdin, and the result as an error status.
func ExecuteWithStdin(stdin io.Reader, cmd string, args ...string) (stdout []byte, stderr []byte, err os.Error) {
	command := exec.Command(cmd, args...)
	command.Stdin = stdin
	wout := NewByteWriter()
	werr := NewByteWriter()
	command.Stdout = wout
	command.Stderr = werr
	err = command.Run()
	return wout.Data(), werr.Data(), err
}
