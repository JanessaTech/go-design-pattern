package observer

import "fmt"

type iSubject interface {
	registerObserver(o iObserver)
	notifyAll(msg string)
}
type subject struct {
	observers []iObserver
}

func (s *subject) registerObserver(o iObserver) {
	s.observers = append(s.observers, o)
}
func (s *subject) notifyAll(msg string) {
	for _, o := range s.observers {
		o.update(msg)
	}
}

type iObserver interface {
	update(msg string)
}
type observer struct {
	id int
}

func (o *observer) update(msg string) {
	fmt.Printf("received msg {%s} from observer #%d\n", msg, o.id)
}

func Main() {
	subject := subject{observers: []iObserver{}}
	o1 := observer{id: 1}
	o2 := observer{id: 2}
	o3 := observer{id: 3}
	o4 := observer{id: 4}
	subject.registerObserver(&o1)
	subject.registerObserver(&o2)
	subject.registerObserver(&o3)
	subject.registerObserver(&o4)
	subject.notifyAll("Hello JanessaTech")
}
