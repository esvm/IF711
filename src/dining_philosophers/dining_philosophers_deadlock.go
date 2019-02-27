package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id        int
	RightFork *Fork
	LeftFork  *Fork
}

// Eat ... makes the philosopher eat with deadlocks. If forks are locked, the philosopher wait
func (p *Philosopher) Eat() {
	for {
		print(p.id, "is thinking")
		print(p.id, fmt.Sprintf("is trying to get fork #%d", p.id))

		p.LeftFork.Lock()

		print(p.id, fmt.Sprintf("get #%d", p.id))
		print(p.id, fmt.Sprintf("is trying to get fork #%d", (p.id+1)%count))

		p.RightFork.Lock()

		print(p.id, fmt.Sprintf("get #%d", (p.id+1)%count))
		print(p.id, "is eating")

		sleep()

		p.LeftFork.Unlock()
		p.RightFork.Unlock()

		print(p.id, "finished eating")

		sleep()
	}
}

func print(id int, doing string) {
	fmt.Printf("The philosopher %d %s\n", id, doing)
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}

// Dinner ... start the meet
func Dinner() {
	forks := make([]*Fork, count)
	philosophers := make([]*Philosopher, count)

	for i := 0; i < count; i++ {
		forks[i] = &Fork{}
	}

	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id:        i,
			RightFork: forks[(i+1)%count],
			LeftFork:  forks[i],
		}

		go philosophers[i].Eat()
	}

	fmt.Scanln()
}

var count int

func main() {
	count, _ = strconv.Atoi(os.Args[1])
	Dinner()
}
