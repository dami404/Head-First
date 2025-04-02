package main

import (
	"fmt"
	"head-first/geo"
)

func task1() {
	var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)

	fmt.Println()
}

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func task2() {
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(s)

	fmt.Println()
}

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func task3() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)

	fmt.Println()
}

type part struct {
	description string
	count       int
}

func doublePack(p *part) {
	p.count *= 2
}
func task4() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)

	fmt.Println()

}

func task5() {

	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}

type Employee struct {
	Name   string
	Salary float64
}

func task6() {

}
func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
}
