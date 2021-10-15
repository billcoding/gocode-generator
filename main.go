package main

import (
	"fmt"
	"github.com/billcoding/gocode-generator/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
