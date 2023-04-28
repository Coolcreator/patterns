package notification

import "fmt"

// Service представляет собой сервис уведомлений
type Service interface {
	NotifyDealer(productID string)
}

type notificationService struct{}

// NotifyDealer уведомляет дилера
func (*notificationService) NotifyDealer(productID string) {
	fmt.Printf("[notificationService] notifying dealer about sale of product %s\n", productID)
}

// NewNotificationService создает экземпляр сервиса уведомлений
func NewNotificationService() Service {
	return &notificationService{}
}
