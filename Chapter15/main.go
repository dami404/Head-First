package main

import (
	"fmt"
)

func task1() {

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

type my struct {
	a int
}

func my_func(m *map[string]my) {
	for _, v := range *m {
		v.a = 9
	}
}

func task5() {
	var m map[string]my = map[string]my{
		"foo": my{a: 1},
		"bar": my{a: 2},
	}

	// for _, v := range m {
	// 	v.a = 9
	// }
	my_func(&m)

	fmt.Println(m)
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
}
