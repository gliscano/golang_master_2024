package main

import (
	"fmt"
	"runtime"
	"time"
)

func controlStructure() {
	/* ================ IF */
	/* timeNow := time.Now()
	hour := timeNow.Hour()

	if hour < 12 {
		fmt.Println("¡Mañana!")
	} else if hour < 17 {
		fmt.Println("¡tarde!")
	} else {
		fmt.Println("¡Noche!")
	} */

	// Declaración dentro del if. Lo de arriba es igual a:
	if timeNow := time.Now(); timeNow.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if timeNow.Hour() < 17 {
		fmt.Println("¡tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	/* ================ Switch */

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
		break // En Go no es necesario agregar break para salir del case.
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> MacOS")
	default:
		fmt.Println("Go run -> otherOS")
	}

	// Switch with condition
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("¡Mañana!")
	case t.Hour() < 17:
		fmt.Println("¡tarde!")
	default:
		fmt.Println("¡noche!")
	}

	/* ================ For */
	for i := 0; i <= 10; i++ {
		if i == 3 {
			continue // Avanza a la próxima iteración sin ejecutar el index actual, en este caso, 3
		}
		fmt.Println(i)
		if i == 8 {
			break
		}
	}
}
