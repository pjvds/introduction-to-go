package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	messages := make(chan string)
	go talk("robert", messages)
	go talk("tomas", messages)

	timeout := time.After(5 * time.Second) // HL
	for {
		select { // HL
		case message := <-messages:
			fmt.Println(message)
		case <-timeout: // HL
			fmt.Println("Boring, I'm leaving")
			return
		}
	}
}

func talk(name string, messages chan string) {
	for {
		messages <- name
		time.Sleep(time.Duration(rand.Intn(2500)) * time.Millisecond)
	}
}

// END OMIT
