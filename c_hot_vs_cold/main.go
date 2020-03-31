package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	hot()
	cold()
}

func hot() {
	ob := rxgo.Defer([]rxgo.Producer{
		func(_ context.Context, next chan<- rxgo.Item) {
			for i := 0; i < 3; i++ {
				next <- rxgo.Of(i)
			}
		},
	})

	for item := range ob.Observe() {
		fmt.Println("hot observable 1 ::", item.V)
	}

	for item := range ob.Observe() {
		fmt.Println("hot observable 2 ::", item.V)
	}
}

func cold() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	observable := rxgo.FromChannel(ch)

	// First Observer
	for item := range observable.Observe() {
		fmt.Println("cold observable 1 ::", item.V)
	}

	// Second Observer. This will not execute
	for item := range observable.Observe() {
		fmt.Println("cold observable 2 ::", item.V)
	}
}
