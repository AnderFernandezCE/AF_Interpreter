package main

import (
	"af/src/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello User! This is the AF programming language!")
	fmt.Println("Type commands")
	//TODO: fmt.Println("Type \"credits\" for more information.")
	repl.Start(os.Stdin, os.Stdout)
}
