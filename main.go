package main

import (
	"fmt"
	"os"
    "flag"

	"github.com/RaidTheWeb/PolyStruct/repl"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [inputfile]\n", os.Args[0])
    flag.PrintDefaults()
    os.Exit(2)
}

func main() {
    flag.Usage = usage
    flag.Parse()

    args := flag.Args()
    if len(args) < 1 {
        fmt.Println("PolyStruct Interpreter ( build: 2020.1-go )")
	repl.Start(os.Stdin, os.Stdout, false)
    }
    fmt.Printf("opening %s\n", args[0]);
    file, err := os.Open(args[0])
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %s", err)
    }
    repl.Start(file, os.Stdout, true)
}
