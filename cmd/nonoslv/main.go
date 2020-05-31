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
	fmt.Printf("%v", *input)
}
