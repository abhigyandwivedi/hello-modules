package main

import (
	"fmt"
	"time"
)

func randEmit(ch chan int) {
	id := 0

	for {
		ch <- id
		id++
		time.Sleep(500 * time.Millisecond)
	}
}

func emit(ch chan string) {
	words := []string{"the", "quick", "brown", "fox"}

	for _, word := range words {
		ch <- word
		time.Sleep(20 * time.Millisecond)

	}
	close(ch)
}

func unitemit(ch chan string) {
	ch <- "acme"
	close(ch)
}

func main() {
	/*	fmt.Println("Hello World")
		ch := make(chan string)
		ch2 := make(chan string)

		go emit(ch)
		go unitemit(ch2)

		for word := range ch {
			fmt.Print(" ", word)
		}
		close(ch)

		fmt.Println(<-ch2)
		str, ok := <-ch2
		fmt.Println(str, ok)
	*/
	ch3 := make(chan int)
	go randEmit(ch3)

	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

}
