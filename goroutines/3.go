package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	go talk("robert") // HL
	go talk("tomas")  // HL

	time.Sleep(10 * time.Second)
}

func talk(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v: %v\n", name, i)
		time.Sleep(time.Duration(rand.Intn(2500)) * time.Millisecond)
	}
}

// END OMIT
