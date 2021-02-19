package model

import (
	"fmt"
	"go_web/src/webapp/chapter04_db/utils"
)

// User Struct
type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

// AddUser M1
func (u *User) AddUser() (err error) {
	// SQL Query
	sqlQuery := "insert into users(username,password,email) values(?,?,?)"

	// Prepare
	inStmt, err := utils.Db.Prepare(sqlQuery)
	if err != nil {
		fmt.Println("AddUser : SQL Prepare Error -", err)
		return
	}

	// Execute
	_, err = inStmt.Exec("admin", "000000", "admin@zzhio.com")
	if err != nil {
		fmt.Println("AddUser : SQL Execute Error -", err)
		return
	}

	return
}

// AddUserNoPre AddUser without prepare
func (u *User) AddUserNoPre() (err error) {
	// SQL Query
	sqlQuery := "insert into users(username,password,email) values(?,?,?)"

	// Execute
	_, err = utils.Db.Exec(sqlQuery, "root", "root", "root@zzhio.com")
	if err != nil {
		fmt.Println("AddUser : SQL Prepare Error -", err)
		return
	}

	return
}

// GetUserById
func (u *User) GetUserById() (user *User, err error) {
	sqlQuery := "select id,username,password,email from users where id = ?"

	// get row
	row := utils.Db.QueryRow(sqlQuery, u.Id)

	// declare variables
	var id int
	var username, password, email string
	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("GetUserById Error -", err)
	}

	user = &User{
		Id:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return
}

// GetUsers
func (u *User) GetUsers() (users []*User, err error) {
	sqlQuery := "select id,username,password,email from users"

	// get rows
	rows, err := utils.Db.Query(sqlQuery)
	if err != nil {
		fmt.Println("GetUsers Error -", err)
	}

	for rows.Next() {
		// declare variables
		var id int
		var username, password, email string
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println("GetUsers Rows Loop Error -", err)
		}
		users = append(users, &User{
			Id:       id,
			Username: username,
			Password: password,
			Email:    email,
		})
	}

	return
}
