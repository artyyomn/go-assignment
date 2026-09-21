package user

type Repository interface{
	Create(*User) error
	FindUser(string) (*User, error)
}
