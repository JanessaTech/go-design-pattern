package factory

import "fmt"

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) Saying() {
	fmt.Println("my name is ", d.name, ", I am a dog")
}

func NewGog(name string) Animal {
	dog := Dog{name: name}
	return &dog
}
