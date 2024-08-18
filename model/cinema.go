package model

type Cinema struct {
	Id          int    //影院id
	Name        string //影院名（品牌+影院名）
	Brand       string //品牌
	Address     string //影院地址
	City        string //影院所属城市
	Area        string //影院所属地区
	Types       string //影厅类型
	Serve       string //影院服务（可改签，可退票）(可能多个服务)
	PhoneNumber string //影院的电话号码
	ImgPath     string //影院图片
	//Playback         []*CinemaPlayback //该影院所有放映信息的切片
	//SameDateAndMovie []*CinemaPlayback //该影院的某电影在某一日的放映信息
}

// AllDate 存放放映日期
type AllDate struct {
	Day string
}
