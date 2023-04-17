package facade

import "fmt"

type productService struct{}

func NewProductService() *productService {
	return &productService{}
}

func (p *productService) productAvailable(productID string) bool {
	fmt.Printf("[productService] checking availability of product %s\n", productID)

	return true
}

func (p *productService) assignProductToUser(productID string, userID string) {
	fmt.Printf("[productService] assigning product %s to user %s\n", productID, userID)
}
