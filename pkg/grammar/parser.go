package grammar

import (
	"errors"
	"fmt"
)

// Context stores all necessary info from poll.
type Context struct {
	// Variables[id] is true iff option with id is selecteed
	Variables map[int]bool
}

// Term is base interface for all nodes
type Term interface {
	String() string
	Eval(context *Context) bool
}

type and struct {
	lhs Term
	rhs Term
}

func (a *and) Eval(context *Context) bool {
	return a.lhs.Eval(context) && a.rhs.Eval(context)
}

func (a *and) String() string {
	return fmt.Sprintf("(%s & %s)", a.lhs, a.rhs)
}

type or struct {
	lhs Term
	rhs Term
}

func (o *or) Eval(context *Context) bool {
	return o.lhs.Eval(context) || o.rhs.Eval(context)
}

func (o *or) String() string {
	return fmt.Sprintf("(%s | %s)", o.lhs, o.rhs)
}

type constant struct {
	id int
}

func (c *constant) Eval(context *Context) bool {
	return context.Variables[c.id]
}

func (c *constant) String() string {
	return fmt.Sprintf("%d", c.id)
}

var errNotImplemented = errors.New("not implemented yet")
var errNotReachable = errors.New("not reachable")

func parseTerm(tokens []Token) (Term, []Token, error) {
	if len(tokens) == 0 {
		return nil, nil, nil
	}

	errWrongType := errors.New("Expected open brace or sign at start of token")

	switch tokens[0].tokenType {
	case idToken:
		return &constant{id: tokens[0].intValue}, tokens[1:], nil

	case braceToken:
		if tokens[0].intValue != +1 {
			return nil, nil, errWrongType
		}

		val, rest, err := Parse(tokens[1:])
		if err != nil {
			return nil, nil, err
		}

		if len(rest) == 0 || rest[0].tokenType != braceToken || rest[0].intValue != -1 {
			return nil, nil, errors.New("no matching ')' for '('")
		}

		return val, rest[1:], nil

	case signToken:
		return nil, nil, errWrongType

	default:
		return nil, nil, errNotReachable
	}
}

// Parse builds term out of given tokens
func Parse(tokens []Token) (Term, []Token, error) {
	lhs, rest, err := parseTerm(tokens)
	if err != nil {
		return nil, nil, err
	}

	if len(rest) == 0 || rest[0].tokenType != signToken {
		return lhs, rest, nil
	}

	switch rest[0].intValue {
	case int('&'):
		rhs, rest, err := Parse(rest[1:])
		if err != nil {
			return nil, nil, err
		}
		return &and{lhs, rhs}, rest, nil

	case int('|'):
		rhs, rest, err := Parse(rest[1:])
		if err != nil {
			return nil, nil, err
		}
		return &or{lhs, rhs}, rest, nil

	default:
		return nil, nil, errNotReachable
	}
}
