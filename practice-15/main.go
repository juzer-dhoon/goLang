package main

import (
	"fmt"
	"time"
)

type CheckOrderStatus int

const (
	OrderPlaced CheckOrderStatus = iota
	OrderPreparing
	OrderOnTheWay
	OrderDelivered
)

func (s CheckOrderStatus) String() string {
	switch s {
	case OrderPlaced:
		return "Order Placed"
	case OrderPreparing:
		return "Preparing"
	case OrderOnTheWay:
		return "On The Way"
	case OrderDelivered:
		return "Delivered"
	default:
		return "Unknown Status"
	}
}

func main() {
	var status CheckOrderStatus = OrderPlaced
	fmt.Println("Juzer's Order Status...")

	for status <= OrderDelivered {
		fmt.Println(status)
		time.Sleep(2 * time.Second)
		status++
	}

	fmt.Println("\n🎉 Enjoy your order!")
}
