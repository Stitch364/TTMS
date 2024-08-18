package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//
//// GetCinema 查询所有的影院信息，并简易展示影院列表
//func GetCinema() (*model.Page, error) {
//	sqlStr := "select * from cinema"
//	rows, err := utils.Db.Query(sqlStr)
//	if err != nil {
//		fmt.Println("查询所有影院信息失败：", err)
//		return nil, err
//	}
//	//创建切片储存影院信息
//
//	var cinemaList []*model.Cinema
//	for rows.Next() {
//		//创建储存单项cinema
//		cinema := new(model.Cinema)
//		err := rows.Scan(&cinema.Id, &cinema.Name, &cinema.Brand, &cinema.Address, &cinema.City, &cinema.Area, &cinema.Types, &cinema.Serve, &cinema.PhoneNumber, &cinema.ImgPath)
//		if err != nil {
//			fmt.Println("遍历读取所有影院信息出错：", err)
//			return nil, err
//		}
//		//添加到切片中
//		cinemaList = append(cinemaList, cinema)
//	}
//	//返回包含所有影院信息的切片
//	page := &model.Page{
//		//将影院切片存到page中
//		Cinemas: cinemaList,
//	}
//	return page, nil
//}

// GetCinemaByLimit 根据筛选条件查询所有的影院信息，并简易展示影院列表
func GetCinemaByLimit(brand, types, area, services, movieId, pageNumber string) (*model.Page, error) {
	var sqlStr, sqlStrTemp string
	var rows *sql.Rows
	var err error
	var cinemaids []string
	var totalRecord int64
	var totalPages int64
	//设置每页只显示6条记录
	var pageSize int64
	pageSize = 2
	var iPageNum int64

	if movieId != "" {
		//查询哪些影院放映该电影，查出影院id
		sqlStrTemp = "select distinct cinemaid from playback where movie_id = ? order by cinemaid"
		row, err := utils.Db.Query(sqlStrTemp, movieId)
		if err != nil {
			fmt.Println("GetCinemaByLimit函数中根据电影id查影院时出错", err)
			return nil, err
		}

		for row.Next() {
			var cinemaid string
			err = row.Scan(&cinemaid)
			if err != nil {
				fmt.Println("GetCinemaByLimit函数中根据id查电影，储存时出错：", err)
				return nil, err
			}
			//含有影院id的切片
			cinemaids = append(cinemaids, cinemaid)
		}
		//fmt.Println("放映该电影的影院id为：", cinemaids)
		//fmt.Println("处理后", cinemaids)
		//fmt.Println("-----------------------")

		placeholders := make([]string, len(cinemaids))
		args := make([]interface{}, len(cinemaids))
		args2 := make([]interface{}, len(cinemaids))

		for i, id := range cinemaids {
			placeholders[i] = "?"
			args[i] = id
			args2[i] = id
		}

		//判断该电影有没有排片信息
		if len(args) != 0 {
			// 生成IN条件的占位符
			inCondition := strings.Join(placeholders, ",")
			//将页码转换为in64类型
			iPageNum, _ = strconv.ParseInt(pageNumber, 10, 64)

			// 其他条件
			args = append(args, brand, brand, types, types, area, area, services, services, (iPageNum-1)*pageSize, pageSize)
			args2 = append(args2, brand, brand, types, types, area, area, services, services)
			//获取数据库中符合筛选条件的影院的总记录数

			//fmt.Println("查询条件1：", args)
			//fmt.Println("查询条件2：", args2)
			//动态生成占位符
			sqlStr = fmt.Sprintf("select count(*) from cinema where  (cinema_id in (%s)) and (? = '' or brand = ?) and (? = '' or types = ?) and (? = '' or area = ?) and (? = '' or serve = ?) ", inCondition)
			row2 := utils.Db.QueryRow(sqlStr, args2...)
			//执行
			//row2 := utils.Db.QueryRow(sqlStr)
			err = row2.Scan(&totalRecord)
			if err != nil {
				fmt.Println("GetCinemaByLimit函数接收总记录数时出错：", err)
				return nil, err
			}
			//接收总页数
			if totalRecord%pageSize == 0 {
				totalPages = totalRecord / pageSize
			} else {
				totalPages = totalRecord/pageSize + 1
			}
			//动态生成占位符
			sqlStr = fmt.Sprintf("select * from cinema where  (cinema_id in (%s)) and (? = '' or brand = ?) and (? = '' or types = ?) and (? = '' or area = ?) and (? = '' or serve = ?) limit ?,? ", inCondition)
			//fmt.Println(sqlStr)
			rows, err = utils.Db.Query(sqlStr, args...) //
			if err != nil {
				fmt.Println("查询所有影院信息失败：", err)
				return nil, err
			}
		} else {
			//该影院没有排片信息
			sqlStr = "select * from cinema where cinema_id = -1"
			rows, err = utils.Db.Query(sqlStr)
			if err != nil {
				fmt.Println("查询所有影院信息失败：", err)
				return nil, err
			}
		}
	} else {
		//将页码转换为in64类型
		iPageNum, _ = strconv.ParseInt(pageNumber, 10, 64)
		//获取数据库中符合筛选条件的影院的总记录数

		sqlStr = "select count(*) from cinema where (? = '' or brand = ?) and (? = '' or types = ?) and (? = '' or area = ?) and (? = '' or serve = ?)"
		row3 := utils.Db.QueryRow(sqlStr, brand, brand, types, types, area, area, services, services)
		//执行
		//row2 := utils.Db.QueryRow(sqlStr)
		err = row3.Scan(&totalRecord)
		if err != nil {
			fmt.Println("GetCinemaByLimit函数接收总记录数时出错：", err)
			return nil, err
		}
		//接收总页数
		if totalRecord%pageSize == 0 {
			totalPages = totalRecord / pageSize
		} else {
			totalPages = totalRecord/pageSize + 1
		}

		sqlStr = "select * from cinema where (? = '' or brand = ?) and (? = '' or types = ?) and (? = '' or area = ?) and (? = '' or serve = ?) limit ?,?"
		rows, err = utils.Db.Query(sqlStr, brand, brand, types, types, area, area, services, services, (iPageNum-1)*pageSize, pageSize)
		if err != nil {
			fmt.Println("查询所有影院信息失败：", err)
			return nil, err
		}
	}

	//fmt.Println("查询的所有条件为：", brand, types, area, services, cinemaids)

	////加入了含有该电影的影院条件
	//sqlStr = "select * from cinema where (? = '' or brand = ?) and (? = '' or types = ?) and (? = '' or area = ?) and (? = '' or serve = ?) and (? = '' or cinema_id in (?))"
	//rows, err = utils.Db.Query(sqlStr, brand, brand, types, types, area, area, services, services, cinemaids, cinemaids)
	//
	////分界线
	//if err != nil {
	//	fmt.Println("查询所有影院信息失败：", err)
	//	return nil, err
	//}
	//创建切片储存影院信息

	var cinemaList []*model.Cinema
	//if !rows.Next() {
	//	fmt.Println("查询为空！")
	//}

	for rows.Next() {
		//创建储存单项cinema
		cinema := new(model.Cinema)
		err := rows.Scan(&cinema.Id, &cinema.Name, &cinema.Brand, &cinema.Address, &cinema.City, &cinema.Area, &cinema.Types, &cinema.Serve, &cinema.PhoneNumber, &cinema.ImgPath)
		if err != nil {
			fmt.Println("遍历读取所有影院信息出错：", err)
			return nil, err
		}
		//fmt.Println("查询到的数据为：", cinema)

		//添加到切片中
		cinemaList = append(cinemaList, cinema)
	}

	//返回包含所有影院信息的切片
	page := &model.Page{
		//将影院切片存到page中
		Cinemas:     cinemaList,
		PageNumber:  iPageNum,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalRecord: totalRecord,
	}
	return page, nil
}

