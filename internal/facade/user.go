package facade

import "fmt"

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

func (us *userService) ValidateUser(userID string) bool {
	fmt.Printf("[userService] validating user %s\n", userID)

	return true
}
