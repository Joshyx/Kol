package evaluator

import (
	"kol/object"
	"kol/token"
)

func evalPrefixExpression(operator string, right object.Object, pos token.Position) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right, pos)
	case "-":
		return evalMinusPrefixOperatorExpression(right, pos)
	default:
		return newError("unknown operator: %s%s", pos, operator, right.Type())
	}
}
func evalBangOperatorExpression(right object.Object, pos token.Position) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return newError("unkown operator: !%s", pos, right.Type())
	}
}

func evalMinusPrefixOperatorExpression(right object.Object, pos token.Position) object.Object {
	switch right.Type() {
	case object.INTEGER_OBJ:
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	case object.FLOAT_OBJ:
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	default:
		return newError("unknown operator: -%s", pos, right.Type())
	}
}
