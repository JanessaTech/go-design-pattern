package adapter

func call(target target) {
	target.NewMethod()
}

// for the case where methods for the old class needs to be adjusted to the new methods
func Main() {
	adaptee := adaptee{}
	adaptor := &adaptor{adaptee: adaptee}
	call(adaptor)
}
