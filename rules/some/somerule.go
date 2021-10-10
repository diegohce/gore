package somerule

import (
	"fmt"
	"gorengine/rules"
)

func SomeRule() rules.Rule {

	r := rules.BaseRule{}

	r.SetName("SomeRule")
	r.SetTriggers("hello_world")
	r.SetConditions(someCondition{})
	r.SetEffects(someEffect{})

	return &r
}

type someCondition struct{}

func (someCondition) Check(r rules.Rule) (bool, error) {
	userName, err := r.Globals().R("name")
	if err != nil {
		return false, err
	}
	if userName == "" {
		return false, nil
	}
	return true, nil
}

type someEffect struct{}

func (someEffect) Apply(r rules.Rule) error {
	userName, err := r.Globals().R("name")
	if err != nil {
		return err
	}

	r.Globals().W("message",
		fmt.Sprintf("Hello, %s!", userName))

	return nil
}

func init() {
	rules.RegisterRuleFactory(SomeRule)
}
