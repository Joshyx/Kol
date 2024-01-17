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
}
