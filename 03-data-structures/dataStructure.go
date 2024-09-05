package main

import "fmt"

func arrays() {
	var matriz [5]int
	matriz[0] = 10

	// Declarar e inicializar con cantidad fija
	var matriz2 = [5]int{10, 20, 30, 40, 50}

	fmt.Println("matriz: ", matriz, "\n matriz2: ", matriz2)

	// Declarar e inicializar con cantidad desconocida
	var matriz3 = [...]int{10, 20, 30, 40, 50}

	for i := 0; i < len(matriz3); i++ {
		fmt.Println(i, ":", matriz3[i])
	}

	for index, value := range matriz3 {
		fmt.Printf("Índice: %d, Valor: %d \n", index, value)
	}

	// Declarar una matriz bidimensional. Cada elemento {} representa una fila y cada elemento interno (número o otro) representa una columna
	var matriz4 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz4)
}

func slices() {
	diasFullSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sábado"}

	diasSemana := diasFullSemana[0:5]
	// Agregar elementos
	diasSemana = append(diasSemana, "Viernes", "Sabado")
	diasSemana = append(diasSemana, "otro día de fin de semana :-)")

	// eliminar elementos
	// [desde : hasta no incluido] -> [:2] -> desde 0 hasta 2 (no incluido)
	diasSemana = append(diasSemana[:2], diasSemana[3:]...)

	fmt.Println(diasSemana, " Total días: ", len(diasSemana), " Capacidad de la semana", cap(diasSemana))

	// Make function
	nombres := make([]string, 5)
	nombres[0] = "Guille"
	nombres[1] = "Day"
	nombres[2] = "Deisy"

	apellidos := make([]string, 5, 10)
	apellidos[0] = "Rojas"
	apellidos[1] = "Fuentes"
	apellidos[2] = "Fuentes"

	// Copy function (destino, fuente)
	resultCopy := copy(apellidos, nombres)

	fmt.Println("\nNombres: ", nombres, " Total de Nombres: ", len(nombres), " Capacidad de Nombres", cap(nombres))
	fmt.Println("\nApellidos: ", apellidos, " Total de apellidos: ", len(apellidos), " Capacidad de apellidos", cap(apellidos))
	fmt.Println("\nResultCopy: ", resultCopy)
}

func main() {
	// Arrays
	// arrays()
	slices()
}
