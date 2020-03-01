package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return

		//Using default selector instead of case c<-x causes deadlock
		//This is because default sends the message even though
		//the receiver is not ready to receive it.
		//Case selector c<-x ensures that the message is added only when
		//receiver is ready.

		default:
			fmt.Println("Sending:", x)
			c <- x
			x, y = y, x+y

			/*	case c <- x:
				//default:
				fmt.Println("Sending:", x)
				x, y = y, x+y */
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			//fmt.Println("test",i)
			//time.Sleep(time.Millisecond * 500)
		}
		fmt.Println("Sending quit")
		quit <- 0
	}()
	fibonacci(c, quit)
}
