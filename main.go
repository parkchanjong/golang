package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [5]string{"park", "jong", "lee"}
	for _, person := range people {
		go isPerson(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isPerson(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}