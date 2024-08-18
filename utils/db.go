package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库的包

var (
	Db  *sql.DB
	err error
)

func init() {
	// DSN (Data Source Name) 是包含数据库连接信息的字符串
	dsn := "root:413814@tcp(localhost:3306)/ttms?charset=utf8mb4&parseTime=True&loc=Local"
	//      用户名：密码@协议（主机名:端口）/数据库名称？设置字符集&让驱动解析时间类型字段为time.Time类型&设置时区为本地时区
	// 创建数据库连接
	//创建数据库对象
	//第一个参数：驱动程序名称，第二个参数告诉驱动程序如何访问基础数据存储
	//sql.Open不会建立于数据库的任何链接，也不会验证驱动程序的任何链接参数
	//与基础数据存储区的第一个实际连接将延迟到第一次需要时建立
	Db, err = sql.Open("mysql", dsn)
	//错误处理
	if err != nil {
		panic(err.Error())
	}

	// 检查连接是否有效
	//.Ping必要时建立连接
	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println("Connected to the database!")
	}
}
