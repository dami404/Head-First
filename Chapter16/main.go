package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func viewHandler(resp http.ResponseWriter, req *http.Request) {
	// message := []byte("signature list goes here")
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(resp, nil)
	check(err)
}

func task1() {
	// http.HandleFunc("/guestbook", viewHandler)
	// err := http.ListenAndServe("localhost:8080", nil)
	// log.Fatal(err)
	fmt.Println()
}

func printIfOdd(m int) {
	if m%2 == 0 {
		fmt.Println(m)
	}
}

func task2() {
	numbers := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	for _, n := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printIfOdd(n)

		}()
		wg.Wait()
	}
	fmt.Println()
}

func task3() {
	s := "12345"
	runeS := []rune(s)[3:]
	fmt.Println(string(runeS))
}

func a(c chan int) {
	c <- 1
	c <- 2
}

func b(c chan int) {
	c <- 3
	c <- 4
}

func task4() {
	c1 := make(chan int, 23)
	c2 := make(chan int, 2)
	go a(c1)
	go b(c2)

	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c2)
	fmt.Println(<-c2)

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
