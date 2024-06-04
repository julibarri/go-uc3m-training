package conjunto

import (
	"flag"
	"fmt"
)

var conjunto bool

func init () {
	flag.BoolVar(&conjunto, "conjunto", false, "Run Conjunto set.")
}

func RunConjuntoActions() {
	fmt.Println("Running conjunto...")
	var conjunto Set
	conjunto.Añadir(1)
	fmt.Println("- Añadiendo 1")
	conjunto.Añadir(3)
	fmt.Println("- Añadiendo 3")
	if (conjunto.Existe(2)) {
		fmt.Println("- 2 existe")
	} else {
		fmt.Println("- 2 no existe")
	}
	if (conjunto.Existe(3)) {
		fmt.Println("- 3 existe")
	} else {
		fmt.Println("- 3 no existe")
	}
}
