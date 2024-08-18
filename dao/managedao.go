package dao

import (
	"TTMS/model"
	. "TTMS/utils"
	"database/sql"
	"errors"
)

// AddManage 插入管理员信息
func AddManage(mname1 string, mpassword1 string, mkey string) error {
	query := "INSERT INTO manage(mname, mpassword,mkey) VALUES (?, ?, ?)"
	//fmt.Println(username1, password1, email1)
	_, err := Db.Exec(query, mname1, mpassword1, mkey)
	if err != nil {
		return err
	}
	//fmt.Println("Data inserted successfully")
	return nil
}

// CheckManageNameAndpassword  根据管理员名,密码从数据库中查询一条记录
// 用于验证账号密码是否输入正确
func CheckManageNameAndpassword(mname1 string, mpassword string) (*model.Manage, error) {
	// 在指定列中指定查询
	query := "SELECT * FROM manage WHERE mname = ? AND mpassword = ?"

	row := Db.QueryRow(query, mname1, mpassword)

	manage1 := &model.Manage{}
	err := row.Scan(&manage1.ID, &manage1.MName, &manage1.MPassword, &manage1.MKey)
	//fmt.Println(err)
	if err != nil {
		return nil, err
	}
	////检查点
	//fmt.Println(user)
	return manage1, nil
}

// CheckManageName 根据管理员名从数据库中查询一条记录（管理员用户密码）
func CheckManageName(mname1 string) (*model.Manage, error) {
	// 在指定列中指定查询
	query := "SELECT * FROM manage WHERE mname = ?"

	row := Db.QueryRow(query, mname1)

	manage1 := &model.Manage{}
	err := row.Scan(&manage1.ID, &manage1.MName, &manage1.MPassword, &manage1.MKey)
	if errors.Is(err, sql.ErrNoRows) {
		//没有找到用户信息（未注册）
		return nil, err
	}
	////检查点
	//fmt.Println(password)
	//用户已经注册
	return manage1, nil
}
