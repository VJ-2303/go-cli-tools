package hello

import (
	"fmt"
	"io"
	"os"
)

var Output io.Writer = os.Stdout

func Print() {
	fmt.Fprintln(Output, "Hello, World")
}
