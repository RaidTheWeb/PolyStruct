package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RaidTheWeb/PolyStruct/evaluator"
	"github.com/RaidTheWeb/PolyStruct/lexer"
	"github.com/RaidTheWeb/PolyStruct/object"
	"github.com/RaidTheWeb/PolyStruct/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, isfile) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
        if isfile != true {
		    fmt.Printf(PROMPT)
        }
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {

	for _, msg := range errors {
		io.WriteString(out, "ERROR: "+msg+"\n")
	}
}
