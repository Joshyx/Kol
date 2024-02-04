package cli

import (
	"bufio"
	"fmt"
	"io"
	"kol/evaluator"
	"kol/lexer"
	"kol/object"
	"kol/parser"
	"os"
	"os/user"
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
	result := evaluator.Eval(program, env)
	if result.Type() == object.ERROR_OBJ {
		fmt.Println(result.Inspect())
	}
}
func StartInterpretedRepl() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Kol programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	scanner := bufio.NewScanner(os.Stdin)
	env := object.NewEnvironment()
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
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Type() != object.VOID_OBJ {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}
}
