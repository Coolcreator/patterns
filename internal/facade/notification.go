package facade

import "fmt"

type notificationService struct{}

func NewNotificationService() *notificationService {
	return &notificationService{}
}

func (ns *notificationService) NotifyDealer(productID string) {
	fmt.Printf("[notificationService] notifying dealer about sale of product %s\n", productID)
}
