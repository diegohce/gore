package main

import (
	"fmt"
	"plugin"

	"gorengine/engine"
	"gorengine/rules"
	"gorengine/rules/ankorules"
	//_ "gorengine/rules/some"
)

func rulesFromPlugin() {

	p, err := plugin.Open("plugins/bin/arule.so")
	if err != nil {
		fmt.Println(err, "loading plugin")
		return
	}

	sym, err := p.Lookup("RuleFactory")
	if err != nil {
		fmt.Println(err, "looking for symbol 'RuleFactory'")
		return
	}

	//factoryFN, ok := sym.(rules.RuleFactoryFunc)
	factoryFN, ok := sym.(func() rules.Rule)
	if !ok {
		fmt.Println("symbol is not RuleFactoryFunc")
		//return
	} else {
		rules.RegisterRuleFactory(factoryFN)
		fmt.Println("registered rule factory")
		return
	}

	sym, err = p.Lookup("Rule")
	if err != nil {
		fmt.Println(err, "looking for symbol 'Rule'")
		return
	}

	r, okk := sym.(*rules.Rule)
	if !okk {
		fmt.Println("symbol is not Rule")
		return
	} else {
		rules.RulesList = append(rules.RulesList, *r)
		fmt.Println("rule appended")
	}

}

func main() {

	ankorules.LoadFromDefault()
	//rulesFromPlugin()

	e := engine.NewEngine()
	e.LoadRules(rules.RulesList...)

	//Some fictional events
	events := []engine.Event{"init", "start", "do_stuff", "end"}

	/*go e.Start()

	for _, ev := range events {
		e.Events <- engine.Event(ev)
	}
	e.Stop()

	*/

	e.Run(events)

	e.Reset()
}
