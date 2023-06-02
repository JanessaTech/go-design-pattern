package command

import "fmt"

type command interface {
	action()
}

type realExecutor struct{}

func (realExecutor *realExecutor) m1() {
	fmt.Println("m1 is executing ...")
}
func (realExecutor *realExecutor) m2() {
	fmt.Println("m2 is executing ...")
}
func (realExecutor *realExecutor) m3() {
	fmt.Println("m3 is executing ...")
}

type command1 struct {
	real realExecutor
}

func (command1 *command1) action() {
	command1.real.m1()
}

type command2 struct {
	real realExecutor
}

func (command2 *command2) action() {
	command2.real.m2()
}

type command3 struct {
	real realExecutor
}

func (command3 *command3) action() {
	command3.real.m3()
}

func Main() {
	menu := map[int]command{}
	realExecutor := realExecutor{}
	menu[1] = &command1{real: realExecutor}
	menu[2] = &command2{real: realExecutor}
	menu[3] = &command3{real: realExecutor}

	menu[1].action()
	menu[2].action()
	menu[3].action()
}
