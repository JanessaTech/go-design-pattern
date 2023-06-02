package adapter

import "fmt"

type adaptee struct {
}

func (adaptee *adaptee) customMethod() {
	fmt.Println("adaptee.customMethod is called")
}
