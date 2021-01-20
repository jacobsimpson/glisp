package tokenizer

import (
	"fmt"
	"io"
)

type Token struct {
	Raw string
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
			return &Token{Raw: string(r)}
		} else if r == ')' {
			return &Token{Raw: string(r)}
		} else if r == '\'' {
			return &Token{Raw: string(r)}
		} else if isDigit(r) {
			return readInteger(scanner, r)
		} else if isSymbolRune(r) {
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
			return &Token{Raw: result}
		}
	}

	return nil
}

func isSymbolRune(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z'
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
			return &Token{Raw: result}
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
			return &Token{Raw: result}
		}
		result += string(r)
	}
	return nil
}

func (t *Tokenizer) Token() *Token {
	return t.token
}
