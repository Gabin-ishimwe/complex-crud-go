package user

import (
	"database/sql"
	"fmt"

	"github.com/gabin-ishimwe/complex-crud-go/types"
)

const (
	QUERY_FETCH_USER_BY_EMAIL = "SELECT * FROM users WHERE email = ?"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Scan each row
func scanRowIntoUser(row *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query(QUERY_FETCH_USER_BY_EMAIL, email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
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
