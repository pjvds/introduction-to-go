package main

import (
	"errors"
	"fmt"
)

// START OMIT
func main() {
	secret, err := TellMe("please")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("secret: %v", secret)
}

func TellMe(password string) (secret string, err error) {
	if password != "java sucks" {
		err = errors.New("incorrect password")
		return
	}

	secret = `top 100 blocked users on HPC are man
              top 100 most favorited are woman`
	return
}

// END OMIT
