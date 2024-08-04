package app

type Mark string

const (
	X           Mark = "X"
	O           Mark = "O"
	EmptyString Mark = ""
)

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

type Game struct {
	PlayerOne *Player
	PlayerTwo *Player
	Board     *Board
}
