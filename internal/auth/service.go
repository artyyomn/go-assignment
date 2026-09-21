package auth

import "github.com/artyyomn/go-assignment/internal/user"


type Service struct{
	user user.Repository
}

func NewService(user user.Repository) *Service{
	return &Service{
		user: user,
	}
}
