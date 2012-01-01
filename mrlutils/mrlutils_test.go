package mrlutils

import (
	"testing"
	"strings"
)

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
