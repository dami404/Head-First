// Циклы, чтение файлов
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func task1() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

func task2() {

}

func task3() {

}

func isGuessed(randomInt, number int) bool {
	if randomInt > number {
		fmt.Println("Oops. Your guess was LOW")
		return false
	} else if randomInt < number {
		fmt.Println("Oops. Your guess was HIGH")
		return false
	} else {
		fmt.Println("Thats right!")
		return true
	}
}

func main() {
	randomInt := rand.Intn(100)
	fmt.Println(randomInt)
	fmt.Println("Enter a random number from 0 to 100: ")
	reader := bufio.NewReader(os.Stdin)

	for i := 5; i >= 0; i-- {

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		userNumber, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if isGuessed(randomInt, userNumber) {
			break
		} else if i == 0 {
			fmt.Println("Sorry. You didnt guess my number. It was:", randomInt)
			break
		} else {
			fmt.Println("Attempts left:", i)
		}
	}

}
