package main

import "fmt"

func main() {

	var Forks = createForks()

	var phil1 = Philosopher{number: 1, left: &Forks[0], right: &Forks[1], outgoing: make(chan string), incoming: make(chan string)}
	var phil2 = Philosopher{number: 2, left: &Forks[1], right: &Forks[2], outgoing: make(chan string), incoming: make(chan string)}
	var phil3 = Philosopher{number: 3, left: &Forks[2], right: &Forks[3], outgoing: make(chan string), incoming: make(chan string)}
	var phil4 = Philosopher{number: 4, left: &Forks[3], right: &Forks[4], outgoing: make(chan string), incoming: make(chan string)}
	var phil5 = Philosopher{number: 5, left: &Forks[4], right: &Forks[0], outgoing: make(chan string), incoming: make(chan string)}

	go phil1.receiver()
	go phil2.receiver()
	go phil3.receiver()
	go phil4.receiver()
	go phil5.receiver()

	for {
		select {
		case msg1 := <-phil1.outgoing:
			fmt.Println(msg1)
		case msg2 := <-phil2.outgoing:
			fmt.Println(msg2)
		case msg3 := <-phil3.outgoing:
			fmt.Println(msg3)
		case msg4 := <-phil4.outgoing:
			fmt.Println(msg4)
		case msg5 := <-phil5.outgoing:
			fmt.Println(msg5)
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
