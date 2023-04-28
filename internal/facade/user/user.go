package user

import "fmt"

// Service представляет собой сервис пользователей
type Service interface {
	ValidateUser(userID string) bool
}

type userService struct{}

// ValidateUser валидирует пользователя
func (*userService) ValidateUser(userID string) bool {
	fmt.Printf("[userService] validating user %s\n", userID)

	return true
}

// NewUserService создает новый экземпляр сервиса пользователей
func NewUserService() Service {
	return &userService{}
}
