package data

import (
	"database/sql"
	"fmt"

	"github.com/evlbit/notesmd/internal/types"
)

func (s *DataStore) CreateUser(user types.User) error {
	_, err := s.db.Exec(
		`INSERT INTO users (name, email, password)
		VALUES (?,?,?)`,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *DataStore) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query(
		`SELECT *
		FROM users
		WHERE email = ?`,
		email,
	)

	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = rowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *DataStore) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query(
		`SELECT *
		FROM users
		WHERE id = ?`,
		id,
	)

	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = rowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func rowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Updated,
		&user.Created,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
