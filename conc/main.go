package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func sum(start int, end int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	c <- sum
}

func main1() {

	N := 10
	channelsNum := 2
	n := N / channelsNum

	wg := &sync.WaitGroup{}
	c := make(chan int)

	for range channelsNum {
		start := N - n + 1
		end := N
		N -= n
		wg.Add(1)
		go sum(start, end, c, wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	totalSum := 0
	for partialSum := range c {
		totalSum += partialSum
	}
	fmt.Print(totalSum)

}

func generateNums(ch chan int) {
	fmt.Println("-3")

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	// close(ch)
}

func doubleNums(ch1, ch2 chan int) {
	fmt.Println("-2")

	for num := range ch1 {
		ch2 <- num * 2
	}
	// close(ch2)
}

func main2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)

	go generateNums(ch1)
	go doubleNums(ch1, ch2)
	fmt.Println("-1")
	for n := range ch2 {
		fmt.Println(n)
	}
}

func main3() {
	counter := 20
	wg := sync.WaitGroup{}

	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			fmt.Println(ii * ii)
		}(i)
	}

	wg.Wait()
}

func main4() {

	writes := 1000
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i

		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}

func main5() {

	writes := 1000

	storage := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		// i := i

		go func(ii int) {
			defer wg.Done()

			// storage[i] = i
			storage.Store(ii, ii)
		}(i)
	}

	wg.Wait()
	// fmt.Println(storage)

	storage.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true // if false, Range stops
	})
}

func main6() {
	storage := make(map[int]int, 1000)

	wg := sync.WaitGroup{}
	reads := 1000
	writes := 1000
	mu := sync.RWMutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}
	wg.Add(reads)
	for i := 0; i < reads; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.RLock()
			defer mu.RUnlock()
			_, _ = storage[i]
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}

func main7() {
	alreadyStored := make(map[int]struct{})
	mu := sync.RWMutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}
	// 3, 4, 5, 0, 4, 9, 8, 6, 6, 5, 5, 4, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.RLock()
			_, ok := alreadyStored[doubles[i]]
			mu.RUnlock()

			if !ok {
				mu.Lock()
				alreadyStored[doubles[i]] = struct{}{}
				mu.Unlock()

				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait()
	close(uniqueIDs)

	for val := range uniqueIDs {
		fmt.Println(val)
	}
	// go func() {
	// 	close(uniqueIDs)
	// }()

	fmt.Printf("len of ids: %d\n", len(uniqueIDs)) // 0, 1, 2, 3, 4 ...
	fmt.Println(uniqueIDs)
}

func main8() {
	ch := make(chan int)

	go func() {
		<-ch
	}()

	ch <- 1

	fmt.Println(runtime.NumGoroutine())

}

// what happens here?
func main9() {
	fmt.Println(runtime.NumGoroutine())
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("iter:", i)
		}
		fmt.Println(runtime.NumGoroutine())

	}()
	select {}
	fmt.Println("finish")
}

// add timeout to avoid long waiting
func main10() {
	// rand.Seed(time.Now().Unix())

	chanForResp := make(chan resp)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

	go RPCCall(ctx, chanForResp)

	result := <-chanForResp

	fmt.Println("id:", result.id)
	fmt.Println("error:", result.err)

}

type resp struct {
	id  int
	err error
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("timeout expired"),
		}
	// sleep 0-4 sec
	case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		ch <- resp{
			id: rand.Int(),
		}
	}

}

// avoid deadlock
func main11() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("default")
	}
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	// ch3 := syncMerge(ch1, ch2) // bad sync way of merging
	ch3 := merge[int](ch1, ch2)
	for val := range ch3 {
		fmt.Println(val)
	}
	close(ch3)
}

func merge[T any](chans ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{}
	for _, singleCh := range chans {
		singleCh := singleCh
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range singleCh {
				result <- val
			}
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func syncMerge[T any](chans ...chan T) chan T {
	l := 0
	for _, singleCh := range chans {
		l += len(singleCh) // 3
	}

	result := make(chan T, l)
	for _, singleCh := range chans {
		for val := range singleCh {
			result <- val
		}
	}
	close(result)

	return result
}
