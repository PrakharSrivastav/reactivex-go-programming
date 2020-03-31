package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

var (
	e1 = errors.New("error 1")
	e2 = errors.New("error 2")
)

func main() {
	observable := rxgo.Just(1, 2, e1, 3, 4, 5, e2, 6)().
		Map(times10).
		Filter(greaterThan30)

	// observable.ForEach() api is useful if you are used to callbacks.
	// ForEach() is non-blocking, but returns a channel which closes
	// once the observable is complete.
	done := observable.ForEach(onNext, onError, onComplete, rxgo.WithErrorStrategy(rxgo.ContinueOnError))
	<-done
}

func onNext(i interface{}) {
	fmt.Println("items ::", i)
}

func onError(err error) {
	fmt.Println("error :: ", err)
}

func onComplete() {
	fmt.Println("complete")
}

func times10(ctx context.Context, i interface{}) (interface{}, error) {
	return i.(int) * 10, nil
}

func greaterThan30(i interface{}) bool {
	return i.(int) > 30
}
