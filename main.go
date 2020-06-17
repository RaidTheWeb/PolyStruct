package main

import (
	"fmt"
	"os"

	"github.com/RaidTheWeb/PolyStruct/repl"
)

func main() {
	fmt.Println("PolyStruct Interpreter ( build: 2020.1-go )")
	repl.Start(os.Stdin, os.Stdout)
}
