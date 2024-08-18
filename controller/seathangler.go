package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// GetTicket 获取座位信息
func GetTicket(w http.ResponseWriter, r *http.Request) {
	playbackId := r.FormValue("playbackId")
	seat, err := dao.GetTicketById(playbackId)
	if err != nil {
		fmt.Println("GetTicket函数在调用GetTicketById函数时出错了：", err)
		return
	}
	var arr [6][10]bool
	for i := 0; i < 6; i++ {
		for j := 0; j < 10; j++ {
			switch i {
			case 0:
				arr[i][j], err = strconv.ParseBool(string(seat.A[j]))
				if err != nil {
					fmt.Println("GetTicket函数在转换类型时出错：", err)
					return
				}
			case 1:
				arr[i][j], err = strconv.ParseBool(string(seat.B[j]))
				if err != nil {
					fmt.Println("GetTicket函数在转换类型时出错：", err)
					return
				}
			case 2:
				arr[i][j], err = strconv.ParseBool(string(seat.C[j]))
				if err != nil {
					fmt.Println("GetTicket函数在转换类型时出错：", err)
					return
				}
			case 3:
				arr[i][j], err = strconv.ParseBool(string(seat.D[j]))
				if err != nil {
					fmt.Println("GetTicket函数在转换类型时出错：", err)
					return
				}
			case 4:
				arr[i][j], err = strconv.ParseBool(string(seat.E[j]))
				if err != nil {
					fmt.Println("GetTicket函数在转换类型时出错：", err)
					return
				}
			case 5:
				arr[i][j], err = strconv.ParseBool(string(seat.F[j]))
				if err != nil {
					fmt.Println("GetTicket函数在转换类型时出错：", err)
					return
				}
			}
		}
	}
	//检查
	//fmt.Println(arr[0])
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/buy.html"))
	err = t.Execute(w, arr)
	if err != nil {
		fmt.Println("GetTicket函数在打开模板时出错：", err)
		return
	}
}

// GetTicketByName2 根据座位名，场次id判断某座位是否可购买，两次检查，一次修改
func GetTicketByName2(w http.ResponseWriter, r *http.Request) {
	//场次id
	id := r.FormValue("Id")
	//座位信息
	sert := r.FormValue("seat")
	//购买的票数
	num := r.FormValue("num")
	//用  ,  去切割
	sert2 := strings.Split(sert, ",")
	//fmt.Println("购买的位置：", sert2)
	//切割后得到例:  0-1  0-2   0-3 的字符串切片
	//循环查询用户购买的所有座位
	for _, v := range sert2 {
		vv := strings.Split(v, "-")
		//fmt.Println("单个位置：", vv)
		//vv[0] 为 "0"
		//vv[1] 为 "1"
		abc := []byte(vv[0])[0]
		//将0转为A
		abc += 17
		Abc := string(abc) + vv[1]
		//整个过程将0-1转化为A1
		fmt.Println("转换后的座位：", Abc, id)
		canbuy, err := dao.GetTicketByName(Abc, id)
		if err != nil {
			fmt.Println("GetTicketByName2函数中的GetTicketByName在查询时出错：", err)
			return
		}
		if !canbuy {
			//第一次检查就失败，提示，直接返回到选座界面
			t := template.Must(template.ParseFiles("views/pages/cart/buylose.html"))
			err = t.Execute(w, id)
			if err != nil {
				fmt.Println("GetTicketByName2函数在打开购买失败模板时出错：", err)
				return
			}
			return
		}
	}
	//第一次检查成功
	//进入预检和修改
	//还是遍历每个座位
	var inf string
	var err error
	for _, v := range sert2 {
		vv := strings.Split(v, "-")
		//vv[0] 为 "0"
		//vv[1] 为 "1"
		abc := []byte(vv[0])[0]
		//将0转为A
		abc += 17
		Abc := string(abc) + vv[1]
		//fmt.Println("要购买的座位为：", Abc)
		//整个过程将0-1转化为A1
		//加锁
		inf, err = dao.UpdateTicket(Abc, id, "1")
		if err != nil {
			fmt.Println("GetTicketByName2函数中的UpdateTicket在查询时出错：", err)
			return
		}
		if inf == "您选择的票已售完！" {
			//没抢到
			break
		}
	}
	fmt.Println(inf)
	if inf == "您选择的票已售完！" {
		//预检失败
		t := template.Must(template.ParseFiles("views/pages/cart/buylose.html"))
		err = t.Execute(w, id)
		if err != nil {
			fmt.Println("GetTicketByName2函数在打开购买失败模板时出错：", err)
			return
		}
		return
	} else {
		//抢票成功,交给payment函数处理
		//10min付款限制
		//去我的订单页面
		func(id, sert, num string) {
			//建立订单的结构体
			//场次id，座位信息，票数
			//增加订单信息
			page := new(model.Page)
			//判断登录
			Username, flag := dao.IsLogin(r)
			if flag {
				page.IsLogin = true
				page.Username = Username
			}
			time1 := time.Now().Format("2006-01-02 15:04:05")
			price, err := dao.GetPriceById(id)

			if err != nil {
				fmt.Println("查询价格失败：", err)
				return
			}
			inum, _ := strconv.Atoi(num)
			prices := price * float64(inum)
			str := strconv.FormatFloat(prices, 'f', 2, 64)
			goods := &model.Goods{
				UserName: Username,
				Ticket:   id,
				Seat:     sert,
				BuyTime:  time1,
				Money:    str,
				State:    "待付款",
			}
			err = dao.AddGoods(*goods)
			if err != nil {
				fmt.Println("增加订单信息", err)
				return
			}
			//打开我的订单
			GetGoods(w, r)
			//定时10min
			//main2()

			return
		}(id, sert, num)
	}
}

// 付款10min的定时器
func main2(name, playbackid string) {
	condition := make(chan bool, 1)
	// 设置10分钟的定时器
	timer := time.NewTimer(5 * time.Second)
	var wg sync.WaitGroup

	go func() {
		//检查用户是否付款

		// 假设5秒后付款
		time.Sleep(2 * time.Second)

		// 发送信号，表示条件已满足
		condition <- true
	}()

	wg.Add(1)
	// 启动另一个goroutine来等待定时器或条件满足
	go func() {
		defer wg.Done()
		select {
		case <-condition:
			fmt.Println("条件已满足，取消定时器")
			if !timer.Stop() {
				//timer.Stop()返回的布尔值表示定时器是否在停止之前已经超时
				// 从定时器通道中读取以消耗已经发送的值
				<-timer.C
			}
		case <-timer.C:
			fmt.Println("10分钟已到，条件未满足，执行操作")
			// 这里执行你需要的操作
			//将座位恢复
			_, err := dao.UpdateTicket(name, playbackid, "0")
			if err != nil {
				fmt.Println("座位复原出错了！", err)
				return
			}

		}
	}()

	//// 等待定时器或条件满足的goroutine完成
	wg.Wait()
}
