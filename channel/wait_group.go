package channel

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 适用于不关心并发操作结果，等待一组并发操作完成的好方法
func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Goroutine one")
		time.Sleep(time.Second)
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine two")
		time.Sleep(2 * time.Second)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All goroutine done.")
}

func WaitGroup2() {
	var count int = 5
	var wg sync.WaitGroup

	dotingInGoroutine := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Do some thing")
		time.Sleep(time.Second)
	}

	wg.Add(count)
	for i := 0; i < count; i++ {
		go dotingInGoroutine(&wg)
	}
	wg.Wait()
	fmt.Println(count, " goroutines done")
}
