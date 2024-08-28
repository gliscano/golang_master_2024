package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world. From code")
	fmt.Println(quote.Hello(), "From third party module")
	fmt.Println(quote.Go(), "From third party module")

	// constAndVariables()
	mathOperations()
}
