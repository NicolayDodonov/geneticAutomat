package model

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type World struct {
	Height      int
	Width       int
	Map         [][]Cell
	ArrayEntity []Entity
	Statistic
}

type Cell struct {
	Poison int
	Food   bool
	Wall   bool
	*Entity
}

type Statistic struct {
	CountOfEntity int
	CountOfFood   int
	CountOfPoison int
	WorldAge      int
}

func CreateWorld(height, width, population int) World {
	world := World{
		height,
		width,
		make([][]Cell, height),
		make([]Entity, population*2),
		Statistic{
			population,
			0,
			0,
			0,
		},
	}
	for x := 0; x < height; x++ {
		world.Map[x] = make([]Cell, width)
	}
	world.Clear()
	for i := 0; i < population; i++ {
		world.ArrayEntity[i] = CreateEntity(
			rand.Intn(world.Height-2)+1,
			rand.Intn(world.Width-2)+1,
			RandomDNA())
		world.ArrayEntity[i].Id += i + 1
		world.ArrayEntity[i].Hp += 1
		world.UpdateEntityCell(world.ArrayEntity[i].Coordinates, &world.ArrayEntity[i])
	}
	return world
}

func (w *World) GetDataCell(coordinates Coordinates) (*Cell, error) {
	if coordinates.X >= 0 && coordinates.X < w.Width &&
		coordinates.Y >= 0 && coordinates.Y < w.Height {
		return &w.Map[coordinates.X][coordinates.Y], nil
	} else {
		return nil, errors.New(fmt.Sprintf("Can't get info in %+v", coordinates))
	}
}

func (w *World) ChangeCellFood(coordinates Coordinates, dFood bool) error {
	if coordinates.X >= 0 && coordinates.X < w.Width &&
		coordinates.Y >= 0 && coordinates.Y < w.Height {
		w.Map[coordinates.X][coordinates.Y].Food = dFood
		return nil
	} else {
		return errors.New(fmt.Sprintf("Can't set food in %+v", coordinates))
	}
}

// TODO:убрать get из названия
// TODO:убрать sum
func (w *World) GetCountFood() int {

	sumFood := 0
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			if w.Map[x][y].Food {
				sumFood++
			}
		}
	}
	return sumFood
}

// TODO:Полностью переписать
func (w *World) GenerateFood(foodChance int) {
	for x := 1; x < w.Width-1; x++ {
		for y := 1; y < w.Height-1; y++ {
			if (w.Map[x][y].Poison < 50) &&
				(rand.Intn(10) > foodChance) {
				w.Map[x][y].Food = true
				w.CountOfFood++
			}
		}
	}
}

func (w *World) ChangeCellPoison(coordinates Coordinates, dPoison int) error {
	if coordinates.X >= 0 && coordinates.X < w.Width &&
		coordinates.Y >= 0 && coordinates.Y < w.Height {

		w.Map[coordinates.X][coordinates.Y].Poison += dPoison

		if w.Map[coordinates.X][coordinates.Y].Poison > 100 {
			w.Map[coordinates.X][coordinates.Y].Poison = 100

		} else if w.Map[coordinates.X][coordinates.Y].Poison < 0 {
			w.Map[coordinates.X][coordinates.Y].Poison = 0

		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("Can't set poison in %+v", coordinates))
	}
}

// TODO:убрать get из названия
// TODO:убрать sum
func (w *World) GetCountPoison() int {
	sumPoison := 0
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			sumPoison += w.Map[x][y].Poison
		}
	}
	return sumPoison
}

// TODO:убрать get из названия
func (w *World) GetPercentPoison() float64 {
	maxPoison := w.Height * w.Width * 100
	var count int
	for x := 0; x < w.Height; x++ {
		for y := 0; y < w.Width; y++ {
			count += w.Map[x][y].Poison
		}
	}
	return math.Round((float64(count) / float64(maxPoison)) * 100)
}

// TODO: переписать?
func (w *World) UpdateEntityCell(coordinates Coordinates, entity *Entity) {
	w.Map[coordinates.X][coordinates.Y].Entity = entity
}

func (w *World) Clear() {
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			w.Map[x][y].Wall = false
			w.Map[x][y].Food = false
			w.Map[x][y].Poison = 0
			w.Map[x][y].Entity = nil
		}
	}
	w.GenerateBorderWalls()
}

func (w *World) GenerateBorderWalls() {
	for x := 0; x < w.Width; x++ {
		w.Map[x][0].Wall = true
		w.Map[x][w.Width-1].Wall = true
	}
	for y := 0; y < w.Height; y++ {
		w.Map[0][y].Wall = true
		w.Map[w.Height-1][y].Wall = true
	}
}

func (w *World) SortEntityByAge() {
	for i := 0; i < len(w.ArrayEntity)-1; i++ {
		for j := 0; j < len(w.ArrayEntity)-i-1; j++ {
			if w.ArrayEntity[j].Age < w.ArrayEntity[j+1].Age {
				w.ArrayEntity[j], w.ArrayEntity[j+1] = w.ArrayEntity[j+1], w.ArrayEntity[j]
			}
		}
	}
}
