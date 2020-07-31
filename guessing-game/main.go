package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min
	// fmt.Println("The secret number is", secretNumber) // test for debugging

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Please input your guess\n(Enter Q anytime to quit)")

	attempts := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if strings.ToUpper(input) == "Q" {
			fmt.Println("You've made", attempts, "attempt(s) at guessing the secret number. Are you sure you want to quit? Y/N")
			yesNo, err1 := reader.ReadString('\n')
			if err1 != nil {
				fmt.Println("An error occured while reading input. Please try again")
				break
			}
			yesNo = strings.TrimSuffix(yesNo, "\n")

			if strings.ToUpper(yesNo) == "Y" {
				fmt.Println("OK, play again soon!")
				return
			} else {
				fmt.Println("OK, please guess an integer value between 1 and 100")
				continue
			}
		} else if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		attempts++
		fmt.Println("Your guess is", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("your guess is smaller than the secret number. Try again")
		} else {
			fmt.Println("Correct, you Legend! You guessed right after", attempts, "attempts")
			break
		}
	}
}
