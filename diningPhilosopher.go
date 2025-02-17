package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

/*
the dining philosopher’s problem with the following constraints/modifications.
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

const NUM_PHILO = 5
const NUM_CHOPS = 5
const MAX_EAT_COUNT = 3
const MAX_CONCURRENT_EATERS = 2

var (
	wg sync.WaitGroup
)

type ChopS struct {
	mu sync.Mutex
}

func (c *ChopS) Lock() {
	c.mu.Lock()
}

func (c *ChopS) Unlock() {
	c.mu.Unlock()
}

type Host struct {
	philosopherChannels []chan bool
	permissionCount     int
}

func (h *Host) listenPermissionRequest() {
	for i := 0; i < NUM_PHILO*MAX_EAT_COUNT; i++ {
		isPermissionGranted := false
		for !isPermissionGranted {
			select {
			case msg := <-h.philosopherChannels[0]:
				isPermissionGranted = h.handlePermission(0, msg)
			case msg := <-h.philosopherChannels[1]:
				isPermissionGranted = h.handlePermission(1, msg)
			case msg := <-h.philosopherChannels[2]:
				isPermissionGranted = h.handlePermission(2, msg)
			case msg := <-h.philosopherChannels[3]:
				isPermissionGranted = h.handlePermission(3, msg)
			case msg := <-h.philosopherChannels[4]:
				isPermissionGranted = h.handlePermission(4, msg)
			}
		}
	}

	for h.permissionCount != 0 {
		select {
		case msg := <-h.philosopherChannels[0]:
			h.handlePermission(0, msg)
		case msg := <-h.philosopherChannels[1]:
			h.handlePermission(1, msg)
		case msg := <-h.philosopherChannels[2]:
			h.handlePermission(2, msg)
		case msg := <-h.philosopherChannels[3]:
			h.handlePermission(3, msg)
		case msg := <-h.philosopherChannels[4]:
			h.handlePermission(4, msg)
		}
	}

	wg.Done()
}

func (h *Host) handlePermission(philoIndex int, request bool) bool {
	if request {
		if h.permissionCount < MAX_CONCURRENT_EATERS {
			h.permissionCount++
			h.philosopherChannels[philoIndex] <- true // Approve eating
			return true
		} else {
			h.philosopherChannels[philoIndex] <- false // Deny eating
			return false
		}
	} else {
		h.permissionCount-- // Philosopher has finished eating
		return false
	}
}

type Philo struct {
	name            string
	leftCS, rightCS *ChopS
	eatCount        int
	hostChannel     chan bool
}

func (p *Philo) eat() {
	for p.eatCount < MAX_EAT_COUNT {

		p.hostChannel <- true

		if <-p.hostChannel {
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Printf("starting to eat %s\n", p.name)
			fmt.Printf("finishing eating %s\n", p.name)
			p.eatCount++

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			p.hostChannel <- false
		}
	}
	wg.Done()
}

func main() {

	CSticks := make([]*ChopS, NUM_CHOPS)
	philos := make([]*Philo, NUM_PHILO)
	channels := make([]chan bool, NUM_PHILO)

	for i := 0; i < NUM_PHILO; i++ {
		channels[i] = make(chan bool)
	}

	var host *Host

	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("Creating Chop Sticks")
	for i := 0; i < NUM_CHOPS; i++ {
		CSticks[i] = new(ChopS)
	}

	fmt.Println("Creating Host")
	host = &Host{channels, 0}
	fmt.Println("Running Host")

	wg.Add(1)
	go host.listenPermissionRequest()

	fmt.Println("Creating Philosophers")
	for i := 0; i < NUM_PHILO; i++ {
		philos[i] = &Philo{strconv.Itoa(i + 1), CSticks[i], CSticks[(i+1)%5], 0, channels[i]}
	}

	fmt.Println("Running Philosophers")

	fmt.Println(strings.Repeat("-", 20))

	for i := 0; i < NUM_PHILO; i++ {
		wg.Add(1)
		go philos[i].eat()
	}

	wg.Wait()
}
