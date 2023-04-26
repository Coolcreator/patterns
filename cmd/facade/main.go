package main

import (
	"patterns/internal/facade/notification"
	"patterns/internal/facade/order"
	"patterns/internal/facade/payment"
	"patterns/internal/facade/product"
	"patterns/internal/facade/user"
)

func main() {
	u := user.NewUserService()
	prt := product.NewProductService()
	pmt := payment.NewPaymentService()
	n := notification.NewNotificationService()
	of := order.NewOrderFacade(u, prt, pmt, n)
	of.PlaceOrder("testUserID", "testProductID")
}
