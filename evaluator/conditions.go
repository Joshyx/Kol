package evaluator

import (
	"kol/ast"
	"kol/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}
	if condition.Type() != object.BOOLEAN_OBJ {
		return newError("%s is not of type BOOLEAN and can't be used in an if-statement", condition.Type())
	}

	if condition == TRUE {
		return Eval(ie.Consequence, env)
	} else if condition == FALSE && ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}
