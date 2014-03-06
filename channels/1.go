package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	messages := make(chan string) // HL
	go talk("robert", messages)
	go talk("tomas", messages)

	for message := range messages { // HL
		fmt.Println(message)
	}
}

func talk(name string, messages chan string) {
	for {
		messages <- name // HL
		time.Sleep(time.Duration(rand.Intn(2500)) * time.Millisecond)
	}
}

// END OMIT
