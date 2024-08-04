package bot

import (
	"fmt"
	"math/rand"

	"github.com/shaileshhb/tic-tac-toe/app"
)

type GameBot struct {
	board  *app.Board
	Player *app.Player
}

func InitializeBot(board *app.Board, mark app.Mark) *GameBot {
	return &GameBot{
		board:  board,
		Player: app.NewPlayer("3TBot", mark),
	}
}

func (bot *GameBot) MakeMove() bool {

	if bot.board.IsBoardFull() {
		return false
	}

	for {
		cellNumber := bot.findBestMove()
		fmt.Println(cellNumber)
		if bot.board.MarkCell(cellNumber, bot.Player.Mark) {
			return true
		}
	}
}

func (bot *GameBot) findBestMove() uint8 {
	// Seed the random number generator to ensure different results each run
	randomNumber := rand.Intn(int(bot.board.Size * bot.board.Size))
	return uint8(randomNumber)
}
