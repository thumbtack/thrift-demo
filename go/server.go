package main

import (
	"math/rand"
	"net/http"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"

	"thrift-examples/gen-go/lib/thrift/demo"
)

// DemoService struct will have interface defined in demo.DemoService interface
type DemoService struct{}

// Calculate calculates result based on input
func (s *DemoService) Calculate(input *demo.CalculateInput) (float64, error) {
	switch input.Op {
	case demo.Operation_ADD:
		return input.Number1 + input.Number2, nil
	case demo.Operation_SUB:
		return input.Number1 - input.Number2, nil
	case demo.Operation_MUL:
		return input.Number1 * input.Number2, nil
	case demo.Operation_DIV:
		if input.Number2 == 0 {
			return 0, &demo.InputException{Message: "Division by zero"}
		}
		return input.Number1 / input.Number2, nil
	}
	return 0, &demo.InputException{Message: "Invalid operation"}
}

// RandomGenerator returns N pseudo-random numbers in range from min to max
func (s *DemoService) RandomGenerator(min int64, max int64, n int64) ([]int64, error) {
	// non secure :)
	rand.Seed(time.Now().UnixNano())
	l := make([]int64, n)
	for i := 0; i < int(n); i++ {
		l[i] = int64(rand.Intn(int(max+1-min)) + int(min))
	}
	return l, nil
}

func main() {
	// create instance of service
	service := &DemoService{}

	// Thrift bootstrapping:
	// create a processor that will handle the requests
	processor := demo.NewDemoServiceProcessor(service)

	// select binary protocol
	binaryF := thrift.NewTBinaryProtocolFactoryDefault()

	// http will route all requests to / to Thrift handler function which needs processor and binary
	// thrift factory
	http.HandleFunc("/", http.HandlerFunc(thrift.NewThriftHandlerFunc(processor, binaryF, binaryF)))

	// end of Thrift bootstrapping

	// start http server
	http.ListenAndServe(":8080", nil)
}
