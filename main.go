package main

import (
	"bufio"
	"fmt"
	"glisp/evaluator"
	"glisp/parser"
	"glisp/tokenizer"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "glisp <file>\n")
		os.Exit(1)
	}
	src := os.Args[1]
	f, err := os.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read %q: %+v\n", src, err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	t := tokenizer.New(r)

	program, err := parser.Parse(t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse: %+v\n", err)
		os.Exit(1)
	}

	global := evaluator.NewEnvironment()
	evaluator.Eval(program, global)
}
