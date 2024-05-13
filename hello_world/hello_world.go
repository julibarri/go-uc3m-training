
package hello_world

import (
	"fmt"
)

func init () {
	fmt.Println("Press x for Hello World")
}

func HelloWorld() int {
	const (
		hello = "\U0001f44b"
		world = "\U0001f30e"
	)
	fmt.Println(" "+hello+"    "+world)
	fmt.Println("Hello World!")
	return 1
}
