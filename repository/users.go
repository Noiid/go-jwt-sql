package repository

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) FetchUserLogin(username, password string) (*User, error) {
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == GetMD5Hash(password) {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("login failed")
}

func (u *UserRepository) SelectAll() ([]User, error) {
	preparedStatement := `SELECT id, name, user_name, password FROM users`

	rows, err := u.db.Query(preparedStatement)

	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
