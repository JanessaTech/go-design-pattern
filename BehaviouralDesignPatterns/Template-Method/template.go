package templatemethod

import "fmt"

type iTemplate interface {
	m1()
	m2()
}

type common struct {
	iTemplate iTemplate
}

func (c *common) do() {
	c.iTemplate.m1()
	c.iTemplate.m2()
}

type concrete1 struct {
}

func (concrete1 *concrete1) m1() {
	fmt.Println("m1...")
}
func (concrete1 *concrete1) m2() {
	fmt.Println("m2...")
}

func Main() {
	concrete1 := concrete1{}
	c := common{iTemplate: &concrete1}
	c.do()

}

//https://golangbyexample.com/all-design-patterns-golang/
