package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
)

func AddGoods(goods model.Goods) error {
	sqlStr := "insert into user_goods(userid, username, ticket, seat, but_time,money,status) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, goods.UserId, goods.UserName, goods.Ticket, goods.Seat, goods.BuyTime, goods.Money, goods.State)
	if err != nil {
		fmt.Println("增加goods内容失败：", err)
		return err
	}
	return nil
}

// GetGood 获取所有订单
func GetGood(username string) ([]*model.Goods, error) {
	fmt.Println(username)
	sqlStr := "select id,username,ticket,seat,but_time,money,status from user_goods where username=?"
	rows, err := utils.Db.Query(sqlStr, username)
	if err != nil {
		fmt.Println("GetGood查询出错了：", err)
		return nil, err
	}
	var goodses []*model.Goods
	for rows.Next() {
		var goods model.Goods
		err := rows.Scan(&goods.Id, &goods.UserName, &goods.Ticket, &goods.Seat, &goods.BuyTime, &goods.Money, &goods.State)
		if err != nil {
			fmt.Println("GetGood函数查询出错：", err)
			return nil, err
		}
		goodses = append(goodses, &goods)
	}
	return goodses, nil
}

// GetPriceById 根据id查价格
func GetPriceById(id string) (float64, error) {
	sqlStr := "select price from playback where id=?"
	row := utils.Db.QueryRow(sqlStr, id)

	var price float64
	err := row.Scan(&price)
	if err != nil {
		fmt.Println("GetPriceById函数查价格出错了: ", err)
	}
	return price, err
}
