package main

import(
	"fmt"
	"time"
)

func main() {
	ts := time.Now()
	fmt.Printf("%s",time.Now().Sub(ts))
}