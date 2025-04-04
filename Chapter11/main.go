package main

import (
	"fmt"
)

type Address struct {
	city   string
	street string
	house  int
}

type Initials struct {
	name    string
	surname string
}

type Info interface {
	GetInfo()
}

func (a Address) GetInfo() {
	fmt.Println("city: ", a.city)
	fmt.Println("street: ", a.street)
	fmt.Println("house: ", a.house)
}

func (i *Initials) GetInfo() {
	fmt.Println("name: ", i.name)
	fmt.Println("surname: ", i.surname)
}

func task1() {

	addr := Info(
		&Address{"SPB", "Pushkina", 2},
	)

	inits := Info(
		&Initials{"Dima", "Pushkin"},
	)

	var addr2 Info
	addr2 = Address{"SPB", "Spasskaya", 32}

	addr.GetInfo()
	fmt.Println()

	inits.GetInfo()
	fmt.Println()

	addr2.GetInfo()
	fmt.Println()

	fmt.Println()
}

type TypePlayer string

type TypePlayer2 string

type Player interface {
	Play(song string)
}

func (t *TypePlayer) Play(song string) {
	fmt.Println("TypePlayer is playing", song)
}

func (t *TypePlayer2) Play(song string) {
	fmt.Println("TypePlayer2 is playing", song)
}

func playList(device Player, song string) {
	device.Play(song)
}

func task2() {
	tp := TypePlayer2("1")
	var player Player = &tp
	playList(player, "234")
	fmt.Println()
}

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

func task3() {
	var vehicle Vehicle = Car("Toyoda Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
	fmt.Println()
}

type Truck1 string

func (t Truck1) Accelerate1() {
	fmt.Println("Speeding up")
}
func (t Truck1) Brake1() {
	fmt.Println("Stopping")
}
func (t Truck1) Steer1(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck1) LoadCargo1(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle1 interface {
	Accelerate1()
	Brake1()
	Steer1(string)
}

func TryVehicle(vehicle Vehicle1) {
	vehicle.Accelerate1()
	vehicle.Steer1("left")
	vehicle.Steer1("right")
	vehicle.Brake1()
	truck, ok := vehicle.(Truck1)
	if ok {
		truck.LoadCargo1("test cargo")
	}
}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func task4() {
	TryVehicle(Truck1("Fnord F180"))

	fmt.Println()

}

type AnyData interface{}

func CPrint(a AnyData) AnyData {
	num, ok := a.(int)
	if ok {
		return num * 100
	}
	return a
}

func task5() {

	fmt.Println(CPrint(1))
	fmt.Println(CPrint(1.3))
	fmt.Println(CPrint("2adfb"))

}
func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
}
