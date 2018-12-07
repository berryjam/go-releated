package main

import "fmt"

type PeopleIt interface {
	Show()
}

type StudentImpl struct{}

func (stu *StudentImpl) Show() {

}

func live() PeopleIt {
	var stu *StudentImpl
	if stu == nil {
		fmt.Println("stu is nil")
	}
	return stu
}

func main() {
	t := live()
	if t == nil {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
}
