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
func GetSession(sessionId string) (session *model.Session, err error) {
	sqlQuery := "select session_id, username, user_id from sessions where session_id = ?"
	stmt, err := utils.Db.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(sessionId)
	session = &model.Session{}
	_ = row.Scan(&session.SessionId, &session.Username, &session.UserId)
	return
}
func DeleteSession(sessionId string) (err error) {
	sqlQuery := "delete from sessions where session_id = ?"
	_, err = utils.Db.Exec(sqlQuery, sessionId)
	return
}