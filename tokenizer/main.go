package tokenizer

import (
	"fmt"
	"io"
)

type TokenType int

const (
	Symbol TokenType = iota
	String
	Integer
	OpenParen
	CloseParen
)

type Token struct {
	Type  TokenType
	Value string
	Raw   string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token {Raw: %q}", t.Raw)
}

type Tokenizer struct {
	scanner io.RuneScanner
	token   *Token
}

func New(s io.RuneScanner) *Tokenizer {
	return &Tokenizer{scanner: s}
}

func (t *Tokenizer) Next() bool {
	t.token = readToken(t.scanner)
	return t.token != nil
}

func readToken(scanner io.RuneScanner) *Token {
	for {
		r, _, err := scanner.ReadRune()
		if err != nil {
			return nil
		}

		if r == '(' {
			return &Token{Type: OpenParen, Raw: string(r), Value: string(r)}
		} else if r == ')' {
			return &Token{Type: CloseParen, Raw: string(r), Value: string(r)}
		} else if r == '\'' {
			return &Token{Type: String, Raw: string(r), Value: string(r)}
		} else if isDigit(r) {
			return readInteger(scanner, r)
		} else if isSymbolFirstRune(r) {
			return readSymbol(scanner, r)
		} else if isStringDelimiter(r) {
			return readString(scanner, r)
		}
	}

	return nil
}

func isDigit(r rune) bool {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func readInteger(scanner io.RuneScanner, first rune) *Token {
	result := string(first)
	for {
		r, _, err := scanner.ReadRune()
		if err != nil {
			return nil
		}

		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result += string(r)
		default:
			scanner.UnreadRune()
			return &Token{Type: Integer, Raw: result, Value: result}
		}
	}

	return nil
}

func isSymbolFirstRune(r rune) bool {
	return !isWhitespace(r) && !isDigit(r) && !isStringDelimiter(r)
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func readSymbol(scanner io.RuneScanner, first rune) *Token {
	result := string(first)
	for {
		r, _, err := scanner.ReadRune()
		if err != nil {
			return nil
		}

		if isWhitespace(r) || r == '(' || r == ')' {
			scanner.UnreadRune()
			return &Token{Type: Symbol, Raw: result, Value: result}
		}
		result += string(r)
	}

	return nil
}

func isStringDelimiter(r rune) bool { return r == '"' }

func readString(scanner io.RuneScanner, first rune) *Token {
	result := string(first)
	for {
		r, _, err := scanner.ReadRune()
		if err != nil {
			return nil
		}
		if r == first {
			result += string(r)
			return &Token{Type: String, Raw: result, Value: result[1 : len(result)-1]}
		}
		result += string(r)
	}
	return nil
}

func (t *Tokenizer) Token() *Token {
	return t.token
}
