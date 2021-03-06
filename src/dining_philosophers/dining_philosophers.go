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
	Locked bool
}

type Philosopher struct {
	id        int
	RightFork *Fork
	LeftFork  *Fork
}

// Eat ... makes the philosopher eat with no deadlocks. If forks are locked, the philosopher wait
func (p *Philosopher) Eat() {
	for {
		print(p.id, "is thinking")
		for {
			p.LeftFork.Lock()
			if p.RightFork.Locked {
				p.LeftFork.Unlock()
			} else {
				p.RightFork.Lock()
				p.RightFork.Locked = true
				break
			}

			sleep(p.id, 100)
		}

		print(p.id, "is eating")
		sleep(p.id, 100)

		print(p.id, "finished eating")
		p.LeftFork.Unlock()
		p.RightFork.Unlock()
		p.RightFork.Locked = false

		sleep(p.id, 100)
	}
}

func print(id int, doing string) {
	fmt.Printf("The philosopher %d %s\n", id, doing)
}

func sleep(id int, amount int) {
	time.Sleep(time.Millisecond * time.Duration(amount*(id+1)))
}

// Dinner ... start the meet
func Dinner() {
	forks := make([]*Fork, count)
	philosophers := make([]*Philosopher, count)

	for i := 0; i < count; i++ {
		forks[i] = &Fork{
			Locked: false,
		}
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
