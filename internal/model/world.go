package model

import (
	"math/rand"
)

type World struct {
	Height      int
	Width       int
	Map         [][]Cell
	ArrayEntity []Entity
	Statistic
}

type Statistic struct {
	CountOfEntity int
	CountOfFood   int
	WorldAge      int
	AvgOfPoison   float64
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
	world.ClearWorld()
	for i := 0; i < population; i++ {
		world.ArrayEntity[i] = CreateEntity(
			rand.Intn(world.Height-1)+1,
			rand.Intn(world.Width-1)+1,
			RandomDNA())
		world.ArrayEntity[i].Id += i + 1
		world.ArrayEntity[i].Hp += 1
		world.UpdateEntityCell(world.ArrayEntity[i].Coordinates, &world.ArrayEntity[i])
	}
	return world
}

func (w *World) insertNewEntity(entity Entity) {
	NotInsert := true
	for i := 0; i < len(w.ArrayEntity); i++ {
		if w.ArrayEntity[i].Hp == -1 || w.ArrayEntity[i].Id == 0 {
			w.ArrayEntity[i] = entity
			NotInsert = false
			break
		}
	}
	if NotInsert {
		w.ArrayEntity = append(w.ArrayEntity, entity)
	}
}

func (w *World) GetDataCell(coordinates Coordinates) *Cell {
	return &w.Map[coordinates.X][coordinates.Y]
}

func (w *World) SetFoodCell(coordinates Coordinates, dFood bool) {
	w.Map[coordinates.X][coordinates.Y].Food = dFood
}

func (w *World) SetPoisonCell(coordinates Coordinates, dPoison int) {
	w.Map[coordinates.X][coordinates.Y].Poison += dPoison
}

func (w *World) UpdateEntityCell(coordinates Coordinates, entity *Entity) {
	w.Map[coordinates.X][coordinates.Y].Entity = entity
}

func (w *World) ClearWorld() {
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			w.Map[x][y].Wall = false
			w.Map[x][y].Food = false
			w.Map[x][y].Poison = 0
			w.Map[x][y].Entity = nil
		}
	}
	w.GenerateWalls()
}

func (w *World) GenerateWalls() {
	for x := 0; x < w.Width; x++ {
		w.Map[x][0].Wall = true
		w.Map[x][w.Width-1].Wall = true
	}
	for y := 0; y < w.Height; y++ {
		w.Map[0][y].Wall = true
		w.Map[w.Height-1][y].Wall = true
	}
}

func (w *World) GenerateFood(foodChance int) {
	for x := 1; x < w.Width-1; x++ {
		for y := 1; y < w.Height-1; y++ {
			if (w.Map[x][y].Poison < 50) &&
				(rand.Intn(10) > foodChance) {
				w.Map[x][y].Food = true
			}
		}
	}
}

func (w *World) SortEntityByAge() {
	var leng = len(w.ArrayEntity)
	for i := 0; i < leng-1; i++ {
		for j := 0; j < leng-i-1; j++ {
			if w.ArrayEntity[j].Age > w.ArrayEntity[j+1].Age {
				w.ArrayEntity[j], w.ArrayEntity[j+1] = w.ArrayEntity[j+1], w.ArrayEntity[j]
			}
		}
	}
}
