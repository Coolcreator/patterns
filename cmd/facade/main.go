package main

import "patterns/internal/facade"

func main() {
	orderFacade := facade.NewOrderFacade()
	orderFacade.PlaceOrder("test-user", "test-product")
}
