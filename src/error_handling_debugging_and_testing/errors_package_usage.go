package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func wrapWithStack() error {
	err := createError()
	// do this when error comes from external source (go lib or vendor)
	return errors.Wrap(err, "wrapping an error with stack")
}

func wrapWithoutStack() error {
	err := createError()
	// do this when error comes from internal Fabric since it already has stack trace
	return errors.WithMessage(err, "wrapping an error without stack")
}

func createError() error {
	return errors.New("original error")
}

func main() {
	err := createError()
	fmt.Printf("print error without stack: %s\n\n", err)
	fmt.Printf("print error with stack: %+v\n\n", err)
	err = wrapWithoutStack()
	fmt.Printf("%+v\n\n", err)
	err = wrapWithStack()
	fmt.Printf("%+v\n\n", err)
}
