package dao

import (
	"TTMS/model"
	. "TTMS/utils"
	"database/sql"
	"errors"
)

// AddUser 插入用户信息
func AddUser(username1 string, password1 string, email1 string) error {
	query := "INSERT INTO user(username, password,email) VALUES (?, ?, ?)"
	//fmt.Println(username1, password1, email1)
	_, err := Db.Exec(query, username1, password1, email1)
	if err != nil {
		return err
	}
	//fmt.Println("Data inserted successfully")
	return nil
}

// CheckUserNameAndpassword  根据用户名,密码从数据库中查询一条记录（用户密码）
func CheckUserNameAndpassword(username1 string, password string) (*model.User, error) {
	// 在指定列中指定查询
	query := "SELECT userid,username,password,email FROM user WHERE username = ? AND password = ?"

	row := Db.QueryRow(query, username1, password)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	//fmt.Println(err)
	if err != nil {
		return nil, err
	}
	////检查点
	//fmt.Println(user)
	return user, nil
}

// CheckUserName 根据用户名从数据库中查询一条记录（用户密码）
func CheckUserName(username1 string) (*model.User, error) {
	// 在指定列中指定查询
	query := "SELECT userid,username,password,email FROM user WHERE username = ?"

	row := Db.QueryRow(query, username1)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		//没有找到用户信息（未注册）
		return nil, err
	}
	////检查点
	//fmt.Println(password)
	//用户已经注册
	return user, nil
}
