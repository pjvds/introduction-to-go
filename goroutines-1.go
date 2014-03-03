package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	count("robert")
}

func count(name string) {
	for i := 1; ; i++ {
		fmt.Printf("%v: %v\n", name, i)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}

// END OMIT
