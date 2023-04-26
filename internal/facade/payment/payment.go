package payment

import "fmt"

// Service представляет собой сервис оплаты
type Service interface {
	MakePayment(userID, productID string)
}

type paymentService struct{}

// MakePayment производит оплату
func (*paymentService) MakePayment(userID string, productID string) {
	fmt.Printf("[paymentService] user %s making payment for product %s\n", userID, productID)
}

// NewPaymentService создает новый экземпляр сервиса оплаты
func NewPaymentService() Service {
	return &paymentService{}
}
