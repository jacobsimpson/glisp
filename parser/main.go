package parser

import (
	"fmt"
	"glisp/tokenizer"
	"strconv"
)

type Expression interface {
	IsAtom() bool
	Left() Expression
	Right() Expression
	String() string
}

type Integer struct {
	value int
}

func (i *Integer) IsAtom() bool      { return true }
func (i *Integer) Left() Expression  { return nil }
func (i *Integer) Right() Expression { return nil }
func (i *Integer) String() string    { return fmt.Sprintf("%d", i.value) }

type String struct {
	Value string
}

func (s *String) IsAtom() bool      { return true }
func (s *String) Left() Expression  { return nil }
func (s *String) Right() Expression { return nil }
func (s *String) String() string    { return fmt.Sprintf("%q", s.Value) }

type Symbol struct {
	Name string
}

func (s *Symbol) IsAtom() bool      { return true }
func (s *Symbol) Left() Expression  { return nil }
func (s *Symbol) Right() Expression { return nil }
func (s *Symbol) String() string    { return s.Name }

type SExpression struct {
	left  Expression
	right Expression
}

func (s *SExpression) IsAtom() bool      { return false }
func (s *SExpression) Left() Expression  { return s.left }
func (s *SExpression) Right() Expression { return s.right }
func (s *SExpression) String() string {
	return fmt.Sprintf("(%v . %v)", s.left, s.right)
}

func Parse(tokens *tokenizer.Tokenizer) ([]Expression, error) {
	results := []Expression{}

	for tokens.Next() {
		token := tokens.Token()
		switch token.Type {
		case tokenizer.OpenParen:
			l, err := list(tokens)
			if err != nil {
				return nil, err
			}
			results = append(results, l)
		case tokenizer.CloseParen:
			return nil, fmt.Errorf("found %q without an open", ")")
		default:
			return nil, fmt.Errorf("found token outside of list")
		}
	}

	return results, nil
}

func list(tokens *tokenizer.Tokenizer) (Expression, error) {
	var head, last *SExpression
	for tokens.Next() {
		token := tokens.Token()
		switch token.Type {
		case tokenizer.OpenParen:
			e, err := list(tokens)
			if err != nil {
				return nil, err
			}
			head, last = appendExpression(head, last, e)
		case tokenizer.CloseParen:
			return head, nil
		case tokenizer.Symbol:
			head, last = appendExpression(head, last, &Symbol{token.Value})
		case tokenizer.Integer:
			i, _ := strconv.Atoi(token.Value)
			head, last = appendExpression(head, last, &Integer{i})
		case tokenizer.String:
			head, last = appendExpression(head, last, &String{token.Value})
		default:
			return nil, fmt.Errorf("found token outside of list")
		}
	}
	return nil, fmt.Errorf("extra stuff")
}

func appendExpression(head, last *SExpression, cell Expression) (*SExpression, *SExpression) {
	if head == nil {
		head = &SExpression{left: cell}
		last = head
		return head, last
	}
	s := &SExpression{left: cell}
	last.right = s
	return head, s
}
