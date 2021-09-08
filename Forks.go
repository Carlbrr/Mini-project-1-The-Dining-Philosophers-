package main

type Fork struct {
	timesUsed int
	inUse     bool
	outgoing  chan string
	incoming  chan string
}
