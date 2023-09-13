package main

import (
	"io"
	"fmt"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
