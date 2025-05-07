package main

import (
	"fmt"
	"time"
)

// Enums

type OrderStatus string

const (
	Recieved  OrderStatus = "Recieved"
	Confirmed             = "Confirmed"
	Prepared              = "Prepared"
	Deliverd              = "Deliverd"
)

// Embedding

type customer struct {
	name string
}

// order struct

type order struct {
	id     string
	amount float32
	status OrderStatus
	customer
	createdAt time.Time // nanoseconds precision
}

// we can use *pointer in go structs whenever we need to modify the values else we can read values without star

func (o *order) changeStatus(status OrderStatus) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func setOrder(id string, amount float32, status OrderStatus, customer customer) *order {

	myOrder := order{
		id:       id,
		amount:   amount,
		status:   status,
		customer: customer,
	}

	return &myOrder
}

func main() {

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"English", true}

	// fmt.Println(language)

	order := setOrder("1", 100, Deliverd, customer{name: "Kid Krish"})

	order.customer.name = "Robin Hood"

	fmt.Println(order.status)

	// If we don't pass any value it will set default zero value

	// string = "", int => 0, bool => false

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 500,
	// 	status: "recieved",
	// }

	// myOrder.changeStatus("paid")

	// fmt.Println(myOrder.getAmount())

	// fmt.Println(myOrder.status)

	fmt.Printf("Structs")
}
