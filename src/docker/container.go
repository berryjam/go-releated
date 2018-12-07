package main

import "fmt"

type Container interface {
	BaseContainer

	Processes()
}

type BaseContainer interface {
	ID() string
}

type MacContainer struct {

}

func (container MacContainer) ID() string {
	return "MacContainer"
}

func main() {
	container := MacContainer{}
	fmt.Printf("%s\n", container.ID())
}