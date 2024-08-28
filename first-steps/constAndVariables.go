package main

import (
	"fmt"
	"math"
)

// declaración de constante
const Pi float32 = 3.14

// declaración multiples de constante con tipo de dato inferido
const (
	x = 100
	y = 0b1010 // valor de 10 en Binario
	z = 0o12   // valor de 10 en Octal
	w = 0xFF   // valor de 255 en Hexadecimal
)

/* const (
	Domingo = 1
	Lunes = 2
	Martes = 3
	...
) */

// Lo de arriba es igual a usar iota. iota es un identificador predeclarado. Está indexado a cero.
const (
	Domingo = iota + 1
	Lunes
	Martes
)

func constAndVariables() {
	// Declaración de forma larga y asiganción de variables separada
	/* var firstName, lastName string
	var age int

	firstName = "Guillermo"
	lastName = "Rojas"
	age = 35 */

	// Lo anterior es igual a
	/* var (
		firstName string = "Guillermo"
		lastName  string = "Rojas"
		age       int    = 35
	) */

	// Lo anterior es igual a
	/* var firstName, lastName, age = "Guillermo", "Rojas", 35 */

	// NOTA: En Go, no se considera idiomático el uso de variable larga cuando se inicia el valor
	// NOTA: Los nombres de las variables distinguen mayúsculas y minúsculas.
	// NOTA: En Go, si una variable comienza con una letra mayúscula, se puede acceder a dicha variable fuera del paquete en el que se declaró (recibió el valor exported). Ejemplo: var Email string
	// NOTA: Los nombres de las variables no pueden comenzar por números
	// NOTA: Los guiones bajos no son convencionales.

	// Declaración y asignación de forma corta y tipo de dato inferido
	firstName, lastName := "Guillermo", "Rojas"
	age := 35

	fmt.Println(firstName, lastName, age)
	fmt.Println("Pi", Pi)
	fmt.Println(x, y, z, w)
	fmt.Println("Martes", Martes)

	/* ============== Tipos de datos básicos */
	/*
		integer8 int8   // int8 es un conjunto de enteros con signo de 8-bit. Range: -128 through 127.
		integer16 int16 // int16 es un conjunto de enteros con signo de 16-bit. Range: -32768 through 32767
		integer32 int32 // int32 es un conjunto de enteros con signo de 32-bit. Range: -2147483648 through 2147483647.
		integer64 int64 // int64 es un conjunto de enteros con signo de 64-bit. Range: -9223372036854775808 through 9223372036854775807.

		integerUint uint // uint es un tipo entero positivos contiene 8, 16, 32 y 64 bits.
		integer int      // La cantidad de numeros dependerá de la arquitectura del Sistema operativo

		decimal32 float32 // Datos decimales de 32-bit
		decimal64 float64 // Datos decimales muy grandes de 64-bit

		valueBool bool 		// true and false

		b byte = 'a' 			//  valor expresado en byte 97
		r rune = '❤' 			 // valor expresado en unicode 10084
		valueStr string 	// cadena de carácteres
	*/

	fmt.Println("Int (Minimo y Máximo)", math.MinInt64, math.MaxInt64)
	fmt.Println("MaxUint8", math.MaxUint8)
	fmt.Println("MaxFloat32", math.MaxFloat32)

	var a byte = 'a'
	fmt.Println(a)

	var r rune = '❤'
	fmt.Println(r)

	// Cadena de texto con simbolos especiales (Comillas y salto de línea)
	name := "Guillermo Rojas \t (\"gliscano\")\n"
	fmt.Println(name)
}
