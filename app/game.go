package app

func NewGame(playerOne, playerTwo *Player, board *Board) *Game {
	return &Game{
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
		Board:     board,
	}
}
