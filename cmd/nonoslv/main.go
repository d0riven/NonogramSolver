package main

import (
	"github.com/d0riven/NonogramSolver/pkg/nonoslv"
)

func main() {
	input, err := nonoslv.ReadFromYaml("sample.yaml")
	if err != nil {
	    panic(err)
	}
	history, err := nonoslv.Solve(input)
	if err != nil {
		panic(err)
	}
	history.Print()
}
