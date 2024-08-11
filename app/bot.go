package app

import (
	"math"
)

func InitializeBot(board *Board, mark Mark) *GameBot {
	return &GameBot{
		board:  board,
		Player: NewPlayer("Bot", mark),
	}
}

func (bot *GameBot) MakeMove() bool {
	if bot.board.IsBoardFull() {
		return false
	}

	cellNumber := bot.getBestMove()
	return bot.board.MarkCell(cellNumber, bot.Player.Mark)
}

func (bot *GameBot) getBestMove() uint8 {
	bestScore := math.Inf(-1)
	bestMove := -1
	board := *bot.board

	for i := 0; i < (int(board.Size * board.Size)); i++ {
		if board.Cell[i].Mark == EmptyString {
			board.Cell[i].Mark = bot.Player.Mark
			score := bot.minimax(&board, false, math.Inf(-1), math.Inf(1))
			board.Cell[i].Mark = EmptyString
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}

	return uint8(bestMove)
}

// X: 1 -> maximizing player
// O: -1 -> minimizing player
// Tie: 0
const (
	MaximizeX float64 = 1
	MinimizeO float64 = -1
	Tie       float64 = 0
)

func (bot *GameBot) minimax(board *Board, isMaximizing bool, alpha, beta float64) float64 {
	result := board.CheckWin()
	if result {
		if isMaximizing {
			score := MinimizeO * (float64(bot.getNumberOfSquares(board) + 1))
			return score
		}

		score := MaximizeX * (float64(bot.getNumberOfSquares(board) + 1))
		return score
	}

	if board.IsBoardFull() {
		return Tie
	}

	if isMaximizing {
		bestScore := math.Inf(-1)

		for i := 0; i < (int(board.Size * board.Size)); i++ {
			if board.Cell[i].Mark == EmptyString {
				board.Cell[i].Mark = bot.Player.Mark
				score := bot.minimax(board, false, alpha, beta)
				board.Cell[i].Mark = EmptyString
				bestScore = math.Max(bestScore, score)
				alpha = math.Max(alpha, score)
				if beta <= alpha {
					break
				}
			}
		}
		return bestScore
	}

	bestScore := math.Inf(1)

	humanPlayerMark := X
	if bot.Player.Mark == X {
		humanPlayerMark = O
	}

	for i := 0; i < (int(board.Size * board.Size)); i++ {
		if board.Cell[i].Mark == EmptyString {
			board.Cell[i].Mark = humanPlayerMark
			score := bot.minimax(board, true, alpha, beta)
			board.Cell[i].Mark = EmptyString
			bestScore = math.Min(bestScore, score)
			beta = math.Min(beta, score)
			if beta <= alpha {
				break
			}
		}
	}

	return bestScore
}

func (bot *GameBot) getNumberOfSquares(board *Board) int {
	counter := 0

	for _, cell := range board.Cell {
		if cell.Mark == EmptyString {
			counter++
		}
	}

	return counter
}
