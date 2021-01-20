package main

import (
	"bufio"
	"fmt"
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

	for t.Next() {
		fmt.Printf("%s\n", t.Token())
	}
}
