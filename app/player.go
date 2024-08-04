package app

func NewPlayer(name string, mark Mark) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}
}
