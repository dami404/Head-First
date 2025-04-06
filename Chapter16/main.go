package main

import (
	"fmt"
)

func addSlice(s []int, e int) {
	s = append(s, e)
}

func addNum(s []int) {
	s = append(s, 4)
}

func task1() {
	s := make([]int, 2, 2)
	fmt.Println(s)

	addSlice(s, 1)
	fmt.Println(s)

	fmt.Println()
}

func task2() {
	s := []int{1, 2, 3}
	fmt.Println(s)

	addNum(s[0:2])
	fmt.Println(s)

	fmt.Println()
}

func task3() {
	x := []int{}         // len cap vals
	x = append(x, 0)     // 1 1 0
	x = append(x, 1)     // 2 2 0,1
	x = append(x, 2)     // 3 4 0,1,2
	y := append(x, 3)    // 4 4 0,1,2,3
	z := append(x, 4)    // 4 4 0,1,2,4
	fmt.Println(x, y, z) // 0,1,2,4 0,1,2,3 0,1,2,4

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
