package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jp-ferrei/minefield/game"
)

func main() {

	difficulty := readLine("choose the difficulty \n")
	nDifficulty, err := strconv.ParseInt(difficulty, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	var board *game.Board

	switch nDifficulty {
	case 1:
		fmt.Println("Beginner 1, 9x9")
		board = game.NewGame(game.BEGINNER.Get())
	case 2:
		fmt.Println("Intermediary 2, 12x12")
		board = game.NewGame(game.INTERMEDIARY.Get())
	case 3:
		fmt.Println("hard 3, 24x24")
		board = game.NewGame(game.HARD.Get())
	default:
		fmt.Println("wrong number")
	}
	board.String()
	for !board.GameOver {
		getAttempt(board)
		board.String()
	}

	if board.GameOver {
		fmt.Printf("you played %d turns ", board.Turns)
	}
}

func readLine(phrase string, a ...any) string {
	fmt.Printf(phrase, a...)

	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	text := reader.Text()
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func getAttempt(board *game.Board) {
	newAttempt := readLine("insira as coordenadas x , y separadas por virgula do campo desejado de %d a %d \n \n", 1, board.Width-2)
	nums := strings.Split(newAttempt, ",")
	coords := []int{}
	for _, v := range nums {
		nV, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		coords = append(coords, int(nV))
	}
	x, y := coords[0], coords[1]
	if len(coords) > 2 {
		board.MarkCell(x, y)
		fmt.Println("mark")
		return
	}
	board.CheckCell(x, y)
}
