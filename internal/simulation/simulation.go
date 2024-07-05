package simulation

import (
	"geneticAutomat/internal/console"
	"geneticAutomat/internal/model"
	"time"
)

const population int = 10

func Run() {
	//TODO: init logger

	//TODO: init console print
	var printer console.Console = console.Console{
		[]byte("H.E "),
	}
	_ = printer
	//TODO: create world
	var world model.World = model.CreateWorld(25, 25, population)
	world.GenerateWalls()
	world.GenerateFood(2)

	//TODO: create goroutine of simulation
	for age := 0; age < 120; age++ {
		world.CountOfEntity = 0
		for i := 0; i < len(world.ArrayEntity); i++ {
			if world.ArrayEntity[i].Hp > 0 {
				world.ArrayEntity[i].RunDNA(&world)
				world.CountOfEntity++
			} else if world.ArrayEntity[i].Hp != -1 {
				world.ArrayEntity[i].Hp = -1
				world.UpdateEntityCell(world.ArrayEntity[i].Coordinates, nil)
			}
		}
		time.Sleep(2 * time.Millisecond)
		//TODO: print world Frame
		printer.Print(&world, age)
	}
}