// GetAllDate 查询某影院都有哪些天放映
func GetAllDate(cinemaid string) ([]*model.AllDate, error) {
	sqlStr := "select distinct data from playback where cinemaid = ?"
	rows, err := utils.Db.Query(sqlStr, cinemaid)
	if err != nil {
		fmt.Println("GetAllDate函数查询时出错了：", err)
		return nil, err
	}
	var alldate []*model.AllDate
	for rows.Next() {
		date := new(model.AllDate)
		err = rows.Scan(&date.Day)
		if err != nil {
			fmt.Println("GetAllDate函数在读取时出错了：", err)
			return nil, err
		}
		alldate = append(alldate, date)
	}
	////检验
	//fmt.Println(alldate)
	return alldate, nil
}

// GetCinemaById 根据影院Id查询影院详细信息
func GetCinemaById(id string) (*model.Page, error) {
	sqlStr := "select * from cinema where cinema_id=?"
	row := utils.Db.QueryRow(sqlStr, id)
	//创建单个cinema
	cinema := new(model.Cinema)
	err := row.Scan(&cinema.Id, &cinema.Name, &cinema.Brand, &cinema.Address, &cinema.City, &cinema.Area, &cinema.Types, &cinema.Serve, &cinema.PhoneNumber, &cinema.ImgPath)
	if err != nil {
		fmt.Println("GetCinemaById函数查询数据库读取数据时出错了：", err)
		return nil, err
	}
	//创建cinema切片，并将查询到的信息添加到切片中
	var cinemaList []*model.Cinema
	cinemaList = append(cinemaList, cinema)
	//创建page并将切片赋值给Cinema1
	page := &model.Page{
		Cinema1: cinemaList,
	}
	return page, nil
}

