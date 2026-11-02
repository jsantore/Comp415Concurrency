package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		go fmt.Println("i =", 100/i)
	}
	time.Sleep(time.Second * 1)
}
