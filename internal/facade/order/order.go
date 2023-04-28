package order

type userService interface {
	ValidateUser(userID string) bool
}

type productService interface {
	ProductAvailable(productID string) bool
	AssignProductToUser(productID string, userID string)
}

type paymentService interface {
	MakePayment(userID, productID string)
}

type notificationService interface {
	NotifyDealer(productID string)
}

// Facade представляет собой фасад сервиса заказов
type Facade interface {
	PlaceOrder(userID, productID string)
}

type orderFacade struct {
	user         userService
	product      productService
	payment      paymentService
	notification notificationService
}

// PlaceOrder размещает заказ
func (f *orderFacade) PlaceOrder(userID, productID string) {
	if f.user.ValidateUser(userID) && f.product.ProductAvailable(productID) {
		f.product.AssignProductToUser(productID, userID)
		f.payment.MakePayment(userID, productID)
		f.notification.NotifyDealer(productID)
	}
}

// NewOrderFacade создает новый экземпляр фасада сервиса заказов
func NewOrderFacade(
	user userService,
	product productService,
	payment paymentService,
	notification notificationService) Facade {
	return &orderFacade{
		user:         user,
		product:      product,
		payment:      payment,
		notification: notification,
	}
}
