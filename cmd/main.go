package main

import (
	"fmt"
	"geneticAutomat/internal/simulation"
)

func main() {
	simulation.RunTrain(1000, 10)
	fmt.Scanln()
}
