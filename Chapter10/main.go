package main

import (
	"fmt"
	"head-first/geo"
	"log"
)

func task1() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
	fmt.Println()
}

type Date struct {
	Year string
	Day  string
}

func (d *Date) SetYear(y string) {
	d.Year = y
}

func (d *Date) SetDay(dy string) {
	d.Day = dy
}

func New(y, dy string) *Date {
	return &Date{y, dy}
}

func task2() {
	day := "14 february"
	year := "2003"
	// d := (&Date{}).New(year, day)
	d := New(year, day)
	// d.New(year, day)
	fmt.Println(*d)
}

func task3() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(28)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
	fmt.Println()
}

func task4() {
	str := "12 45"
	fmt.Println(len(str))

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
