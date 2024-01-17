package evaluator

import (
	"kol/object"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"str": {
		Fn: func(args ...object.Object) object.Object {
			var result string

			for _, arg := range args {
				result += string(arg.Inspect())
			}

			return &object.String{Value: result}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				result, err := strconv.ParseInt(arg.Value, 10, 64)

				if err != nil {
					return newError("Could not parse '%s' to a string", arg.Value)
				}
				return &object.Integer{Value: result}
			default:
				return newError("argument to `int` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
	"remove": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `remove` must be ARRAY, got %s",
					args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("index must be a INTEGER, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array).Elements
			index := args[1].(*object.Integer).Value
			length := len(arr)

			if index >= int64(length) {
				return newError("Index %d to big for array of size %d", index, length)
			}

			newArray := make([]object.Object, length-1, length-1)

			copy(newArray, append(arr[:index], arr[index+1:]...))

			return &object.Array{Elements: newArray}
		},
	},
}
