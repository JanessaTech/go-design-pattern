package factory

import "fmt"

type Cat struct {
	name string
}

func (c *Cat) SetName(name string) {
	c.name = name
}

func (c *Cat) Saying() {
	fmt.Println("my name is ", c.name, ", I am a cat")
}

func NewCat(name string) Animal {
	cat := Cat{name: name}
	return &cat
}
