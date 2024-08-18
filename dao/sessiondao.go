package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"net/http"
)

// AddSession 向数据库中添加Session
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions (session_id,username,user_id) values (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionId, sess.UserName, sess.UserId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// DeleteSession 删除数据库中的Session
func DeleteSession(sessionId string) error {
	sqlStr := "delete from sessions where session_id=?"
	_, err := utils.Db.Exec(sqlStr, sessionId)
	if err != nil {
		return err
	}
	return nil
}

// GetSession 根据session_id查Session
func GetSession(sessionId string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id=?"
	row := utils.Db.QueryRow(sqlStr, sessionId)
	var sess model.Session
	err := row.Scan(&sess.SessionId, &sess.UserName, &sess.UserId)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

// IsLogin 判断是否登录
func IsLogin(r *http.Request) (string, bool) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookievalue := cookie.Value
		session, _ := GetSession(cookievalue)
		if session != nil {
			return session.UserName, true
		}
	}
	return "", false
}
