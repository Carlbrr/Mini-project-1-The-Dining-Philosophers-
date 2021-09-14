package main

import (
	"strconv"
	"time"
)

type Philosopher struct {
	timesEaten int
	left       *Fork
	leftFree   bool
	right      *Fork
	rightFree  bool
	number     int
	outgoing   chan string
	incoming   chan string
}

func (p Philosopher) eat() {
	p.left.Lock()
	p.right.Lock()
	p.outgoing <- "Eating"

	philosophers[p.number-1%5].incoming <- "LeftNotFree"
	philosophers[p.number+1%5].incoming <- "RightNotFree"

	p.outgoing <- "Philosopher " + strconv.Itoa(p.number) + " has eaten " + strconv.Itoa(p.timesEaten)
	time.Sleep(time.Millisecond * 500)
	p.timesEaten++
	p.left.Unlock()
	p.right.Unlock()
	p.outgoing <- "Thinking"

	philosophers[p.number-1%5].incoming <- "LeftFree"
	philosophers[p.number+1%5].incoming <- "RightFree"
}

func (p Philosopher) receiver() {
	for {
		start := <-p.incoming
		if start == "LeftFree" {
			p.leftFree = true
		} else if start == "LeftNotFree" {
			p.leftFree = false
		} else if start == "RightFree" {
			p.rightFree = true
		} else if start == "RightNotFree" {
			p.rightFree = false
		} else if start == "timesEaten" {
			p.outgoing <- strconv.Itoa(p.timesEaten)
		}

		p.eatIfPossible()
	}
}

func (p Philosopher) eatIfPossible() {
	if p.leftFree && p.rightFree {
		p.eat()
	}
}
