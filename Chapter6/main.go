// Слайсы
package main

import (
	"fmt"
)

func task1() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string{"a", "b", "c"}
	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}

func task2() {
	arr := []int{1, 2}
	slc := arr
	slc[0] = 3
	fmt.Println(arr)
	fmt.Println(slc)

	arr[0] = 4
	fmt.Println(arr)
	fmt.Println(slc)
}

func task3() {
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	s5 := append(s4, "s4", "s4")
	s6 := append(s5, "s4", "s4")
	s7 := append(s6, "s4", "s4")
	s8 := append(s7, "s4", "s4")

	fmt.Println(s1, s2, s3, s4, s5, s6, s7, s8)
	s4[0] = "PP"
	s5[0] = "XX"
	fmt.Println(s1, s2, s3, s4, s5, s6, s7, s8)

}

func task4() {
	var slc []int
	fmt.Printf("%#v", slc)

}

func task5(numbers ...int) {

	fmt.Println(numbers[1])
}
func main() {
	task1()
	fmt.Println()

	task2()
	fmt.Println()

	task3()
	fmt.Println()

	task4()
	fmt.Println()
	var a int = 1
	var b int = 2

	task5(a, b)

	fmt.Println(sum(9, 7))
	fmt.Println(sum(4, 2, 1))

}

func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
