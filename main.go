package main

import (
	"fmt"
	"log"

	"github.com/shaileshhb/tic-tac-toe/app"
)

func main() {
	playerone := getPlayerInfo()
	board := app.NewBoard(3)
	botmark := app.X
	fmt.Println("Your mark:", app.O)
	fmt.Println("Bot has mark:", botmark)
	bot := app.InitializeBot(board, botmark)

	game := app.NewGame(playerone, bot, board)
	game.Play()
}

func getPlayerInfo() *app.Player {
	// fmt.Print("Enter your name: ")
	// name := getUserInput()

	return &app.Player{
		Name: "John doe",
		Mark: app.O,
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
