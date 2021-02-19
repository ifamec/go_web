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
