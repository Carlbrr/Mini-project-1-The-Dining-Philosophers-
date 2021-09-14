package main

import "time"

type Philosopher struct {
	timesEaten int
	left       *Fork
	right      *Fork
	number     int
	outgoing   chan string
	incoming   chan string
}

func (p Philosopher) eat() {
	p.left.Lock()
	p.right.Lock()
	p.outgoing <- "Eating"
	time.Sleep(time.Millisecond * 500)
	p.timesEaten++
	p.left.Unlock()
	p.right.Unlock()
	p.outgoing <- "Thinking"
}
