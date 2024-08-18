package model

type MPage struct {
	Single       []*Movie //存放单个电影信息的
	AllMovies    []*Movie //所有电影（或者根据条件筛选的）（其实还是存放所有热门电影或者所有即将上映电影）
	PageNumber1  int64    //当前页
	PageSize1    int64    //每页显示的条数
	TotalPages1  int64    //总页数，通过计算得到
	TotalRecord1 int64    //总记录数，用过查询数据库得到
	IsLogin      bool     //是否登录
	Username     string   //用户名

	//影院相关
	Cinemas          []*Cinema          //放所有影院的信息，简易展示
	SameDateAndMovie []*CinemaPlayback  //该影院的某电影在某一日的放映信息，只展示某一日的某部电影
	MoveAndImgpath   []*MovieAndImgpath //因为影片一天会有很多场，记录每天都有什么电影就行
}
