package mrlutils

import (
	"testing"
	"strings"
	"bytes"
)

func ok(cond bool, t *testing.T, msg string, msgParams ...interface{}) {
	if !cond {
		t.Errorf(msg, msgParams...)
	}
}

func TestLines1(t *testing.T) {
	var s []byte = []byte("foo\nbar\r\nbaz")
	split := SplitLines(s)
	ok(len(split) == 3, t, "Bad split length")
	for i, val := range []string{"foo", "bar", "baz"} {
		ok(bytes.Equal(split[i], []byte(val)), t, "Bad line %s", split[i])
	}
}

func TestLines2(t *testing.T) {
	var s []byte = []byte("foo\nbar\r\n\nbaz\n")
	split := SplitLines(s)
	ok(len(split) == 4, t, "Bad split length")
	for i, val := range []string{"foo", "bar", "", "baz"} {
		ok(bytes.Equal(split[i], []byte(val)), t, "Bad line %s", split[i])
	}
}

func TestLines3(t *testing.T) {
	var s []byte = []byte("\n")
	split := SplitLines(s)
	ok(len(split) == 1, t, "Bad split length")
}

func TestLines4(t *testing.T) {
	var s []byte
	split := SplitLines(s)
	ok(len(split) == 0, t, "Bad split length")
}

func TestExecuteStdout(t *testing.T) {
	stdout, stderr, err := Execute("echo", "foo")
	if string(stdout) != "foo\n" {
		t.Errorf("bad stdout <%s>", string(stdout))
	}
	if string(stderr) != "" {
		t.Errorf("bad stderr")
	}
	if err != nil {
		t.Errorf("bad error status")
	}
}

func TestExecuteStderr(t *testing.T) {
	stdout, stderr, err := Execute("cat", "doesnotexist")
	if string(stdout) != "" {
		t.Errorf("bad stdout <%s>", string(stdout))
	}
	if !strings.HasSuffix(string(stderr), "No such file or directory\n") {
		t.Errorf("bad stderr <%s>", string(stderr))
	}
	if err == nil {
		t.Errorf("bad error status")
	}
}

func TestExecuteBashStderr(t *testing.T) {
	stdout, stderr, err := Execute("bash", "-c", "echo foo 1>&2")
	if string(stdout) != "" {
		t.Errorf("bad stdout <%s>", string(stdout))
	}
	if !strings.HasSuffix(string(stderr), "foo\n") {
		t.Errorf("bad stderr <%s>", string(stderr))
	}
	if err != nil {
		t.Errorf("bad error status")
	}
}

func TestExecuteErr(t *testing.T) {
	stdout, stderr, err := Execute("xyzabcfoobar")
	if string(stdout) != "" {
		t.Errorf("bad stdout")
	}
	if string(stderr) != "" {
		t.Errorf("bad stderr")
	}
	if err == nil {
		t.Errorf("bad error status")
	}
}
