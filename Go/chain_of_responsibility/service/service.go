package service

import "fmt"

// Request is a request.
/** 这里request不是必需，一般这里传入的处理对象的某些属性可以控制责任链的调用归属
如果没有request，那么责任链的调用一般是按照责任链的顺序依次调用，这个时候可以使用错误类型
或者Handle()的返回值状态属性来控制责任链的调用进行或者结束
*/
type Request struct {
	Type string
}

// Handler is an interface for handling requests.
type Handler interface {
	Handle(request *Request)
	SetNext(Handler) Handler
}

// HandlerA is a handler for requests of type A.
type HandlerA struct {
	next Handler
}

// HandlerA.Handle implements Handler.Handle.
func (h *HandlerA) Handle(request *Request) {
	if request.Type == "A" {
		fmt.Println("HandlerA.Handle")
		doSomething(request.Type)
		return
	}
	if h.next != nil {
		h.next.Handle(request)
	}
}

// HandlerA.SetNext implements Handler.SetNext.
func (h *HandlerA) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// HandlerB is a handler for requests of type B.
type HandlerB struct {
	next Handler
}

// HandlerB.Handle implements Handler.Handle.
func (h *HandlerB) Handle(request *Request) {
	if request.Type == "B" {
		fmt.Println("HandlerB.Handle")
		doSomething(request.Type)
	}
	if h.next != nil {
		h.next.Handle(request)
	}
}

// HandlerB.SetNext implements Handler.SetNext.
func (h *HandlerB) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// HandlerC is a handler for requests of type C.
type HandlerC struct {
	next Handler
}

// HandlerC.Handle implements Handler.Handle.
func (h *HandlerC) Handle(request *Request) {
	if request.Type == "C" {
		fmt.Println("HandlerC.Handle")
		doSomething(request.Type)
	}
	if h.next != nil {
		h.next.Handle(request)
	}
}

// HandlerC.SetNext implements Handler.SetNext.
func (h *HandlerC) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// HandlerD is a handler for requests of type D.
type HandlerD struct {
	next Handler
}

// HandlerD.Handle implements Handler.Handle.
func (h *HandlerD) Handle(request *Request) {
	if request.Type == "D" {
		fmt.Println("HandlerD.Handle")
		doSomething(request.Type)
	}
	if h.next != nil {
		h.next.Handle(request)
	}
}

// HandlerD.SetNext implements Handler.SetNext.
func (h *HandlerD) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// doSomething is a function that does something.
func doSomething(val string) {
	fmt.Printf("val: %s", val)
	fmt.Println()
}
