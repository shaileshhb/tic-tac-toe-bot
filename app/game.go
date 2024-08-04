package app

import (
	"fmt"
	"log"
)

func NewGame(playerOne *Player, bot *GameBot, board *Board) *Game {
	return &Game{
		PlayerOne: playerOne,
		Bot:       bot,
		Board:     board,
	}
}

func (g *Game) Play() {
	for {
		cellNumber := getCellNumber()
		g.Board.MarkCell(cellNumber, g.PlayerOne.Mark)

		g.Board.ShowBoard()
		if g.Board.CheckWin() {
			fmt.Println("====================================")
			fmt.Println("Hurray!! You have won!")
			fmt.Println("====================================")
			return
		}

		if g.Board.IsBoardFull() {
			fmt.Println("====================================")
			fmt.Println("It's a tie!")
			fmt.Println("====================================")
			return
		}

		botMove := g.Bot.MakeMove()
		if !botMove {
			log.Fatal("some error occurred while bot was playing.")
		}

		g.Board.ShowBoard()
		if g.Board.CheckWin() {
			fmt.Println("====================================")
			fmt.Println("Better luck next time! Bot has won this round.")
			fmt.Println("====================================")
			return
		}

		if g.Board.IsBoardFull() {
			fmt.Println("====================================")
			fmt.Println("It's a tie!")
			fmt.Println("====================================")
			return
		}
	}
}

func getCellNumber() uint8 {
	var i int
	fmt.Print("Enter cell number: (number must be between 1 to 9): ")
	_, err := fmt.Scanf("%d\n", &i)
	if err != nil {
		log.Fatal("invalid input. Please enter a number.")
	}

	if i < 1 || i > 9 {
		log.Fatal("number must be between 1 and 9.")
	}

	return uint8(i - 1)
}
