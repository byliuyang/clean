package lexer

import (
	"regexp"
)

type TokenType int

type Literal string

const (
	identifier TokenType = iota

	breakKeyword
	defaultKeyword
	funcKeyword
	interfaceKeyword
	caseKeyword
	deferKeyword
	cleanKeyword
	mapKeyword
	structKeyword
	elseKeyword
	gotoKeyword
	packageKeyword
	switchKeyword
	constKeyword
	fallthroughKeyword
	ifKeyword
	rangeKeyword
	typeKeyword
	continueKeyword
	forKeyword
	importKeyword
	returnKeyword
	varKeyword

	additionOp
	subtractionOp
	multiplicationOp
	divisionOp
	modulusOp
	bitwiseAndOp
	bitwiseOrOp
	bitwiseXorOp
	leftShiftOp
	rightShiftOp
	plusEqualOp
	multiplyEqualOp
	divideEqualOp
	modularEqualOp
	andEqualOp
	orEqualOp
	xorEqualOp
	leftShiftEqualOp
	rightShiftEqualOp
	andOp
	orOp
	incrementOp
	decrementOp
	equalOp
	lessThanOp
	greaterThanOp
	assignOp
	notOp
	notEqualOp
	lessEqualOp
	greaterEqualOp

	leftParen
	rightParen
	leftBracket
	rightBracket
	leftCurly
	rightCurly
	semicolon
	comma
	dot

	integerLiter
	floatingPointLiter
	stringLiter

	eof
	illegal
)

type Token struct {
	Type    TokenType
	Literal Literal
}

type Lexer struct {
	input   string
	currPos int
	readPos int
	ch      byte
}

func New(input string) *Lexer {
	l := Lexer{
		input: input,
	}
	l.readChar()
	return &l
}

func NewToken(t TokenType, lit byte) Token {
	return Token{
		Type:    t,
		Literal: Literal(lit),
	}
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.currPos = l.readPos
	l.readPos++
}

func (l *Lexer) isWhiteSpace() bool {
	return l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'
}

func (l *Lexer) eatWhiteSpace() {
	for l.isWhiteSpace() {
		l.readChar()
	}
}

func (l *Lexer) isDelimiter() bool {
	return l.isWhiteSpace() ||
		l.ch == '+' ||
		l.ch == '-' ||
		l.ch == '*' ||
		l.ch == '/' ||
		l.ch == '%' ||
		l.ch == '&' ||
		l.ch == '|' ||
		l.ch == '^' ||
		l.ch == '<' ||
		l.ch == '>' ||
		l.ch == '=' ||
		l.ch == '!' ||
		l.ch == '(' ||
		l.ch == '[' ||
		l.ch == '{' ||
		l.ch == ',' ||
		l.ch == ')' ||
		l.ch == ']' ||
		l.ch == '}' ||
		l.ch == ';' ||
		l.ch == 0
}

var keywords = map[string]TokenType{
	"break":       breakKeyword,
	"default":     defaultKeyword,
	"func":        funcKeyword,
	"interface":   interfaceKeyword,
	"case":        caseKeyword,
	"defer":       deferKeyword,
	"clean":       cleanKeyword,
	"map":         mapKeyword,
	"struct":      structKeyword,
	"else":        elseKeyword,
	"goto":        gotoKeyword,
	"package":     packageKeyword,
	"switch":      switchKeyword,
	"const":       constKeyword,
	"fallthrough": fallthroughKeyword,
	"if":          ifKeyword,
	"range":       rangeKeyword,
	"type":        typeKeyword,
	"continue":    continueKeyword,
	"for":         forKeyword,
	"import":      importKeyword,
	"return":      returnKeyword,
	"var":         varKeyword,
}

func (l *Lexer) readWord() string {
	start := l.currPos
	for !l.isDelimiter() {
		l.readChar()
	}
	return l.input[start:l.currPos]
}

func isKeyword(word string) bool {
	_, ok := keywords[word]
	return ok
}

func isIdentifier(word string) bool {
	reg := regexp.MustCompile("^(([_a-zA-Z]|\\p{L})([_a-zA-Z0-9]|\\p{L})*)$")
	return reg.MatchString(word)
}

func isInteger(word string) bool {
	reg := regexp.MustCompile("^(0|[0-9]+|(0x[0-9a-zA-Z]+))$")
	return reg.MatchString(word)
}

func isFloatingPoint(word string) bool {
	reg := regexp.MustCompile("^([0-9]*(\\.[0-9]*)?((e|E)[0-9]+)?)$")
	return reg.MatchString(word)
}

func keywordType(word string) TokenType {
	tokenType, _ := keywords[word]
	return tokenType
}

func (l *Lexer) NextToken() Token {
	var token Token
	l.eatWhiteSpace()
	switch l.ch {
	case '+':
		token = NewToken(additionOp, l.ch)
	case '-':
		token = NewToken(subtractionOp, l.ch)
	case '*':
		token = NewToken(multiplicationOp, l.ch)
	case '/':
		token = NewToken(divisionOp, l.ch)
	case '%':
		token = NewToken(modulusOp, l.ch)
	case '&':
		token = NewToken(bitwiseAndOp, l.ch)
	case '|':
		token = NewToken(bitwiseOrOp, l.ch)
	case '^':
		token = NewToken(bitwiseXorOp, l.ch)
	case '<':
		token = NewToken(lessThanOp, l.ch)
	case '>':
		token = NewToken(greaterThanOp, l.ch)
	case '=':
		token = NewToken(assignOp, l.ch)
	case '!':
		token = NewToken(notOp, l.ch)
	case '(':
		token = NewToken(leftParen, l.ch)
	case '[':
		token = NewToken(leftBracket, l.ch)
	case '{':
		token = NewToken(leftCurly, l.ch)
	case ',':
		token = NewToken(comma, l.ch)
	case '.':
		token = NewToken(dot, l.ch)
	case ')':
		token = NewToken(rightParen, l.ch)
	case ']':
		token = NewToken(rightBracket, l.ch)
	case '}':
		token = NewToken(rightCurly, l.ch)
	case ';':
		token = NewToken(semicolon, l.ch)
	case 0:
		token = NewToken(eof, l.ch)
	default:
		word := l.readWord()
		if isKeyword(word) {
			token.Type = keywordType(word)
		} else if isIdentifier(word) {
			token.Type = identifier
		} else if isInteger(word) {
			token.Type = integerLiter
		} else if isFloatingPoint(word) {
			token.Type = floatingPointLiter
		} else {
			token.Type = illegal
		}

		token.Literal = Literal(word)
		return token
	}
	l.readChar()
	return token
}
