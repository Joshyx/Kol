package main

import (
	"fmt"
	"kol/cli"
	"os"
	"os/user"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		startRepl()
	} else {
		startInterpreter(args[0])
	}
}

func startInterpreter(fileName string) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Encountered Error: %s\n", err)
		return
	}

	text := string(content)

	cli.StartInterpreter(text)
}

func startRepl() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Kol programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	cli.StartRepl(os.Stdin, os.Stdout)
}

func printParserErrors(errors []string) {
	for _, msg := range errors {
		fmt.Println("\t" + msg + "\n")
	}
}
