package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/d0riven/NonogramSolver/pkg/nonoslv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Printf("%v\n", args)
	if len(args) != 1 {
		showUsage()
		return
	}
	if _, err := os.Stat(args[0]); err != nil {
		showUsage()
		return
	}
	input, err := nonoslv.ReadFromYaml(args[0])
	if err != nil {
	    panic(err)
	}
	history, err := nonoslv.Solve(input)
	if err != nil {
		panic(err)
	}
	history.Print()
}

func showUsage() {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("failed to get caller")
	}
	fmt.Printf("go run %s inputs/5x5/snake.yaml\n", filepath.Base(file))
}
