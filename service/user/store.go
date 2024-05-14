package user

import (
	"database/sql"
	"fmt"

	"github.com/sajin-shrestha/ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// func (s *Store) GetUserByEmail(email string) (*types.User, error) {
// 	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	u := new(types.User)

// 	for rows.Next() {
// 		u, err := scanRowIntoUser(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if u.ID == 0 {
// 		return nil, fmt.Errorf("user not found")
// 	}

// 	return u, nil
// }

// func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
// 	user := new(types.User)
// 	err := rows.Scan(
// 		&user.ID,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Email,
// 		&user.Password,
// 		&user.CreatedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// uncomment GetUserByEmail and scanRowIntoUser function above if it throws an error and comment below code
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := &types.User{} // Initialize u outside the loop

	for rows.Next() {
		err := scanRowIntoUser(rows, u)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows, user *types.User) error {
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}
