package main

import (
	"fmt"
)

type Population int

func task1() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congratulations, Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println()
}

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}
func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func task2() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	fmt.Println()
}

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}
func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func task3() {
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", l, l.ToMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters\n", ml, ml.ToLiters())
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
