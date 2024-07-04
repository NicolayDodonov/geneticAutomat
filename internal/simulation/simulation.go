package simulation

import (
	"geneticAutomat/internal/console"
	"geneticAutomat/internal/model"
	"math/rand"
	"time"
)

const population int = 10

func Run() {
	//TODO: init logger

	//TODO: init console print
	var console console.Console = console.Console{
		[]byte("H.E "),
	}

	//TODO: create world
	var world model.World = model.CreateWorld(25, 25)
	world.GenerateWalls()
	world.GenerateFood(2)

	//TODO: create start population
	arrayEntity := make([]model.Entity, population)
	for i := 0; i < population; i++ {
		arrayEntity[i] = model.CreateEntity(rand.Intn(world.Height-1)+1, rand.Intn(world.Width-1)+1,
			model.RandomDNA())
		world.UpdateEntityCell(arrayEntity[i].Coordinates, &arrayEntity[i])
	}
	//TODO: create goroutine of simulation
	for age := 0; age < 120; age++ {
		world.CountOfEntity = 0
		for i := 0; i < population; i++ {
			if arrayEntity[i].Hp > 0 {
				world.CountOfEntity++
				model.RunDNA(&arrayEntity[i], &world)
			} else if !(arrayEntity[i].Hp == -1) {
				world.UpdateEntityCell(arrayEntity[i].Coordinates, nil)
				world.SetPoisonCell(arrayEntity[i].Coordinates, 10)
				arrayEntity[i].Hp = -1
			}
		}
		time.Sleep(2 * time.Millisecond)
		//TODO: print world Frame
		console.Print(&world, age)
	}
}
