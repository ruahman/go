package generics

import (
	"fmt"
)

// no type check with this approach
func print(output ...interface{}) {
	fmt.Println(output)
}

type Type interface {
	int | string
}

// make sure all types are the same
func print_generic[T Type](output ...T) {
	fmt.Println(output)
}

// any is nicer than interface{}
func print_any[T any](output ...T) {
	fmt.Println(output)
}

func Demo() {
	fmt.Println("----- generics -----")
	print(1, 2, 3, 4, 5)
	print(1, "2", 3.3)
	print_generic(6, 7, 8, 9)

	// can't do this since Type only work for int and string
	// print_generic(6.6, 7.7, 8.8, 9.9)
	print_generic("6", "7", "8", "9")

	print_any(6.6, 7.7, 8.8, 9.9)

	// this causes a problem
	// print_generic(6, 7.7, "8", 9)
}
