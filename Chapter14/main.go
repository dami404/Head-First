package main

import (
	"fmt"
	"head-first/prose"
)

func task1() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	fmt.Println()
}

func task2() {

	fmt.Println()
}

func task3() {

	fmt.Println()
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
