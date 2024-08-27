package console

import (
	"atomicgo.dev/cursor"
	"fmt"
	"geneticAutomat/internal/model"
)

type Console struct {
	Alphabet []byte
}

func (c *Console) Print(world *model.World, counter int) {
	var canvas [][]byte = make([][]byte, world.Height)
	for x := 0; x < world.Height; x++ {
		canvas[x] = make([]byte, world.Width+1)
		for y := 0; y < world.Width; y++ {
			if world.Map[x][y].Wall {
				canvas[x][y] = c.Alphabet[1]
			} else if world.Map[x][y].Food {
				canvas[x][y] = c.Alphabet[2]
			} else {
				canvas[x][y] = c.Alphabet[0]
			}
		}
		canvas[x][world.Width] = byte('\n')
	}
	for i := 0; i < len(world.ArrayEntity); i++ {
		if world.ArrayEntity[i].IsLive {
			canvas[world.ArrayEntity[i].X][world.ArrayEntity[i].Y] = 'e'
		}
	}
	for i := 0; i < len(canvas); i++ {
		fmt.Print(string(canvas[i]))
	}

	fmt.Println("Age:", world.WorldAge, "Life:", world.CountOfEntity,
		"              \nPoison:", world.AvgOfPoison, "Food", world.CountOfFood, "              ")
	cursor.Up(world.Height + 2)
}
