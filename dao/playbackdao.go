package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
)

// SelectPlayback 根据影院Id查找排片信息
func SelectPlayback(id string) (*model.CinemaPlayback, error) {
	sqlStr := "select * from playback where id=?"
	row := utils.Db.QueryRow(sqlStr, id)
	oneplayback := new(model.CinemaPlayback)
	err := row.Scan(&oneplayback.Id, &oneplayback.CinemaId, &oneplayback.Date, &oneplayback.Time, &oneplayback.Duration, &oneplayback.MovieId, &oneplayback.MovieName, &oneplayback.Price, &oneplayback.Number, &oneplayback.ImgPath, &oneplayback.Hall)
	if err != nil {
		fmt.Println("SelectPlayback函数查询信息是出错了：", err)
		return nil, err
	}
	return oneplayback, nil
}

//增加排片信息

//修改排片信息

//自动删除排片信息
//默认过期时间为当前时间的1小时后，将当前时间倒退1小时
//例9点开场的电影，10点删
