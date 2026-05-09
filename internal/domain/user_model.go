package domain

type User struct {
	ID              int64
	Name            string
	Email, Password string `binding:"required"`
}
