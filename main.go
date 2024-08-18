package main

import (
	. "TTMS/controller"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("🐱‍💻👍")
	//处理静态资源
	//index页面找css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//其他页面找css
	http.Handle("/views/static/", http.StripPrefix("/views/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/views/pages/", http.StripPrefix("/views/pages/", http.FileServer(http.Dir("views/pages"))))

	//用户端
	//主界面
	http.HandleFunc("/", IndexHandler)
	//登录
	http.HandleFunc("/login", Login)
	//注销
	http.HandleFunc("/logout", Logout)
	//注册
	http.HandleFunc("/register", Register)
	//电影详情页
	http.HandleFunc("/information", Information)
	//根据四个条件筛选电影
	http.HandleFunc("/GetMovieByLimit", GetMovieByLimit2)
	////所有影院
	//http.HandleFunc("/allcinema", GetCinema2)
	//影院详情页
	http.HandleFunc("/cinema", GoCinema)
	//根据四个筛选条件筛选影院(5个，一个是电影id在隐藏域中)
	http.HandleFunc("/GetCinemaByLimit", GetCinemaByLimit2)
	//买票界面
	http.HandleFunc("/GetTicket", GetTicket)
	//查询选择的座位是否可买
	http.HandleFunc("/Check", GetTicketByName2)
	//我的订单
	http.HandleFunc("/GetGoods", GetGoods)
	//添加评论
	http.HandleFunc("/AddComments", AddComments)

	//管理员
	//管理员登录
	http.HandleFunc("/manageLogin", ManageLogin)
	//管理员注册
	http.HandleFunc("/manageRegister", ManageRegister)
	//管理员注销
	http.HandleFunc("/manageLogout", ManageLogout)
	//管理界面-筛选电影（查找）
	http.HandleFunc("/ManageGetMovieByLimit", ManageGetMovieByLimit2)
	//管理-电影信息的页面
	http.HandleFunc("/ManageInformation", ManageInformation2)
	//管理-增加电影信息（增加，修改，删除按钮在一个页面）
	http.HandleFunc("/AddMovie", AddMovie2)
	//管理-删除电影信息
	http.HandleFunc("/DeleteMovie", DeleteMovie2)
	//管理-修改电影信息
	http.HandleFunc("/UpdateMovie", UpdateMovie2)
	//管理界面-筛选影院（查找）
	http.HandleFunc("/ManageGetCinemaByLimit", ManageGetCinemaByLimit2)
	//管理-影院信息的页面
	http.HandleFunc("/GoManageCinema", GoManageCinema)
	//管理-增加影院
	http.HandleFunc("/AddCinema", AddCinema2)
	//管理-删除影院
	http.HandleFunc("/DeleteCinema", DeleteCinema2)
	//管理-修改影院信息
	http.HandleFunc("/UpdateCinema", UpdateCinema2)
	////管理-增加排片信息
	//http.HandleFunc("/AddCinema", AddCinema2)
	////管理-更改排片信息
	//http.HandleFunc("/UpdateCinema", UpdateCinema2)
	////手动删除排片信息

	//自动删除排片信息（删除过期的：开场1小时后）
	//http.HandleFunc("/DeleteCinema", DeleteCinema2)
	//评分后更改评分

	//创建路由
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
