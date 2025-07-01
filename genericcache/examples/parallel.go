package examples

import (
	"fmt"
	"time"
)

func printEverySecond(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	// Run two goroutines
	go printEverySecond("Hello")
	go printEverySecond("World")

	var input string
	fmt.Scanln(&input)
}
