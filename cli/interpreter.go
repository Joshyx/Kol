package cli

import (
	"fmt"
	"kol/evaluator"
	"kol/lexer"
	"kol/object"
	"kol/parser"
)

func StartInterpreter(input string) {

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println("Errors while parsing file")
		for _, e := range p.Errors() {
			fmt.Println(e)
		}
	}

	env := object.NewEnvironment()
	evaluator.Eval(program, env)
}
