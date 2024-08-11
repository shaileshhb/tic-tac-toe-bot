package app

import (
	"fmt"
)

func NewBoard(size uint8) *Board {
	cell := make([]Cell, size*size)

	return &Board{
		Size: size,
		Cell: cell,
	}
}

func (b *Board) MarkCell(cellNumber uint8, sign Mark) bool {
	if cellNumber > (b.Size*b.Size) || b.IsCellOccupied(cellNumber) {
		return false
	}

	b.Cell[cellNumber].Mark = sign
	return true
}

func (b *Board) IsCellOccupied(cellNumber uint8) bool {
	return b.Cell[cellNumber].Mark != EmptyString
}

func (b *Board) CheckWin() bool {
	return b.checkRow() || b.checkColumn() || b.checkDiagonal()
}

func (b *Board) IsBoardFull() bool {
	for i := 0; i < (int(b.Size * b.Size)); i++ {
		if b.Cell[i].Mark == EmptyString {
			return false
		}
	}
	return true
}

func (b *Board) ShowBoard() {
	fmt.Println("Board:")
	for i := 0; i < int(b.Size); i++ {
		for j := 0; j < int(b.Size); j++ {
			cell := &b.Cell[i*int(b.Size)+j]
			if cell.Mark == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(string(cell.Mark))
			}

			if j < int(b.Size)-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < int(b.Size)-1 {
			fmt.Println("------")
		}
	}
}

func (b *Board) checkRow() bool {
	matchFound := false

	for i := 0; i < int(b.Size*b.Size); i += int(b.Size) {
		if b.Cell[i].Mark != EmptyString &&
			b.Cell[i].Mark == b.Cell[i+1].Mark {
			for j := i + 1; j < (i + int(b.Size)); j++ {
				if b.Cell[i].Mark == b.Cell[j].Mark {
					matchFound = true
					continue
				}
				matchFound = false
			}
			if matchFound {
				return true
			}
		}
	}

	return false
}

func (b *Board) checkColumn() bool {
	matchFound := false

	for i := 0; i < int(b.Size); i++ {
		if b.Cell[i].Mark != EmptyString &&
			b.Cell[i].Mark == b.Cell[i+int(b.Size)].Mark {
			for j := i; j < int(b.Size*b.Size); j += int(b.Size) {
				if b.Cell[i].Mark == b.Cell[j].Mark {
					matchFound = true
					continue
				}
				matchFound = false
				break
			}
			if matchFound {
				return true
			}
		}
	}

	return false
}

func (b *Board) checkDiagonal() bool {
	matchFound := false
	startIndex := 0

	// left diagonal
	for i := int(b.Size) + 1; i < int(b.Size*b.Size); i += int(b.Size) + 1 {
		if b.Cell[startIndex].Mark != EmptyString &&
			b.Cell[startIndex].Mark == b.Cell[i].Mark {
			matchFound = true
			continue
		}
		matchFound = false
		break
	}

	if matchFound {
		return true
	}

	// right diagonal
	startIndex = int(b.Size) - 1

	for i := (startIndex * 2); i < int(b.Size*b.Size)-startIndex; i += startIndex {
		if b.Cell[startIndex].Mark != EmptyString &&
			b.Cell[startIndex].Mark == b.Cell[i].Mark {
			matchFound = true
			continue
		}
		matchFound = false
		break
	}

	return matchFound
}
