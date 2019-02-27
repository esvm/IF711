package dining_philosophers

import (
	"fmt"
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
	for i := 0; i < 5; i++ {
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

			time.Sleep(time.Millisecond * 100 * time.Duration(p.id+1))
		}

		print(p.id, "is eating")
		time.Sleep(time.Millisecond * time.Duration(p.id+1))

		p.LeftFork.Unlock()
		p.RightFork.Unlock()
		p.RightFork.Locked = false

		print(p.id, "finished eating")
		time.Sleep(time.Millisecond * time.Duration(p.id+1))
	}

	waitGroup.Done()
}

// EatDeadlock ... makes the philosopher eat with deadlock. If forks are locked, the philosopher wait
func (p *Philosopher) EatDeadlock() {
	for i := 0; i < 5; i++ {
		print(p.id, "is thinking")
		p.LeftFork.Lock()
		p.RightFork.Lock()

		print(p.id, "is eating")
		time.Sleep(time.Millisecond * time.Duration(p.id+1))

		p.LeftFork.Unlock()
		p.RightFork.Unlock()

		print(p.id, "finished eating")
		time.Sleep(time.Millisecond * time.Duration(p.id+1))
	}

	waitGroup.Done()
}

func print(id int, doing string) {
	fmt.Printf("The philosopher %d %s\n", id, doing)
}

var waitGroup sync.WaitGroup

// Dinner ... start the meet
func Dinner() {
	count := 5
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

		waitGroup.Add(1)
		go philosophers[i].Eat()
	}

	waitGroup.Wait()
}
