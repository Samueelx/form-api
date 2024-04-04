package types

import "time"

/**All Application types*/
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	UpdateUserById(firstName, lastName, age, town, gender string, id int) error
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
	Age       string `json:"age" validate:"required"`
	Town      string `json:"town" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       string    `json:"age"`
	Town      string    `json:"town"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserUpdateData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       string `json:"age"`
	Town      string `json:"town"`
	Gender    string `json:"gender"`
}
