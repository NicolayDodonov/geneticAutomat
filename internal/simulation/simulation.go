package simulation

import (
	"fmt"
	"geneticAutomat/internal/world"
	"math/rand"
)

const population int = 1

func Run() {
	//TODO: init logger

	//TODO: init console print

	//TODO: create world
	var model world.World = world.CreateWorld(64, 64)
	model.GenerateWalls()
	model.GenerateFood(2)

	//TODO: create start population
	arrayEntity := make([]world.Entity, population)
	for i := 0; i < population; i++ {
		arrayEntity[i] = world.CreateEntity(rand.Intn(model.Height), rand.Intn(model.Width),
			world.RandomDNA())
	}
	//TODO: create goroutine of simulation
	for age := 0; age < 120; age++ {
		countLive := 0
		for i := 0; i < population; i++ {
			if arrayEntity[i].Hp > 0 {
				countLive++
				world.RunDNA(&arrayEntity[i], &model)
			} else if !(arrayEntity[i].Hp == -1) {
				model.UpdateEntityCell(arrayEntity[i].Coordinates, nil)
				model.SetPoisonCell(arrayEntity[i].Coordinates, 10)
				arrayEntity[i].Hp = -1
			}
		}
		if age%10 == 0 {
			model.CountOfEntity = countLive
			fmt.Print(countLive)
			fmt.Println(" age: ", age)
		}
	}
}
