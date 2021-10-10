package ankorules

import (
	"fmt"
	"gorengine/rules"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"

	_ "github.com/mattn/anko/packages"
)

type externalRule struct {
	rules.BaseRule
	ank *env.Env
}

type externalCondition struct {
	script string
}

func (c externalCondition) Check(r rules.Rule) (bool, error) {

	ank := r.(*externalRule).ank
	ank.Define("rule", r)

	out, err := vm.Execute(ank, nil, c.script)
	if err != nil {
		return false, err
	}

	tuple, ok := out.([]interface{})
	if !ok {
		return false, fmt.Errorf("%s: script must return 2 values (bool, error)", r.Name())
	}

	ret := tuple[0].(bool)
	if !ret && tuple[1] != nil {
		err = tuple[1].(error)
	}
	return ret, err
}

type externalEffect struct {
	script string
}

func (e externalEffect) Apply(r rules.Rule) error {

	ank := r.(*externalRule).ank
	ank.Define("rule", r)

	out, err := vm.Execute(ank, nil, e.script)
	if err != nil {
		return err
	}
	if out != nil {
		scritptErr, ok := out.(error)
		if !ok {
			return fmt.Errorf("%s: script return type must be of type 'error'", r.Name())
		}
		return scritptErr
	}

	return nil
}
