package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"time"
)

func AddText(text *model.Text) error {
	text.Time = time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "INSERT INTO comments_1(movie_id, username, text, object,time,objectid) VALUES (?,?,?,?,?,?);"
	_, err := utils.Db.Exec(sqlStr, text.Movieid, text.Username, text.Text, text.Object, text.Time, text.ObjectId)
	if err != nil {
		fmt.Println("AddText函数增加评论时出错了", err)
		return err
	}
	return nil
}

// GetText 获取所有评论
func GetText(movieid string) ([]*model.Text, error) {
	var texts []*model.Text
	sqlStr := "SELECT * FROM comments_1 WHERE movie_id=? order by time desc;"
	rows, err := utils.Db.Query(sqlStr, movieid)
	if err != nil {
		fmt.Println("获取所有评论时出错了：", err)
		return nil, err
	}
	for rows.Next() {
		text := new(model.Text)
		err = rows.Scan(&text.Id, &text.Movieid, &text.Username, &text.Text, &text.Object, &text.Time, &text.ObjectId)
		if err != nil {
			fmt.Println("读取所有评论时出错了：", err)
			return nil, err
		}
		texts = append(texts, text)
	}
	return texts, nil
}

//// GetText 获取所有评论（顶级评论和回复评论）
//func GetText(movieid string) ([]*model.Text, error) {
//	var texts []*model.Text
//	//sqlStr := "SELECT * FROM comments_1 WHERE movie_id=? order by time desc;"
//
//	sqlStr := `
//WITH RECURSIVE CommentTree AS (
//  SELECT
//		c.Id,
//		c.movie_id,
//		c.username,
//		c.text,
//		c.object,
//		c.time,
//		c.objectid,
//      1 AS depth
//  FROM
//      comments_1 c
//  WHERE
//      c.objectid IS NULL AND movie_id = ?
//
//  UNION ALL
//
//  SELECT
//      c.id,
//      c.movie_id,
//      c.username,
//      c.text,
//      c.object,
//      c.time,
//      c.objectid,
//      ct.depth + 1
//  FROM
//      comments_1 c
//  INNER JOIN CommentTree ct ON c.objectid = ct.id
//)
//SELECT * FROM CommentTree
//ORDER BY  depth, time;
//`
//
//	rows, err := utils.Db.Query(sqlStr, movieid)
//	if err != nil {
//		fmt.Println("获取所有评论时出错了：", err)
//		return nil, err
//	}
//	for rows.Next() {
//		text := new(model.Text)
//		err = rows.Scan(&text.Id, &text.Movieid, &text.Username, &text.Text, &text.Object, &text.Time)
//		if err != nil {
//			fmt.Println("读取所有评论时出错了：", err)
//			return nil, err
//		}
//		texts = append(texts, text)
//	}
//	fmt.Println("texts:", texts)
//	return texts, nil
//}
