package bridge

import "fmt"

type color interface {
	getName() string
}

type red struct{}

func (red *red) getName() string {
	return "red"
}

type blue struct{}

func (blue *blue) getName() string {
	return "blue"
}

type shape interface {
	draw()
}

type circle struct {
	color color
}

func (circle *circle) draw() {
	fmt.Println("drawing a circle with", circle.color.getName(), "color")
}

type square struct {
	color color
}

func (square *square) draw() {
	fmt.Println("drawing a square with", square.color.getName(), "color")
}

func callDraw(shape shape) {
	shape.draw()
}

// use this pattern when there are more than independent factors
func Main() {
	redColor := red{}
	buleColor := blue{}

	redSquare := square{color: &redColor}
	blueSquare := square{color: &buleColor}
	redCircle := circle{color: &redColor}
	blueCircle := circle{color: &buleColor}

	callDraw(&redSquare)
	callDraw(&blueSquare)
	callDraw(&redCircle)
	callDraw(&blueCircle)
}
