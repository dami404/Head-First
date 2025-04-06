package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func task1() {
	url := "https://yandex.ru"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(body))

	fmt.Println()
}

func letter(l string) {
	for i := range 20 {
		fmt.Println(i, ":", l)
	}
}

func task2() {
	// go letter("a")
	// go letter("b")
	time.Sleep(2000)
	fmt.Println()
}

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)
	}
}

func task3() {
	// go repeat("x")
	// go repeat("y")
	time.Sleep(time.Second)
	fmt.Println()
}

func task4() {
	var myChannel chan int
	myChannel = make(chan int)
	fmt.Println(myChannel)
	fmt.Println()

}

func abc(c chan int) {
	c <- 1
	fmt.Println("Blocked in abc - 6")

	c <- 3
	fmt.Println("Blocked in abc - 7")

}

func def(c chan int) {
	c <- 2
	fmt.Println("Blocked in def - 5")

	c <- 4
	fmt.Println("Blocked in def - 10")

}
func task5() {
	c1 := make(chan int)
	c2 := make(chan int)

	go abc(c1)
	fmt.Println("Blocked in main - 1")

	go def(c2)
	fmt.Println("Blocked in main - 2")

	fmt.Println(<-c1)
	fmt.Println("Blocked in main - 3")

	fmt.Println(<-c2)
	fmt.Println("Blocked in main - 4")

	fmt.Println(<-c1)
	fmt.Println("Blocked in main - 8")

	fmt.Println(<-c2)
	fmt.Println("Blocked in main - 9")

	fmt.Println()
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}
func task6() {
	fmt.Println()

	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)
}

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}
func even(channel chan int) {
	channel <- 2
	channel <- 4
}
func task7() {
	channelA := make(chan int)
	channelB := make(chan int)
	go odd(channelA)
	go even(channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
}

func responseSize(url string, channel chan Website) {
	// fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Website{url, len(body)}
}

type Website struct {
	url  string
	size int
}

func task8() {
	urls := []string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	size := make(chan Website)

	for _, url := range urls {
		go responseSize(url, size)
		fmt.Println(<-size)

	}

	time.Sleep(5 * time.Second)
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()

}
