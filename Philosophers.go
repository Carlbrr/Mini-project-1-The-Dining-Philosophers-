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
	p.left.locker.Lock()
	p.right.locker.Lock()
	p.outgoing <- "Eating"
	time.Sleep(time.Millisecond * 500)
	p.timesEaten++
	p.left.locker.Unlock()
	p.right.locker.Unlock()
	p.outgoing <- "Thinking"
}
