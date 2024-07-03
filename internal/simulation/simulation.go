package simulation

import (
	"geneticAutomat/internal/entity"
	"geneticAutomat/internal/world"
	"math/rand"
)

func run() {
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

}
