package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
)

//// AddGoods2 增加订单
//func AddGoods2(w http.ResponseWriter, r *http.Request) {
//	page := new(model.Page)
//	//判断登录
//	Username, flag := dao.IsLogin(r)
//	if flag {
//		page.IsLogin = true
//		page.Username = Username
//	}
//	goods := &model.Goods{
//		UserName: Username,
//		Ticket:   r.FormValue(),
//		Seat:     r.FormValue(),
//		BuyTime:  r.FormValue(),
//	}
//}

// GetGoods 显示所有订单
func GetGoods(w http.ResponseWriter, r *http.Request) {
	page := new(model.Page)
	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}
	goods1, err := dao.GetGood(Username)
	if err != nil {
		fmt.Println("GetGoods函数出错了：", err)
		return
	}
	//订单切片
	page.Goods = goods1
	//打开我的订单
	t := template.Must(template.ParseFiles("views/pages/cart/buysuccess.html"))

	err = t.Execute(w, page)
	if err != nil {
		fmt.Println("解析information出错：", err)
		return
	}
}
