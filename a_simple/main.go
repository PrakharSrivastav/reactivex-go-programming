package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5, 6)(). // create
		Map(times10). // operate
		Filter(greaterThan30) // operate

	for item := range observable.Observe() { // observe
		fmt.Println("items are ::", item.V)
	}
}

// times10 func multiplies each entry by 10
func times10(ctx context.Context, i interface{}) (interface{}, error) {
	return i.(int) * 10, nil
}

// greaterThan30 func filters values > 30
func greaterThan30(i interface{}) bool {
	return i.(int) > 30
}
