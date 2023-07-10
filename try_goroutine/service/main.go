package service

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func RoutineDemo() {
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

func ChannelDemo() {
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

func SelectDemo() {
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

func numConsumer(numChan chan int, quit chan bool, no int) {
	fmt.Printf("consumer %v in place\n", no)
	for {
		select {
		case num := <-numChan:
			fmt.Printf("consumer %v get %v\n", no, num)
			time.Sleep(1 * time.Second)
		case <-quit:
			fmt.Printf("consumer %v quit\n", no)
			return
		}
	}
}

func ChannelSize() {
	maxChanLen := 10
	c := make(chan int, maxChanLen)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	go func() {
		for {
			randNum := r.Intn(1000)
			c <- randNum
			fmt.Printf("produce %v into c, now len: %v\n", randNum, len(c))
			time.Sleep(time.Duration(randNum) * time.Millisecond)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	consumerQuit := make(chan bool)
	consumerNum := 0
	time.Sleep(1 * time.Second)
	for {
		select {
		case <-time.Tick(1 * time.Second):
			fmt.Printf("now c len: %v, now consumers: %v\n", len(c), consumerNum)
			fmt.Println(strings.Repeat("*", 40))
			fmt.Println(strings.Repeat("n", len(c)))
			fmt.Println(strings.Repeat("c", consumerNum))
			fmt.Println(strings.Repeat("*", 40))
			if len(c) > maxChanLen/2 {
				consumerNum += 1
				fmt.Printf("start consumer %v\n", consumerNum)
				go numConsumer(c, consumerQuit, consumerNum)
			} else if len(c) <= maxChanLen/2 && consumerNum > 1 {
				fmt.Println("stop a consumer")
				consumerQuit <- true
				consumerNum -= 1
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func GenDemo() {
	gen := func(ctx context.Context) chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("gen inner goroutine receive Done")
					time.Sleep(200 * time.Millisecond)
					//return // returning not to leak the goroutine
				case dst <- n:
					n++
				default:
					fmt.Println("waiting to produce")
					time.Sleep(200 * time.Millisecond)
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		time.Sleep(1 * time.Second)
		if n == 3 {
			break
		}
	}
	cancel()
	time.Sleep(1 * time.Second)
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
	fmt.Println("start gen")
	out := make(chan int)
	go func() {
		defer func() { fmt.Println("gen stop"); close(out) }()
		for _, n := range nums {
			time.Sleep(time.Second)
			fmt.Println("gen")
			select {
			case out <- n:
			case <-done:
				fmt.Println("receive done")
				return
			}
		}
	}()
	return out
}

func DoneDemo() {
	done := make(chan struct{})
	c := gen(done, 1, 2, 3)
	fmt.Println(<-c)
	close(done)
	time.Sleep(2 * time.Second)
}
