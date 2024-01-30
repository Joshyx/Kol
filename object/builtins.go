package object

import (
	"fmt"
	"strconv"
)

var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"println",
		&Builtin{func(args ...Object) Object {
			for _, o := range args {
				fmt.Print(o.Inspect())
			}
			fmt.Println()
			return nil
		},
		},
	},
	{
		"len",
		&Builtin{func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
	{
		"str",
		&Builtin{func(args ...Object) Object {
			var result string

			for _, arg := range args {
				result += string(arg.Inspect())
			}

			return &String{Value: result}
		},
		},
	},
	{
		"int",
		&Builtin{func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *String:
				result, err := strconv.ParseInt(arg.Value, 10, 64)

				if err != nil {
					return newError("Could not parse '%s' to a int", arg.Value)
				}
				return &Integer{Value: result}
			default:
				return newError("argument to `int` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
	{
		"float",
		&Builtin{func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *String:
				result, err := strconv.ParseFloat(arg.Value, 10)

				if err != nil {
					return newError("Could not parse '%s' to a float", arg.Value)
				}
				return &Float{Value: result}
			default:
				return newError("argument to `float` not supported, got %s",
					args[0].Type())
			}
		},
		},
	},
	{
		"push",
		&Builtin{func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			newElements := make([]Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &Array{Elements: newElements}
		},
		},
	},
	{
		"remove",
		&Builtin{func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `remove` must be ARRAY, got %s",
					args[0].Type())
			}
			if args[1].Type() != INTEGER_OBJ {
				return newError("index must be a INTEGER, got %s",
					args[0].Type())
			}

			arr := args[0].(*Array).Elements
			index := args[1].(*Integer).Value
			length := len(arr)

			if index >= int64(length) {
				return newError("Index %d to big for array of size %d", index, length)
			}

			newArray := make([]Object, length-1, length-1)

			copy(newArray, append(arr[:index], arr[index+1:]...))

			return &Array{Elements: newArray}
		},
		},
	},
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}
func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
