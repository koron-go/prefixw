/*
Package prefixw provides a writer which puts with prefix to each lines.
*/
package prefixw

import (
	"bytes"
	"io"
	"sync"
)

// Writer implements io.Writer with prefix each lines.
type Writer struct {
	mu sync.Mutex    // mutex to write
	w  io.Writer     // base io.Writer
	p  []byte        // prefix bytes, not changed.
	c  *bytes.Buffer // carried data, which not end with '\n'
}

// New creates a new prefix Writer.
func New(w io.Writer, prefix string) *Writer {
	return &Writer{
		w: w,
		p: []byte(prefix),
	}
}

// Write writes data to base Writer with prefix.
func (w *Writer) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.w == nil {
		return 0, io.EOF
	}
	size := len(p)

	// combined carried data to output.
	if w.c != nil {
		w.c.Write(p)
		p = w.c.Bytes()
		w.c = nil
	}

	b := new(bytes.Buffer)
	for len(p) > 0 {
		n := bytes.IndexByte(p, '\n')
		if n < 0 {
			w.c = new(bytes.Buffer)
			w.c.Write(p)
			break
		}
		b.Write(w.p)
		b.Write(p[:n+1])
		p = p[n+1:]
	}

	if b.Len() > 0 {
		n, err := b.WriteTo(w.w)
		if err != nil {
			return int(n), err
		}
	}
	return size, nil
}

func (w *Writer) flush() error {
	if w.c == nil {
		return nil
	}
	b := new(bytes.Buffer)
	b.Write(w.p)
	w.c.WriteTo(b)
	w.c = nil
	b.WriteByte('\n')
	_, err := b.WriteTo(w.w)
	return err
}

// Close flushes buffered data and closes Writer.
func (w *Writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.w == nil {
		// no errors for second or more close.
		return nil
	}
	err := w.flush()
	w.w = nil
	return err
}

var _ io.WriteCloser = &Writer{}
