package evaluator

import (
	"kol/ast"
	"kol/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	result, err := isTrue(ie.Condition, env)
	if err != nil {
		return err
	}

	if result {
		return Eval(ie.Consequence, env)
	} else if !result && ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalForExpression(ie *ast.ForExpression, env *object.Environment) object.Object {
	result, err := isTrue(ie.Condition, env)
	if err != nil {
		return err
	}

	var obj object.Object
	for result {
		obj = Eval(ie.Consequence, env)

		result, err = isTrue(ie.Condition, env)
		if err != nil {
			return err
		}
	}
	if ie.Alternative != nil {
		obj = Eval(ie.Alternative, env)
	} else {
		return NULL
	}
	return obj
}
func isTrue(ex ast.Expression, env *object.Environment) (bool, object.Object) {
	condition := Eval(ex, env)
	if isError(condition) {
		return false, condition
	}
	if condition.Type() != object.BOOLEAN_OBJ {
		return false, newError("%s is not of type BOOLEAN and can't be used as a condition", ex.GetPosition(), condition.Type())
	}
	return condition == TRUE, nil
}
