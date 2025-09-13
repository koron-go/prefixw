package prefixw_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/koron-go/prefixw"
)

// TestClose tests about Close().
//   - double close
//   - close without carried data
//   - Write() after Close()
func TestClose(t *testing.T) {
	b := &bytes.Buffer{}
	w := prefixw.New(b, "PREFIX ")
	if err := w.Close(); err != nil {
		t.Fatalf("1st close failed: %s", err)
	}
	if _, err := w.Write([]byte("hello\n")); !errors.Is(err, io.EOF) {
		t.Fatalf("write after close should fail:\nwant=io.EOF\ngot=%s", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("2nd close failed: %s", err)
	}
	if s := b.String(); s != "" {
		t.Fatalf("unexpected output:\nwant=%q\ngot=%q", "", s)
	}
}

type limitedBuffer struct {
	buf []byte
}

func (lb *limitedBuffer) Write(p []byte) (int, error) {
	n := len(p)
	if n > len(lb.buf) {
		n = len(lb.buf)
	}
	copy(lb.buf, p[:n])
	return n, nil
}

func TestInsufficientWriter(t *testing.T) {
	lb := &limitedBuffer{buf: make([]byte, 16)}
	w := prefixw.New(lb, "[PREFIX] ")
	n, err := w.Write([]byte("Hello World with prefixw\n"))
	exp := "[PREFIX] Hello W"
	if out := string(lb.buf); out != exp {
		t.Errorf("unexpected output:\nwant=%s\ngot=%s", exp, out)
	}
	if n != 16 {
		t.Errorf("unexpected written bytes: want=16 got=%d", n)
	}
	if err == nil {
		t.Fatalf("unexpected success")
	}
	if !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("unexpected error:\nwant=%s\ngot=%s", io.ErrShortWrite, err)
	}
}
