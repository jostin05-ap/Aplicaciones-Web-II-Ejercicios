package main

import "fmt"

func main() {
	var producto1 float64 = 25.50
	var producto2 float64 = 40.00
	var producto3 float64 = 15.75

	total := producto1 + producto2 + producto3
	promedio := total / 3
	descuento := total * 0.85

	fmt.Printf("producto 1 :$%.2f\n", producto1)
	fmt.Printf("producto 2 :$%.2f\n", producto2)
	fmt.Printf("producto 3 :$%.2f\n", producto3)
	fmt.Printf("TOTAL:$%.2f\n", total)
	fmt.Printf("PROMEDIO:$%.2f\n", promedio)
	fmt.Printf("TOTAL CON EL 15%% DESCUENTO:$%.2f\n", descuento)

}
