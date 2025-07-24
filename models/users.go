package models

import (
	"github.com/justmamadou/rest-api-golang/db"
	"github.com/justmamadou/rest-api-golang/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Signup() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}
