package main

import (
	"fmt"
	"strconv"
	"time"
)

var fork0 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
var fork1 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
var fork2 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
var fork3 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}
var fork4 = Fork{timesUsed: 0, outgoing: make(chan int), incoming: make(chan int)}

var phil0 = Philosopher{left: &fork0, right: &fork1, incoming: make(chan int), outgoing: make(chan int)}
var phil1 = Philosopher{left: &fork1, right: &fork2, incoming: make(chan int), outgoing: make(chan int)}
var phil2 = Philosopher{left: &fork2, right: &fork3, incoming: make(chan int), outgoing: make(chan int)}
var phil3 = Philosopher{left: &fork3, right: &fork4, incoming: make(chan int), outgoing: make(chan int)}
var phil4 = Philosopher{left: &fork0, right: &fork4, incoming: make(chan int), outgoing: make(chan int)} //Venstrehåndet

func main() {

	go phil0.eat()
	go phil1.eat()
	go phil2.eat()
	go phil3.eat()
	go phil4.eat()

	go phil0.philReceiver()
	go phil1.philReceiver()
	go phil2.philReceiver()
	go phil3.philReceiver()
	go phil4.philReceiver()

	go fork0.forkReciever()
	go fork1.forkReciever()
	go fork2.forkReciever()
	go fork3.forkReciever()
	go fork4.forkReciever()

	go sendRequests()

	for {
		select {
		case msg1 := <-phil0.outgoing:
			if msg1 == -1 {
				fmt.Println("1 - Philosopher 1 is thinking")
			} else {
				phil0.incoming <- 2 //2 is like asking; how many times have you eaten?
				timesEaten := <-phil0.outgoing
				fmt.Println("1 - Philosopher 1 is eating " + strconv.Itoa(timesEaten))

			}
		case msg2 := <-phil1.outgoing:
			if msg2 == -1 {
				fmt.Println("2 - Philosopher 2 is thinking")
			} else {
				phil1.incoming <- 2
				timesEaten := <-phil1.outgoing
				fmt.Println("2 - Philosopher 2 is eating " + strconv.Itoa(timesEaten))

			}

		case msg3 := <-phil2.outgoing:
			if msg3 == -1 {
				fmt.Println("3 - Philosopher 3 is thinking")
			} else {
				phil2.incoming <- 2
				timesEaten := <-phil2.outgoing
				fmt.Println("3 - Philosopher 3 is eating " + strconv.Itoa(timesEaten))
			}

		case msg4 := <-phil3.outgoing:
			if msg4 == -1 {
				fmt.Println("4 - Philosopher 4 is thinking")
			} else {
				phil3.incoming <- 2
				timesEaten := <-phil3.outgoing
				fmt.Println("4 - Philosopher 4 is eating " + strconv.Itoa(timesEaten))
			}

		case msg5 := <-phil4.outgoing:
			if msg5 == -1 {
				fmt.Println("5 - Philosopher 5 is thinking")
			} else {
				phil4.incoming <- 2
				timesEaten := <-phil4.outgoing
				fmt.Println("5 - Philosopher 5 is eating " + strconv.Itoa(timesEaten))
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

func sendRequests() {
	for {
		fmt.Println("-------------------------------------------")
		phil0.incoming <- 1
		phil1.incoming <- 1
		phil2.incoming <- 1
		phil3.incoming <- 1
		phil4.incoming <- 1
		fork0.incoming <- 1
		fork1.incoming <- 1
		fork2.incoming <- 1
		fork3.incoming <- 1
		fork4.incoming <- 1

		time.Sleep(2500 * time.Millisecond)
	}
}
