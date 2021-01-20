package parser

import (
	"glisp/tokenizer"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []Expression
	}{
		{
			input: `(display "hello")`,
			want:  []Expression{&SExpression{left: &Symbol{"display"}, right: &SExpression{left: &String{"hello"}}}},
		},
		{
			input: `(car     1
			2                 3)`,
			want: []Expression{&SExpression{left: &Symbol{"car"}, right: &SExpression{left: &Integer{1}, right: &SExpression{left: &Integer{2}, right: &SExpression{left: &Integer{3}}}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := Parse(tokenizer.New(strings.NewReader(test.input)))
			assert.Nil(err)
			assert.Equal(test.want, got)
		})
	}
}
