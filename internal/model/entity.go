package model

import (
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

func (e *Entity) RunDNA(w *World) {
	switch e.Array[e.PointerDNA] {
	case command["move"]:
		relativeCord := makeTurn(e.turn)
		absoluteCord := Sum(relativeCord, e.Coordinates) //абсолютные координаты
		if (absoluteCord.X < w.Width) &&
			(absoluteCord.Y < w.Height) &&
			(absoluteCord.X >= w.Width) &&
			(absoluteCord.Y >= w.Height) && //TODO: заменить эту проверку функцией
			!(w.GetDataCell(absoluteCord).Wall) &&
			(w.GetDataCell(absoluteCord).Entity == nil) {

			w.UpdateEntityCell(e.Coordinates, nil)
			e.Coordinates = Sum(relativeCord, e.Coordinates)
			w.UpdateEntityCell(e.Coordinates, e)
		}
	case command["turnLeft"]:
		e.turn--
		if e.turn < 0 {
			e.turn = 7
		}
	case command["turnRight"]:
		e.turn++
		if e.turn > 7 {
			e.turn = 0
		}
	case command["look"]:
		//if empty e.Pointer +=0
		if w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Wall {
			e.PointerDNA += 1
		} else if w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Food {
			e.PointerDNA += 2
		} else if w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Entity != nil {
			e.PointerDNA += 3
		}
	case command["get"]:

		if w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Wall {
			e.Hp -= 5
		} else if w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Food {
			e.Hp += 10
			w.SetFoodCell(Sum(makeTurn(e.turn), e.Coordinates), false)
			w.CountOfFood--
		} else if w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Entity != nil {
			e.attack(w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Entity, w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)))
		}
	case command["recycling"]:
		poisonLevel := w.GetDataCell(Sum(makeTurn(e.turn), e.Coordinates)).Poison
		var dPoison = 0
		if poisonLevel > 75 {
			dPoison = 20
		} else if poisonLevel > 50 {
			dPoison = 10
		} else if poisonLevel > 25 {
			dPoison = 5
		} else if poisonLevel > 5 {
			dPoison = 1
		}
		w.SetPoisonCell(Sum(makeTurn(e.turn), e.Coordinates), -dPoison)
		e.Hp += dPoison
		e.PointerDNA++
	case command["reproduction"]:
		//TODO: REMAKE THIS!!!!!!!!!!
		//w.insertNewEntity(CreateEntity(makeTurn(e.turn).X, makeTurn(e.turn).Y, Mutation(&e.DNA, 2)))
		TODO()
	case command["jump dna"]:
		dPointerDNA := e.PointerDNA + 1
		if dPointerDNA >= longDNA {
			dPointerDNA -= longDNA
		}
		e.PointerDNA += e.Array[dPointerDNA]
	}

	e.PointerDNA++
	if e.PointerDNA >= longDNA {
		e.PointerDNA -= longDNA
	}
	if e.PointerDNA < 0 {
		e.PointerDNA += longDNA
	}
	e.Hp--
	e.Age++
	w.SetPoisonCell(e.Coordinates, 1)
	if e.Hp <= 0 {
		e.IsLive = false
	}
}

func makeTurn(turn int) Coordinates {
	cordTurn := Coordinates{
		0,
		0,
	}
	switch turn {
	case 0:
		cordTurn.Y--
	case 1:
		cordTurn.X++
		cordTurn.Y--
	case 2:
		cordTurn.X++
	case 3:
		cordTurn.X++
		cordTurn.Y++
	case 4:
		cordTurn.Y++
	case 5:
		cordTurn.X--
		cordTurn.Y++
	case 6:
		cordTurn.X--
	case 7:
		cordTurn.X--
		cordTurn.Y--
	}
	return cordTurn
}

func (me *Entity) attack(another *Entity, cell *Cell) {
	if another.IsLive {
		me.Hp += another.Hp
		another.Hp = 0
		another.IsLive = false
		cell.Entity = nil
	}
}

func TODO() {

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

func (e *Entity) Mutation(count int) {
	for i := 0; i < count; i++ {
		e.DNA.Array[rand.Intn(longDNA-1)] = rand.Intn(8)
	}
}
