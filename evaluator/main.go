package evaluator

import (
	"fmt"
	"glisp/parser"
)

type Environment struct {
	parent    *Environment
	variables map[string]Binding
}

type Binding interface{}

type BuiltInBinding struct{}

type SpecialFormBinding struct{}

type ExpressionBinding struct {
	Expression parser.Expression
}

func NewEnvironment() *Environment {
	return &Environment{
		variables: map[string]Binding{},
	}
}

// LookupVariable checks this environment to find the binding for a variable.
// If the variable is unbound, it will return a nil result.
func (e *Environment) LookupBinding(name string) Binding {
	if v, ok := e.variables[name]; ok {
		return v
	}
	if e.parent == nil {
		return nil
	}
	return e.parent.LookupBinding(name)
}

func Eval(program []parser.Expression, environment *Environment) {
	for _, expression := range program {
		eval(expression, environment)
	}
}

func eval(expression parser.Expression, environment *Environment) {
	fmt.Printf("Doing %+v\n", expression)
	switch e := expression.(type) {
	case *parser.Symbol:
		evalSymbol(e)
	}
}

func evalSymbol(s *parser.Symbol, environment *Environment) {
	_ = environment.LookupBinding(e.Name)
}

func evalSExpression(e *parser.SExpression, environment *Environment) {
	e.Left()
}
