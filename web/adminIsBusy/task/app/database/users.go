package database

import (
	"KaifTravel/models"
	"KaifTravel/utils"
)

func UserExists(username string) (bool, error) {
	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?;", username).Scan(&rowCount)
	if err != nil {
		return false, err
	}
	return rowCount == 1, nil
}

func CreateUser(username string, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT username, password FROM users WHERE username = ?;", username).Scan(&user.Username, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}
