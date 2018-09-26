package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T)  {
	tests := []Token {
		{
			Type: additionOp,
			Literal:Literal('+'),
		},
		{
			Type: subtractionOp,
			Literal:Literal('-'),
		},
		{
			Type: multiplicationOp,
			Literal:Literal('*'),
		},
		{
			Type: divisionOp,
			Literal:Literal('/'),
		},
		{
			Type: modulusOp,
			Literal:Literal('%'),
		},
		{
			Type: bitwiseAndOp,
			Literal:Literal('&'),
		},
		{
			Type: bitwiseOrOp,
			Literal:Literal('|'),
		},
		{
			Type: bitwiseXorOp,
			Literal:Literal('^'),
		},
		{
			Type: lessThanOp,
			Literal:Literal('<'),
		},
		{
			Type: greaterThanOp,
			Literal:Literal('>'),
		},
		{
			Type: assignOp,
			Literal:Literal('='),
		},
		{
			Type: notOp,
			Literal:Literal('!'),
		},
		{
			Type: leftParen,
			Literal:Literal('('),
		},
		{
			Type: rightParen,
			Literal:Literal(')'),
		},
		{
			Type: leftBracket,
			Literal:Literal('['),
		},
		{
			Type: rightBracket,
			Literal:Literal(']'),
		},
		{
			Type: leftCurly,
			Literal:Literal('{'),
		},
		{
			Type: rightCurly,
			Literal:Literal('}'),
		},
		{
			Type: comma,
			Literal:Literal(','),
		},
		{
			Type: dot,
			Literal:Literal('.'),
		},
		{
			Type: semicolon,
			Literal:Literal(';'),
		},
		{
			Type: illegal,
			Literal:Literal('?'),
		},
		{
			Type: eof,
			Literal:Literal(0),
		},
	}

	source := `
+ - * / % & | ^ < > = ! () [] {} , . ; ?
`
	lexer := New(source)
	for _, test := range tests {
		assert.Equal(t, test, lexer.NextToken())
	}
}
