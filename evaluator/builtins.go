package evaluator

import (
	"kol/object"
)

var builtins = map[string]*object.Builtin{
	"println": object.GetBuiltinByName("println"),
	"len":     object.GetBuiltinByName("len"),
	"str":     object.GetBuiltinByName("str"),
	"int":     object.GetBuiltinByName("int"),
	"push":    object.GetBuiltinByName("push"),
	"remove":  object.GetBuiltinByName("remove"),
}
