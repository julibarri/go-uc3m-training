package helloworld

import (
	"flag"
	"fmt"
)

const (
	HELLO = "\U0001f44b"
	WORLD = "\U0001f30e"
)

var helloWorld bool

func init () {
	flag.BoolVar(&helloWorld, "hello-world", false, "Run or not Hello World.")
}

func HelloWorld() int {
	if !helloWorld {
		return 0
	}
	fmt.Println(" "+HELLO+"    "+WORLD)
	fmt.Println("Hello World!")
	return 1
}
