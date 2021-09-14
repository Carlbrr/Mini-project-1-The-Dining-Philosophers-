package main

import (
	"strconv"
	"sync"
)

type Fork struct {
	timesUsed int
	inUse     bool
	outgoing  chan string
	incoming  chan string
	locker    sync.Mutex
}

func (f Fork) Lock() {
	f.locker.Lock()
	f.inUse = true
	f.timesUsed++
	//f.outgoing <- "Lifted"
}

func (f Fork) Unlock() {
	f.inUse = false
	//f.outgoing <- "Put down"
	f.locker.Lock()
}

func (f Fork) forkReciever() {
	for {
		call := <-f.incoming
		if call == "timesUsed" {

			f.outgoing <- strconv.Itoa(f.timesUsed)

		} else if call == "inUse" {

			f.outgoing <- strconv.FormatBool(f.inUse)

		} else if call == "lock" {

			f.Lock()

		} else if call == "unlock" {

			f.Unlock()

		}

	}
}
