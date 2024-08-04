package app

import (
	"errors"
)

type Mark string

const (
	X           Mark = "X"
	O           Mark = "O"
	EmptyString Mark = ""
)

func (m Mark) Validate() error {
	switch m {
	case X:
		return nil
	case O:
		return nil
	case EmptyString:
		return nil
	}
	return errors.New("invalid mark specified")
}

type Player struct {
	Name string
	Mark Mark
}

type Board struct {
	Size uint8
	Cell []Cell
}

type Cell struct {
	Mark Mark
}

type GameBot struct {
	board  *Board
	Player *Player
}

type Game struct {
	PlayerOne *Player
	PlayerTwo *Player
	Bot       *GameBot
	Board     *Board
}
