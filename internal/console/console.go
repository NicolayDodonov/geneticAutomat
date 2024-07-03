package console

import (
	"geneticAutomat/internal/entity"
)

type console interface {
	clearConsole(typeOS string)
	printBorder(startX, startY, endX, endY int)
	printEntity(array *[]entity.Entity)
	printPoisonLevel()
}

func clearConsole(typeOS string) {

}

func printBorder(startX, startY, endX, endY int) {

}

func printEntity(array *[]entity.Entity) {

}

func printPoisonLevel() {

}
