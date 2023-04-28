package product

import "fmt"

// Service представляет собой сервис товаров
type Service interface {
	ProductAvailable(productID string) bool
	AssignProductToUser(productID string, userID string)
}

type productService struct{}

// ProductAvailable проверяет доступен ли продукт
func (*productService) ProductAvailable(productID string) bool {
	fmt.Printf("[productService] checking availability of product %s\n", productID)

	return true
}

// AssignProductToUser присваивает товар пользователю
func (*productService) AssignProductToUser(productID string, userID string) {
	fmt.Printf("[productService] assigning product %s to user %s\n", productID, userID)
}

// NewProductService создает новый экземпляр сервиса товаров
func NewProductService() Service {
	return &productService{}
}
