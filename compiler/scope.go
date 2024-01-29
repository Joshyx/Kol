package compiler

import "kol/code"

type CompilationScope struct {
	instructions        code.Instructions
	lastInstruction     EmittedInstruction
	previousInstruction EmittedInstruction
}
type EmittedInstruction struct {
	Opcode   code.Opcode
	Position int
}

func (c *Compiler) enterScope() {
	scope := CompilationScope{
		instructions:    code.Instructions{},
		lastInstruction: EmittedInstruction{}, previousInstruction: EmittedInstruction{},
	}
	c.scopes = append(c.scopes, scope)
	c.scopeIndex++
	c.symbolTable = NewEnclosedSymbolTable(c.symbolTable)
}
func (c *Compiler) leaveScope() code.Instructions {
	instructions := c.currentInstructions()
	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--

	c.symbolTable = c.symbolTable.Outer
	return instructions
}
