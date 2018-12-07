package main

import "fmt"

type PeopleInterface interface {
	Speak(string) string
}

type Student struct{}

func (std Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo PeopleInterface = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
