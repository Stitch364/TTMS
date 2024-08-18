package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

// GetMovieIdByNameAndUpdateId1  管理端设置
// GetMovieIdByNameAndUpdateId1  从movie表中查找对应name的电影的id，并修改ranking1表中的id
func GetMovieIdByNameAndUpdateId1(name string) error {
	//1.在movie表中找到对应name的id
	sqlStr := "select movie_id from movie where name=?"
	row := utils.Db.QueryRow(sqlStr, name)
	//查找到的id
	var movieId string
	err := row.Scan(&movieId)
	if err != nil {
		fmt.Println("查找id失败，读取错误：", err)
		return err
	}
	//2.修改ranking表中对应name的id
	sqlStr = "update ranking1 set movie_name=? where ranking_id=?"
	_, err = utils.Db.Exec(sqlStr, name, movieId)
	if err != nil {
		fmt.Println("修改ranging1中的id失败：", err)
		return err
	}
	return nil
}

// GetMovieIdByNameAndUpdateId2 从expect_movie表中查找对应name的电影的id，并修改ranking2表中的id
func GetMovieIdByNameAndUpdateId2(name string) error {
	//1.在movie表中找到对应name的id
	sqlStr := "select movie_id from movie where name=?"
	row := utils.Db.QueryRow(sqlStr, name)
	//查找到的id
	var movieId string
	err := row.Scan(&movieId)
	if err != nil {
		fmt.Println("查找id失败，读取错误：", err)
		return err
	}
	//2.修改ranking表中对应name的id
	sqlStr = "update ranking2 set movie2_name=? where rang2_id=?"
	_, err = utils.Db.Exec(sqlStr, name, movieId)
	if err != nil {
		fmt.Println("修改ranging1中的id失败：", err)
		return err
	}
	return nil
}

// AddMovie 增加电影信息
func AddMovie(movie *model.Movie, status string) error {
	var sqlStr string
	//判断电影是否已经添加过，防止页面刷新的表单重复提交
	if status == "now" {
		sqlStr = "select name from movie where name = ?"
	} else {
		sqlStr = "select name from expect_movie where name = ?"
	}
	//执行查询
	row := utils.Db.QueryRow(sqlStr, movie.Name)
	var name string
	err := row.Scan(&name)
	if !errors.Is(err, sql.ErrNoRows) {
		//数据重复
		//fmt.Println("重复的数据")
		return nil
	}

	//增加前修改id
	if status == "now" {
		sqlStr = "select max(movie_id) from movie"
	} else {
		sqlStr = "select max(movie_id) from expect_movie"
	}

	row = utils.Db.QueryRow(sqlStr)
	var movieId string
	err = row.Scan(&movieId)
	if err != nil {
		fmt.Println("查询最大id时出错：", err)
		return err
	}
	imaxid, _ := strconv.Atoi(movieId)
	//修改id
	movie.Id = imaxid + 1

	if status == "now" {

		sqlStr = "insert into movie(movie_id,name, nickname, describes, type, area, time, year, starring, score, grosser, award, awardtimes, img_path) values (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)  "

	} else {
		sqlStr = "insert into expect_movie(movie_id,name, nickname, describes, type, area, time, year, starring, score, grosser, award, awardtimes, img_path) values (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)  "
	}
	////sql插入语句
	//sqlStr = "insert into movie(name, nickname, describes, type, area, time, year, starring, score, grosser, award, awardtimes, img_path) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)  "
	//执行
	_, err = utils.Db.Exec(sqlStr, movie.Id, movie.Name, movie.NickName, movie.Introduce, movie.Type, movie.Area, movie.Times, movie.Year, movie.Starring, movie.Score, movie.Grossed, movie.Award, movie.AwardTimes, movie.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMovie 根据电影id删除电影
func DeleteMovie(movieId string) error {
	var sqlStr string

	imovieId, _ := strconv.Atoi(movieId)
	//根据id判断该电影在哪一个表里
	if imovieId <= 1000 {
		sqlStr = "delete from movie where movie_id=?"
	} else {
		sqlStr = "delete from expect_movie where movie_id=?"
	}

	_, err := utils.Db.Exec(sqlStr, movieId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMovie 修改电影信息
func UpdateMovie(movie *model.Movie, status string) error {
	var sqlStr string
	if status == "now" {
		sqlStr = "UPDATE movie SET name = ?, nickname = ?, describes = ?, type = ?, area = ?, time = ?, year = ?, starring = ?, score = ?, grosser = ?, award = ?, awardtimes = ?, img_path = ? WHERE movie_id = ?"
	} else {
		sqlStr = "UPDATE expect_movie SET name = ?, nickname = ?, describes = ?, type = ?, area = ?, time = ?, year = ?, starring = ?, score = ?, grosser = ?, award = ?, awardtimes = ?, img_path = ? WHERE movie_id = ?"
	}
	//执行
	_, err := utils.Db.Exec(sqlStr, movie.Name, movie.NickName, movie.Introduce, movie.Type, movie.Area, movie.Times, movie.Year, movie.Starring, movie.Score, movie.Grossed, movie.Award, movie.AwardTimes, movie.ImgPath, movie.Id)
	if err != nil {
		return err
	}
	return nil

}
