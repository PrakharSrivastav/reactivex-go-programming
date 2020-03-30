package source

import (
	"github.com/reactivex/rxgo/v2"
	"time"
)
import "syreclabs.com/go/faker"

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	ID        string
}

type CustomerChannel struct {
	Ch chan rxgo.Item
}

func New(_ DBConnection) CustomerChannel {
	return CustomerChannel{
		Ch: make(chan rxgo.Item, 10),
	}
}

func (cc CustomerChannel) Read() {
	go func() {
		for _ = range time.Tick(time.Second) {
			for i := 0; i < 1000; i++ {
				cc.Ch <- rxgo.Of(Customer{
					ID:        faker.Code().Isbn10(),
					FirstName: faker.Name().FirstName(),
					LastName:  faker.Name().LastName(),
					Age:       faker.RandomInt(22, 53),
				})
			}
		}
		close(cc.Ch)
	}()
}
