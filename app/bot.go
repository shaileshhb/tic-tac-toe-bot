package app

import (
	"math/rand"
)

func InitializeBot(board *Board, mark Mark) *GameBot {
	return &GameBot{
		board:  board,
		Player: NewPlayer("3TBot", mark),
	}
}

func (bot *GameBot) MakeMove() bool {

	if bot.board.IsBoardFull() {
		return false
	}

	for {
		cellNumber := bot.findBestMove()
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
