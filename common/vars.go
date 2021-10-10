package common

import (
	"errors"
)

var (
	ErrUndefinedVariable = errors.New("undefined variable")
)

type WithVariables struct {
	vars map[string]interface{}
}

func NewWithVariables() WithVariables {
	return WithVariables{
		vars: make(map[string]interface{}),
	}
}

func (h *WithVariables) W(name string, value interface{}) {
	h.vars[name] = value
}

func (h *WithVariables) R(name string) (interface{}, error) {

	v, ok := h.vars[name]
	if !ok {
		return nil, ErrUndefinedVariable
	}
	return v, nil
}

func (h *WithVariables) D(name string) {
	_, ok := h.vars[name]
	if ok {
		delete(h.vars, name)
	}
}

func (h *WithVariables) Reset() {
	h.vars = make(map[string]interface{})
}

type VariableHandler interface {
	W(name string, value interface{})
	R(name string) (interface{}, error)
	D(name string)
	Reset()
}
