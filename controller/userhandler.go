package controller

import (
	. "TTMS/dao"
	"TTMS/model"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

// MD5StrSalted 处理器包
// MD5StrSalted 加盐（固定值）加密 413814
func MD5StrSalted(password []byte) (string, error) {
	salt := []byte("413814")
	return fmt.Sprintf("%x", md5.Sum(append(salt, password...))), nil
}

// Login 登录
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	_, flag := IsLogin(r)
	if flag {
		//去首页
		IndexHandler(w, r)
	} else { //获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//加密
		md5password, err := MD5StrSalted([]byte(password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //400
		}
		//fmt.Println(username, password)
		//调用userdao中验证用户名的密码的方法
		user, _ := CheckUserNameAndpassword(username, md5password)
		//fmt.Println(user)
		//用户账号和密码都正确
		if user != nil && user.ID > 0 {
			//用户名和密码正确
			//跳转到登录成功页面
			//生成UUID并转换为字符串
			uuid0 := uuid.New().String()
			//fmt.Println(uuid0)
			//创建一个Session
			sess := &model.Session{
				SessionId: uuid0,
				UserName:  user.Username,
				UserId:    user.ID,
			}
			//将session保存到数据库中
			err = AddSession(sess)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			//创建一个cookie.让它与Session相关联
			cookie := http.Cookie{
				Name:  "user",
				Value: uuid0,
			}
			//将Cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_to_index.html"))
			err := t.Execute(w, nil)
			if err != nil {
				return
			}
		} else {
			//用户名和密码不正确，继续回到登录页面
			//fmt.Println("用户名或密码错误！")
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
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

// Register 注册
func Register(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	MD5Password, err := MD5StrSalted([]byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
	}
	email := r.PostFormValue("email")
	//调用userdao中验证用户名的密码的方法
	user, _ := CheckUserName(username)
	if user != nil && user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		err := t.Execute(w, "用户名已存在！")
		if err != nil {
			return
		}
	} else {
		//用户名可用，将用户数据保存到数据库中
		err := AddUser(username, MD5Password, email)
		//处理错误
		if err != nil {
			return
		}
		//fmt.Println("注册成功！")
		t := template.Must(template.ParseFiles("views/pages/user/register_to_login.html"))
		err = t.Execute(w, "")
		if err != nil {
			return
		}
	}
}

// Logout 注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {

		//获取cookie的value
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		err := DeleteSession(cookieValue)
		if err != nil {
			return
		}
		//设置cookie失效
		cookie.MaxAge = -1
	}
	//去首页
	IndexHandler(w, r)
}
