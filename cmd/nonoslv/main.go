package main

import (
	"fmt"

	"github.com/d0riven/NonogramSolver/pkg/nonoslv"
)

func main() {
	input, err := nonoslv.ReadFromYaml("sample.yaml")
	if err != nil {
	    panic(err)
	}
	fmt.Printf("input: %v\n", *input)
	history, err := nonoslv.Solve(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("history: %v\n", *history)
}
