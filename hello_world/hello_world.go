
package hello_world

import (
	"flag"
	"fmt"
)

execute := false

func init () {
	flag.BoolVar(&helloWorld, "hello-world", false, "Run or not Hello World.")
	flag.Parse()
	if helloWorld {
		execute = true
	}
}

func HelloWorld() int {
	if !execute {
		return 0
	}
	const (
		hello = "\U0001f44b"
		world = "\U0001f30e"
	)
	fmt.Println(" "+hello+"    "+world)
	fmt.Println("Hello World!")
	return 1
}
