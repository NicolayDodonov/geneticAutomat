package model

import (
	"geneticAutomat/internal/slogger"
)

const (
	left      = 1
	right     = -1
	lengthDNA = 64
	maxGene
)
const (
	move = iota
	rotatedLeft
	rotatedRight
	look
	get
	recycling
	reproduction
	jump
)

type Entity struct {
	Id     int
	Age    int
	Hp     int
	turn   int
	IsLive bool
	Coordinates
	DNA
}

func (e *Entity) RunDNA(w *World) {
	if !e.IsLive {
		slogger.LogEntityInfo.Debug("Run DNA in die entity")
		return
	}

	for countRun := 0; countRun < 10; {
		switch e.Array[e.PointerDNA] {
		case move:
			e.move(w)
			countRun += 5
		case rotatedLeft:
			e.rotation(left)
			countRun++
		case rotatedRight:
			e.rotation(right)
			countRun++
		case look:
			e.look(w)
			countRun += 2
		case get:
			e.get(w)
			countRun += 5
		case recycling:
			e.recycling(w)
			countRun += 5
		case reproduction:
			e.reproduction()
			countRun += 12
		default:
			e.jump()
			countRun++
		}
		e.PointerDNA++
		e.loopPointer()
	}
	e.Hp--
	e.Age++

	if e.Hp <= 0 && e.IsLive {
		e.IsLive = false
		w.UpdateEntityCell(e.Coordinates, nil)
		slogger.LogEntityInfo.Debug("I'm dying!!!!", "id", e.Id)

		if err := w.ChangeCellPoison(e.Coordinates, 10); err != nil {
			slogger.LogErrors.Error("Error update final RunDNA", err)
		}

		slogger.LogEntityInfo.Debug(
			"End RunDNA",
			"id", e.Id,
			"LiveStatus:", e.IsLive,
			"Hp", e.Hp,
			"Age", e.Age,
			"Coords", e.Coordinates)
		return
	}

	if err := w.ChangeCellPoison(e.Coordinates, 1); err != nil {
		slogger.LogErrors.Error("Error update final RunDNA", err)
	}

	slogger.LogEntityInfo.Debug(
		"End RunDNA",
		"id", e.Id,
		"LiveStatus:", e.IsLive,
		"Hp", e.Hp,
		"Age", e.Age,
		"Coords", e.Coordinates)

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
		slogger.LogErrors.Error("Move has failed", err)
	}
}

func (e *Entity) rotation(turnCount int) {
	e.turn += turnCount
	if e.turn > 7 {
		e.turn = 0
	}
	if e.turn < 0 {
		e.turn = 7
	}
	slogger.LogEntityInfo.Debug("rotation", "id", e.Id, "turnNow", e.turn)
}

func (e *Entity) look(w *World) {
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
}

func (e *Entity) get(w *World) {

	attack := func(another *Entity, cell *Cell) {
		if another.IsLive {
			e.Hp += another.Hp
			another.Hp = 0
			another.IsLive = false
			cell.Entity = nil
		}
	}

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
			attack(cell.Entity, cell)
			slogger.LogEntityInfo.Debug("Get Attack",
				"id", e.Id,
				"id victim", cell.Entity.Id)
		}
	} else {
		slogger.LogErrors.Error("Get is Fall", err)
	}

}

func (e *Entity) recycling(w *World) {
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
}

func (e *Entity) reproduction() {
	//TODO: reproduction
}

func (e *Entity) jump() {
	dPointerDNA := (e.PointerDNA + e.Array[e.PointerDNA]) % lengthDNA
	e.PointerDNA += dPointerDNA
	slogger.LogEntityInfo.Debug("jump dna", "id", e.Id, "pointer set", e.PointerDNA)
}

func (e *Entity) loopPointer() {
	if e.PointerDNA >= lengthDNA {
		e.PointerDNA -= lengthDNA
	}
	if e.PointerDNA < 0 {
		e.PointerDNA += lengthDNA
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
