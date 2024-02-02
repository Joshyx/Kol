package evaluator

import (
	"kol/object"
	"kol/token"
)

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}
func applyFunction(fn object.Object, args []object.Object, pos token.Position) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		if len(fn.Parameters) != len(args) {
			return newError("Wrong number of arguments: want=%d, got=%d", pos, len(fn.Parameters), len(args))
		}
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		if result := fn.Fn(args...); result != nil {
			return result
		}
		return NULL
	default:
		return newError("not a function: %s", pos, fn.Type())
	}
}
