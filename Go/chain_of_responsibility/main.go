package main

import "chain_of_responsibility/service"

func main() {
	requestA := service.Request{Type: "A"}
	requestB := service.Request{Type: "B"}
	requestC := service.Request{Type: "C"}
	requestD := service.Request{Type: "D"}
	startHandler := service.HandlerA{}
	startHandler.
		SetNext(&service.HandlerB{}).
		SetNext(&service.HandlerC{}).
		SetNext(&service.HandlerD{})

	startHandler.Handle(&requestA)
	startHandler.Handle(&requestB)
	startHandler.Handle(&requestC)
	startHandler.Handle(&requestD)
}
