package data

import (
	"database/sql"
	"fmt"
)

type User struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

func CreateUser(handle string, password string) (User, error) {
	if usable, err := CheckHandleUsable(handle); err != nil {
		return User{}, err
	} else if !usable {
		return User{}, fmt.Errorf("DB: Handle '%s' is not usable", handle)
	}

	newUser := User{
		Handle:   handle,
		Password: password,
	}

	query := `INSERT INTO users (handle, password) VALUES (?, ?)`
	_, err := DB.Exec(query, newUser.Handle, newUser.Password)

	if err != nil {
		return User{}, err
	}

	fmt.Println("DB: User created")
	return newUser, nil
}

func GetUserByHandle(handle string) (User, error) {
	var user User
	query := `SELECT handle, password FROM users WHERE handle=?`
	err := DB.QueryRow(query, handle).Scan(&user.Handle, &user.Password)

	return user, err
}

func CheckHandleUsable(handle string) (bool, error) {
	var count uint

	query := `SELECT * FROM users WHERE handle=?`
	query = fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery_alias", query)

	err := DB.QueryRow(query, handle).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if count > 0 {
		return false, nil
	} else {
		return true, nil
	}
}
