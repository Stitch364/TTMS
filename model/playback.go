package model

//根据电影名，影院名查找该影院所有日期的该电影的播放时间

// CinemaPlayback 电影放映时间信息
type CinemaPlayback struct {
	Id        int     //某一次放映的Id
	CinemaId  int     //电影院ID
	Date      string  //放映日期
	Time      string  //放映时间
	Duration  string  //电影时长
	MovieId   int     //放映电影的Id
	MovieName string  //放映电影的名字
	Price     float64 //价格
	Number    int     //剩余票的数量（售票时要加锁）
	ImgPath   string  //电影的图片地址
	Hall      string  //影厅
}
