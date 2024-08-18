package model

type Goods struct {
	Id         int
	UserId     int
	UserName   string
	Ticket     string //场次的id
	Seat       string //座位信息
	BuyTime    string //购买时间
	Money      string
	State      string //订单状态
	CinemaName string
	MovieName  string
}
