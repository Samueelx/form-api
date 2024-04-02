package user

import (
	"database/sql"
	"fmt"

	"github.com/Samueelx/form-api/types"
)

/**A Repository where you fetch user details*/

type Store struct {
	db *sql.DB /**Use this to make queries*/
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	/**Create an empty user*/
	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}

/**Scans the row and returns a pointer to User type*/
func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	/**Create an empty user*/
	user := new(types.User)
	/**Fill up the user with data coming from the database*/
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Age,
		&user.Town,
		&user.Gender,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
