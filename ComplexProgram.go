package main

import "fmt"
import "math/cmplx"

func main() {
	var a complex128
	a = cmplx.Acos(-2 + 12i)

	fmt.Println("The value of a:", a)
}
