package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"kol/ast"
	"kol/code"
	"kol/token"
	"strings"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

const (
	INTEGER_OBJ           = "INTEGER"
	FLOAT_OBJ             = "FLOAT"
	BOOLEAN_OBJ           = "BOOLEAN"
	VOID_OBJ              = "VOID"
	RETURN_VALUE_OBJ      = "RETURN_VALUE"
	BREAK_VALUE_OBJECT    = "BREAK_VALUE"
	ERROR_OBJ             = "ERROR"
	FUNCTION_OBJ          = "FUNCTION"
	STRING_OBJ            = "STRING"
	BUILTIN_OBJ           = "BUILTIN"
	ARRAY_OBJ             = "ARRAY"
	HASH_OBJ              = "HASH"
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION_OBJ"
	CLOSURE_OBJ           = "CLOSURE"
)

func TypeFromString(input string) (ObjectType, bool) {
	switch input {
	case "int":
		return INTEGER_OBJ, true
	case "float":
		return FLOAT_OBJ, true
	case "bool":
		return BOOLEAN_OBJ, true
	case "str":
		return STRING_OBJ, true
	case "fn":
		return FUNCTION_OBJ, true
	case "array":
		return ARRAY_OBJ, true
	case "map":
		return HASH_OBJ, true
	case "void":
		return VOID_OBJ, true
	default:
		return ERROR_OBJ, false
	}
}

type Object interface {
	Type() ObjectType
	Inspect() string
}
type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Float struct {
	Value float64
}

func (f *Float) Inspect() string  { return fmt.Sprintf("%v", f.Value) }
func (f *Float) Type() ObjectType { return FLOAT_OBJ }

func IsNumber(object Object) bool {
	return object.Type() == INTEGER_OBJ || object.Type() == FLOAT_OBJ
}
func GetNumber(obj Object) float64 {
	if obj.Type() == INTEGER_OBJ {
		return float64(obj.(*Integer).Value)
	} else if obj.Type() == FLOAT_OBJ {
		return obj.(*Float).Value
	} else {
		return 0
	}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Void struct{}

func (n *Void) Type() ObjectType { return VOID_OBJ }
func (n *Void) Inspect() string  { return "void" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type BreakValue struct {
	Value Object
}

func (bv *BreakValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (bv *BreakValue) Inspect() string  { return bv.Value.Inspect() }

type Error struct {
	Message  string
	Position *token.Position
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string {
	if e.Position != nil {
		return fmt.Sprintf("Error at %d:%d: %s", e.Position.Line, e.Position.Column, e.Message)
	}
	return "ERROR: " + e.Message
}

type Function struct {
	Parameters []*ast.FunctionParameter
	ReturnType *ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.Ident.String()+" "+p.Type.String())
	}
	out.WriteString("fun")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.ReturnType.String())
	out.WriteString(" {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
func (h *Hash) Type() ObjectType { return HASH_OBJ }

type CompiledFunction struct {
	Instructions  code.Instructions
	NumLocals     int
	NumParameters int
}

func (cf *CompiledFunction) Type() ObjectType { return COMPILED_FUNCTION_OBJ }
func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}

type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) Type() ObjectType { return CLOSURE_OBJ }
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
