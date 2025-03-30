// Пакеты
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"head-first/Chapter3"
)

const k = 1

func task1() {

	fmt.Println()
}

func task2() {

	fmt.Println()
}

func task3() {
	fmt.Print("enter:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("entered:", input)
}

func main() {
	Chapter3.Main()
	task1()
	task2()

	task3()
}
