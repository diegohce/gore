package ankorules

import (
	"fmt"
	"gorengine/rules"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

const (
	defaultRulesDir = "./rules"
)

type ruleFile struct {
	Name       string
	Triggers   []string
	Conditions []struct {
		Code string
	}
	Effects []struct {
		Code string
	}
}

func LoadFromDefault() {
	Load(defaultRulesDir)
}

func Load(dirname string) {

	ent, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}

	ankoEnv := setupEnv()

	for _, f := range ent {

		r := ruleFile{}

		filename := path.Join(dirname, f.Name())

		if _, err := toml.DecodeFile(filename, &r); err != nil {
			fmt.Println(err)
		}

		rule := &externalRule{}
		rule.ank = ankoEnv

		rule.SetName(r.Name)
		rule.SetTriggers(r.Triggers...)

		for _, c := range r.Conditions {
			rule.SetConditions(externalCondition{script: c.Code})
		}
		for _, e := range r.Effects {
			rule.SetEffects(externalEffect{script: e.Code})
		}

		rules.RulesList = append(rules.RulesList, rule)
	}
}
