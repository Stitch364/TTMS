package controller

import (
	. "TTMS/dao"
	"TTMS/model"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

//// MD5StrSalted 处理器包
//// MD5StrSalted 加盐（固定值）加密 413814
//func MD5StrSalted(password []byte) (string, error) {
//	salt := []byte("413814")
//	return fmt.Sprintf("%x", md5.Sum(append(salt, password...))), nil
//}

// ManageLogin 管理员登录
func ManageLogin(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	_, flag := MIsLogin(r)
	if flag {
		//管理去首页
		IndexHandler(w, r)
	} else { //获取用户名和密码
		mname := r.PostFormValue("mname")
		mpassword := r.PostFormValue("mpassword")
		//加密
		md5password, err := MD5StrSalted([]byte(mpassword))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //400
		}
		//fmt.Println(username, password)
		//调用userdao中验证用户名的密码的方法
		manage, _ := CheckManageNameAndpassword(mname, md5password)
		//fmt.Println(user)
		//用户账号和密码都正确
		if manage != nil && manage.ID > 0 {
			//用户名和密码正确
			//跳转到登录成功页面
			//生成UUID并转换为字符串
			uuid0 := uuid.New().String()
			//fmt.Println(uuid0)
			//创建一个Session
			sess := &model.MSession{
				SessionId: uuid0,
				MName:     manage.MName,
				MId:       manage.ID,
			}
			//将session保存到数据库中
			err = AddMSession(sess)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			//创建一个cookie.让它与Session相关联
			cookie := http.Cookie{
				Name:  "manage",
				Value: uuid0,
			}
			//将Cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/manager/mlogin_to_manage.html"))
			err := t.Execute(w, nil)
			if err != nil {
				return
			}
		} else {
			//用户名和密码不正确，继续回到登录页面
			//fmt.Println("用户名或密码错误！")
			t := template.Must(template.ParseFiles("views/pages/manager/mlogin.html"))
			//fmt.Println("模板解析错误")
			err := t.Execute(w, "用户名或密码不正确！")
			//fmt.Println("模板运行错误：", err)
			if err != nil {
				return
			}
			//Login(w, r)
		}
	}
}

// ManageRegister 注册
func ManageRegister(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	mname := r.PostFormValue("mname")
	mpassword := r.PostFormValue("mpassword")
	MD5Password, err := MD5StrSalted([]byte(mpassword))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
	}
	mkey := r.PostFormValue("mkey")
	//调用userdao中验证用户名的密码的方法
	user, _ := CheckUserName(mname)
	if user != nil && user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/manager/mregister.html"))
		err := t.Execute(w, "用户名已存在！")
		if err != nil {
			return
		}
	} else {
		//用户名可用，将用户数据保存到数据库中
		err := AddManage(mname, MD5Password, mkey)
		//处理错误
		if err != nil {
			return
		}
		//fmt.Println("注册成功！")
		t := template.Must(template.ParseFiles("views/pages/manager/mregister_to_manage.html"))
		err = t.Execute(w, "")
		if err != nil {
			return
		}
	}
}

// ManageLogout 注销的函数
func ManageLogout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("manage")
	if cookie != nil {

		//获取cookie的value
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		err := DeleteMSession(cookieValue)
		if err != nil {
			return
		}
		//设置cookie失效
		cookie.MaxAge = -1
	}
	//去管理首页
	IndexHandler(w, r)
}
