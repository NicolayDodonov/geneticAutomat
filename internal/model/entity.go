package model

import (
	"geneticAutomat/internal/slogger"
	"math/rand"
	"strconv"
)

const longDNA = 64

var command = map[string]int{
	"move":         0,
	"turnLeft":     1,
	"turnRight":    2,
	"look":         3,
	"get":          4,
	"recycling":    5,
	"reproduction": 6,
	"jump dna":     7,
}

type Entity struct {
	Id     int
	Age    int
	Hp     int
	turn   int
	IsLive bool
	Coordinates
	DNA
}

type DNA struct {
	PointerDNA int
	Array      [longDNA]int
}

func TODO() {
}

// TODO:Разобрать свич на отдельные функции
func (e *Entity) RunDNA(w *World) {
	for countRun := 0; countRun < 10; {
		switch e.Array[e.PointerDNA] {
		case command["move"]:
			e.move(w)

			countRun += 5
		case command["turnLeft"]:
			e.turn--
			if e.turn < 0 {
				e.turn = 7
			}
			slogger.LogEntityInfo.Debug("TurnLeft", "id", e.Id, "turnNow", e.turn)

			countRun++
		case command["turnRight"]:
			e.turn++
			if e.turn > 7 {
				e.turn = 0
			}
			slogger.LogEntityInfo.Debug("TurnRight", "id", e.Id, "turnNow", e.turn)

			countRun++
		case command["look"]:
			cell, err := w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates))
			if err == nil {
				//if empty e.Pointer +=0
				if cell.Wall {
					e.PointerDNA += 1
					slogger.LogEntityInfo.Debug("Look on Wall", "id", e.Id)
				} else if cell.Food {
					e.PointerDNA += 2
					slogger.LogEntityInfo.Debug("Look on Food", "id", e.Id)
				} else if cell.Entity != nil &&
					cell.Entity.IsLive {
					e.PointerDNA += 3
					slogger.LogEntityInfo.Debug("Look on Entity", "id", e.Id)
				}
			} else {
				slogger.LogErrors.Error("Look is Fall", err)
			}

			countRun += 2
		case command["get"]:
			cell, err := w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates))
			if err == nil {
				if cell.Wall {
					e.Hp -= 5
					slogger.LogEntityInfo.Debug("Get Wall", "id", e.Id, "You a dumb?", "Yes")
				} else if cell.Food {
					e.Hp += 10
					errUpdate := w.ChangeCellFood(Sum(makeTurn(e.turn), e.Coordinates), false)
					if errUpdate != nil {
						slogger.LogEntityInfo.Debug("Get Food", "id", e.Id)
					} else {
						slogger.LogErrors.Error("Error Update", errUpdate)
					}
				} else if cell.Entity != nil {
					e.attack(cell.Entity, cell)
					slogger.LogEntityInfo.Debug("Get Attack", "id", e.Id, "id victim",
						cell.Entity.Id)
				}
			} else {
				slogger.LogErrors.Error("Get is Fall", err)
			}
			
			countRun += 5
		case command["recycling"]:
			cell, err := w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates))
			var dPoison = 0

			if err == nil {
				if cell.Poison > 75 {
					dPoison = 20
				} else if cell.Poison > 50 {
					dPoison = 10
				} else if cell.Poison > 25 {
					dPoison = 5
				} else if cell.Poison > 5 {
					dPoison = 1
				}
				errUpdate := w.ChangeCellPoison(Sum(makeTurn(e.turn), e.Coordinates), -dPoison)
				if errUpdate != nil {
					e.Hp += dPoison
					slogger.LogEntityInfo.Debug("recycling", "id", e.Id, "poisonLevel", cell.Poison,
						"coord", Sum(makeTurn(e.turn), e.Coordinates))
				} else {
					slogger.LogErrors.Error("Error Update", errUpdate)
				}
			} else {
				slogger.LogErrors.Error("Recycling is Fall", err)
			}
			countRun += 5
		case command["reproduction"]:
			//TODO: REMAKE THIS!!!!!!!!!!
			//w.insertNewEntity(CreateEntity(makeTurn(e.turn).X, makeTurn(e.turn).Y, Mutation(&e.DNA, 2)))
			TODO()
			countRun += 12
		default:
			dPointerDNA := (e.PointerDNA + e.Array[e.PointerDNA]) % longDNA
			e.PointerDNA += dPointerDNA
			slogger.LogEntityInfo.Debug("jump dna", "id", e.Id, "pointer set", e.PointerDNA)
			countRun++
		}
		e.PointerDNA++
		if e.PointerDNA >= longDNA {
			e.PointerDNA -= longDNA
		}
		if e.PointerDNA < 0 {
			e.PointerDNA += longDNA
		}
	}

	e.Hp--
	e.Age++
	errUpdate := w.ChangeCellPoison(e.Coordinates, 1)
	if errUpdate == nil {
		slogger.LogErrors.Error("Error update final RunDNA", errUpdate)
	}
	if e.Hp <= 0 {
		e.IsLive = false
		slogger.LogEntityInfo.Debug("I'm dying!!!!", "id", e.Id)
	}
	slogger.LogEntityInfo.Debug("End RunDNA", "id", e.Id, "LiveStatus:", e.IsLive, "Hp", e.Hp,
		"Age", e.Age, "Coords", e.Coordinates)
}

func (e *Entity) move(w *World) {
	relativeCord := makeTurn(e.turn)
	absoluteCord := Sum(relativeCord, e.Coordinates) //абсолютные координаты
	cell, err := w.GetDataCell(absoluteCord)
	if err == nil {
		if (absoluteCord.X < w.Width) &&
			(absoluteCord.Y < w.Height) &&
			(absoluteCord.X >= w.Width) &&
			(absoluteCord.Y >= w.Height) && //TODO: заменить эту проверку функцией
			!(cell.Wall) &&
			(cell.Entity == nil) {

			w.UpdateEntityCell(e.Coordinates, nil)
			e.Coordinates = Sum(relativeCord, e.Coordinates)
			w.UpdateEntityCell(e.Coordinates, e)
		}
		slogger.LogEntityInfo.Debug("Move", "id", e.Id, "past", Del(e.Coordinates, relativeCord),
			"new", e.Coordinates)
	} else {
		slogger.LogErrors.Error("Move is Fall", err)
	}
}

func (me *Entity) attack(another *Entity, cell *Cell) {
	if another.IsLive {
		me.Hp += another.Hp
		another.Hp = 0
		another.IsLive = false
		cell.Entity = nil
	}
}

func CreateEntity(x, y int, dna DNA) Entity {
	return Entity{
		0,
		0,
		100,
		0,
		true,
		Coordinates{
			x,
			y,
		},
		dna,
	}
}

func (dna DNA) GoString() (stringDNA string) {
	for i := 0; i < len(dna.Array); i++ {
		stringDNA += strconv.Itoa(dna.Array[i]) + ", "
	}
	return stringDNA
}
func RandomDNA() DNA {
	var dna DNA
	for i := 0; i < longDNA; i++ {
		dna.Array[i] = rand.Intn(longDNA - 1)
	}
	dna.PointerDNA = rand.Intn(longDNA - 1)
	return dna
}

func (dna1 *DNA) SetDNA(dna2 DNA) {
	*dna1 = dna2
}

func (e *Entity) MutationDNA(count int) {
	for i := 0; i < count; i++ {
		e.DNA.Array[rand.Intn(longDNA-1)] = rand.Intn(8)
	}
}
