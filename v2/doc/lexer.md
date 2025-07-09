# Go Lexer Package Specification for Toy Pascal

This document outlines the design and API for the lexer and token packages in
Go, which together form the lexical analysis stage of the Toy Pascal compiler.

## 1. Package Structure

The lexer functionality will be split into two packages to maintain a clean
separation of concerns:

* token/: Defines the data structures that represent a token. This package will
  have no dependencies on other parts of the compiler.  

* lexer/: Contains the logic for reading source code and producing tokens. It
  will depend on the token package.

compiler/  
├── lexer/  
│   ├── lexer.go  
│   └── lexer\_test.go  
└── token/  
    └── token.go

## 2. token Package Specification

### token.go

This file defines the types of tokens that can appear in the Toy Pascal
language.

#### TokenType

A string-based type alias to represent the category of a token.

package token

type TokenType string

#### Token Constants

Constants for every possible TokenType.

const (  
    // Special tokens  
    ILLEGAL = "ILLEGAL" // A token/character we don't know about  
    EOF     = "EOF"     // "End of File"

    // Identifiers \+ literals  
    IDENT = "IDENT" // add, foobar, x, y, ...  
    INT   = "INT"   // 1343456

    // Operators  
    ASSIGN   = ":="  
    PLUS     = "+"  
    MINUS    = "-"  
    ASTERISK = "\*"  
    SLASH    = "/"  
    CARET    = "^"  
    DOT      = "."

    LT     = "<"  
    GT     = ">"  
    EQ     = "="  
    NOT_EQ = "<>"  
    LTE    = "<="  
    GTE    = ">="

    // Delimiters  
    COMMA     = ","  
    SEMICOLON = ";"  
    COLON     = ":"  
    LPAREN    = "("  
    RPAREN    = ")"

    // Keywords  
    FUNCTION  = "FUNCTION"  
    PROCEDURE = "PROCEDURE"  
    TYPE      = "TYPE"  
    VAR       = "VAR"  
    RECORD    = "RECORD"  
    IF        = "IF"  
    ELSE      = "ELSE"  
    THEN      = "THEN"  
    WHILE     = "WHILE"  
    DO        = "DO"  
    BEGIN     = "BEGIN"  
    END       = "END"  
    NIL       = "NIL"  
)

#### Token Struct

Represents a single token, containing its type, the literal text from the
source, and its position for error reporting.

    type Token struct {  
        Type    TokenType  
        Literal string  
        Line    int  
        Column  int  
    }

#### **Keyword Map**

A map to quickly look up whether an identifier is a keyword.

    var keywords = map[string]TokenType{  
        "function":  FUNCTION,  
        "procedure": PROCEDURE,  
        "type":      TYPE,  
        "var":       VAR,  
        "record":    RECORD,  
        "if":        IF,  
        "else":      ELSE,  
        "then":      THEN,  
        "while":     WHILE,  
        "do":        DO,  
        "begin":     BEGIN,  
        "end":       END,  
        "nil":       NIL,  
    }
    
    // LookupIdent checks the keywords map for an identifier.  
    // If found, it returns the keyword's token type. Otherwise, it returns IDENT.  
    func LookupIdent(ident string) TokenType {  
        if tok, ok := keywords[ident]; ok {  
            return tok  
        }  
        return IDENT  
    }

## 3. lexer Package Specification

### lexer.go

This file contains the Lexer struct and the core logic for producing tokens.

#### Lexer Struct

The Lexer holds the state required for scanning the input string.

    package lexer
    
    import "compiler/token"
    
    type Lexer struct {  
        input        string // The source code being scanned  
        position     int    // Current position in input (points to current char)  
        readPosition int    // Current reading position in input (after current char)  
        ch           byte   // Current char under examination  
        line         int    // Current line number  
        column       int    // Current column number  
    }

#### Public Functions

* New(input string) Lexer: A constructor function that initializes a new Lexer
  with the provided source code and sets its initial state.  

* NextToken() token.Token: The primary method of the lexer. It reads the input,
  identifies the next complete token, and returns it. It is responsible for:  

  * Skipping whitespace.  
  * Skipping comments (//...).  
  * Parsing identifiers and keywords.  
  * Parsing integer literals.  
  * Parsing single and multi-character operators (e.g., =, :=, <>).

## 4. `lexer_test.go` Specification

A table-driven test will be created to provide comprehensive coverage of the
lexer's functionality.

* TestNextToken: This will be the main test function.  
* Test Cases: The test will include a wide array of source code snippets to validate:  
  * Correct tokenization of all operators and delimiters.  
  * Correct tokenization of all keywords.  
  * Handling of single-line comments.  
  * Correct parsing of identifiers and integers.  
  * Proper handling of whitespace.  
  * A complete, simple Toy Pascal program snippet to test the lexer in a more integrated scenario.  
  * Accurate tracking of line and column numbers.

