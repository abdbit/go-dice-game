package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxScore := 30
	totalTurns := 13
	currentTurns := 0

	playerScore := 0
	botScore := 0

	fmt.Println("You have only got '13' turns and if you run out of turns the person with the highest score will win, Goodluck!")

	for currentTurns < totalTurns {
		fmt.Println("\nTurn number: " + strconv.Itoa(currentTurns+1))

		userInput := myInput("Wanna roll the dice? (y/n): ")

		if strings.Contains("yYyesYesYES", userInput) == true {
			playerDiceRoll := rand.IntN(7)
			playerScore = playerScore + playerDiceRoll
			fmt.Printf("Player Rolled: %d || Player Score: %d\n", playerDiceRoll, playerScore)
			checkWinner(playerScore, botScore, maxScore)

			fmt.Println("Bot turn...")
			time.Sleep(3 * time.Second)

			botDiceRoll := rand.IntN(6)
			botScore = botScore + botDiceRoll
			fmt.Printf("Bot Rolled: %d ||     Bot Score: %d\n", botDiceRoll, botScore)
			checkWinner(playerScore, botScore, maxScore)

			currentTurns = currentTurns + 1

		} else {
			fmt.Println("Please enter only 'y' or 'Y' or 'yes' or 'Yes' or 'YES'")
		}
	}

	if playerScore > botScore {
		fmt.Println("You won!")
	} else if botScore > playerScore {
		fmt.Println("Hahaha I won!")
	} else {
		fmt.Println("It's a draw :(")
	}
}

func myInput(prompt string) string {
	fmt.Printf("%s", prompt)

	var userInput string
	fmt.Scan(&userInput)

	return userInput
}

func checkWinner(playerOneScore int, playerTwoScore int, maxScore int) {
	if playerOneScore >= maxScore || playerTwoScore >= maxScore {
		if playerOneScore > playerTwoScore {
			fmt.Println("player won!")
			time.Sleep(5 * time.Second)
			os.Exit(0)
		} else if playerTwoScore > playerOneScore {
			fmt.Println("Bot won!")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		} else {
			fmt.Println("It's a draw :(")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	}
}
