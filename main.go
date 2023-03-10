package main

import (
	"fmt"
	"sync"
)

func main() {
	balance := 100
	var myWaitGroup sync.WaitGroup
	var myMutext sync.Mutex

	myWaitGroup.Add(2)
	go deposit(&balance, 10, &myMutext, &myWaitGroup)
	withdraw(&balance, 50, &myMutext, &myWaitGroup)

	myWaitGroup.Wait()
	fmt.Println(balance)
}
func deposit(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance += amount
	myMutex.Unlock()
	myWaitGroup.Done()
}
func withdraw(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance += amount
	myMutex.Unlock()
	myWaitGroup.Done()
}
