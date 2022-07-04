package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	Width    int
	Height   int
	Cells    [][]*Cell
	Mines    int
	Turns    int
	GameOver bool
}

func NewGame(board *Board) *Board {
	for column := 0; column < board.Width; column++ {
		board.Cells[column] = make([]*Cell, board.Height)
		for row := 0; row < board.Height; row++ {
			if row == 0 || column == 0 || column == board.Width-1 || row == board.Height-1 {
				board.Cells[column][row] = nil
				continue
			}
			board.Cells[column][row] = &Cell{
				X: row,
				Y: column,
			}
		}
	}
	board.insertMines()
	return board
}

func (b *Board) insertMines() {
	rand.Seed(time.Now().Unix())
	cont := 0
	for i := 0; i < b.Mines; i++ {
		cell := b.Cells[rand.Intn(b.Width-2)+1][rand.Intn(b.Height-2)+1]
		if !cell.hasMine {
			cell.hasMine = true
			hood := b.GetCellsAround(cell.X, cell.Y)
			cont++
			for _, v := range hood {
				if v != nil {
					v.MinesAround++
				}
			}
		} else {
			i--
		}

	}

	fmt.Println(cont)
}

func (b *Board) GetCellsAround(x, y int) (cells []*Cell) {
	hood := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, v := range hood {
		cells = append(cells, b.Cells[y+v[0]][x+v[1]])
	}

	return cells
}

func (b *Board) String() {

	for column := 1; column < b.Width-1; column++ {
		fmt.Println()
		for row := 1; row < b.Height-1; row++ {
			cell := b.Cells[column][row]
			if !cell.IsVisible && !cell.IsMarked {
				fmt.Print("[]", " ")
				continue
			}
			if cell.hasMine {
				fmt.Print(" * ", " ")
				continue
			}
			if cell.IsMarked {
				fmt.Print(" X", " ")
				continue
			}
			fmt.Printf(" %d ", cell.MinesAround)
		}
	}
	fmt.Printf("number of mines: %d, Turns: %d \n", b.Mines, b.Turns)
	fmt.Println()
}

func (b *Board) CheckCell(x, y int) error {
	b.Turns++
	var cell *Cell

	if !b.IsValidCoordinates(x, y) {
		return errors.New("coordinates invalid")
	}

	cell = b.Cells[y][x]
	cell.ShowCell()
	if cell.hasMine {
		b.gameOver()
	}
	cells := b.GetCellsAround(x, y)
	for _, c := range cells {
		if c != nil && !c.IsVisible && !c.hasMine {
			fmt.Println("neibohodd", c)
			c.ShowCell()
			if c.MinesAround > 0 {
				return nil
			}
			b.CheckCell(c.X, c.Y)
		}
		// }
	}

	fmt.Println(cell.X, cell.Y)
	if cell.hasMine {
		b.gameOver()
	}

	return nil
}

func (b *Board) MarkCell(x, y int) error {

	if !b.IsValidCoordinates(x, y) {
		return errors.New("coordinates invalid")
	}

	b.Cells[y][x].MarkCell()
	return nil
}

func (b *Board) IsValidCoordinates(x, y int) bool {

	if x > 0 && x < b.Width-1 && y > 0 && y < b.Height-1 {
		return true
	}
	return false
}

func (b *Board) gameOver() {
	for column := 1; column < b.Width-1; column++ {
		for row := 1; row < b.Height-1; row++ {
			b.Cells[column][row].IsVisible = true
		}
	}
	b.GameOver = true
}

func (b *Board) ShowAll() {

	for column := 1; column < b.Width-1; column++ {
		fmt.Println()
		for row := 1; row < b.Height-1; row++ {
			cell := b.Cells[column][row]
			if cell.hasMine {

				fmt.Print(" * ")
			} else {
				fmt.Print(" ", cell.MinesAround, " ")
			}
		}
	}
}
