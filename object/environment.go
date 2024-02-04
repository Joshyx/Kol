package object

func NewEnvironment() *Environment {
	s := make(map[string]Variable)
	return &Environment{store: s, outer: nil}
}
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

type Variable struct {
	Value   Object
	Mutable bool
}
type Environment struct {
	store map[string]Variable
	outer *Environment
}

func (e *Environment) Get(name string) (Variable, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) Set(name string, obj Object) Object {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Set(name, obj)
	}
	e.store[name] = Variable{Value: obj, Mutable: val.Mutable}
	return obj
}
func (e *Environment) SetValue(name string, val Variable) Object {
	e.store[name] = val
	return val.Value
}
func (e *Environment) HasValue(name string) bool {
	_, ok := e.store[name]
	return ok
}
