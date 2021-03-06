package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func AddSession(session *model.Session) (err error) {
	sqlQuery := "insert into sessions values(?,?,?)"
	_, err = utils.Db.Exec(sqlQuery, session.SessionId, session.Username, session.UserId)
	return
}
func DeleteSession(sessionId string) (err error) {
	sqlQuery := "delete from sessions where session_id = ?"
	_, err = utils.Db.Exec(sqlQuery, sessionId)
	return
}