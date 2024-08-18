package model

// Page 首页Page 结构，可能亦适用于所有电影界面
type Page struct {
	//电影相关
	Single       []*Movie    //存放单个电影信息的
	Movies       []*Movie    //热门电影
	ExpectMovies []*Movie    //最受期待的电影
	AllMovies    []*Movie    //所有电影（或者根据条件筛选的）（其实还是存放所有热门电影或者所有即将上映电影）
	MovieRank1   []*Ranking1 //今日票房电影排行
	MovieRank2   []*Ranking2 //最受期待电影排行
	PageNumber   int64       //当前页
	PageSize     int64       //每页显示的条数
	TotalPages   int64       //总页数，通过计算得到
	TotalRecord  int64       //总记录数，用过查询数据库得到
	IsLogin      bool        //是否登录
	Username     string      //用户名

	//影院相关
	Cinemas []*Cinema //放所有影院的信息，简易展示
	Cinema1 []*Cinema //其实只放一个电影院的信息,用于展示详细电影信息
	//	Playback         []*CinemaPlayback //该影院所有放映信息的切片
	AllDate          []*AllDate         //该影院都哪些天有电影
	SameDateAndMovie []*CinemaPlayback  //该影院的某电影在某一日的放映信息，只展示某一日的某部电影
	MoveAndImgpath   []*MovieAndImgpath //因为影片一天会有很多场，记录每天都有什么电影就行

	//MovieId string //电影Id
	//用户订单相关
	Goods []*Goods
	//电影评论
	Comments []*Text
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNumber > 1
}

// IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNumber < p.TotalPages
}

// GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNumber - 1
	}
	return 1
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNumber + 1
	}
	return p.TotalPages

}
