package model

import (
	"math/rand"
)

type World struct {
	Height int
	Width  int
	Map    [][]Cell
	Statistic
}

type Statistic struct {
	CountOfEntity int
	CountOfFood   int
	AvgOfPoison   float64
}

func (w *World) GetDataCell(coordinates Coordinates) Cell {
	return w.Map[coordinates.X][coordinates.Y]
}

func (w *World) SetFoodCell(coordinates Coordinates, dfood bool) {
	w.Map[coordinates.X][coordinates.Y].Food = dfood
}

func (w *World) SetPoisonCell(coordinates Coordinates, dPoison int) {
	w.Map[coordinates.X][coordinates.Y].Poison += dPoison
}

func (w *World) UpdateEntityCell(coordinates Coordinates, entity *Entity) {
	w.Map[coordinates.X][coordinates.Y].Entity = entity
}

func CreateWorld(height, width int) World {
	world := World{
		height,
		width,
		make([][]Cell, height),
		Statistic{
			0,
			0,
			0,
		},
	}
	for x := 0; x < height; x++ {
		world.Map[x] = make([]Cell, width)
	}
	world.ClearWorld()
	return world
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
