package evaluator

import (
	"kol/object"
)

func evalInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	switch {
	case object.IsNumber(left) && object.IsNumber(right):
		return evalNumberInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalNumberInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := object.GetNumber(left)
	rightVal := object.GetNumber(right)

	switch operator {
	case "+":
		if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
			return &object.Integer{Value: int64(leftVal + rightVal)}
		} else {
			return &object.Float{Value: leftVal + rightVal}
		}
	case "-":
		if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
			return &object.Integer{Value: int64(leftVal - rightVal)}
		} else {
			return &object.Float{Value: leftVal - rightVal}
		}
	case "*":
		if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
			return &object.Integer{Value: int64(leftVal * rightVal)}
		} else {
			return &object.Float{Value: leftVal * rightVal}
		}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
