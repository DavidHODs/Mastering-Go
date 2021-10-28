package chapternine

import (
	"fmt"
	"time"
)

func writeChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func Channel() {
	c := make(chan int)

	go writeChannel(c, 10)
	time.Sleep(1 * time.Second)
}

func writeChan(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func Chan() {
	c := make(chan int)
	go writeChan(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read", <- c)
	time.Sleep(1 * time.Second)

	_, ok := <- c
	if ok {
		fmt.Println("Channel is opened")
	} else {
		fmt.Println("Channel is closed")
	}
}