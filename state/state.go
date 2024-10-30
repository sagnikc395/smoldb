package state

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sagnikc395/smoldb/errors"
)

type InputBuffer struct {
	buffer        []byte
	buffer_length int
	input_length  int
}

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{
		buffer:        make([]byte, 0),
		buffer_length: 0,
		input_length:  0,
	}
}

func (ib *InputBuffer) ReadInput() {
	reader := bufio.NewReader(os.Stdin)
	bytes_read, _, _ := reader.ReadLine()
	// ib.buffer = bytes_read
	// ib.buffer_length = len(bytes_read)

	if len(bytes_read) <= 0 {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(errors.EXIT_FAILURE)
	}

	ib.input_length = len(bytes_read) - 1
	ib.buffer[len(bytes_read[:])-1] = 0
}
