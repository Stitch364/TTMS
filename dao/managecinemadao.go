package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

// ManageGetCinemaByLimit 根据筛选条件查询所有的影院信息，并简易展示影院列表
func ManageGetCinemaByLimit(brand, types, area, services, pageNumber string) (*model.Page, error) {
	var sqlStr string
	var rows *sql.Rows
	var err error

	//将页码转换为in64类型
	iPageNum, _ := strconv.ParseInt(pageNumber, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr = "select count(*) from cinema where (? = '' or brand = ?) and (? = '' or types = ?) and (? = '' or area = ?) and (? = '' or serve = ?)"

	//接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, brand, brand, types, types, area, area, services, services)
	err = row.Scan(&totalRecord)
	if err != nil {
		fmt.Println("GetMovieByLimit函数接收总记录数时出错：", err)
		return nil, err
	}
	//fmt.Println(totalRecord)
	//设置每页只显示6条记录
	var pageSize int64
	pageSize = 2
	//接收总页数
	var totalPages int64
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

	//创建切片储存影院信息
	var cinemaList []*model.Cinema
	for rows.Next() {
		//创建储存单项cinema
		cinema := new(model.Cinema)
		err := rows.Scan(&cinema.Id, &cinema.Name, &cinema.Brand, &cinema.Address, &cinema.City, &cinema.Area, &cinema.Types, &cinema.Serve, &cinema.PhoneNumber, &cinema.ImgPath)
		if err != nil {
			fmt.Println("遍历读取所有影院信息出错：", err)
			return nil, err
		}
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

// ManageGetCinemaById 根据影院Id查询影院详细信息
func ManageGetCinemaById(id string) (*model.Page, error) {
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

// AddCinema 增加影院信息
func AddCinema(cinema *model.Cinema) error {
	//查看是否重复提价
	var sqlStr string
	sqlStr = "select phonenumber from cinema where phonenumber = ?"
	//执行查询
	row := utils.Db.QueryRow(sqlStr, cinema.PhoneNumber)
	var temp string
	err := row.Scan(&temp)
	if !errors.Is(err, sql.ErrNoRows) {
		//数据重复
		//fmt.Println("重复的数据")
		return nil
	}

	//直接加到数据库就行
	sqlStr = "insert into cinema(name,brand,address,city,area,types,serve,phonenumber,imgpath) values(?,?,?,?,?,?,?,?,?)"

	_, err = utils.Db.Exec(sqlStr, cinema.Name, cinema.Brand, cinema.Address, cinema.City, cinema.Area, cinema.Types, cinema.Serve, cinema.PhoneNumber, cinema.ImgPath)
	if err != nil {
		fmt.Println("AddCinema添加数据出错了：", err)
		return err
	}
	return nil
}

// UpdateCinema 修改影院信息
func UpdateCinema(cinema *model.Cinema) error {
	var sqlStr string
	sqlStr = "UPDATE cinema SET name = ?,brand = ?,address = ?,area = ?,types = ?,serve = ?,phonenumber = ?,imgpath= ? WHERE cinema_id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cinema.Name, cinema.Brand, cinema.Address, cinema.Area, cinema.Types, cinema.Serve, cinema.PhoneNumber, cinema.ImgPath, cinema.Id)
	if err != nil {
		fmt.Println("UpdateCinema函数出错了：", err)
		return err
	}
	return nil

}

// DeleteCinema 删除影院信息，根据影院id删
func DeleteCinema(cinemaId string) error {
	var sqlStr string

	imovieId, _ := strconv.Atoi(cinemaId)

	sqlStr = "delete from cinema where cinema_id = ?"

	_, err := utils.Db.Exec(sqlStr, imovieId)
	if err != nil {
		return err
	}
	return nil
}
