package main

import (
	"fmt"
	kol "kol/cli"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Kol",
		Usage: "A compiler/interpreter for the best language",
		Commands: []*cli.Command{
			{
				Name:    "interpret",
				Aliases: []string{"i"},
				Usage:   "Start Interpreter",
				Action: func(cCtx *cli.Context) error {
					startInterpreter(cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "compile",
				Aliases: []string{"c"},
				Usage:   "Start Compiler",
				Action: func(cCtx *cli.Context) error {
					startCompiler(cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startInterpreter(fileName string) {
	if len(fileName) == 0 {
		kol.StartInterpretedRepl()
		return
	}
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Encountered Error: %s\n", err)
		return
	}

	text := string(content)

	kol.StartInterpreter(text)
}
func startCompiler(fileName string) {
	if len(fileName) == 0 {
		kol.StartCompiledRepl()
		return
	}
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Encountered Error: %s\n", err)
		return
	}

	text := string(content)

	kol.StartCompiler(text)
}
