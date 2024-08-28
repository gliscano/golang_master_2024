package main

import (
	"fmt"
	"math"
	"strconv"
)

func mathOperations() {
	/*================== data conversions*/
	fmt.Println("================== data conversions")

	var integer16 int16 = 50
	var integer32 = 100

	var convertionInt32 = int32(integer16)

	fmt.Println(convertionInt32 + int32(integer32))

	str := "100"

	// String to Number
	integer, _ := strconv.Atoi(str)
	fmt.Println(integer + integer32)

	// int to string
	n := 42
	str = strconv.Itoa(n)
	fmt.Println(str)

	/*================== fmt package > https://pkg.go.dev/fmt*/
	fmt.Println("================== fmt package > https://pkg.go.dev/fmt")

	name := "Guillermo"
	age := 35

	fmt.Printf("Hola, me llamo %s y tengo %d años. (fmt.Printf) \n", name, age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años. (fmt.Sprintf)", name, age)
	fmt.Println(greeting)

	var name2 string
	var age2 int

	fmt.Println("Ingrese su nombre")
	fmt.Scanln(&name2)
	fmt.Println("Ingrese su edad")
	fmt.Scanln(&age2)

	fmt.Printf("Hola, me llamo %s y tengo %d años. (fmt.Scan) \n", name2, age2)

	/*================== Math operations https://pkg.go.dev/math*/
	fmt.Println("================== Math operations https://pkg.go.dev/math")

	var base, altura float64
	const precision = 2

	fmt.Println("Calcular el área y perímetro de un triángulo")
	fmt.Println("Ingrese la base del triángulo")
	fmt.Scanln(&base)
	fmt.Println("Ingrese la altura del triángulo")
	fmt.Scanln(&altura)

	hipotenusa := math.Sqrt(math.Pow(base, 2) + math.Pow(altura, 2))
	area := (base * altura) / 2
	perimetro := base + altura + hipotenusa

	fmt.Printf("\nÁrea del tringulo: %.*f", precision, area)
	fmt.Printf("\nPerímetro del tringulo: %.*f", precision, perimetro)
}
