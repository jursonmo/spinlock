package main

import (
	"fmt"
	"sync"

	"github.com/jursonmo/spinlock"
)

func dosomething() int {
	n := 10000
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func main() {
	wg := sync.WaitGroup{}
	gnum := 100
	loop := 10000
	counter := 0
	l := spinlock.NewSpinLock()
	wg.Add(gnum)
	for i := 0; i < gnum; i++ {
		go func() {
			for j := 0; j < loop; j++ {
				l.Lock()
				dosomething()
				counter++
				l.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if counter != gnum*loop {
		fmt.Println(counter, "spinlock test fail")
		panic("fail")
	}
}
