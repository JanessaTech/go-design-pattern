package factory

import "fmt"

type dog struct {
	name string
}

func (d *dog) setName(name string) {
	d.name = name
}

func (d *dog) saying() {
	fmt.Println("my name is ", d.name, ", I am a dog")
}

func NewGog(name string) animal {
	dog := dog{name: name}
	return &dog
}
