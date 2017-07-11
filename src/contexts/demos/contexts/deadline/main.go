package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Value int
	Err   error
}

func main() {
	// launch 10k goroutines, trying to get random
	// to return "150"
	deadline, cf := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cf()
	result := Random(deadline, 150, 10000)
	if result.Err != nil {
		fmt.Println("error:", result.Err)
	}
	fmt.Println(result.Value)

}

func Random(ctx context.Context, target, count int) Result {
	rand.Seed(82)
	c := make(chan Result, count)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	search := func(target int) {
		sleep := time.Duration(rand.Intn(10))
		select {
		case <-ctx.Done():
			return
		case <-time.After(sleep):
			r := rand.Intn(1000)
			if r == target {
				c <- Result{Value: r}
				return
			}
			fmt.Printf(".")
		}
	}
	for i := 0; i < count; i++ {
		go search(target)
	}
	select {
	case <-ctx.Done():
		return Result{Err: ctx.Err()}
	case r := <-c:
		return r
	}

}
