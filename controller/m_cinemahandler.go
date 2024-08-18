package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// ManageGetCinemaByLimit2 根据筛选条件查询所有影院信息
func ManageGetCinemaByLimit2(w http.ResponseWriter, r *http.Request) {
	//获取筛选影院的4个条件
	//先置空，再获取
	brand, types, area, services := " ", " ", " ", " "
	//获取四个条件
	brand = r.FormValue("brand")
	types = r.FormValue("types")
	area = r.FormValue("area")
	services = r.FormValue("services")

	pageNumber := r.FormValue("pageNumber")
	if pageNumber == "" || pageNumber == "0" {
		pageNumber = "1"
	}

	page, err := dao.ManageGetCinemaByLimit(brand, types, area, services, pageNumber)
	//返回page
	if err != nil {
		fmt.Println("ManageGetCinemaByLimit函数出错了：", err)
		return
	}
	//页面还没写
	t := template.Must(template.ParseFiles("views/pages/manager/manageallcinema.html"))

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

// GoManageCinema 去影院详情页
// 获取影院信息
func GoManageCinema(w http.ResponseWriter, r *http.Request) {
	//调整排片电影的图片
	dao.SetImgpathByTable()

	//获取影院id
	cinemaId := r.FormValue("cinema_id")
	//通过id去查影院的信息，并创建page变量
	page, err := dao.ManageGetCinemaById(cinemaId)
	if err != nil {
		fmt.Println("GoCinema函数在查影院信息出错了：", err)
		return
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/managecinema.html"))

	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("GoCinema在打开模板文件时出错了：", err)
		return
	}
}

// AddCinema2 增加电影信息
func AddCinema2(w http.ResponseWriter, r *http.Request) {
	//先获取前端传来的所有信息

	cinema := &model.Cinema{
		Name:        r.FormValue("Name"),
		Brand:       r.FormValue("Brand"),
		Address:     r.FormValue("Address"),
		Area:        r.FormValue("Area"),
		City:        "西安",
		Types:       r.FormValue("Types"),
		Serve:       r.FormValue("Serve"),
		PhoneNumber: r.FormValue("PhoneNumber"),
		ImgPath:     "views/static/img/c1.png",
	}

	//修改信息
	//fmt.Println(movie.ImgPath)
	err := dao.AddCinema(cinema)
	if err != nil {
		fmt.Println("增加影院信息时出错：", err)
		return
	}
	//去管理影院的页面
	ManageGetCinemaByLimit2(w, r)
}

// UpdateCinema2 修改影院信息
func UpdateCinema2(w http.ResponseWriter, r *http.Request) {
	//先获取前端传来的所有信息
	cinemaId, _ := strconv.Atoi(r.FormValue("Id"))

	cinema := &model.Cinema{
		Id:          cinemaId,
		Name:        r.FormValue("Name"),
		Brand:       r.FormValue("Brand"),
		Address:     r.FormValue("Address"),
		Area:        r.FormValue("Area"),
		City:        "西安",
		Types:       r.FormValue("Types"),
		Serve:       r.FormValue("Serve"),
		PhoneNumber: r.FormValue("PhoneNumber"),
		ImgPath:     "views/static/img/c1.png",
	}

	//修改信息
	//fmt.Println(movie.ImgPath)
	err := dao.UpdateCinema(cinema)
	if err != nil {
		fmt.Println("修改影院信息时出错：", err)
		return
	}
	//去管理影院的页面
	ManageGetCinemaByLimit2(w, r)
}

// DeleteCinema2 删除电影信息
func DeleteCinema2(w http.ResponseWriter, r *http.Request) {
	cinemaid := r.FormValue("cinema_id")
	err := dao.DeleteCinema(cinemaid)
	if err != nil {
		fmt.Println("DeleteCinema2函数删除影院信息出错：", err)
		return
	}
	//去管理影院的页面
	ManageGetCinemaByLimit2(w, r)
}
