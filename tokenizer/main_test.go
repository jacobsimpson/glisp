package tokenizer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		data string
		want []string
	}{
		{
			name: "empty input",
			data: "",
			want: []string{},
		},
		{
			name: "parens",
			data: ")(",
			want: []string{")", "("},
		},
		{
			name: "ignoring whitespace",
			data: `
		)  (	`,
			want: []string{")", "("},
		},
		{
			name: "with a number",
			data: ")123(",
			want: []string{")", "123", "("},
		},
		{
			name: "with a symbol",
			data: "(999 this)",
			want: []string{"(", "999", "this", ")"},
		},
		{
			name: "with a string",
			data: `(999 this     "or that")`,
			want: []string{"(", "999", "this", `"or that"`, ")"},
		},
		{
			name: "with a quoted list",
			data: `(sum '(1 2 3))`,
			want: []string{"(", "sum", "'", "(", "1", "2", "3", ")", ")"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokenizer := New(strings.NewReader(test.data))
			got := []string{}
			for tokenizer.Next() {
				got = append(got, tokenizer.Token().Raw)
			}
			assert.Equal(test.want, got)
		})
	}
}
