package console

import (
	"atomicgo.dev/cursor"
	"fmt"
	"geneticAutomat/internal/model"
)

type Console struct {
	Array []byte
}

func (c *Console) Print(world *model.World, age int) {
	var canvas [][]byte = make([][]byte, world.Height)
	for x := 0; x < world.Height; x++ {
		canvas[x] = make([]byte, world.Width+1)
		for y := 0; y < world.Width; y++ {
			if world.Map[x][y].Entity != nil {
				canvas[x][y] = c.Array[2]
			} else if world.Map[x][y].Food {
				canvas[x][y] = c.Array[1]
			} else if world.Map[x][y].Wall {
				canvas[x][y] = c.Array[0]
			} else {
				canvas[x][y] = c.Array[3]
			}
		}
		canvas[x][world.Width] = byte('\n')
		fmt.Print(string(canvas[x]))
	}
	fmt.Println("Age: ", age, "Number of Life: ", world.CountOfEntity)
	cursor.Up(world.Height + 1)
}
