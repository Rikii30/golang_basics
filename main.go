package main

import (
	"fmt"
	"math/rand"
	"time"
)

// import (
// 	"fmt"
// 	"tutorial/myfunc"
// )

// func main(){
// 	fmt.Println("hello world")
// }

// ---------------------------------------------------------------

func main() {

	var choices = []string{"rock", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())
	userScore := 0
	computerScore := 0

	fmt.Println("welcome to rock, paper, scissors game")

	for {
		userChoice := ""
		fmt.Println("enter a choice: ")
		fmt.Scanln(&userChoice)

		isValidChoice := false
		var choice string
		for _, choice = range choices {
			if userChoice == choice {
				isValidChoice = true
				break
			}
		}
		if !isValidChoice {
			fmt.Println("Invalid Choice")
			continue
		}
		computerChoice := choices[rand.Intn(len(choices))]

		if userChoice == computerChoice {
			fmt.Println("its a tie")
		} else if userChoice == "rock" && computerChoice == "scissors" || userChoice == "paper" && computerChoice == "rock" || userChoice == "scissors" && computerChoice == "paper" {
			fmt.Println("user win!!")
			userScore++
		} else {
			fmt.Println("Computer win!!")
			computerScore++
		}
		fmt.Println("user score is: ", userScore, "computer score is: ", computerScore)

		var ask_user string
		fmt.Println("Do you want to play (yes)")
		fmt.Scanln(&ask_user)

		if ask_user == "yes" {
			continue
		} else {
			fmt.Println("Thanks for playing")
			break
		}
	}

	// var my_test bool = true //without "=true" the output will be false
	// fmt.Println(my_test)
	// ----------------------------------------------------------------
	// var my_number int = 10
	// fmt.Println(my_number)
	// ----------------------------------------------------------------
	// var my_float float64 = 20
	// fmt.Println(my_float)
	// ----------------------------------------------------------------
	// var my_text string = "5"
	// fmt.Println(my_text)
	// ----------------------------------------------------------------
	// var x,y int = 10,15
	// var y int = 15

	// fmt.Println(x)
	// fmt.Println(y)
	// ----------------------------------------------------------------

	// x, y := 10, 50
	// fmt.Println(x)
	// fmt.Println(y)
	// ----------------------------------------------------------------
	// var (
	// 	x = 10
	// 	y = 90.8
	// 	z = "riki"
	// )
	// fmt.Println(x, y, z)

	// fmt.Println(y)
	// fmt.Println(z)
	// ----------------------------------------------------------------
	// Constants
	// ----------
	// const num = 8
	// fmt.Println(num)
	// ----------------------------------------------------------------
	// ARRAYS
	// fixed size cannot be changed during run time
	// -------
	// var arr []int
	// fmt.Println(arr)
	// ----------------------------------------------------------------
	// var arr = [3]int{1, 2, 3}
	// fmt.Println(arr)
	// ----------------------------------------------------------------
	// var arr = [3]int{1, 2, 3, 4}
	// fmt.Println(arr)
	// ----------------------------------------------------------------
	// var arr = [3]int{1, 2, 3}
	// fmt.Println(arr[0])
	// ----------------------------------------------------------------
	// Slices
	// dynamically sized
	// -------
	// var my_slice []int
	// var second_slice = []int{1, 2, 3}
	// my_slice = append(my_slice, second_slice...)
	// my_slice = append(my_slice, 4)
	// fmt.Println(my_slice)
	// ----------------------------------------------------------------
	// output := hello()
	// fmt.Println(output)
	// ----------------------------------------------------------------
	// CREATE A FUNCTION
	// a := 10
	// b := 20
	// result := myfunc.Sum(a, b)
	// fmt.Println(result)
	// ----------------------------------------------------------------
	// LOOPS
	// for,while,range
	// -----------------
	// FOR
	// ----
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }
	// WHILE
	// ------
	// i := 0
	// for i < 10 {
	// 	fmt.Println("while loop: ", i)
	// 	i++
	// }
	// RANGE
	// ------
	// var my_array = [5]int{1, 2, 3, 4, 5}
	// index := 0
	// value := 0
	// for index, value = range my_array {
	// 	fmt.Println("index: ", index)
	// 	fmt.Println("value: ", value)
	// }
	// ----------------------------------------------------------------
	// i := 7
	// j := 4
	// if i == j {
	// 	fmt.Println("i and j are equal")
	// } else if j == 6 {
	// 	fmt.Println("They are not equal")
	// } else {
	// 	fmt.Println("all false")
	// }
	// ----------------------------------------------------------------
	// SWITCH STATEMENT
	// var my_array = [3]int{10, 20, 30}
	// switch my_array[2] {
	// case 10:
	// 	fmt.Println("num: 10")
	// case 20:
	// 	fmt.Println("num: 20")
	// case 30:
	// 	fmt.Println("num: 30")
	// default:
	// 	fmt.Println("none of them are true")
	// }

}

// func hello() string {
// 	return "hello"
// }
