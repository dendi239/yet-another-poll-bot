package grammar

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	idToken     = "idToken"
	signToken   = "signToken"
	braceToken  = "braceToken"
	negateToken = "negateToken"
)

// Token is stucture that represents one of: brace, identificator, operator-sign
type Token struct {
	tokenType string
	intValue  int
}

// Tokenize builds sequence of tokens out of given string
func Tokenize(s string) (tokens []Token, err error) {
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}
		n := len(tokens)

		switch {
		case unicode.IsSpace(c):
			continue

		case unicode.IsDigit(c):
			if len(tokens) > 0 && tokens[n-1].tokenType == idToken {
				tokens[n-1].intValue *= 10
				tokens[n-1].intValue += int(c - '0')
			} else {
				tokens = append(tokens, Token{idToken, int(c - '0')})
			}

		case c == '&' || c == '|':
			tokens = append(tokens, Token{signToken, int(c)})

		case c == '(' || c == ')':
			braceType := +1
			if c == ')' {
				braceType = -1
			}

			tokens = append(tokens, Token{braceToken, braceType})

		case c == '!':
			tokens = append(tokens, Token{negateToken, 0})

		default:
			err = fmt.Errorf("parse error: unknown token %c", c)
			return
		}
	}

	return
}

// TokenizeImplication builds two sequence of tokens out of patter A => B
func TokenizeImplication(message string) (t1 []Token, t2 []Token, err error) {
	strs := strings.Split(message, "=>")

	if len(strs) != 2 {
		if len(strs) <= 1 {
			err = errors.New("\"=>\" not found")
		} else {
			err = errors.New("Too many \"=>\" found")
		}
		return
	}

	s1, s2 := strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1])

	t1, err = Tokenize(s1)
	if err != nil {
		err = fmt.Errorf("failed to tokenize %v with %v", s1, err)
		return nil, nil, err
	}

	t2, err = Tokenize(s2)
	if err != nil {
		err = fmt.Errorf("failed to tokenize %v with %v", s2, err)
		return nil, nil, err
	}

	return
}
