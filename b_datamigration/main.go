package main

import (
	"context"
	"fmt"
	"github.com/PrakharSrivastav/reactivex-go-programming/b_datamigration/sink"
	"github.com/PrakharSrivastav/reactivex-go-programming/b_datamigration/source"
	"github.com/PrakharSrivastav/reactivex-go-programming/b_datamigration/transform"
	"github.com/reactivex/rxgo/v2"
	"log"
	"strings"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// initialize source and sink
	src := createSource(ctx)
	dst := createSink(ctx)

	// create an observable
	ob := rxgo.
		FromChannel(src.Ch, rxgo.WithPublishStrategy()).
		Filter(filterNamesWithA).
		Map(mapToCustomerWithAddress, rxgo.WithPool(32)).
		Map(mapToUpperCase).
		BufferWithTimeOrCount(rxgo.WithDuration(time.Millisecond*500), 5)

	// connect to observer
	ob.Connect()
	src.Read()


	//<-ob.ForEach(onNextFunc(dst), onError, onComplete)
	for items := range ob.Observe() {
		for _, item := range items.V.([]interface{}) {
			dst.Write(item.(sink.CustomerWithAddress))
		}
	}
}

func onComplete() {
	fmt.Println("done")
}

func onError(err error) {
	fmt.Println(err)
}

func onNextFunc(dst sink.Sink) func(items interface{}) {
	return func(items interface{}) {
		for _, item := range items.([]interface{}) {
			dst.Write(item.(sink.CustomerWithAddress))
		}
	}
}

func filterNamesWithA(i interface{}) bool {
	cust := i.(source.Customer)
	return strings.HasPrefix(cust.FirstName, "A")
}

func createSink(ctx context.Context) sink.Sink {
	connDest, err := sink.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}
	dest := sink.New(connDest)
	return dest
}

func createSource(ctx context.Context) source.CustomerChannel {
	connSource, err := source.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}
	src := source.New(connSource)
	return src
}

func mapToCustomerWithAddress(_ context.Context, i interface{}) (interface{}, error) {
	c := i.(source.Customer)
	return sink.CustomerWithAddress{
		Name:    transform.GetName(c.FirstName, c.LastName),
		Address: transform.GetAddress(c.ID),
	}, nil
}

func mapToUpperCase(_ context.Context, i interface{}) (interface{}, error) {
	c := i.(sink.CustomerWithAddress)
	c.Name = strings.ToUpper(c.Name)
	return c, nil
}
