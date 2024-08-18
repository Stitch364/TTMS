package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"strconv"
	"sync"
)

var mutex sync.RWMutex //定义锁变量

// GetTicketById 获取座位信息
// GetTicketById 根据场次id查出座位情况
func GetTicketById(Id string) (*model.Seat, error) {
	sqlStr := "select * from ticket where playbackid = ?"
	row := utils.Db.QueryRow(sqlStr, Id)

	ABC := new(model.Seat)
	err := row.Scan(&ABC.Id, &ABC.PlaybackId, &ABC.A, &ABC.B, &ABC.C, &ABC.D, &ABC.E, &ABC.F)
	if err != nil {
		fmt.Println("GetTicketById查询出错了：", err)
		return nil, err
	}
	return ABC, nil
}

// GetTicketByName 根据座位名,场次id查询该座位是否可买
func GetTicketByName(name, playbackid string) (bool, error) {
	//加读锁
	mutex.RLock()
	//解锁
	defer mutex.RUnlock()
	Ticketrow := string(name[0])
	column, _ := strconv.Atoi(string(name[1]))
	var inf string
	var sqlStr string
	switch Ticketrow {
	case "A":
		sqlStr = "select A from ticket where playbackid = ?"
	case "B":
		sqlStr = "select B from ticket where playbackid = ?"
	case "C":
		sqlStr = "select C from ticket where playbackid = ?"
	case "D":
		sqlStr = "select D from ticket where playbackid = ?"
	case "E":
		sqlStr = "select E from ticket where playbackid = ?"
	case "F":
		sqlStr = "select F from ticket where playbackid = ?"
	}

	//fmt.Println("查询条件：", Ticketrow, playbackid, column)
	err := utils.Db.QueryRow(sqlStr, playbackid).Scan(&inf)
	//fmt.Println(inf)
	if err != nil {
		fmt.Println("查询座位是否可买时出错了：", err)
		return false, err
	}

	//time.Sleep(3 * time.Second)
	//fmt.Println(column)
	if string(inf[column]) == "0" {
		//fmt.Println("座位可以买！")
		return true, nil
	} else {
		//fmt.Println("座位不可买！")
		return false, nil
	}
}

// UpdateTicket 更改座位信息
// 三个参数为 座位名，场次名 ，更改类型（1改0，0改1）
func UpdateTicket(name, playbackid, ty string) (string, error) {
	//加锁
	mutex.Lock()
	//解锁
	defer mutex.Unlock()

	//ABCDEF
	Ticketrow := string(name[0])
	//0~10
	column, _ := strconv.Atoi(string(name[1]))
	//column := int(name[1])
	//储存这一行信息的字符串
	var inf string
	var sqlStr string
	switch Ticketrow {
	case "A":
		sqlStr = "select A from ticket where playbackid = ?"
	case "B":
		sqlStr = "select B from ticket where playbackid = ?"
	case "C":
		sqlStr = "select C from ticket where playbackid = ?"
	case "D":
		sqlStr = "select D from ticket where playbackid = ?"
	case "E":
		sqlStr = "select E from ticket where playbackid = ?"
	case "F":
		sqlStr = "select F from ticket where playbackid = ?"
	}

	err := utils.Db.QueryRow(sqlStr, playbackid).Scan(&inf)
	if err != nil {
		fmt.Println("查询座位是否可买时出错了：", err)
		return "", err
	}
	//判断该座位的信息
	temp := []byte(inf)
	fmt.Println("temp：", temp, column)
	//修改前预检
	fmt.Println(string(temp[column]))
	if string(temp[column]) == "0" {
		//如果该位置还能买
		temp[column] = byte(49)
		// 0 改为 1
		//买票成功
	} else {
		//座位不能买
		//改为1
		if ty == "0" {
			//如果要改为可买，就改
			temp[column] = byte(48)
		} else {
			//ty == "1" 且 座位本身就是 1
			return "您选择的票已售完！", nil
		}

	}

	//更改数据库座位信息
	inf = string(temp)
	fmt.Println(inf)
	//fmt.Println("转换后的座位信息", inf)
	//写入
	switch Ticketrow {
	case "A":
		sqlStr = "update ticket set A = ? where playbackid = ?"
	case "B":
		sqlStr = "update ticket set B = ? where playbackid = ?"
	case "C":
		sqlStr = "update ticket set C = ? where playbackid = ?"
	case "D":
		sqlStr = "update ticket set D = ? where playbackid = ?"
	case "E":
		sqlStr = "update ticket set E = ? where playbackid = ?"
	case "F":
		sqlStr = "update ticket set F = ? where playbackid = ?"
	}

	_, err = utils.Db.Exec(sqlStr, inf, playbackid)
	if err != nil {
		fmt.Println("UpdateTicket函数在更改座位信息时出错了：", err)
		return "", err
	}
	//跳转到付款界面
	return "待付款。", nil
}

//根据场次的Id查找该场次电影的全部信息
