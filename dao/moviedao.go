package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"database/sql"
	"fmt"
	"strconv"
)

// GetPageMovie 获取首页电影信息
func GetPageMovie() (*model.Page, error) {
	//首页只有一页，规定每项最多展示9部电影

	//获取数据库中电影的总记录数
	sqlStr := "select count(*) from movie"
	//接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	err := row.Scan(&totalRecord)
	if err != nil {
		return nil, err
	}

	//查出数据库中热门电影
	var movies []*model.Movie
	sqlStr = "select * from movie limit 0,7"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		movie := new(model.Movie)
		err := rows.Scan(&movie.Id, &movie.Name, &movie.NickName, &movie.Introduce, &movie.Type, &movie.Area, &movie.Times, &movie.Year, &movie.Starring, &movie.Score, &movie.Grossed, &movie.Award, &movie.AwardTimes, &movie.ImgPath)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	//查出数据库中最受期待电影
	var expectmovies []*model.Movie
	sqlStr = "select * from expect_movie limit 0,7"
	exRows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for exRows.Next() {
		movie := new(model.Movie)
		err := exRows.Scan(&movie.Id, &movie.Name, &movie.NickName, &movie.Introduce, &movie.Type, &movie.Area, &movie.Times, &movie.Year, &movie.Starring, &movie.Score, &movie.Grossed, &movie.Award, &movie.AwardTimes, &movie.ImgPath)
		if err != nil {
			return nil, err
		}
		expectmovies = append(expectmovies, movie)
	}

	//查出数据库中今日票房排行
	var moviesranking1s []*model.Ranking1
	sqlStr = "select * from ranking1"
	RowsR1, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for RowsR1.Next() {
		movie := new(model.Ranking1)
		err := RowsR1.Scan(&movie.Id, &movie.Name, &movie.Grossed, &movie.Cls, &movie.Movieid)
		if err != nil {
			return nil, err
		}
		moviesranking1s = append(moviesranking1s, movie)
	}

	//查出数据库中最受期待排行
	var moviesranking2s []*model.Ranking2
	sqlStr = "select * from ranking2"
	RowsR2, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for RowsR2.Next() {
		movie := new(model.Ranking2)
		err := RowsR2.Scan(&movie.Id, &movie.Name, &movie.Expect, &movie.Cls, &movie.Movieid)
		if err != nil {
			return nil, err
		}
		moviesranking2s = append(moviesranking2s, movie)
	}

	page := &model.Page{
		Movies:       movies,
		ExpectMovies: expectmovies,
		MovieRank1:   moviesranking1s,
		MovieRank2:   moviesranking2s,
		TotalRecord:  totalRecord,
	}

	return page, nil
}

// GetMovieById 根据电影Id查询电影详细信息
func GetMovieById(id string) (*model.Movie, error) {
	//将id转为int类型
	iid, _ := strconv.Atoi(id)
	var sqlStr string
	if iid <= 1000 {
		//如果id在<=1000，在表一查
		sqlStr = "select * from movie where movie_id=?"
	} else {
		//否则在表二查
		sqlStr = "select * from expect_movie where movie_id=?"
	}

	row := utils.Db.QueryRow(sqlStr, id)
	//两个表电影类型统一为Movie类型
	movie := new(model.Movie)
	err := row.Scan(&movie.Id, &movie.Name, &movie.NickName, &movie.Introduce, &movie.Type, &movie.Area, &movie.Times, &movie.Year, &movie.Starring, &movie.Score, &movie.Grossed, &movie.Award, &movie.AwardTimes, &movie.ImgPath)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

//// GetMovie 获取所有电影信息
//func GetMovie() ([]*model.Movie, error) {
//	var movies []*model.Movie
//	sqlStr := "select * from movie"
//	rows, err := utils.Db.Query(sqlStr)
//	if err != nil {
//		return nil, err
//	}
//	for rows.Next() {
//		movie := new(model.Movie)
//		err := rows.Scan(&movie.Id, &movie.Name, &movie.NickName, &movie.Introduce, &movie.Type, &movie.Area, &movie.Times, &movie.Year, &movie.Starring, &movie.Score, &movie.Grossed, &movie.Award, &movie.AwardTimes, &movie.ImgPath)
//		if err != nil {
//			return nil, err
//		}
//		movies = append(movies, movie)
//	}
//
//	return movies, nil
//}

// GetMovieByLimit 已分页
// GetMovieByLimit 根据筛选条件查询所有的电影信息，并简易展示电影列表
func GetMovieByLimit(status, genre, region, year, pageNumber string) (*model.Page, error) {
	var sqlStr string
	var rows *sql.Rows
	var err error
	genre2 := "%" + genre + "%"
	year2 := year + "%"

	//将页码转换为in64类型
	iPageNum, _ := strconv.ParseInt(pageNumber, 10, 64)
	//获取数据库中电影的总记录数
	if status == "upcoming" {
		sqlStr = "select count(*) from expect_movie where (? = '' or type like ?) and (? = '' or area = ?) and (? = '' or  year like ?)"
	} else {
		sqlStr = "select count(*) from movie where (? = '' or type like ?) and (? = '' or area = ?) and (? = '' or year like ?)"
	}

	//接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, genre, genre2, region, region, year, year2)
	err = row.Scan(&totalRecord)
	if err != nil {
		fmt.Println("GetMovieByLimit函数接收总记录数时出错：", err)
		return nil, err
	}
	//设置每页只显示6条记录
	var pageSize int64
	pageSize = 6
	//接收总页数
	var totalPages int64
	if totalRecord%pageSize == 0 {
		totalPages = totalRecord / pageSize
	} else {
		totalPages = totalRecord/pageSize + 1
	}

	//fmt.Println("总记录数为：", totalRecord)
	//fmt.Println("获取到的条件为：", status, genre, region, year)
	//根据电影上映状态分表查询
	if status == "upcoming" {
		//status = upcoming  即将上映
		sqlStr = "select * from expect_movie where (? = '' or type like ?) and (? = '' or area = ?) and (? = '' or year like ?) limit ?,?"
	} else {
		//正在热映
		sqlStr = "select * from movie where (? = '' or type like ?) and (? = '' or area = ?) and (? = '' or year like ?) limit ?,?"
	}
	rows, err = utils.Db.Query(sqlStr, genre, genre2, region, region, year, year2, (iPageNum-1)*pageSize, pageSize)
	//分界线

	if err != nil {
		fmt.Println("查询电影信息失败：", err)
		return nil, err
	}
	//创建切片储存影院信息

	var movies []*model.Movie
	//if !rows.Next() {
	//	fmt.Println("查询为空！")
	//}

	for rows.Next() {
		//创建储存单项cinema
		movie := new(model.Movie)
		err := rows.Scan(&movie.Id, &movie.Name, &movie.NickName, &movie.Introduce, &movie.Type, &movie.Area, &movie.Times, &movie.Year, &movie.Starring, &movie.Score, &movie.Grossed, &movie.Award, &movie.AwardTimes, &movie.ImgPath)
		if err != nil {
			fmt.Println("遍历读取所有影院信息出错：", err)
			return nil, err
		}
		//fmt.Println("查询到的数据为：", movie)

		//添加到切片中
		movies = append(movies, movie)
	}

	//返回包含所有影院信息的切片
	page := &model.Page{
		//将影院切片存到page中
		AllMovies:   movies,
		PageNumber:  iPageNum,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalRecord: totalRecord,
	}
	return page, nil
}
