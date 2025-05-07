package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
}

// we can use *pointer in go structs whenever we need to modify the values else we can read values without star

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	myOrder := order{
		id:     "1",
		amount: 500,
		status: "recieved",
	}

	myOrder.changeStatus("paid")

	fmt.Println(myOrder.getAmount())

	fmt.Println(myOrder.status)

	fmt.Printf("Structs")
}
