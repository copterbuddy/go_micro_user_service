package service

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserService interface {
	Create(email string, password string, name string) (UserResponse, error)
	GetAll() ([]UserResponse, error)
}
