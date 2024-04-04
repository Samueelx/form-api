package user

import (
	"database/sql"
	"fmt"
	"log"

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
	fmt.Printf("Querying the database: %s\n", email)
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

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
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
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
func (s *Store) UpdateUserById(firstName, lastName, age, town, gender string, id int) error {
	result, err := s.db.Query("UPDATE users SET firstName = ?, lastName = ?, age = ?, town = ?, gender = ? WHERE id = ?", firstName, lastName, age, town, gender, id)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		return err
	}

	log.Printf("Rows data after updating: %v\n", result)
	return nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password, age, town, gender) VALUES(?, ?, ?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.Age, user.Town, user.Gender)
	if err != nil {
		return err
	}

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
		&user.Password,
		&user.Age,
		&user.Town,
		&user.Gender,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
