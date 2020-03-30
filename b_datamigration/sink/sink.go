package sink

import "fmt"

type CustomerWithAddress struct {
	Name    string
	Address string
}

type Sink string

func New(_ FileConnection) Sink {
	return ""
}

func (s Sink) Write(cust CustomerWithAddress) {
	fmt.Printf("CustomerWithAddress :: %+v \n", cust)
}
