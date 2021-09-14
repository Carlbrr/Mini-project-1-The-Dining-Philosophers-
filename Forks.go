package main

import "sync"

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
	f.outgoing <- "Lifted"
}

func (f Fork) Unlock() {
	f.inUse = false
	f.outgoing <- "Put down"
	f.locker.Lock()
}
