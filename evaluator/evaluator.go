package evaluator

import (
	"fmt"
	"kol/ast"
	"kol/object"
	"kol/token"
)

var (
	VOID  = &object.Void{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.BooleanLiteral:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right, node.GetPosition())
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right, node.GetPosition())
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ForExpression:
		return evalForExpression(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.BreakStatement:
		val := Eval(node.BreakValue, env)
		if isError(val) {
			return val
		}
		return &object.BreakValue{Value: val}
	case *ast.LetStatement:
		if env.HasValue(node.Name.Value) {
			return newError("Variable %s can't be redefined", node.GetPosition(), node.Name.Value)
		}
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.SetValue(node.Name.Value, object.Variable{Value: val, Mutable: node.Mutable})
	case *ast.ReassignStatement:
		prevVar, existing := env.Get(node.Name.Value)
		if !existing {
			return newError("Variable %s isn't defined", node.GetPosition(), node.Name.Value)
		}
		if !prevVar.Mutable {
			return newError("Variable %s isn't mutable", node.GetPosition(), node.Name.Value)
		}
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		if val.Type() != prevVar.Value.Type() {
			return newError("Can't change type of variable from %s to %s", node.GetPosition(), prevVar.Value.Type(), val.Type())
		}
		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		params := node.Parameters
		returnType := node.ReturnType
		body := node.Body
		return &object.Function{Parameters: params, ReturnType: returnType, Env: env, Body: body}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args, node.GetPosition())
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	}

	return nil
}
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()

			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ || rt == object.BREAK_VALUE_OBJECT {
				return result
			}
		}
	}
	return result
}
func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val.Value
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	return newError("identifier not found: "+node.Value, node.GetPosition())
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func newError(format string, pos token.Position, a ...interface{}) *object.Error {
	msg := fmt.Sprintf(format, a...)
	return &object.Error{Message: msg, Position: &pos}
}
func newUnpositionedError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...), Position: nil}
}
