package strategy

import "fmt"

type iCalculator interface {
	cacl(op1, op2 int) int
}

type Context struct {
	iCalculator iCalculator
}

func (c *Context) calc(op1 int, op2 int) {
	res := c.iCalculator.cacl(op1, op2)
	fmt.Println("res =", res)
}

type Add struct{}

func (Add *Add) cacl(op1 int, op2 int) int {
	return op1 + op2
}

type Subtract struct{}

func (Subtract *Subtract) cacl(op1 int, op2 int) int {
	return op1 - op2
}

func Main() {
	add := Add{}
	subtract := Subtract{}
	c := Context{iCalculator: &add}
	c.calc(1, 2)
	c.iCalculator = &subtract
	c.calc(1, 2)
}
