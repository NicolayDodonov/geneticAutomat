package model

type Cell struct {
	Poison int
	Food   bool
	Wall   bool
	*Entity
}

type Coordinates struct {
	X int
	Y int
}

func Sum(a, b Coordinates) Coordinates {
	SumCord := Coordinates{
		a.X + b.X,
		a.Y + b.Y,
	}
	return SumCord
}

func Del(a, b Coordinates) Coordinates {
	DelCord := Coordinates{
		a.X - b.X,
		a.Y - b.Y,
	}
	return DelCord
}
