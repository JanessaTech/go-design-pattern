package chainresponsibility

import "fmt"

type iRequest interface {
	execute()
}

type Request struct {
	next iRequest
	name string
}

func (Request *Request) execute() {
	fmt.Println(Request.name, "is executing ...")
	if Request.next != nil {
		Request.next.execute()
	}
}

func Main() {
	req1 := Request{name: "req1"}
	req2 := Request{next: &req1, name: "req2"}
	req3 := Request{next: &req2, name: "req3"}

	req3.execute()

}
