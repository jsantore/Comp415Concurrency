package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		go badWrite("output.txt", fmt.Sprintf("the number for this write is %d", 100/i))
		fmt.Println(fmt.Sprintf("the number for this write is %d", 100/i))
	}
	time.Sleep(10 * time.Second) //sleep to make sure go routines run before program ends.
}

func badWrite(fileName string, text string) {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", text)) //first param is number of bytes written
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	file.Close()

}

func betterWrite(fileName string, receiver <-chan string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
}
