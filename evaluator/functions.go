package evaluator

import (
	"kol/object"
	"kol/token"
)

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) (*object.Environment, *object.Error) {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		typ, _ := object.TypeFromString(param.Type.Value)
		if typ != args[paramIdx].Type() {
			return nil, newError("Parameter %d not valid: Expected %s but got %s", param.Type.GetPosition(), paramIdx+1, typ, args[paramIdx].Type())
		}
		env.Set(param.Ident.Value, args[paramIdx])
	}
	return env, nil
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
		extendedEnv, err := extendFunctionEnv(fn, args)
		if err != nil {
			return err
		}
		evaluated := Eval(fn.Body, extendedEnv)
		returnValue := unwrapReturnValue(evaluated)
		typ, _ := object.TypeFromString(fn.ReturnType.Value)
		if returnValue.Type() != typ {
			return newError("Returned type %s doesn't match expected type %s", fn.Body.GetPosition(), returnValue.Type(), fn.ReturnType.Value)
		}
		return returnValue
	case *object.Builtin:
		if result := fn.Fn(args...); result != nil {
			return result
		}
		return VOID
	default:
		return newError("not a function: %s", pos, fn.Type())
	}
}
