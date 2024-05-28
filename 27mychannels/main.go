package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myChannel := make(chan int, 2)

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(2)

	//read
	go func(ch <-chan int, wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		val, isChannelOpen := <-myChannel
		// fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		mut.Unlock()
		wg.Done()
	}(myChannel, wg, mut)

	//write
	go func(ch chan<- int, wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		close(myChannel)
		// myChannel <- 5
		// myChannel <- 6
		mut.Unlock()
		wg.Done()
	}(myChannel, wg, mut)

	wg.Wait()
}
