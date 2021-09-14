package main

import (
	"fmt"
	"strconv"
)

func main() {

	var Forks = createForks()

	var phil1 = Philosopher{number: 1, left: &Forks[0], right: &Forks[1], outgoing: make(chan string), incoming: make(chan string)}
	var phil2 = Philosopher{number: 2, left: &Forks[1], right: &Forks[2], outgoing: make(chan string), incoming: make(chan string)}
	var phil3 = Philosopher{number: 3, left: &Forks[2], right: &Forks[3], outgoing: make(chan string), incoming: make(chan string)}
	var phil4 = Philosopher{number: 4, left: &Forks[3], right: &Forks[4], outgoing: make(chan string), incoming: make(chan string)}
	var phil5 = Philosopher{number: 5, left: &Forks[4], right: &Forks[0], outgoing: make(chan string), incoming: make(chan string)}

	for {
		select {
		case msg1 := <-phil1.outgoing:
			if msg1 == "Eating" {
				fmt.Println("phil1 eating:" + strconv.Itoa(phil1.timesEaten))
			} else {

			}
		case msg2 := <-phil2.outgoing:
			if msg2 == "Eating" {
				fmt.Println("phil2 eating: " + strconv.Itoa(phil2.timesEaten))
			} else {

			}

		case msg3 := <-phil3.outgoing:
			if msg3 == "Eating" {
				fmt.Println("phil3 eating: " + strconv.Itoa(phil3.timesEaten))
			} else {

			}

		case msg4 := <-phil4.outgoing:
			if msg4 == "Eating" {
				fmt.Println("phil4 eating: " + strconv.Itoa(phil4.timesEaten))
			} else {
				if phil3.timesEaten > phil5.timesEaten {
					phil5.eat()
				} else {
					phil3.eat()
				}
			}

		case msg5 := <-phil5.outgoing:
			if msg5 == "Eating" {
				fmt.Println("phil5 eating: " + strconv.Itoa(phil5.timesEaten))
			} else {
				if phil4.timesEaten > phil1.timesEaten {
					phil1.eat()
				} else {
					phil4.eat()
				}
			}

		}
	}

}

func createForks() []Fork {
	Forks := make([]Fork, 5)
	for i := 0; i < 5; i++ {
		Forks[i] = Fork{timesUsed: 0, outgoing: make(chan string), incoming: make(chan string), inUse: false}
	}
	return Forks
}
