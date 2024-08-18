package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	//fmt.Println("文件解析成功！")
	//获取所有电影信息
	page, err := dao.GetPageMovie()
	//fmt.Println("电影信息获取成功！")
	//fmt.Println(page, err)

	//确认是否登录
	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}

	//fmt.Println(err)
	if err != nil {
		//电影信息获取失败，打开空模板
		err := t.Execute(w, nil)
		if err != nil {
			return
		}
	}
	err = t.Execute(w, page)
	if err != nil {
		return
	}
}

// Information 去电影的详情页
func Information(w http.ResponseWriter, r *http.Request) {
	//获取电影的Id
	movieId := r.FormValue("movie_Id")
	//fmt.Println(movieId)
	//调用用Id查找电影信息的函数，查找出电影的所有信息
	movieinformation, err := dao.GetMovieById(movieId)
	if err != nil {
		fmt.Println("用id查电影信息出错：", err)
		return
	}

	//创建movie切片，并将查到的电影信息加进去
	var movies []*model.Movie
	movies = append(movies, movieinformation)
	//赋值切片
	page := &model.Page{
		Single: movies,
	}

	//确认是否登录
	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}

	//fmt.Println("登录状态：", page.IsLogin)

	texts, err := dao.GetText(movieId)
	if err != nil {
		fmt.Println("处理器获取评论出错：", err)
		return
	}
	page.Comments = texts
	//page.MovieId = movieId

	t := template.Must(template.ParseFiles("views/pages/cart/information.html"))

	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("解析information出错：", err)
		return
	}
}

func GetMovieByLimit2(w http.ResponseWriter, r *http.Request) {
	//获取筛选的四个条件
	//先置空
	status, genre, region, year := "now", " ", " ", " "
	//获取条件
	status = r.FormValue("status")
	genre = r.FormValue("genre")
	region = r.FormValue("region")
	year = r.FormValue("year")
	pageNumber := r.FormValue("pageNumber")
	if pageNumber == "" || pageNumber == "0" {
		pageNumber = "1"
	}

	page, err := dao.GetMovieByLimit(status, genre, region, year, pageNumber)
	if err != nil {
		fmt.Println("GetMovieByLimit2根据4个条件获取电影错误", err)
		return
	}
	//确认是否登录
	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}

	t := template.Must(template.ParseFiles("views/pages/cart/allmovie1.html"))
	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("GetMovieByLimit2函数在打开模板时出错：", err)
		return
	}
}
