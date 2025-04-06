package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func task1() {
	file, err := os.Open("data.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var sum float64

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	fmt.Println(sum)
}

func scanDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if !file.IsDir() {
			fmt.Println(filePath)
		} else {
			scanDir(filePath)
		}
	}

}

/*
opening
closing
ref is empty
.snack
*/

func task2() {
	scanDir("../")
	fmt.Println()
}

func getRecover() {
	recover()
}

func throwPanic() {
	defer getRecover()
	panic("Lets panic!")
}
func task3() {
	throwPanic()
	fmt.Println()
}

func task5() {

	fmt.Println()
}
func main() {
	task1()
	task2()
	task3()
	// task4()
	task5()
}
