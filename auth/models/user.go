package models

type AuthUser struct {
	Username     string
	PasswordHash string
	UserId       int
}
