package main

import (
	"fmt"
)

type language interface {
	getGreeting() string
}

type english struct{}

func (english) getGreeting() string {
	return "Hello there"
}

type spanish struct{}

func (spanish) getGreeting() string {
	return "Ola!"
}

func printGreetings(l language) {
	fmt.Println(l.getGreeting())
}

func main() {
	e := english{}
	printGreetings(e)

	s := spanish{}
	printGreetings(s)
}
