package game

type Difficulty int

const (
	BEGINNER Difficulty = iota + 1
	INTERMEDIARY
	HARD
)

func (d Difficulty) Get() *Board {

	boards := []*Board{
		{
			Width:  11,
			Height: 11,
			Mines:  12,
			Cells:  make([][]*Cell, 11),
		},
		{
			Width:  18,
			Height: 18,
			Mines:  40,
			Cells:  make([][]*Cell, 18),
		},
		{
			Width:  26,
			Height: 26,
			Mines:  99,
			Cells:  make([][]*Cell, 26),
		},
	}

	return boards[d-1]
}
