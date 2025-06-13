package wr

import (
	"testing"
)

type testBuf struct {
	b []byte
}

func (buf *testBuf) Write(data []byte) (int, error) {
	buf.b = data
	return len(buf.b), nil
}

func TestWriteToWriter(t *testing.T) {
	var data []byte = []byte("Hello")
	var buf testBuf

	_, err := Write(data, &buf)
	if err != nil {
		t.Error(err)
	}

	strBuf := string(buf.b)
	strData := string(data)

	if strBuf != strData {
		t.Error()
	}
}
