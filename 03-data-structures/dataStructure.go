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

func maps() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Nota: El orden de los elementos del mapa es alfabetico, desde la inicialización y próximas adiciones.
	// Nota: Map devuelve 2 datos: value y ok (verificación si existe la key)

	colors["black"] = "#000000"
	fmt.Println("colors", colors)

	redValue, redOk := colors["red"]       // Existe la key
	whiteValue, whiteOk := colors["white"] // No existe la key, Resultado: Devuelve un string vacio y ok en false

	fmt.Println(redValue, redOk)
	fmt.Println(whiteValue, whiteOk)

	if value, ok := colors["blue"]; ok {
		fmt.Println("Si existe la key azul y su valor es: ", value)
	} else {
		fmt.Println("No existe la key ", value)
	}

	// Eliminar un elemento del mapa
	delete(colors, "red")

	for key, value := range colors {
		fmt.Printf("\nLa key %s tiene valor de: %s", key, value)
	}
}

// Definición de un tipo de dato
// struct es la representación de un objecto (Similar al modelo de una clase)
type Person struct {
	name  string
	age   int
	email string
}

// Crear un método de Person. Recibe un receptor un puntero de Person
// Un método pertenece a una estructura colocando un receptor y se accede mediante una instancia
func (p *Person) sayHello() {
	fmt.Println("Hola mi nombre es ", p.name)
}

func structFunction() {
	/* var p Person

	p.name = "Guille"
	p.age = 35
	p.email = "guille@gmail.com" */

	p := Person{"Guille", 35, "guille@gmail.com"}
	p.age = 36
	fmt.Println(p)

	p.sayHello()
}

func pointerStructFunction() {
	var x int = 10
	var p *int = &x // con & indica el puntero a la referencia de la memoria de la variable x

	fmt.Println(x)  // >> result: 10
	fmt.Println(p)  // >> result: 0xc00000a0b8 (Referencia de la memoria)
	fmt.Println(&x) // >> result: 0xc00000a0b8 (Referencia de la memoria)

	pointerEdit(&x) // sino se envia &x el cambio se estará realizando en una copia de la variable
	fmt.Println(x)  // >> result: 20
}

func pointerEdit(x *int) {
	*x = 20
}

func main() {
	/* Arrays */
	/* arrays() */
	/* slices() */
	/* maps() */
	/* pointerStructFunction() */
	structFunction()
}
