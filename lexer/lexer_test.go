package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []Token{
		{
			Type:    additionOp,
			Literal: Literal('+'),
		},
		{
			Type:    subtractionOp,
			Literal: Literal('-'),
		},
		{
			Type:    multiplicationOp,
			Literal: Literal('*'),
		},
		{
			Type:    divisionOp,
			Literal: Literal('/'),
		},
		{
			Type:    modulusOp,
			Literal: Literal('%'),
		},
		{
			Type:    bitwiseAndOp,
			Literal: Literal('&'),
		},
		{
			Type:    bitwiseOrOp,
			Literal: Literal('|'),
		},
		{
			Type:    bitwiseXorOp,
			Literal: Literal('^'),
		},
		{
			Type:    lessThanOp,
			Literal: Literal('<'),
		},
		{
			Type:    greaterThanOp,
			Literal: Literal('>'),
		},
		{
			Type:    assignOp,
			Literal: Literal('='),
		},
		{
			Type:    notOp,
			Literal: Literal('!'),
		},
		{
			Type:    leftParen,
			Literal: Literal('('),
		},
		{
			Type:    rightParen,
			Literal: Literal(')'),
		},
		{
			Type:    leftBracket,
			Literal: Literal('['),
		},
		{
			Type:    rightBracket,
			Literal: Literal(']'),
		},
		{
			Type:    leftCurly,
			Literal: Literal('{'),
		},
		{
			Type:    rightCurly,
			Literal: Literal('}'),
		},
		{
			Type:    comma,
			Literal: Literal(','),
		},
		{
			Type:    dot,
			Literal: Literal('.'),
		},
		{
			Type:    semicolon,
			Literal: Literal(';'),
		},
		{
			Type:    illegal,
			Literal: Literal('?'),
		},
		{
			Type: breakKeyword,
			Literal:"break",
		},
		{
			Type: defaultKeyword,
			Literal:"default",
		},
		{
			Type: funcKeyword,
			Literal:"func",
		},
		{
			Type: interfaceKeyword,
			Literal:"interface",
		},
		{
			Type: caseKeyword,
			Literal:"case",
		},
		{
			Type: deferKeyword,
			Literal:"defer",
		},
		{
			Type: cleanKeyword,
			Literal:"clean",
		},
		{
			Type: mapKeyword,
			Literal:"map",
		},
		{
			Type: structKeyword,
			Literal:"struct",
		},
		{
			Type: elseKeyword,
			Literal:"else",
		},
		{
			Type: gotoKeyword,
			Literal:"goto",
		},
		{
			Type: packageKeyword,
			Literal:"package",
		},
		{
			Type: switchKeyword,
			Literal:"switch",
		},
		{
			Type: constKeyword,
			Literal:"const",
		},
		{
			Type: fallthroughKeyword,
			Literal:"fallthrough",
		},
		{
			Type: ifKeyword,
			Literal:"if",
		},
		{
			Type: rangeKeyword,
			Literal:"range",
		},
		{
			Type: typeKeyword,
			Literal:"type",
		},
		{
			Type: continueKeyword,
			Literal:"continue",
		},
		{
			Type: forKeyword,
			Literal:"for",
		},
		{
			Type: importKeyword,
			Literal:"import",
		},
		{
			Type: returnKeyword,
			Literal:"return",
		},
		{
			Type: varKeyword,
			Literal:"var",
		},
		{
			Type:    eof,
			Literal: Literal(0),
		},
	}

	source := `
+ - * / % & | ^ < > = ! () [] {} , . ; ?
break default func interface case defer clean map struct else goto package switch const fallthrough if range type
continue for import return var
`
	lexer := New(source)
	for _, test := range tests {
		assert.Equal(t, test, lexer.NextToken())
	}
}
