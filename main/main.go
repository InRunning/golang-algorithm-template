package main

import (
	"fmt"
	"github.com/go-errors/errors"
)

func main() {
	//err := test1()
	//fmt.Printf(err.(*errors.Error).ErrorStack())
	err := testNoStackErrors()
	err = errors.Wrap(err, 0)
	fmt.Printf(err.(*errors.Error).ErrorStack())
}

func test1() error {
	return test2()
}

func test2() error {
	return errors.Errorf("oh dear")
}
