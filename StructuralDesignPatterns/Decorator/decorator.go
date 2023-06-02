package decorator

import "fmt"

type decorator interface {
	wear()
}

type woman struct {
}

func (woman *woman) wear() {
	fmt.Println("I am naked")
}

type DressDecorator struct {
	decorator decorator
}

func (DressDecorator *DressDecorator) wear() {
	DressDecorator.decorator.wear()
	fmt.Println("I am wearing dress")
}

type ShoesDecorator struct {
	decorator decorator
}

func (ShoesDecorator *ShoesDecorator) wear() {
	ShoesDecorator.decorator.wear()
	fmt.Println("I am wearing shoes")
}

func Main() {
	woman := woman{}
	dress := DressDecorator{decorator: &woman}
	shoes := ShoesDecorator{decorator: &dress}
	shoes.wear()
}
