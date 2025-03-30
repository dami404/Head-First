// Мапы
package main

import (
	"bufio"
	"fmt"
	"head-first/cerrors"
	"log"
	"os"
)

func getWinner(cV map[string]int) (winner string, err error) {
	maxScore := 0
	for name, score := range cV {
		if score > maxScore {
			maxScore = score
			winner = name
		} else if score == maxScore {
			return "", cerrors.ErrorDraw
		}
	}
	return winner, nil
}

func task1() {
	file, err := os.Open("./votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	candidatesVotes := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		candidatesVotes[scanner.Text()]++
	}
	winner, err := getWinner(candidatesVotes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Winner is", winner)
}

func task2() {

	fmt.Println()
}

func task3() {
	err := cerrors.SQLError.TableNotFoundError("1test1")
	log.Fatal(err)
}

func task4() {

	fmt.Println()

}

func task5() {

	fmt.Println()
}
func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
}
