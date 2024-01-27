package channel

import (
	"fmt"
	"sync"
	"time"
)

// RunConcurrencyWithoutLock
// * 协程执行顺序：
// *  1. 从寄存器读取 a 的值；
// *  2. 然后做加法运算；
// *  3. 最后写到寄存器中
func RunConcurrencyWithoutLock() {
	var a = 0
	for i := 0; i < 10; i++ {
		go func() {
			a += 1
			fmt.Println(a)
		}()
	}
	time.Sleep(time.Second)
}

func RunConcurrencyWithLock() {
	var a = 0
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			mutex.Lock()
			a += 1
			fmt.Println(a)
			mutex.Unlock()

			defer wait.Done()
		}()
	}
	time.Sleep(time.Second)
	wait.Wait()
}
