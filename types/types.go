package types

import "time"

/**All Application types*/
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"fistName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Town      string `json:"town"`
	Age       uint64 `json:"age"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       uint64    `json:"age"`
	Town      string    `json:"town"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
}
