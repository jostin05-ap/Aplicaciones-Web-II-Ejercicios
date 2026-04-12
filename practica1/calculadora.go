package main

import "fmt"

func main() {
	var v1, v2 int
	var op string

	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	fmt.Println("Ingresa el primer número")
	fmt.Scanln(&v1)
	fmt.Println("Ingresa el segundo número")
	fmt.Scanln(&v2)

	fmt.Println("Ingresa la operación (+, -, *, /,^,!):")
	fmt.Scanln(&op)

	switch op {
	case "+":

		sum := v1 + v2
		fmt.Printf("%d + %d = %d\n", v1, v2, sum)
		break

	case "-":

		rest := v1 - v2
		fmt.Printf("%d - %.d = %d\n", v1, v2, rest)
		break

	case "*":

		mult := v1 * v2
		fmt.Printf("%d * %d = %.d\n", v1, v2, mult)
		break

	case "/":
		if v2 == 0 {
			fmt.Println("Error: no se puede dividir entre cero")
		} else {
			div := float64(v1) / float64(v2)
			fmt.Printf("%.d / %d = %.2f\n", v1, v2, div)
		}
		break

	case "^":
		au := 1

		for i := 1; i <= v2; i++ {

			au = au * v1

		}
		fmt.Printf("%.d^%d = %.d\n", v1, v2, au)

		break
	case "!":
		au2 := 1
		for i := 1; i <= v1; i++ {

			au2 = au2 * i

		}
		fmt.Printf("%.d!= %.d\n", v1, au2)
		break

	default:
		fmt.Println("Operación no válida")

	}

}
