package scanner

import (
	"github.com/hashicorp/hil/ast"
)

type Token struct {
	Type    TokenType
	Content string
	Pos     ast.Pos
}

//go:generate stringer -type=TokenType
type TokenType rune

const (
	// Raw string data outside of ${ .. } sequences
	LITERAL TokenType = 'o'

	// STRING is like a LITERAL but it's inside a quoted string
	// within a ${ ... } sequence, and so it can contain backslash
	// escaping.
	STRING TokenType = 'S'

	// Other Literals
	INTEGER TokenType = 'I'
	FLOAT   TokenType = 'F'
	BOOL    TokenType = 'B'

	BEGIN    TokenType = '$' // actually "${"
	END      TokenType = '}'
	OQUOTE   TokenType = '“' // Opening quote of a nested quoted sequence
	CQUOTE   TokenType = '”' // Closing quote of a nested quoted sequence
	OPAREN   TokenType = '('
	CPAREN   TokenType = ')'
	OBRACKET TokenType = '['
	CBRACKET TokenType = ']'
	COMMA    TokenType = ','

	IDENTIFIER TokenType = 'i'

	PERIOD  TokenType = '.'
	PLUS    TokenType = '+'
	MINUS   TokenType = '-'
	STAR    TokenType = '*'
	SLASH   TokenType = '/'
	PERCENT TokenType = '%'

	EOF TokenType = '␄'

	// Produced for sequences that cannot be understood as valid tokens
	// e.g. due to use of unrecognized punctuation.
	INVALID TokenType = '�'
)
