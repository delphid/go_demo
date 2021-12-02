package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
    "syscall"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    //routineDemo()
    //time.Sleep(1 * time.Second)
    //channelDemo()
    //time.Sleep(1 * time.Second)
    selectDemo()
}

func routineDemo() {
    f("direct1")

    go f("goroutine1")
    go f("goroutine2")
    go f("goroutine3")

    f("direct2")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}

func channelDemo() {
    messages := make(chan string)
    go func() {
        fmt.Println("start goroutine")
        time.Sleep(2 * time.Second)
        fmt.Println("goroutine slept 2")
        messages <- "ping"
        time.Sleep(2 * time.Second)
        fmt.Println("goroutine slept 4")
    }()
    msg := <-messages
    time.Sleep(1 * time.Second)
    fmt.Println(msg)
}

func selectDemo() {
    c := make(chan int)
	done := make(chan int)
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	go func() {
		for i := 0; i < 10; i++ {
            time.Sleep(1 * time.Second)
			fmt.Println(<-c)
		}
		done <- 0
	}()
	fibonacci(c, done, quit)
}

func fibonacci(c chan int, done chan int, quit chan os.Signal) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("done")
			return
        case <-quit:
            fmt.Println("prepare to quit")
            time.Sleep(2 * time.Second)
            return
            //os.Exit(0)
        default:
            fmt.Println("f blocking")
            time.Sleep(200 * time.Millisecond)
		}
	}
}
