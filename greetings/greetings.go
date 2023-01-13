package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	
	var message2 string
	message2 = fmt.Sprintf("Hi, %v. Welcome!", name)
	if message2.l
	
	return message, nil
}