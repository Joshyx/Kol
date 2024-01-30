package vm

import (
	"fmt"
	"kol/code"
	"kol/object"
)

func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	switch {
	case object.IsNumber(left) && object.IsNumber(right):
		return vm.executeBinaryNumberOperation(op, left, right)
	case right.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return vm.executeBinaryStringOperation(op, left, right)
	default:
		return fmt.Errorf("unsupported types for binary operation: %s %s",
			left.Type(), right.Type())
	}
}
func (vm *VM) executeBinaryNumberOperation(
	op code.Opcode,
	left, right object.Object,
) error {
	leftValue := object.GetNumber(left)
	rightValue := object.GetNumber(right)
	switch op {
	case code.OpAdd:
		if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
			return vm.push(&object.Integer{Value: int64(leftValue + rightValue)})
		} else {
			return vm.push(&object.Float{Value: leftValue + rightValue})
		}
	case code.OpSub:
		if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
			return vm.push(&object.Integer{Value: int64(leftValue - rightValue)})
		} else {
			return vm.push(&object.Float{Value: leftValue - rightValue})
		}
	case code.OpMul:
		if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
			return vm.push(&object.Integer{Value: int64(leftValue * rightValue)})
		} else {
			return vm.push(&object.Float{Value: leftValue * rightValue})
		}
	case code.OpDiv:
		return vm.push(&object.Float{Value: leftValue / rightValue})
	default:
		return fmt.Errorf("unknown integer operator: %d", op)
	}
}
func (vm *VM) executeBinaryStringOperation(
	op code.Opcode,
	left, right object.Object,
) error {
	if op != code.OpAdd {
		return fmt.Errorf("unknown string operator: %d", op)
	}
	leftValue := left.(*object.String).Value
	rightValue := right.(*object.String).Value
	return vm.push(&object.String{Value: leftValue + rightValue})
}
func (vm *VM) executeMinusOperator() error {
	operand := vm.pop()
	switch operand.Type() {
	case object.INTEGER_OBJ:
		value := operand.(*object.Integer).Value
		return vm.push(&object.Integer{Value: -value})
	case object.FLOAT_OBJ:
		value := operand.(*object.Float).Value
		return vm.push(&object.Float{Value: -value})
	default:
		return fmt.Errorf("unsupported type for negation: %s", operand.Type())
	}
}
func (vm *VM) executeComparison(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()
	if object.IsNumber(left) && object.IsNumber(right) {
		return vm.executeNumberComparison(op, left, right)
	}

	switch op {
	case code.OpEqual:
		return vm.push(nativeBoolToBooleanObject(right == left))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObject(right != left))
	default:
		return fmt.Errorf("unknown operator: %d (%s %s)",
			op, left.Type(), right.Type())
	}
}
func (vm *VM) executeBangOperator() error {
	operand := vm.pop()
	switch operand {
	case True:
		return vm.push(False)
	case False:
		return vm.push(True)
	default:
		return fmt.Errorf("Can't convert %s to a boolean", operand.Type())
	}
}
func (vm *VM) executeNumberComparison(
	op code.Opcode,
	left, right object.Object,
) error {
	leftValue := object.GetNumber(left)
	rightValue := object.GetNumber(right)
	switch op {
	case code.OpEqual:
		return vm.push(nativeBoolToBooleanObject(rightValue == leftValue))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObject(rightValue != leftValue))
	case code.OpGreaterThan:
		return vm.push(nativeBoolToBooleanObject(leftValue > rightValue))
	case code.OpGreaterEqualsThan:
		return vm.push(nativeBoolToBooleanObject(leftValue >= rightValue))
	default:
		return fmt.Errorf("unknown operator: %d", op)
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}
func isTrue(obj object.Object) bool {
	return obj == True
}
