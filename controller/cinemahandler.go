package controller

import (
	"TTMS/dao"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// GetCinemaByLimit2 根据筛选条件查询所有影院信息
func GetCinemaByLimit2(w http.ResponseWriter, r *http.Request) {
	//获取筛选影院的4个条件
	//先置空，再获取
	brand, types, area, services := " ", " ", " ", " "
	//获取四个条件
	brand = r.FormValue("brand")
	types = r.FormValue("types")
	area = r.FormValue("area")
	services = r.FormValue("services")
	movieid := r.FormValue("movieId")
	pageNumber := r.FormValue("pageNumber")
	if pageNumber == "" || pageNumber == "0" {
		pageNumber = "1"
	}

	page, err := dao.GetCinemaByLimit(brand, types, area, services, movieid, pageNumber)
	//返回page
	if err != nil {
		fmt.Println("GetCinemaByLimit函数出错了：", err)
		return
	}
	//页面还没写
	t := template.Must(template.ParseFiles("views/pages/cart/allcinema.html"))

	//确认是否登录
	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}
	//传入page
	//page中只有登录和影院信息
	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("打开模板文件失败：", err)
		return
	}
}

// GoCinema 去影院详情页
// 1.获取影院信息
// 2.获取影院今天放映的电影信息
func GoCinema(w http.ResponseWriter, r *http.Request) {
	//调整排片电影的图片
	dao.SetImgpathByTable()

	//1.第一次查询

	//获取影院id
	cinemaId := r.FormValue("cinema_id")
	//通过id去查影院的信息，并创建page变量
	page, err := dao.GetCinemaById(cinemaId)
	if err != nil {
		fmt.Println("GoCinema函数在查影院信息出错了：", err)
		return
	}

	//确认是否登录
	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}

	//获取并处理时间信息
	//默认情况下显示的是今天的
	//获取今天的日期
	now := time.Now().Format("1月2日")
	date1 := r.FormValue("date")
	//不用解析了
	//date为空就显示今天的
	if date1 == "" {
		date1 = now
	}
	//else {
	//	// 解析日期字符串为 time.Time 类型
	//	datetemp, err := time.Parse("2006-01-02", date1)
	//	if err != nil {
	//		fmt.Println("GoCinema函数在解析日期字符串时出错了：", err)
	//		return
	//	}
	//	// 格式化日期为“1月1日”样式
	//	date1 = datetemp.Format("1月2日")
	//}

	//date1为最终处理好的时间

	//2.第二次查询
	//获取电影列表信息并加入page变量中
	//fmt.Println("查询电影列表前传入的数据：", cinemaId, date1)
	movielist, err := dao.Movielist(cinemaId, date1)
	page.MoveAndImgpath = movielist
	//fmt.Println("电影列表：", page.MoveAndImgpath)

	//3.第三次查询

	//一个页面只能显示某一天，某一部电影的信息
	//所以只查某一个影院，某一天，某一部电影
	//首先你得这个影院在某天都有哪些电影
	//cinemaId := r.FormValue("cinema_id")

	//初始的movie为空
	movieid := r.FormValue("movieId")
	//fmt.Println("获取到的movieId:", movieId)
	if movieid == "" {
		page.SameDateAndMovie = nil
	} else {
		//根据影院id,日期，电影id去放映列表的数据库查
		playback2, err := dao.GetPlaybackByIdAndDataAndMovie(cinemaId, date1, movieid)
		if err != nil {
			fmt.Println("GoCinema函数中查询的GetPlaybackByIdAndData函数出错了", err)
		}
		//将某一日的某个电影信息切片加到page中
		page.SameDateAndMovie = playback2
	}

	//4.第四次查询
	//查询该电影院都有哪些天放放映
	alldate, err := dao.GetAllDate(cinemaId)
	if err != nil {
		fmt.Println("第四次查询出错了：", err)
		return
	}
	//写入page
	page.AllDate = alldate

	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/cart/cinema.html"))

	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("GoCinema在打开模板文件时出错了：", err)
		return
	}

}

//func GetPlaybackByIdAndData2(w http.ResponseWriter, r *http.Request) {
//	//默认情况下显示的是今天的
//	//获取今天的日期
//	now := time.Now().Format("01月02")
//	//先获取到影院的id和所选的信息
//	cinemaId := r.FormValue("cinema_id")
//	date1 := r.FormValue("date")
//	//date为空就显示今天的
//
//	var date2 string
//	if date1 == "" {
//		date1 = now
//	} else {
//		// 解析日期字符串为 time.Time 类型
//		datetemp, err := time.Parse("2006-01-02", date1)
//		if err != nil {
//			fmt.Println("GoCinema函数在解析日期字符串时出错了：", err)
//			return
//		}
//
//		// 格式化日期为“1月1日”样式
//		date2 = datetemp.Format("1月2日")
//	}
//	//影院id和查询的日期均准备好了
//	playback2, err := dao.GetPlaybackByIdAndData(cinemaId, date2)
//	if err != nil {
//		fmt.Println("GoCinema函数中查询的GetPlaybackByIdAndData函数出错了", err)
//	}
//	page:=&model.Page{
//		SameDateAndMovie:playback2,
//	}
//
//	//解析模板......
//}
