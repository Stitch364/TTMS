package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"net/http"
)

func AddComments(w http.ResponseWriter, r *http.Request) {
	//增加评论
	//判断登录
	var page model.Page

	Username, flag := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = Username
	}

	text := &model.Text{
		Movieid:  r.FormValue("movie_Id"),
		Username: page.Username,
		Text:     r.FormValue("text"),
		ObjectId: r.FormValue("objectId"),
	}
	Object1 := r.FormValue("object")

	fmt.Println("Object:", Object1)
	if Object1 != "" {
		Object1 = "回复给" + Object1
	}
	text.Object = Object1
	err := dao.AddText(text)
	if err != nil {
		fmt.Println("添加评论失败：", err)
		return
	}
	Information(w, r)
}
