package rules

import (
	"gorengine/common"
)

type RuleFactoryFunc func() Rule

type Condition interface {
	Check(r Rule) (bool, error)
}

type Effect interface {
	Apply(r Rule) error
}

type BaseRule struct {
	common.WithNameField
	common.WithVariables
	globals    *common.WithVariables
	triggers   []string
	conditions []Condition
	effects    []Effect
}

func (r *BaseRule) SetEffects(e ...Effect) {
	r.effects = append(r.effects, e...)
}

func (r *BaseRule) Effects() []Effect {
	return r.effects
}

func (r *BaseRule) SetConditions(c ...Condition) {
	r.conditions = append(r.conditions, c...)
}

func (r *BaseRule) Conditions() []Condition {
	return r.conditions
}

func (r *BaseRule) SetTriggers(t ...string) {
	r.triggers = append(r.triggers, t...)
}

func (r *BaseRule) Triggers() []string {
	return r.triggers
}

func (r *BaseRule) SetGlobals(g *common.WithVariables) {
	r.globals = g
}

func (r *BaseRule) Globals() *common.WithVariables {
	return r.globals
}

func (r *BaseRule) FindTrigger(event string) bool {
	for i := range r.triggers {
		if r.triggers[i] == event {
			return true
		}
	}
	return false
}

var RulesList []Rule

func RegisterRuleFactory(f RuleFactoryFunc) {
	RulesList = append(RulesList, f())
}

type Rule interface {
	common.Namer
	common.VariableHandler
	SetEffects(e ...Effect)
	Effects() []Effect
	SetConditions(c ...Condition)
	Conditions() []Condition
	SetTriggers(t ...string)
	Triggers() []string
	SetGlobals(g *common.WithVariables)
	Globals() *common.WithVariables
	FindTrigger(event string) bool
}
