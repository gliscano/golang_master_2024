package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().UnixNano()) // envia el tiempo en nano segundos (Para versiones de Go anteriores a 1.20).
	play()
}

func play() {
	var numberEntered int
	var attempts int
	const maximumAttempts = 10

	numRandom := rand.Intn(100)

	for attempts < maximumAttempts {
		attempts++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maximumAttempts-attempts+1)
		fmt.Scanln(&numberEntered)

		if numberEntered == numRandom {
			fmt.Printf("¡Felicitaciones, adivinaste el número!")
			playAgain()
			return
		} else if numberEntered < numRandom {
			fmt.Printf("El número a adivinar es mayor.")
		} else if numberEntered > numRandom {
			fmt.Printf("El número a adivinar es menor.")
		}
	}

	fmt.Printf("\nSe acabaron los intentos... El número era %d", numRandom)
	playAgain()
}

func playAgain() {
	var optionSelected string
	fmt.Println("Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&optionSelected)

	switch optionSelected {
	case "s":
		play()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Elección inválida. Inténtalo nuevamente")
		playAgain()
	}
}
