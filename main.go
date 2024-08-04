package main

import (
	"fmt"

	"github.com/shaileshhb/tic-tac-toe/app"
	"github.com/shaileshhb/tic-tac-toe/bot"
)

func main() {
	// playerone := app.NewPlayer("playerone", app.X)
	board := app.NewBoard(3)
	gamebotOne := bot.InitializeBot(board, app.O)
	gamebotTwo := bot.InitializeBot(board, app.X)

	game := app.NewGame(gamebotOne.Player, gamebotTwo.Player, board)
	gamebotOne.MakeMove()
	gamebotTwo.MakeMove()
	gamebotOne.MakeMove()
	gamebotTwo.MakeMove()
	gamebotOne.MakeMove()
	gamebotTwo.MakeMove()

	game.Board.ShowBoard()
	isWin := game.Board.CheckWin()
	fmt.Println(isWin)
}
