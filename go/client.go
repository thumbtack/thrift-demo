package main

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"

	"thrift-examples/gen-go/lib/thrift/demo"
)

func main() {
	// Create HTTP transport
	transport, err := thrift.NewTHttpPostClient("http://localhost:8080/")
	if err != nil {
		fmt.Println("Error creating transport:", err)
		return
	}
	transport.Open()
	defer transport.Close()

	// binary protocol
	binaryF := thrift.NewTBinaryProtocolFactoryDefault()

	// create a client
	service := demo.NewDemoServiceClientFactory(transport, binaryF)

	// at this point, client is ready, and we can make RPC calls
	fmt.Printf("Random numbers: ")
	fmt.Println(service.RandomGenerator(1, 10, 5))

	input := &demo.CalculateInput{
		Number1: float64(5),
		Number2: float64(2),
		Op:      demo.Operation_ADD,
	}
	fmt.Printf("Calculation: ")
	fmt.Println(service.Calculate(input))
}
