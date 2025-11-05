package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Input  io.Reader
	Output io.Writer
}

type option func(*Counter) error

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.Input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.Output = output

		return nil
	}
}

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Counter) Lines() int {
	lines := 0

	input := bufio.NewScanner(c.Input)

	for input.Scan() {
		lines++
	}

	return lines
}

func Main() error {
	c, err := NewCounter()
	if err != nil {
		return err
	}
	fmt.Println(c.Lines())

	return nil
}
