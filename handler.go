package lab2

import (
	"io"
	"bufio"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := new(strings.Builder)
	_, err := io.Copy(buffer, ch.Input)
	if err != nil {
		return err
	}

	result, err := PostfixToInfix(buffer.String())
	if err != nil {
		return err
	}

	w := bufio.NewWriter(ch.Output) 
  _, wErr := w.WriteString(result)
  if wErr != nil {
    return wErr
  }
	w.Flush()
	return nil
}
