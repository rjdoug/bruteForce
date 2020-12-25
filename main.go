package main

import (
	"fmt"
	"time"
	"github.com/blavibutcher/keyboard"
	"log"
	"strconv"
)

func main () {
	fmt.Println("Please enter a 4 digit pin:")
	var input string
	for {
		dummyIn, err := keyboard.GetString()
		if err != nil {
			log.Fatal(err)
		}
		input = dummyIn
		
		if len(input) > 4 {
			fmt.Println("Please enter a 4 digit pin:")
		} else { break }
	}
	inputNum, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	pin := inputNum
	var guessedPin int
	startNano := GetCurrentNano()

	for i := 0; i < 10_000; i++ {
		if pin == i {
			guessedPin = i
			break 
		}
	}
	endNano := GetCurrentNano()
	nano := endNano - startNano
	fmt.Printf("It took %d nanoseconds to crack the pin of %04d\n", nano, guessedPin)
	fmt.Println("There are 1,000,000 nano seconds to a millisecond")

}

// GetCurrentNano returns the current time in nanoseconds
func GetCurrentNano() int64 {
	now := time.Now()                                                 
	return now.UnixNano()
}