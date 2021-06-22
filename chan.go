package main

import (
	"fmt"
	"time"
)

func main() {
	a := chanA()
	b := chanB()
	c := make(chan int)

	go func() {
		defer close(c)
		for {
			if a == nil && b == nil {
				return
			}
			select {
			case x, ok := <-a:
				if !ok {
					fmt.Println("close chan A")
					a = nil
					continue
				}
				c <- x
			case x, ok := <-b:
				if !ok {
					fmt.Println("close chan B")
					b = nil
					continue
				}
				c <- x
			}
		}
	}()

	for x := range c {
		fmt.Println(x)
	}
}

func chanA() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
			time.Sleep(time.Second)
		}
		close(c)
	}()
	return c
}

func chanB() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 5; i++ {
			c <- i
			time.Sleep(time.Second)
		}
		c = nil
		fmt.Println("after send to nil channel")
	}()
	return c
}
