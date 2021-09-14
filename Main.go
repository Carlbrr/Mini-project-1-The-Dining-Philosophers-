package main

import (
	"fmt"
	"strconv"
)

var philosophers []Philosopher

func main() {

	var Forks = createForks()

	var phil0 = Philosopher{number: 0, left: &Forks[0], right: &Forks[1], outgoing: make(chan string), incoming: make(chan string)}
	var phil1 = Philosopher{number: 1, left: &Forks[1], right: &Forks[2], outgoing: make(chan string), incoming: make(chan string)}
	var phil2 = Philosopher{number: 2, left: &Forks[2], right: &Forks[3], outgoing: make(chan string), incoming: make(chan string)}
	var phil3 = Philosopher{number: 3, left: &Forks[3], right: &Forks[4], outgoing: make(chan string), incoming: make(chan string)}
	var phil4 = Philosopher{number: 4, left: &Forks[4], right: &Forks[0], outgoing: make(chan string), incoming: make(chan string)}

	philosophers = []Philosopher{phil0, phil1, phil2, phil3, phil4}
	fmt.Println("Philos init")
	go phil0.receiver()
	fmt.Println("Routine 1")
	go phil1.receiver()
	fmt.Println("Routine 2")
	go phil2.receiver()
	fmt.Println("Routine 3")
	go phil3.receiver()
	fmt.Println("Routine 4")
	go phil4.receiver()
	fmt.Println("Routine 5")

	fmt.Println("Routines started")
	for {
		select {
		case msg1 := <-phil0.outgoing:
			if msg1 == "Eating" {
				fmt.Println("phil1 eating:" + strconv.Itoa(phil0.timesEaten))
				fmt.Println("Fork 1 times used: " + strconv.Itoa(phil0.left.timesUsed))
				fmt.Println("Fork 5 times used: " + strconv.Itoa(phil0.right.timesUsed))
			} else {

			}
		case msg2 := <-phil1.outgoing:
			if msg2 == "Eating" {
				fmt.Println("phil2 eating: " + strconv.Itoa(phil1.timesEaten))
				fmt.Println("Fork 2 times used: " + strconv.Itoa(phil1.left.timesUsed))
				fmt.Println("Fork 1 times used: " + strconv.Itoa(phil1.right.timesUsed))
			} else {

			}

		case msg3 := <-phil2.outgoing:
			if msg3 == "Eating" {
				fmt.Println("phil3 eating: " + strconv.Itoa(phil2.timesEaten))
				fmt.Println("Fork 3 times used: " + strconv.Itoa(phil2.left.timesUsed))
				fmt.Println("Fork 2 times used: " + strconv.Itoa(phil2.right.timesUsed))
			} else {

			}

		case msg4 := <-phil3.outgoing:
			if msg4 == "Eating" {
				fmt.Println("phil4 eating: " + strconv.Itoa(phil3.timesEaten))
				fmt.Println("Fork 4 times used: " + strconv.Itoa(phil3.left.timesUsed))
				fmt.Println("Fork 3 times used: " + strconv.Itoa(phil3.right.timesUsed))
			} else {

			}

		case msg5 := <-phil4.outgoing:
			if msg5 == "Eating" {
				fmt.Println("phil5 eating: " + strconv.Itoa(phil4.timesEaten))
				fmt.Println("Fork 5 times used: " + strconv.Itoa(phil4.left.timesUsed))
				fmt.Println("Fork 4 times used: " + strconv.Itoa(phil4.right.timesUsed))
			} else {

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
