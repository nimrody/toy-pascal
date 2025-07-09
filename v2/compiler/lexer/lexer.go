package lexer

import (
	"nimrody.com/toypascal/v2/compiler/token"
)

// Lexer holds the state of the scanner.
type Lexer struct {
	input        string // The source code being scanned
	position     int    // Current position in input (points to current char)
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // Current char under examination
	line         int    // Current line number
	column       int    // Current column number for the start of the token
}

// New creates a new Lexer and initializes it.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readChar() // Initialize l.ch, l.position, l.readPosition
	return l
}

// readChar gives us the next character and advances our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL", signifies EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// peekChar looks at the next character without consuming it.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken scans the input and returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Store start column before consuming characters for the token
	startColumn := l.column

	switch l.ch {
	case '=':
		tok = l.newToken(token.EQ, l.ch, startColumn)
	case ';':
		tok = l.newToken(token.SEMICOLON, l.ch, startColumn)
	case ':':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.ASSIGN, Literal: literal, Line: l.line, Column: startColumn}
		} else {
			tok = l.newToken(token.COLON, l.ch, startColumn)
		}
	case '(':
		tok = l.newToken(token.LPAREN, l.ch, startColumn)
	case ')':
		tok = l.newToken(token.RPAREN, l.ch, startColumn)
	case ',':
		tok = l.newToken(token.COMMA, l.ch, startColumn)
	case '+':
		tok = l.newToken(token.PLUS, l.ch, startColumn)
	case '-':
		tok = l.newToken(token.MINUS, l.ch, startColumn)
	case '*':
		tok = l.newToken(token.ASTERISK, l.ch, startColumn)
	case '/':
		// Handle comments
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken() // Recursively call NextToken to get the token after the comment
		}
		tok = l.newToken(token.SLASH, l.ch, startColumn)
	case '<':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, Line: l.line, Column: startColumn}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LTE, Literal: literal, Line: l.line, Column: startColumn}
		} else {
			tok = l.newToken(token.LT, l.ch, startColumn)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GTE, Literal: literal, Line: l.line, Column: startColumn}
		} else {
			tok = l.newToken(token.GT, l.ch, startColumn)
		}
	case '^':
		tok = l.newToken(token.CARET, l.ch, startColumn)
	case '.':
		tok = l.newToken(token.DOT, l.ch, startColumn)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
		tok.Column = startColumn
	default:
		if isLetter(l.ch) {
			tok.Line = l.line
			tok.Column = startColumn
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // Early return since readIdentifier advances position
		} else if isDigit(l.ch) {
			tok.Line = l.line
			tok.Column = startColumn
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok // Early return since readNumber advances position
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch, startColumn)
		}
	}

	l.readChar()
	return tok
}

// newToken is a helper to create a new token.
func (l *Lexer) newToken(tokenType token.TokenType, ch byte, column int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: l.line, Column: column}
}

// readIdentifier reads in an identifier and advances the lexer's position.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads in a number and advances the lexer's position.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace consumes all subsequent whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment consumes a single-line comment.
func (l *Lexer) skipComment() {
	// Assumes the first '/' has been seen
	l.readChar() // consume the second '/'
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace() // Skip any whitespace after the comment
}

// isLetter checks if a byte is a letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if a byte is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
