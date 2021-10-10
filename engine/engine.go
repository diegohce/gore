package engine

import (
	"fmt"
	"gorengine/common"
	"gorengine/rules"
)

type Event string
type stopSignal struct{}

type Engine struct {
	Globals common.WithVariables
	Rules   []rules.Rule
	Events  chan Event
	stop    chan stopSignal
}

func NewEngine() *Engine {
	e := &Engine{
		Globals: common.NewWithVariables(),
		stop:    make(chan stopSignal),
	}
	return e
}

func (e *Engine) Start() {

	for {
		select {
		case event := <-e.Events:
			es := []Event{event}
			e.Run(es)

		case <-e.stop:
			close(e.Events)
			return
		}
	}
}

func (e *Engine) Stop() {
	e.stop <- stopSignal{}
}

func (e *Engine) Run(events []Event) {
	var found bool
	var conditionsPass bool

	for _, r := range e.Rules {
		//check triggers
		found = false
		for _, e := range events {
			if found = r.FindTrigger(string(e)); found {
				break
			}
		}
		if !found {
			continue
		}

		r.SetGlobals(&e.Globals)

		//run conditions
		conditionsPass = false
		if len(r.Conditions()) > 0 {
			for _, c := range r.Conditions() {
				if ok, err := c.Check(r); !ok || err != nil {
					if err != nil {
						fmt.Println(err)
					}
					break
				} else {
					conditionsPass = ok
				}
			}
		} else {
			conditionsPass = true
		}

		//if all pass, apply effects
		if !conditionsPass {
			continue
		}
		for _, e := range r.Effects() {
			if err := e.Apply(r); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (e *Engine) LoadRules(rs ...rules.Rule) {
	e.Rules = append(e.Rules, rs...)
	e.Events = make(chan Event)
}

func (e *Engine) Reset() {
	e.Globals.Reset()
	for _, r := range e.Rules {
		r.Reset()
	}
}
