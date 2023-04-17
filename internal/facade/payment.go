package facade

import "fmt"

type paymentService struct{}

func NewPaymentService() *paymentService {
	return &paymentService{}
}

func (ps *paymentService) makePayment(userID string, productID string) {
	fmt.Printf("[paymentService] user %s making payment for product %s\n", userID, productID)
}
