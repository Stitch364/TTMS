package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"net/http"
)

//管理员的Session管理

// AddMSession 向数据库中添加Session
func AddMSession(sess *model.MSession) error {
	sqlStr := "insert into msession (msession_id,mname,mid) values (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionId, sess.MName, sess.MId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// DeleteMSession 删除数据库中的Session
func DeleteMSession(sessionId string) error {
	sqlStr := "delete from msession where msession_id=?"
	_, err := utils.Db.Exec(sqlStr, sessionId)
	if err != nil {
		return err
	}
	return nil
}

// GetMSession 根据session_id查Session
func GetMSession(sessionId string) (*model.MSession, error) {
	sqlStr := "select * from msession where msession_id=?"
	row := utils.Db.QueryRow(sqlStr, sessionId)
	var sess model.MSession
	err := row.Scan(&sess.SessionId, &sess.MName, &sess.MId)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

// MIsLogin 判断是否登录
func MIsLogin(r *http.Request) (string, bool) {
	cookie, _ := r.Cookie("manage")
	if cookie != nil {
		cookievalue := cookie.Value
		session, _ := GetMSession(cookievalue)
		if session != nil {
			return session.MName, true
		}
	}
	return "", false
}
