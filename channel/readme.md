# Concurrency in Go
*resource: https://go.dev/talks/2012/concurrency.slide*


## sync.WaitGroup
WaitGroup 适用于不关心并发操作结果，或能通过其他方式获取操作结果，等待一组并发操作完成的好方法

## Mutex
不同 Goroutine 并发执行时，无法确保执行顺序，存在对共享资源抢占的情况。需要通过锁的机制来实现并发安全。

Mutual exclusion, 互斥操作提供了一种并发安全的方式来表示对共享资源访问的独占

### sync.Mutex

### sync.RWMutex


**Fan-in using select**

```go
func Fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
    go func() {
        for {
            select {
            case s := <-input1: c <-s
            case s := <-input2: c <-s
            }
        }
    }()
    return c
}
```

**Timeout using select**
```go
func main() {
    c := boring("Joe")
    for {
        select {
        case s := <-c: fmt.Println(s)
        case <-time.After(1 * time.Second):
            fmt.Println("You are so slow.")
            return
        }
    }
}
```


**Daisy-chain**
```go
func f(left, right chan int) {
    left <- 1 + <-right
}

func main() {
    const n = 10000
    leftmost := make(chan int)

    left := leftmost
    right := leftmost

    for i = 0; i < n; i++ {
        right := make(chan int)
        go f(left, right)
        left = right
    }

    go func(c chan int) {
        c <- 1
    }(right)

    fmt.Println(<-leftmost)
}
```