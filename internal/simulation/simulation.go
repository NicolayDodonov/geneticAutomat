package simulation

import (
	"fmt"
	"geneticAutomat/internal/model"
	"math/rand"
)

const population int = 1

func Run() {
	//TODO: init logger

	//TODO: init console print

	//TODO: create world
	var world model.World = model.CreateWorld(64, 64)
	world.GenerateWalls()
	world.GenerateFood(2)

	//TODO: create start population
	arrayEntity := make([]model.Entity, population)
	for i := 0; i < population; i++ {
		arrayEntity[i] = model.CreateEntity(rand.Intn(world.Height), rand.Intn(world.Width),
			model.RandomDNA())
	}
	//TODO: create goroutine of simulation
	for age := 0; age < 120; age++ {
		countLive := 0
		for i := 0; i < population; i++ {
			if arrayEntity[i].Hp > 0 {
				countLive++
				model.RunDNA(&arrayEntity[i], &world)
			} else if !(arrayEntity[i].Hp == -1) {
				world.UpdateEntityCell(arrayEntity[i].Coordinates, nil)
				world.SetPoisonCell(arrayEntity[i].Coordinates, 10)
				arrayEntity[i].Hp = -1
			}
		}
		if age%10 == 0 {
			world.CountOfEntity = countLive
			fmt.Print(countLive)
			fmt.Println(" age: ", age)
		}
	}
}
