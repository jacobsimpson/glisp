package parser

import (
	"fmt"
	"glisp/tokenizer"
)

type Expression interface {
	IsAtom() bool
}

type Integer struct {
}

func (i *Integer) IsAtom() bool { return true }

type String struct{}

func (s *String) IsAtom() bool { return true }

func Parse(tokenizer tokenizer.Tokenizer) (Expression, error) {
	for tokenizer.Next() {
		t := tokenizer.Token()
		switch t.Type {
		case tokenizer.OpenParen:
			return list(tokenizer)
		case tokenizer.CloseParen:
			return nil, fmt.Errorf("found %q without an open", ")")
		default:
			return nil, fmt.Errorf("found token outside of list")
		}
	}
}

func list(tokenizer tokenizer.Tokenizer) (Expression, error) {
}
