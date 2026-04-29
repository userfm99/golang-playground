package main

import (
	"fmt"
)

type Message string

func NewMessage() Message {
	return "hello there"
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) greet() {
	fmt.Println(g.Message)
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	e.Greeter.greet()
}

func main() {

	msg := NewMessage()
	greeter := NewGreeter(msg)
	event := NewEvent(greeter)

	event.Start()
}
