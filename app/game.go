package app

import (
	"bufio"
	"fmt"
	"os"
)

func NewGame(player *Player, bot *GameBot, board *Board) *Game {
	return &Game{
		Player: player,
		Bot:    bot,
		Board:  board,
	}
}

func (g *Game) Play() {
	for {
		for {
			botMove := g.Bot.MakeMove()
			if botMove {
				break
			}
		}

		status := g.getBoardStatus(g.Bot.Player)
		if status != InProcess {
			return
		}

		var cellNumber uint8
		for {
			cellNumber = getCellNumber()
			if !g.Board.IsCellOccupied(cellNumber) {
				break
			}
			fmt.Println("This cell is occupied. Please choose another cell")
		}
		g.Board.MarkCell(cellNumber, g.Player.Mark)

		status = g.getBoardStatus(g.Bot.Player)
		if status != InProcess {
			return
		}
	}
}

func getCellNumber() uint8 {
	for {
		var i int
		fmt.Print("Enter cell number: (number must be between 1 to 9): ")
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil {
			fmt.Println("invalid input. Please enter a number.")
			// Clear the input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if i < 1 || i > 9 {
			fmt.Println("number must be between 1 and 9.")
			continue
		}

		return uint8(i - 1)
	}
}

func (g *Game) getBoardStatus(player *Player) Result {
	g.Board.ShowBoard()
	if g.Board.CheckWin() {
		fmt.Println("====================================")
		fmt.Printf("%s has won this round.\n", player.Name)
		fmt.Println("====================================")
		fmt.Println("Press 'Enter' to exit...")

		// Wait for the user to press 'Enter'
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return Win
	}

	if g.Board.IsBoardFull() {
		fmt.Println("====================================")
		fmt.Println("It's a tie!")
		fmt.Println("====================================")
		fmt.Println("Press 'Enter' to exit...")

		// Wait for the user to press 'Enter'
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return Draw
	}

	return InProcess
}
