package model

type LoginResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
