package main

import (
	"fmt"
	"strconv"
	"time"
)

var philosophers []Philosopher

func main() {

	var Forks = createForks()

	var phil0 = Philosopher{number: 0, left: &Forks[0], right: &Forks[1], outgoing: make(chan string), incoming: make(chan string), leftFree: true, rightFree: true}
	var phil1 = Philosopher{number: 1, left: &Forks[1], right: &Forks[2], outgoing: make(chan string), incoming: make(chan string), leftFree: true, rightFree: true}
	var phil2 = Philosopher{number: 2, left: &Forks[2], right: &Forks[3], outgoing: make(chan string), incoming: make(chan string), leftFree: true, rightFree: true}
	var phil3 = Philosopher{number: 3, left: &Forks[3], right: &Forks[4], outgoing: make(chan string), incoming: make(chan string), leftFree: true, rightFree: true}
	var phil4 = Philosopher{number: 4, left: &Forks[4], right: &Forks[0], outgoing: make(chan string), incoming: make(chan string), leftFree: true, rightFree: true}

	philosophers = []Philosopher{phil0, phil1, phil2, phil3, phil4}
	fmt.Println("Philos init")
	fmt.Println("Routine 1")
	go phil0.receiver()
	time.Sleep(time.Second * 1)
	fmt.Println("Routine 2")
	go phil1.receiver()
	fmt.Println("Routine 3")
	time.Sleep(time.Second * 1)
	go phil2.receiver()
	fmt.Println("Routine 4")
	time.Sleep(time.Second * 1)
	go phil3.receiver()
	fmt.Println("Routine 5")
	time.Sleep(time.Second * 1)
	go phil4.receiver()

	time.Sleep(time.Second * 1)

	fmt.Println("Routines started")
	phil0.incoming <- "start"
	for {
		select {
		case msg1 := <-phil0.outgoing:
			if msg1 == "p0 Eating" {
				fmt.Println("phil1 eating:" + strconv.Itoa(phil0.timesEaten))
				fmt.Println("Fork 1 times used: " + strconv.Itoa(phil0.left.timesUsed))
				fmt.Println("Fork 5 times used: " + strconv.Itoa(phil0.right.timesUsed))
			} else {
				fmt.Println("p0 thinking")

			}
		case msg2 := <-phil1.outgoing:
			if msg2 == "p1 Eating" {
				fmt.Println("phil2 eating: " + strconv.Itoa(phil1.timesEaten))
				fmt.Println("Fork 2 times used: " + strconv.Itoa(phil1.left.timesUsed))
				fmt.Println("Fork 1 times used: " + strconv.Itoa(phil1.right.timesUsed))
			} else {
				fmt.Println("p1 thinking")
			}

		case msg3 := <-phil2.outgoing:
			if msg3 == "p2 Eating" {
				fmt.Println("phil3 eating: " + strconv.Itoa(phil2.timesEaten))
				fmt.Println("Fork 3 times used: " + strconv.Itoa(phil2.left.timesUsed))
				fmt.Println("Fork 2 times used: " + strconv.Itoa(phil2.right.timesUsed))
			} else {
				fmt.Println("p2 thinking")
			}

		case msg4 := <-phil3.outgoing:
			if msg4 == "p3 Eating" {
				fmt.Println("phil4 eating: " + strconv.Itoa(phil3.timesEaten))
				fmt.Println("Fork 4 times used: " + strconv.Itoa(phil3.left.timesUsed))
				fmt.Println("Fork 3 times used: " + strconv.Itoa(phil3.right.timesUsed))
			} else {
				fmt.Println("p3 thinking")
			}

		case msg5 := <-phil4.outgoing:
			if msg5 == "p4 Eating" {
				fmt.Println("phil5 eating: " + strconv.Itoa(phil4.timesEaten))
				fmt.Println("Fork 5 times used: " + strconv.Itoa(phil4.left.timesUsed))
				fmt.Println("Fork 4 times used: " + strconv.Itoa(phil4.right.timesUsed))
			} else {
				fmt.Println("p4 thinking")
			}

		}
	}

}

func createForks() []Fork {
	Forks := make([]Fork, 5)
	for i := 0; i < 5; i++ {
		Forks[i] = Fork{timesUsed: 0, outgoing: make(chan string), incoming: make(chan string), inUse: false}
		go Forks[i].forkReciever()
	}
	return Forks
}
