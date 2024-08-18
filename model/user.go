package model

// User 结构体
type User struct {
	ID       int    //用户id
	Username string //用户名
	Password string //用户密码（md5加密）
	Email    string //用户邮箱
}
