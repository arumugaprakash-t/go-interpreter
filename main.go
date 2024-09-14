package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
)

func main() {
	fmt.Println("Hello! Welcome to go interpreter. This evaluates Monkey programming language")
	fmt.Println("Feel free to enter your commands")
	repl.Start(os.Stdin, os.Stdout)
}
