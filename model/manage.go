package model

// Manage 结构体
type Manage struct {
	ID        int    //管理员id
	MName     string //管理员名
	MPassword string //管理员密码（md5加密）
	MKey      string //影院授权key
}
