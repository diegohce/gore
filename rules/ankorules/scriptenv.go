package ankorules

import (
	"errors"
	"fmt"

	"github.com/mattn/anko/env"
)

func setupEnv() *env.Env {
	e := env.NewEnv()

	e.Define("print", fmt.Println)
	e.Define("format", fmt.Sprintf)
	e.Define("error", errors.New)
	e.Define("foreach", forEach)

	return e
}

func forEach(arr []interface{}, fn func(interface{}) error) error {
	for _, i := range arr {
		if err := fn(i); err != nil {
			return err
		}
	}
	return nil
}
