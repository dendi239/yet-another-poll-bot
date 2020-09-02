package grammar

import "fmt"

type or struct {
	lhs Term
	rhs Term
}

type and struct {
	lhs Term
	rhs Term
}

type constant struct {
	id int
}

type negate struct {
	inner Term
}

func (n *negate) Eval(context *Context) bool {
	return !n.inner.Eval(context)
}

func (n *negate) String() string {
	return fmt.Sprintf("!%v", n.inner)
}

func (a *and) Eval(context *Context) bool {
	return a.lhs.Eval(context) && a.rhs.Eval(context)
}

func (a *and) String() string {
	return fmt.Sprintf("(%v & %v)", a.lhs, a.rhs)
}

func (o *or) Eval(context *Context) bool {
	return o.lhs.Eval(context) || o.rhs.Eval(context)
}

func (o *or) String() string {
	return fmt.Sprintf("(%v | %v)", o.lhs, o.rhs)
}

func (c *constant) Eval(context *Context) bool {
	return context.Variables[c.id]
}

func (c *constant) String() string {
	return fmt.Sprintf("%d", c.id)
}