// Movielist 先要查询该电影院有什么电影，展示出有的电影
// 在通过点击页面的电影链接，传电影id过来，才能查电影具体的放映情况
func Movielist(id, data string) ([]*model.MovieAndImgpath, error) {
	sqlStr2 := "select distinct movie,movie_id,imgpath,cinemaid from playback where cinemaid = ? and data = ?"
	rows2, err := utils.Db.Query(sqlStr2, id, data)
	if err != nil {
		fmt.Println("查询该电影院都放映什么电影时出错：", err)
		return nil, err
	}
	var movieList []*model.MovieAndImgpath
	for rows2.Next() {
		movie := new(model.MovieAndImgpath)
		err = rows2.Scan(&movie.Name, &movie.Id, &movie.Imgpath, &movie.Cinema_id)
		if err != nil {
			fmt.Println("写入movieList时出错了", err)
			return nil, err
		}
		//添加到切片
		movieList = append(movieList, movie)
	}
	//将放映的电影列表返回
	return movieList, nil

}

// GetPlaybackByIdAndDataAndMovie 通过影院id,日期查询该影院放映电影的信息
// 返回CinemaPlayback类型切片就行，page变量已经在查询前创建
func GetPlaybackByIdAndDataAndMovie(id, date, movieid string) ([]*model.CinemaPlayback, error) {
	//查询放映列表中该影院某一天的放映信息
	sqlStr := "select * from playback where cinemaid = ? && data = ? && movie_id = ?"
	rows, err := utils.Db.Query(sqlStr, id, date, movieid)
	if err != nil {
		fmt.Println("GetPlaybackByIdAndData函数在查询表示出错了：", err)
		return nil, err
	}
	//查出信息，写入切片
	var duration string
	var playbacks []*model.CinemaPlayback
	for rows.Next() {
		playback := new(model.CinemaPlayback)
		err = rows.Scan(&playback.Id, &playback.CinemaId, &playback.Date, &playback.Time, &duration, &playback.MovieId, &playback.MovieName, &playback.Price, &playback.Number, &playback.ImgPath, &playback.Hall)
		if err != nil {
			fmt.Println("GetPlaybackByIdAndData函数中读取信息时出错了：", err)
			return nil, err
		}
		//计算出电影的散场时间
		iduration, _ := strconv.Atoi(duration)
		temp, err := time.Parse("15：04", playback.Time)
		if err != nil {
			fmt.Println("解析时间格式错误：", err)
			return nil, err
		}
		afterXMinutes := temp.Add(time.Duration(iduration) * time.Minute)
		//fmt.Println("散场时间转换格式前：", afterXMinutes)
		playback.Duration = afterXMinutes.Format("15：04")
		//fmt.Println("开始时间：", playback.Time)
		//fmt.Println("电影时长：", iduration)
		//fmt.Println("散场时间", playback.Duration)
		//time.Sleep(3 * time.Second)
		//添加进切片
		playbacks = append(playbacks, playback)
	}

	//将playbacks切片返回
	return playbacks, nil
}

// SetImgpathByTable 未测试是否有Bug
// SetImgpathByTable 设置palyback表中的电影海报的地址
func SetImgpathByTable() {
	//查放映列表中都有哪些电影
	sqlStr := "select distinct movie_id from playback"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("SetImgpathByTable函数查询时出错了：", err)
	}
	var Ids []*int
	for rows.Next() {
		var Id int
		err = rows.Scan(&Id)
		if err != nil {
			fmt.Println("SetImgpathByTable读取出错了：", err)
			return
		}
		//将查到的id储存起来
		Ids = append(Ids, &Id)
	}
	var sqlStr1, sqlStr2 string
	//遍历所有在影院上映的电影
	for _, Id := range Ids {
		//判断从哪个表里查
		if *Id <= 1000 {
			sqlStr1 = "select img_path from movie where movie_id=?"
		} else {
			sqlStr1 = "select img_path from expect_movie where movie_id=?"
		}
		//查询Id为*Id的电影的图片地址
		row := utils.Db.QueryRow(sqlStr1, *Id)
		//存入ImgPath1中
		var ImgPath1 string
		err = row.Scan(&ImgPath1)
		if err != nil {
			fmt.Println("SetImgpathByTable函数中读取图片路径出错了：", err)
			return
		}
		//更改表playback中电影的图片地址
		sqlStr2 = "update playback set imgpath = ? where movie_id=?"
		_, err = utils.Db.Exec(sqlStr2, ImgPath1, *Id)
		if err != nil {
			fmt.Println("SetImgpathByTable函数中更新图片路径出错了", err)
			return
		}
	}
}
