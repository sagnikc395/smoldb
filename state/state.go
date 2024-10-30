package state

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sagnikc395/smoldb/exitcodes"
)

type InputBuffer struct {
	Buffer       []byte
	BufferLength int
	InputLength  int
}

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{
		Buffer:       make([]byte, 0),
		BufferLength: 0,
		InputLength:  0,
	}
}

func (ib *InputBuffer) ReadInput() {
	reader := bufio.NewReader(os.Stdin)
	bytes_read, _, _ := reader.ReadLine()
	// ib.buffer = bytes_read
	// ib.buffer_length = len(bytes_read)

	if len(bytes_read) <= 0 {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(exitcodes.EXIT_FAILURE)
	}

	ib.InputLength = len(bytes_read) - 1
	ib.Buffer[len(bytes_read[:])-1] = 0
}
