package simulation

import (
	"fmt"
	"geneticAutomat/internal/entity"
	"geneticAutomat/internal/world"
	"math/rand"
)

func Run() {
	//TODO: init logger

	//TODO: init console print

	//TODO: create world
	var model world.World = world.Create(64, 64)
	model.GenerateWalls()
	model.GenerateFood(2)

	//TODO: create start population
	arrayEntity := make([]entity.Entity, 64)
	for i := 0; i < 64; i++ {
		arrayEntity[i] = entity.Create(rand.Intn(model.Height), rand.Intn(model.Width),
			entity.RandomDNA())
	}
	//TODO: create goroutine of simulation
	for age := 0; age < 100000; age++ {
		countLive := 0
		for _, mob := range arrayEntity {
			if mob.Hp > 0 {
				countLive++
				mob.RunDNA(model)
			} else if !(mob.Hp == -1) {
				model.UpdateEntityCell(mob.Coordinates, nil)
				model.SetPoisonCell(mob.Coordinates, 10)
				mob.Hp = -1
			}
		}
		model.CountOfEntity = countLive
		fmt.Println(countLive)
	}
}
