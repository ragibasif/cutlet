// Package token defines all the token types for Cutlet.
package token

import (
	"fmt"
)

type TokenKind int

type Token struct {
	Kind    TokenKind
	Literal string
}

// token list subject to change
const (
	NULL    TokenKind = iota // "null"
	EOF                      // "eof"
	UNKNOWN                  // "unknown"

	// Operators
	ASSIGNMENT     // "="
	EQUALS         // "=="
	NOTEQUALS      // "!="
	PLUS           // "+"
	PLUSEQUALS     // "+="
	PLUSPLUS       // "++"
	MINUS          // "-"
	MINUSEQUALS    // "-="
	MINUSMINUS     // "--"
	MULTIPLY       // "*"
	MULTIPLYEQUALS // "*="
	DIVIDEEQUALS   // "/="
	MODULO         // "%"
	MODULOEQUALS   // "%="

	// Delimiters
	COMMA     // ","
	COLON     // ":"
	SEMICOLON // ";"

	// Logical
	AND // "&&"
	OR  // "||"
	NOT // "!"

	// Bitwise
	BITSAND              // "&"
	BITSANDEQUALS        // "&="
	BITSOR               // "|"
	BITSOREQUALS         // "|="
	BITSCOMPLEMENT       // "~"
	BITSCOMPLEMENTEQUALS // "~="
	BITSRSHIFT           // ">>"
	BITSRSHIFTEQUALS     // ">>="
	BITSLSHIFT           // "<<"
	BITSLSHIFTEQUALS     // "<<="

	// Grouping
	LPAREN   // "("
	RPAREN   // ")"
	LBRACE   // "{"
	RBRACE   // "}"
	LBRACKET // "["
	RBRACKET // "]"

	TRUE  // "true"
	FALSE // "false"

	NUMBER     // int: "[0-9]*" float:"[0-9]*(.[0-9]*([eE]-?[0-9]+)?)"
	INTEGER    // int: "[0-9]*"
	FLOAT      // float:"[0-9]*(.[0-9]*([eE]-?[0-9]+)?)"
	STRING     // "",''
	IDENTIFIER // "[_a-zA-Z][_a-zA-Z0-9]*"

	// Conditional
	LESS          // "<"
	LESSEQUALS    // "<="
	GREATER       // ">"
	GREATEREQUALS // ">="

	// Symbols
	DOT       // "."
	DOTDOT    // ".."
	DOTDOTDOT // "..."
	QUESTION  // "?"
	COMMENTS  //  "// comment to EOL"

	PERCENT    // "%"
	TILDE      // "~"
	CARET      // "^"
	OCTOTHORPE // "#"

	// Keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	SIZEOF
	IN

	// Internal
	END   // show at the end of the list of tokens
	COUNT // total number of tokens
)

var keywords = map[string]TokenKind{
	"fn":      FN,
	"let":     LET,
	"const":   CONST,
	"class":   CLASS,
	"new":     NEW,
	"import":  IMPORT,
	"from":    FROM,
	"if":      IF,
	"else":    ELSE,
	"foreach": FOREACH,
	"while":   WHILE,
	"for":     FOR,
	"export":  EXPORT,
	"typeof":  TYPEOF,
	"sizeof":  SIZEOF,
	"in":      IN,
	"null":    NULL,
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NULL:
		return "null"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IDENTIFIER:
		return "identifier"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOTEQUALS:
		return "notequals"
	case NOT:
		return "not"
	case LESS:
		return "less"
	case GREATER:
		return "greater"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOTDOT:
		return "dotdot"
	case SEMICOLON:
		return "semicolon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUSPLUS:
		return "plusplus"
	case MINUSMINUS:
		return "minusminus"
	case PLUSEQUALS:
		return "plusequals"
	case MINUSEQUALS:
		return "minusequals"
	case PLUS:
		return "plus"
	case PERCENT:
		return "percent"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:
		return "new"
	case IMPORT:
		return "import"
	case FROM:
		return "from"
	case FN:
		return "fn"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case EXPORT:
		return "export"
	case IN:
		return "in"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}

func LookupIdententifier(identifier string) TokenKind {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENTIFIER
}
