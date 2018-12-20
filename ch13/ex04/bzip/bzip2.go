package bzip

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	w   io.WriteCloser
	cmd *exec.Cmd
	wg  sync.WaitGroup
}

func NewWriter(out io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("bzip2")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	w := &writer{w: stdin, cmd: cmd}

	w.wg.Add(1)
	go func(dest io.Writer, src io.Reader) {
		defer w.wg.Done()
		io.Copy(dest, src)
	}(out, stdout)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	return w.w.Write(data)
}

func (w *writer) Close() error {
	if err := w.w.Close(); err != nil {
		return err
	}
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	w.wg.Wait()
	return nil
}
