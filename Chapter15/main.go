package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(resp http.ResponseWriter, req *http.Request) {
	message := []byte("Hellow wrld!")
	_, err := resp.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func task1() {
	// http.HandleFunc("/hello", viewHandler)
	// err := http.ListenAndServe("localhost:8080", nil)
	// log.Fatal(err)
	fmt.Println()
}

func callFunction(passedFunction func()) {
	passedFunction()
}
func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}
func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("This sentence is", false)
}
func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}
func functionA() {
	fmt.Println("function called")
}
func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}
func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}
func task2() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}

// func task2() {

// 	fmt.Println()
// }

func task3() {

	fmt.Println()
}

func task4() {

	fmt.Println()

}

type my struct {
	a int
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
