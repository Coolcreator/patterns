package facade

type orderFacade struct {
	user         *userService
	product      *productService
	payment      *paymentService
	notification *notificationService
}

func NewOrderFacade() *orderFacade {
	return &orderFacade{
		user:         NewUserService(),
		product:      NewProductService(),
		payment:      NewPaymentService(),
		notification: NewNotificationService(),
	}
}

func (of *orderFacade) PlaceOrder(userID, productID string) {
	if of.user.ValidateUser(userID) && of.product.productAvailable(productID) {
		of.product.assignProductToUser(productID, userID)
		of.payment.makePayment(userID, productID)
		of.notification.NotifyDealer(productID)
	}
}
