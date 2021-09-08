package main

type Philosopher struct {
	timesEaten int
	left       *Fork
	right      *Fork
	number     int
	outgoing   chan string
	incoming   chan string
}
