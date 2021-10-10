package main

import (
	"gorengine/rules"
)

type cond struct{}

func (c cond) Check(r rules.Rule) (bool, error) {
	return true, nil
}

type eff struct{}

func (e eff) Apply(r rules.Rule) error {
	r.Globals().W("plugin_message", "hello from plugin")
	return nil
}

func RuleFactory() rules.Rule {
	r := rules.BaseRule{}

	r.SetTriggers("plugin")
	r.SetConditions(cond{})
	r.SetEffects(eff{})

	return &r
}

var Rule = RuleFactory()
