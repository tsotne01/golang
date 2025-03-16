// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"sync"
)

func myFn(ch chan int) {

	for i := 100; i < 200; i++ {
		ch <- i
	}
	close(ch)
}
func myFn2(ch chan int) {
	for i := 0; i < 200; i++ {
		ch <- i
	}
	close(ch)
}
func listenCh(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("--channel 1--")
		fmt.Println(val)
	}
}
func listenCh2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("--channel 2--")
		fmt.Println(val)
	}
}

func main() {

	ch, ch2 := make(chan int, 20), make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go myFn(ch)
	go myFn2(ch2)
	go listenCh(ch, &wg)
	go listenCh2(ch2, &wg)
	wg.Wait()
}
