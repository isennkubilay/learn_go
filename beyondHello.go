package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	//str := "Learn Go!" // string type.
	// s2 := `A "raw" string
	// literal can include line braks`

	// f := 3.14195 // float64
	// c := 3 + 4i  // complex128
	var u uint = 7
	fmt.Println("u:", u)
	var pi float32 = 22. / 7
	fmt.Println(pi)
	// n := byte('\n')
	// fmt.Println((n))
	var a4 [4]int
	fmt.Println(a4)                  // [0 0 0 0]
	a5 := [...]int{3, 1, 5, 10, 100} // [3 1 5 10 100]
	fmt.Println(a5)
	a4_cpy := a4
	a4_cpy[0] = 25
}
