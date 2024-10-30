package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sagnikc395/smoldb/exitcodes"
	"github.com/sagnikc395/smoldb/repl"
	"github.com/sagnikc395/smoldb/state"
)

func main() {
	var ib state.InputBuffer

	for {
		repl.PrintPrompt()
		ib.ReadInput()

		if strings.Compare(string(ib.Buffer), ".exit") == 0 {
			os.Exit(exitcodes.EXIT_SUCCESS)
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", ib.Buffer)
		}
	}
}
