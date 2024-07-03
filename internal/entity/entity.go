package entity

import (
	"geneticAutomat/internal/world"
	"math/rand"
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
	id   int
	age  int
	Hp   int
	turn int
	world.Coordinates
	DNA
}

type DNA struct {
	PointerDNA int
	Array      [longDNA]int
}

func (e Entity) RunDNA(w world.World) {
	switch e.Array[e.PointerDNA] {
	case command["move"]:
		relativeCord := makeTurn(e.turn)
		absoluteCord := world.Sum(relativeCord, e.Coordinates) //абсолютные координаты
		if (absoluteCord.X < w.Width) &&
			(absoluteCord.Y < w.Height) &&
			(absoluteCord.X >= w.Width) &&
			(absoluteCord.Y >= w.Height) && //TODO: заменить эту проверку функцией
			!(w.GetDataCell(absoluteCord).Wall) &&
			(w.GetDataCell(absoluteCord).Entity == nil) {

			w.UpdateEntityCell(e.Coordinates, nil)
			e.Coordinates = world.Sum(relativeCord, e.Coordinates)
			w.UpdateEntityCell(e.Coordinates, &e)
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
		//if enmpty e.Pointer +=0
		if w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Wall {
			e.PointerDNA += 1
		} else if w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Food {
			e.PointerDNA += 2
		} else if w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Entity != nil {
			e.PointerDNA += 3
		}
	case command["get"]:
		if w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Wall {
			e.Hp -= 5
		} else if w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Food {
			e.Hp += 10
			w.SetFoodCell(world.Sum(makeTurn(e.turn), e.Coordinates), false)
		} else if w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Entity != nil {
			e.PointerDNA += 3
		}
	case command["recycling"]:
		poisonLevel := w.GetDataCell(world.Sum(makeTurn(e.turn), e.Coordinates)).Poison
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
		w.SetPoisonCell(world.Sum(makeTurn(e.turn), e.Coordinates), -dPoison)
		e.Hp += dPoison
		e.PointerDNA++
	case command["reproduction"]:
		TODO()
	case command["jump dna"]:
		e.PointerDNA += e.Array[e.PointerDNA+1]
	}

	e.PointerDNA++
	if e.PointerDNA > longDNA {
		e.PointerDNA -= longDNA
	}
	if e.PointerDNA < 0 {
		e.PointerDNA += longDNA
	}
	e.Hp--
	w.SetPoisonCell(e.Coordinates, 1)
}

func makeTurn(turn int) world.Coordinates {
	cordTurn := world.Coordinates{
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

func TODO() {

}

func Create(x, y int, dna DNA) Entity {
	return Entity{
		0,
		0,
		100,
		0,
		world.Coordinates{
			x,
			y,
		},
		dna,
	}
}

func RandomDNA() DNA {
	var dna DNA
	for i := 0; i < longDNA; i++ {
		dna.Array[i] = rand.Intn(longDNA)
	}
	dna.PointerDNA = rand.Intn(longDNA)
	return dna
}

func (dna1 *DNA) SetDNA(dna2 DNA) {
	*dna1 = dna2
}

func (dna DNA) Mutation(count int) {
	for i := 0; i < count; i++ {
		dna.Array[i] = rand.Intn(longDNA)
	}
}
