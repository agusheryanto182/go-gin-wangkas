package user

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	Role         string
}