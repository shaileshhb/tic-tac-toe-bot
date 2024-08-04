package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shaileshhb/tic-tac-toe/app"
)

func main() {
	playerone := getPlayerInfo()
	board := app.NewBoard(3)
	botmark := app.X
	if playerone.Mark == app.X {
		botmark = app.O
	}
	fmt.Println("Great!")
	fmt.Println("Bot has mark:", botmark)
	bot := app.InitializeBot(board, botmark)

	game := app.NewGame(playerone, bot, board)
	game.Play()
}

func getPlayerInfo() *app.Player {
	fmt.Print("Enter your name: ")
	name := getUserInput()

	fmt.Print("Select your marker. (X or O/x or o): ")
	markStr := getUserInput()

	mark := app.Mark(strings.ToUpper(markStr))
	err := mark.Validate()
	if err != nil {
		log.Fatal("Invalid marker. Please enter 'X' or 'O'.")
	}

	return &app.Player{
		Name: name,
		Mark: mark,
	}
}

func getUserInput() string {
	var userInput string
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		log.Fatal(err)
	}

	return userInput
}
