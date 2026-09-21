package auth

import "github.com/artyyomn/go-assignment/internal/user"


type Service struct{
	user user.Repository
	session SessionRespository
}

func NewService(user user.Repository, session SessionRespository) *Service{
	return &Service{
		user: user,
		session: session,
	}
}
