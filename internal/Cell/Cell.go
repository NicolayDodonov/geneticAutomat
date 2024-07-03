package Cell

import "geneticAutomat/internal/entity"

type Cell struct {
	Poison int
	Food   bool
	Wall   bool
	*entity.Entity
}

type Coordinates struct {
	X int
	Y int
}
