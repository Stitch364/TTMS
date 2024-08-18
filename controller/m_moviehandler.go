package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func ManageGetMovieByLimit2(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("ManageGetMovieByLimit2根据5个条件获取电影错误", err)
		return
	}

	//到这个页面必然登录了
	t := template.Must(template.ParseFiles("views/pages/manager/manageallmovie.html"))
	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("ManageGetMovieByLimit2函数在打开模板时出错：", err)
		return
	}
}

// ManageInformation2 去管理电影的信息页
func ManageInformation2(w http.ResponseWriter, r *http.Request) {
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
	//不必确认登录
	t := template.Must(template.ParseFiles("views/pages/manager/manageinformation.html"))

	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("解析information出错：", err)
		return
	}
}

// AddMovie2 增加电影信息
func AddMovie2(w http.ResponseWriter, r *http.Request) {
	//先获取前端传来的所有信息
	times := r.FormValue("Times")
	grossed := r.FormValue("Grossed")
	awardtimes := r.FormValue("AwardTimes")
	status := r.FormValue("status")
	score := r.FormValue("Score")
	//fmt.Println(status)
	//将这三个转换为int类型
	itimes, _ := strconv.Atoi(times)
	igrossed, _ := strconv.Atoi(grossed)
	iawardtimes, _ := strconv.Atoi(awardtimes)
	moviename := r.FormValue("Name")
	fscore, _ := strconv.ParseFloat(score, 64)
	//fmt.Println(imovieId)
	movie := &model.Movie{
		Name:       r.FormValue("Name"),
		NickName:   r.FormValue("NickName"),
		Introduce:  r.FormValue("Introduce"),
		Type:       r.FormValue("Types"),
		Area:       r.FormValue("Area"),
		Times:      itimes,
		Year:       r.FormValue("Year"),
		Starring:   r.FormValue("Starring"),
		Score:      fscore,
		Grossed:    igrossed,
		Award:      r.FormValue("Award"),
		AwardTimes: iawardtimes,
	}

	//// 解析表单
	//err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	//if err != nil {
	//	http.Error(w, "Unable to parse form", http.StatusBadRequest) //400
	//	fmt.Println("解析表单出错：", err)
	//	return
	//}

	// 获取上传的文件地址
	imgpath := r.FormValue("imgpath")
	//打开文件
	file, err := os.Open(imgpath)
	if err != nil {
		fmt.Println("图片地址打开失败：", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件file关闭失败：", err)
			return
		}
	}(file)

	// 生成新的文件名
	newFileName := "img_" + moviename + ".png" // 生成文件名格式统一为  img_+电影名.png
	//储存路径
	filePath := filepath.Join("E:/Program Files/GO/TTMS/views/static/img", newFileName)
	//储存到数据库中的地址
	newpath := "views/static/img/img_" + moviename + ".png"

	if newpath != imgpath {
		// 创建文件
		outFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("创建文件出错：", err)
			return
		}
		//关闭
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				fmt.Println("文件outFile关闭失败：", err)
				return
			}
		}(outFile)

		// 将上传的文件内容复制到新的文件中
		_, err = io.Copy(outFile, file)
		if err != nil {
			fmt.Println("复制文件出错：", err)
			return
		}
		//文件保存成功后保存文件地址
		movie.ImgPath = newpath
	} else {
		movie.ImgPath = imgpath
	}
	//修改信息
	//fmt.Println(movie.ImgPath)
	err = dao.AddMovie(movie, status)
	if err != nil {
		fmt.Println("修改电影信息时出错：", err)
		return
	}
	//去管理电影的页面
	ManageGetMovieByLimit2(w, r)
}

// UpdateMovie2 修改电影信息（几乎和增加一样）
func UpdateMovie2(w http.ResponseWriter, r *http.Request) {
	//先获取前端传来的所有信息
	movieId := r.FormValue("movie_Id")
	times := r.FormValue("Times")
	grossed := r.FormValue("Grossed")
	awardtimes := r.FormValue("AwardTimes")
	status := r.FormValue("status")
	score := r.FormValue("Score")
	//fmt.Println(status)
	//将这三个转换为int类型
	imovieId, _ := strconv.Atoi(movieId)
	itimes, _ := strconv.Atoi(times)
	igrossed, _ := strconv.Atoi(grossed)
	iawardtimes, _ := strconv.Atoi(awardtimes)
	moviename := r.FormValue("Name")
	fscore, _ := strconv.ParseFloat(score, 64)
	//fmt.Println(imovieId)
	movie := &model.Movie{
		Id:         imovieId,
		Name:       r.FormValue("Name"),
		NickName:   r.FormValue("NickName"),
		Introduce:  r.FormValue("Introduce"),
		Type:       r.FormValue("Type"),
		Area:       r.FormValue("Area"),
		Times:      itimes,
		Year:       r.FormValue("Year"),
		Starring:   r.FormValue("Starring"),
		Score:      fscore,
		Grossed:    igrossed,
		Award:      r.FormValue("Award"),
		AwardTimes: iawardtimes,
	}

	// 获取上传的文件地址
	imgpath := r.FormValue("imgpath")
	//打开文件
	file, err := os.Open(imgpath)
	if err != nil {
		fmt.Println("图片地址打开失败：", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件file关闭失败：", err)
			return
		}
	}(file)

	// 生成新的文件名
	newFileName := "img_" + moviename + ".png" // 生成文件名格式统一为  img_+电影名.png
	//filePath := filepath.Join("views/static/img", newFileName) // 存储路径
	filePath := filepath.Join("E:/Program Files/GO/TTMS/views/static/img", newFileName)
	newpath := "views/static/img/img_" + moviename + ".png"

	if newpath != imgpath {
		// 创建文件
		outFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("创建文件出错：", err)
			return
		}
		//关闭
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				fmt.Println("文件outFile关闭失败：", err)
				return
			}
		}(outFile)

		// 将上传的文件内容复制到新的文件中
		_, err = io.Copy(outFile, file)
		if err != nil {
			fmt.Println("复制文件出错：", err)
			return
		}
		//文件保存成功后保存文件地址
		movie.ImgPath = newpath
	} else {
		movie.ImgPath = imgpath
	}
	//修改信息
	//fmt.Println(movie.ImgPath)

	err = dao.UpdateMovie(movie, status)
	if err != nil {
		fmt.Println("修改电影信息时出错：", err)
		return
	}
	//去管理电影的页面
	ManageGetMovieByLimit2(w, r)
}

func DeleteMovie2(w http.ResponseWriter, r *http.Request) {
	movieId := r.FormValue("movie_id")
	err := dao.DeleteMovie(movieId)
	if err != nil {
		fmt.Println("删除电影信息失败：", err)
		return
	}
	//去管理电影的页面
	ManageGetMovieByLimit2(w, r)
}
