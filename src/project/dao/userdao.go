package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

// validate user password
func LoginValidate(username string, password string) (*model.User, error) {
	sqlQuery := "select id,username,password,email from users where username = ? and password = ?"

	row := utils.Db.QueryRow(sqlQuery, username, password)

	u := &model.User{}
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.Email)

	return u, err
}

// validate username is available
func UsernameAvailability(username string) (bool, error) {
	sqlQuery := "select id,username,password,email from users where username = ?"

	row := utils.Db.QueryRow(sqlQuery, username)
	err := row.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	} else {
		return false, nil
	}
}

// create user
func CreateUser(username, password, email string) error {
	sqlQuery := "insert into users(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlQuery, username, password, email)
	return err
}