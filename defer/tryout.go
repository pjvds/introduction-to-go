package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	do()
}

func do() {
	msg := "hello"
	err := fmt.Errorf("error1")
	defer printIt(&msg, &err)
	err = fmt.Errorf("error2")

	msg = "world"

	defer log(time.Now())

	time.Sleep(1 * time.Second)
}

func log(started time.Time) {
	fmt.Printf("took: %v", time.Now().Sub(started))
}

func printIt(msg *string, err *error) {
	fmt.Println(*msg)
	fmt.Println(*err)
}
