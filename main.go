package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

func main() {
	maxScore := 30
	totalTurns := 30
	currentTurns := 0

	playerScore := 0
	botScore := 0

	for currentTurns < totalTurns {
		fmt.Println("\nTurn number: " + strconv.Itoa(currentTurns+1))

		userInput := myInput("Wanna roll the dice? (yes/no): ")

		if strings.Contains("yYyesYesYES", userInput) == true {
			playerDiceRoll := rand.IntN(6)
			botDiceRoll := rand.IntN(6)

			fmt.Printf("Player Rolled: %d || Bot Rolled: %d\n", playerDiceRoll, botDiceRoll)

			playerScore = playerScore + playerDiceRoll
			botScore = botScore + botDiceRoll

			fmt.Printf("Your Score: %d || Bot Score: %d\n", playerScore, botScore)

			if playerScore >= maxScore || botScore >= maxScore {
				if playerScore > botScore {
					fmt.Println("You won!")
				} else if botScore > playerScore {
					fmt.Println("Hahaha I won!")
				} else {
					fmt.Println("It's a draw :(")
				}
			}

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
