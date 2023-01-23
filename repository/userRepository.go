package repository

import (
	"bioskop/model"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func SaveUser(db *sql.DB, user model.User) error {
	sql := `
	INSERT INTO account(
	username, password, customerid)
	VALUES ($1, $2, $3);
	`
	var pass = BeforeSave(user)
	err := db.QueryRow(sql, user.Username, pass, user.CustomerID)
	return err.Err()
}
func BeforeSave(user model.User) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	return user.Password
}

func FindUserByUsername(db *sql.DB, username string) (user model.User, err error) {
	sql := `SELECT * FROM account WHERE username = $1`
	rows, err := db.Query(sql, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.CustomerID)
		if err != nil {
			panic(err)
		}
	}
	return
}

func FindUserById(db *sql.DB, id int) (user model.User, err error) {
	sql := `SELECT * FROM account WHERE id = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.CustomerID)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetAllUser(db *sql.DB) (result []model.User, err error) {
	sql := `SELECT * FROM account`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = model.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.CustomerID)
		if err != nil {
			panic(err)
		}

		result = append(result, user)
	}
	return
}
