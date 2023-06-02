package adapter

type adaptor struct {
	adaptee adaptee
}

func (adaptor *adaptor) NewMethod() {
	adaptor.adaptee.customMethod()
}
