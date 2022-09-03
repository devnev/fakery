package testutils

import (
	"bytes"
	"io"
	"os"
	"sync/atomic"
	"testing"
)

func RecordStderr(t *testing.T) *Recorder {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rec := &Recorder{
		buf:  bytes.NewBuffer(nil),
		done: make(chan struct{}),
	}
	stop := make(chan struct{})
	rec.stop.Store(&stop)
	orig := os.Stderr
	os.Stderr = w

	go func() {
		defer close(rec.done)
		n, err := io.Copy(rec.buf, io.TeeReader(r, orig))
		if err != nil {
			t.Errorf("copying output failed: %v", err)
		} else {
			t.Logf("copying output finished, copied %d bytes", n)
		}
	}()
	go func() {
		<-stop
		os.Stderr = orig
		w.Close()
	}()
	t.Cleanup(func() {
		<-rec.Stop()
	})

	return rec
}

type Recorder struct {
	buf  *bytes.Buffer
	stop atomic.Pointer[chan struct{}]
	done chan struct{}
}

func (r *Recorder) Stop() <-chan struct{} {
	if s := r.stop.Swap(nil); s != nil {
		close(*s)
	}
	return r.done
}

func (r *Recorder) Done() <-chan struct{} {
	return r.done
}

func (r *Recorder) Out() string {
	<-r.done
	return r.buf.String()
}
