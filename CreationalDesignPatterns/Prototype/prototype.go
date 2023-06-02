package prototype

import "fmt"

type fruit interface {
	show()
	clone() fruit
}

type apple struct {
	from string
}

func (a *apple) show() {
	fmt.Println("This is an apple")
}
func (a *apple) clone() fruit {
	return &apple{from: a.from}
}

// when we want to create a new intstane according to the exsiting one
// a new intstane is deeply cloned from the old one but is different from old one
func Main() {
	apple1 := &apple{from: "xian"}
	apple2 := apple1.clone()
	fmt.Println("apple1:", apple1)
	fmt.Printf("addr of apple1=%p\n", apple1)
	fmt.Println("apple2:", apple2)
	fmt.Printf("addr of apple2=%p\n", apple2)
}
