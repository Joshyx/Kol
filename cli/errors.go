package cli

import "fmt"

func printParserErrors(errors []string) {
	for _, msg := range errors {
		fmt.Println("\t" + msg + "\n")
	}
}
