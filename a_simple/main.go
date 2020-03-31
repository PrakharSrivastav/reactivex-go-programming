package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

var e1 = errors.New("error 1")
var e2 = errors.New("error 2")

func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5, e1, 6, e2)(). // create
		Map(times10). // operate
		Filter(greaterThan30) // operate

	// Observe
	for item := range observable.Observe() { // this will stop on first error
		//for item := range observable.Observe(rxgo.WithErrorStrategy(rxgo.ContinueOnError)) { // this will continue all errors
		switch {
		case item.Error():
			fmt.Println("error :: ", item.E)
		default:
			fmt.Println("items ::", item.V)
		}
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
