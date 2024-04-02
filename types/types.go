package types

import "time"

/**All Application types*/
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"fistName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Town      string `json:"town" validate:"required"`
	Age       uint64 `json:"age" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Age       uint64    `json:"age"`
	Town      string    `json:"town"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
}
