package main

import (
	"strings"
)

type TokenType int

const (
    LiteralStr TokenType = iota
    Space
    EscapeSequence
    SingleQuoted
    DoubleQuoted
    BackQuoted
)

type Token struct {
    Literal string
    Type    TokenType
}

type Parser struct {
    input    string
    position int
}

func NewParser(input string) *Parser {
    return &Parser{input: input}
}

// Parse breaks the raw input into tokens, honoring SH quoting rules.
func (p *Parser) Parse() ([]Token, error) {
	var tokens []Token
	runes := []rune(p.input)
	length := len(runes)
	pos := 0

	for pos < length {
			switch r := runes[pos]; {
			case r == ' ':
					// collapse contiguous spaces
					for pos < length && runes[pos] == ' ' {
							pos++
					}
					tokens = append(tokens, Token{" ", Space})

			case r == '\'':
					pos++
					start := pos
					for pos < length && runes[pos] != '\'' {
							pos++
					}
					if pos >= length {
							return nil, &ParseError{"Unterminated single quote"}
					}
					tokens = append(tokens,
							Token{string(runes[start:pos]), SingleQuoted},
					)
					pos++ // skip closing '

			case r == '"':
					pos++
					var sb strings.Builder
					for pos < length && runes[pos] != '"' {
							if runes[pos] == '\\' && pos+1 < length {
									nxt := runes[pos+1]
									if nxt == '$' || nxt == '`' || nxt == '"' || nxt == '\\' {
											sb.WriteRune(nxt)
											pos += 2
											continue
									}
							}
							sb.WriteRune(runes[pos])
							pos++
					}
					if pos >= length {
							return nil, &ParseError{"Unterminated double quote"}
					}
					tokens = append(tokens,
							Token{sb.String(), DoubleQuoted},
					)
					pos++ // skip closing "

			case r == '`':
					pos++
					start := pos
					for pos < length && runes[pos] != '`' {
							pos++
					}
					if pos >= length {
							return nil, &ParseError{"Unterminated back quote"}
					}
					tokens = append(tokens,
							Token{string(runes[start:pos]), BackQuoted},
					)
					pos++ // skip closing `

			case r == '\\':
					if pos+1 >= length {
							return nil, &ParseError{"Dangling escape character"}
					}
					esc := runes[pos+1]
					tokens = append(tokens,
							Token{string(esc), EscapeSequence},
					)
					pos += 2

			default:
					start := pos
					for pos < length && !isSpecialChar(runes[pos]) {
							pos++
					}
					tokens = append(tokens,
							Token{string(runes[start:pos]), LiteralStr},
					)
			}
	}

	return tokens, nil
}


// ParseTokens runs Parse() and then groups tokens into argv-style []string.
func (p *Parser) ParseTokens() ([]string, error) {
		tokens, err := p.Parse()
    if err != nil {
        return nil, err
    }
		
    var args []string
    var current strings.Builder

    flush := func() {
        if current.Len() > 0 {
		        args = append(args, current.String())
            current.Reset()
        }
    }

    for _, tok := range tokens {
        switch tok.Type {
        case Space:
            flush()
        case LiteralStr, EscapeSequence, SingleQuoted, DoubleQuoted, BackQuoted:
            current.WriteString(tok.Literal)
        }
    }
    flush()
    return args, nil
}

func isSpecialChar(r rune) bool {
    return r == ' ' || r == '\'' || r == '"' || r == '`' || r == '\\'
}

type ParseError struct{ Message string }

func (e *ParseError) Error() string { return e.Message }

// func Test() {
// 	input := "echo 'single quoted $HOME' " +
// 	"\"double quoted \\\"\\$PATH\\\"\" " +
// 	"foo\\ bar `date`"
// 	parser := NewParser(input)

//     toks, err := parser.Parse()
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("TOKENS:")
//     for _, t := range toks {
//         fmt.Printf("  %#v\n", t)
//     }

//     args, err := parser.ParseTokens()
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("ARGS:", args)
// }
