package cmd

import (
	"io"
	"io/ioutil"
	"os/exec"
)

type Cmd struct {
	Cmd      *exec.Cmd
	PipeTo   *Cmd
	PipeFrom *Cmd
}

func Command(name string, args ...string) *Cmd {
	return &Cmd{Cmd: exec.Command(name, args...)}
}

func (c *Cmd) Pipe(name string, args ...string) (*Cmd, error) {
	r, err := c.Cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	cmd := Command(name, args...)

	cmd.PipeFrom = c
	cmd.Cmd.Stdin = r

	c.PipeTo = cmd
	return cmd, nil
}

func (c *Cmd) PipeCmd(cmd *Cmd) error {
	r, err := c.Cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.PipeFrom = c
	cmd.Cmd.Stdin = r

	c.PipeTo = cmd

	return nil
}

func (c *Cmd) Start() ([]byte, error) {
	var (
		b   []byte
		err error
		r   io.ReadCloser
	)

	if c.PipeTo == nil {
		r, err = c.stdOut()
		if err != nil {
			return b, err
		}

	}

	if err := c.Cmd.Start(); err != nil {
		return b, err
	}

	if c.PipeTo != nil {
		return c.PipeTo.Start()
	}

	b, err = ioutil.ReadAll(r)
	if err != nil {
		return b, err
	}

	if err = c.Cmd.Wait(); err != nil {
		return b, err
	}

	return b, nil
}

func (c *Cmd) stdOut() (io.ReadCloser, error) {
	if c.PipeTo == nil {
		return c.Cmd.StdoutPipe()
	}

	return c.PipeTo.stdOut()
}
