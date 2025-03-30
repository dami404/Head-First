// База
package main

import (
	"fmt"
	"reflect"
)

func task1() {
	fmt.Println("Cannonball!!!")
	fmt.Println(reflect.TypeOf(1.0))
}

func task2() {
	var originalCount int = 10
	var eatenCount int = 4
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}

func task3() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", int(tax), "dollars.")
	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", int(total), "dollars.")
	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", int(total) <= availableFunds)
}

func main() {
	// task1()
	// task2()
	task3()
}
