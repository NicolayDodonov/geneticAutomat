package console

import (
	"geneticAutomat/internal/entityOLD"
)

type console interface {
	clearConsole(typeOS string)
	printBorder(startX, startY, endX, endY int)
	printEntity(array *[]entityOLD.Entity)
	printPoisonLevel()
}

func clearConsole(typeOS string) {

}

func printBorder(startX, startY, endX, endY int) {

}

func printEntity(array *[]entityOLD.Entity) {

}

func printPoisonLevel() {

}
