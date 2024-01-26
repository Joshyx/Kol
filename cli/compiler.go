package cli

import (
	"bufio"
	"fmt"
	"io"
	"kol/compiler"
	"kol/lexer"
	"kol/parser"
	"kol/vm"
	"os"
)

const PROMPT = ">>"

var in, out = os.Stdin, os.Stdout

func StartCompiler(input string) {

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println("Errors while parsing file")
		for _, e := range p.Errors() {
			fmt.Println(e)
		}
	}

	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
		return
	}
	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		return
	}
	stackTop := machine.StackTop()
	fmt.Println(stackTop.Inspect())
}
func StartCompiledRepl() {

	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(p.Errors())
			continue
		}
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}
		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}
		stackTop := machine.LastPoppedStackElem()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}
