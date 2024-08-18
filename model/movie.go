package model

// Movie 热映电影&即将上映电影
type Movie struct {
	Id         int     //电影id
	Name       string  //电影名
	NickName   string  //别名
	Introduce  string  //电影描述
	Type       string  //电影类型
	Area       string  //区域
	Times      int     //时长（分钟）
	Year       string  //上映日期
	Starring   string  //主演
	Score      float64 //评分
	Grossed    int     //票房
	Award      string  //获奖
	AwardTimes int     //获奖次数
	ImgPath    string  //图片地址
}

//// ExpectMovie 即将上映的电影（最受期待）
//type ExpectMovie struct {
//	Id         int     //电影id
//	Name       string  //电影名
//	NickName   string  //别名
//	Introduce  string  //电影描述
//	Type       string  //电影类型
//	Area       string  //区域
//	Times      int     //时长（分钟）
//	Year       string  //上映日期
//	Starring   string  //主演
//	Score      float64 //评分
//	Expect     int     //想看人数
//	Award      string  //获奖
//	AwardTimes int     //获奖次数
//	ImgPath    string  //图片地址
//}

type Ranking1 struct {
	Id      int    //Id（排名）
	Name    string //电影名
	Grossed int    //票房
	Cls     string //class
	Movieid int    //电影的id
}

type Ranking2 struct {
	Id      int    //Id（排名）
	Name    string //电影名
	Expect  int    //想看人数
	Cls     string //class
	Movieid int    //电影的id
}

type MovieAndImgpath struct {
	Name      string //电影名
	Id        string //电影id
	Imgpath   string //电影图片
	Cinema_id string //影院id
}
