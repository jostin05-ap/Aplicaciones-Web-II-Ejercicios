package main

import "fmt"

func main() {
	var v1, v2 int
	v1 = 10
	v2 = 3

	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	sum := v1 + v2
	rest := v1 - v2
	mult := v1 * v2
	div := float64(v1 / v2)

	fmt.Printf("%.d + %.d = %.d\n", v1, v2, sum)
	fmt.Printf("%.d _ %.d = %.d\n", v1, v2, rest)
	fmt.Printf("%.d * %.d = %.d\n", v1, v2, mult)
	fmt.Printf("%.d / %.d = %.2f\n", v1, v2, div)

}
