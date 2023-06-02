package factory

import "fmt"

type cat struct {
	name string
}

func (c *cat) setName(name string) {
	c.name = name
}

func (c *cat) saying() {
	fmt.Println("my name is ", c.name, ", I am a cat")
}

func NewCat(name string) animal {
	cat := cat{name: name}
	return &cat
}
