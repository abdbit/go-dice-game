package main

import (
	"fmt"
	"math/rand/v2"
)

func helper() {
	fmt.Println("Hello World")

	// myInt := 1 // implitcit declaration, must use colon
	var myInt int = 1 // explicit declaration, no need for colon, value optional
	fmt.Println(myInt)

	// myString := "abc"
	var myString string = "abc"
	fmt.Println(myString)

	// myBoolean := true
	var myBoolean bool = true
	fmt.Println(myBoolean)

	// myIntArray := []int{1, 2, 3, 4, 5}
	var myIntArray []int = []int{1, 2, 3}
	fmt.Println(myIntArray)

	// myIntElement := myIntArray[0]
	var myIntElement int = myIntArray[0]
	fmt.Println(myIntElement)

	myStringArray := []string{"a", "b", "c"}
	// var myStringArray []string = []string{"a", "b", "c"}
	fmt.Println(myStringArray)

	var myStringElement string = myStringArray[0]
	fmt.Println(myStringElement)

	myIntOne := 1
	myIntTwo := 2

	if myIntOne < myIntTwo {
		fmt.Println("One is greater than two.")
	} else {
		fmt.Println("Two is greater than one.")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println(i)
	}

	for false {
		fmt.Println("Hello World")
	}

	mySum := mySumFunction(1, 2)
	fmt.Println(mySum)

	myAns := myMultiplyFunc(5, 4)
	fmt.Println(myAns)

	userInput := input("Input your name: ")
	fmt.Println("Your name is: " + userInput)

	fmt.Printf("Random number: %d", rand.IntN(6))
}

func mySumFunction(numberOne int, numberTwo int) int {
	fmt.Println("inside my beautiful function")

	sum := numberOne + numberTwo

	return sum
}

func myMultiplyFunc(no1 int, no2 int) int {
	ans := no1 * no2

	return ans
}

func input(prompt string) string {
	fmt.Printf("%s", prompt)

	var userInput string
	fmt.Scan(&userInput)

	return userInput
}

// var declaration ✅
// basic types ✅
// printing ✅
// if else ✅
// for ✅
// while ✅
// func ✅
// user input ✅
