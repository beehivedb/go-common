package logs

import (
	"strings"
	"testing"
)

type mock struct {
	message string
}

func (m *mock) Write(p []byte) (n int, err error) {
	m.message = string(p)
	return len(p), nil
}

func TestLogs(t *testing.T) {
	m := new(mock)
	SetOutput(m)
	SetLevel("Info")
	Trace("abc")
	if strings.Contains(m.message, "abc") {
		t.Fail()
	}

	Info("bcd")

	if strings.Contains(m.message, "bcd") == false {
		t.Fail()
	}
}
