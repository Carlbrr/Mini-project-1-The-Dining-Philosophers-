package main

import (
	"fmt"
	"strconv"
)

func main() {

	var fork0 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
	var fork1 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
	var fork2 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
	var fork3 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
	var fork4 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}

	var phil0 = Philosopher{left: &fork0, right: &fork1, outgoing: make(chan int)}
	var phil1 = Philosopher{left: &fork1, right: &fork2, outgoing: make(chan int)}
	var phil2 = Philosopher{left: &fork2, right: &fork3, outgoing: make(chan int)}
	var phil3 = Philosopher{left: &fork3, right: &fork4, outgoing: make(chan int)}
	var phil4 = Philosopher{left: &fork0, right: &fork4, outgoing: make(chan int)} //Venstrehåndet

	go phil0.eat()
	go phil1.eat()
	go phil2.eat()
	go phil3.eat()
	go phil4.eat()

	go fork0.forkReciever()
	go fork1.forkReciever()
	go fork2.forkReciever()
	go fork3.forkReciever()
	go fork4.forkReciever()

	for {
		select {
		case msg1 := <-phil0.outgoing:
			if msg1 == -1 {
				fmt.Println("Philosopher 1 is thinking")
			} else {
				fmt.Println("Philosopher 1 is eating " + strconv.Itoa(msg1))
			}
		case msg2 := <-phil1.outgoing:
			if msg2 == -1 {
				fmt.Println("Philosopher 2 is thinking")
			} else {
				fmt.Println("Philosopher 2 is eating " + strconv.Itoa(msg2))
			}

		case msg3 := <-phil2.outgoing:
			if msg3 == -1 {
				fmt.Println("Philosopher 3 is thinking")
			} else {
				fmt.Println("Philosopher 3 is eating " + strconv.Itoa(msg3))
			}

		case msg4 := <-phil3.outgoing:
			if msg4 == -1 {
				fmt.Println("Philosopher 4 is thinking")
			} else {
				fmt.Println("Philosopher 4 is eating " + strconv.Itoa(msg4))
			}

		case msg5 := <-phil4.outgoing:
			if msg5 == -1 {
				fmt.Println("Philosopher 5 is thinking")
			} else {
				fmt.Println("Philosopher 5 is eating " + strconv.Itoa(msg5))
			}
		case msg6 := <-fork0.outgoing:
			fmt.Println("Fork 0 has been used " + strconv.Itoa(msg6) + " times")
		case msg7 := <-fork1.outgoing:
			fmt.Println("Fork 1 has been used " + strconv.Itoa(msg7) + " times")
		case msg8 := <-fork2.outgoing:
			fmt.Println("Fork 2 has been used " + strconv.Itoa(msg8) + " times")
		case msg9 := <-fork3.outgoing:
			fmt.Println("Fork 3 has been used " + strconv.Itoa(msg9) + " times")
		case msg10 := <-fork4.outgoing:
			fmt.Println("Fork 4 has been used " + strconv.Itoa(msg10) + " times")
		}
	}
}
