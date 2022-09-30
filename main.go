package main

import (
	"fmt"
	"time"
)

func main() {
	go count("park")
	go count("jong")
	time.Sleep(time.Second * 5)
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "count", i)
		time.Sleep(time.Second)
	}
}