package lexer

import (
	"compiler/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
type
  NodePtr = ^Node;
  Node = record
    Data: Integer;
    Next: NodePtr;
  end;

var
  x: Integer;
  listHead: NodePtr;

// This is a comment
procedure Test(a: Integer);
var
  b: Integer;
begin
  b := a + 5;
  if b > 10 then
    x := 1
  else
    x := 2;

  while b <> 0 do
    b := b - 1;
end;

// Operators test
5 * 5;
5 / 5;
5 <= 6;
5 >= 4;
listHead^.Data := 100;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.TYPE, "type", 2, 1},
		{token.IDENT, "NodePtr", 3, 3},
		{token.EQ, "=", 3, 11},
		{token.CARET, "^", 3, 13},
		{token.IDENT, "Node", 3, 14},
		{token.SEMICOLON, ";", 3, 18},
		{token.IDENT, "Node", 4, 3},
		{token.EQ, "=", 4, 8},
		{token.RECORD, "record", 4, 10},
		{token.IDENT, "Data", 5, 5},
		{token.COLON, ":", 5, 9},
		{token.IDENT, "Integer", 5, 11},
		{token.SEMICOLON, ";", 5, 18},
		{token.IDENT, "Next", 6, 5},
		{token.COLON, ":", 6, 9},
		{token.IDENT, "NodePtr", 6, 11},
		{token.SEMICOLON, ";", 6, 18},
		{token.END, "end", 7, 3},
		{token.SEMICOLON, ";", 7, 6},

		{token.VAR, "var", 9, 1},
		{token.IDENT, "x", 10, 3},
		{token.COLON, ":", 10, 4},
		{token.IDENT, "Integer", 10, 6},
		{token.SEMICOLON, ";", 10, 13},
		{token.IDENT, "listHead", 11, 3},
		{token.COLON, ":", 11, 11},
		{token.IDENT, "NodePtr", 11, 13},
		{token.SEMICOLON, ";", 11, 20},

		{token.PROCEDURE, "procedure", 14, 1},
		{token.IDENT, "Test", 14, 11},
		{token.LPAREN, "(", 14, 15},
		{token.IDENT, "a", 14, 16},
		{token.COLON, ":", 14, 17},
		{token.IDENT, "Integer", 14, 19},
		{token.RPAREN, ")", 14, 26},
		{token.SEMICOLON, ";", 14, 27},
		{token.VAR, "var", 15, 1},
		{token.IDENT, "b", 16, 3},
		{token.COLON, ":", 16, 4},
		{token.IDENT, "Integer", 16, 6},
		{token.SEMICOLON, ";", 16, 13},
		{token.BEGIN, "begin", 17, 1},
		{token.IDENT, "b", 18, 3},
		{token.ASSIGN, ":=", 18, 5},
		{token.IDENT, "a", 18, 8},
		{token.PLUS, "+", 18, 10},
		{token.INT, "5", 18, 12},
		{token.SEMICOLON, ";", 18, 13},
		{token.IF, "if", 19, 3},
		{token.IDENT, "b", 19, 6},
		{token.GT, ">", 19, 8},
		{token.INT, "10", 19, 10},
		{token.THEN, "then", 19, 13},
		{token.IDENT, "x", 20, 5},
		{token.ASSIGN, ":=", 20, 7},
		{token.INT, "1", 20, 9},
		{token.ELSE, "else", 21, 3},
		{token.IDENT, "x", 22, 5},
		{token.ASSIGN, ":=", 22, 7},
		{token.INT, "2", 22, 9},
		{token.SEMICOLON, ";", 22, 10},
		{token.WHILE, "while", 24, 3},
		{token.IDENT, "b", 24, 9},
		{token.NOT_EQ, "<>", 24, 11},
		{token.INT, "0", 24, 14},
		{token.DO, "do", 24, 16},
		{token.IDENT, "b", 25, 5},
		{token.ASSIGN, ":=", 25, 7},
		{token.IDENT, "b", 25, 10},
		{token.MINUS, "-", 25, 12},
		{token.INT, "1", 25, 14},
		{token.SEMICOLON, ";", 25, 15},
		{token.END, "end", 26, 1},
		{token.SEMICOLON, ";", 26, 4},

		{token.INT, "5", 29, 1},
		{token.ASTERISK, "*", 29, 3},
		{token.INT, "5", 29, 5},
		{token.SEMICOLON, ";", 29, 6},
		{token.INT, "5", 30, 1},
		{token.SLASH, "/", 30, 3},
		{token.INT, "5", 30, 5},
		{token.SEMICOLON, ";", 30, 6},
		{token.INT, "5", 31, 1},
		{token.LTE, "<=", 31, 3},
		{token.INT, "6", 31, 6},
		{token.SEMICOLON, ";", 31, 7},
		{token.INT, "5", 32, 1},
		{token.GTE, ">=", 32, 3},
		{token.INT, "4", 32, 6},
		{token.SEMICOLON, ";", 32, 7},
		{token.IDENT, "listHead", 33, 1},
		{token.CARET, "^", 33, 9},
		{token.DOT, ".", 33, 10},
		{token.IDENT, "Data", 33, 11},
		{token.ASSIGN, ":=", 33, 16},
		{token.INT, "100", 33, 19},
		{token.SEMICOLON, ";", 33, 22},

		{token.EOF, "", 34, 1},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}

		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - column wrong. expected=%d, got=%d",
				i, tt.expectedColumn, tok.Column)
		}
	}
}
