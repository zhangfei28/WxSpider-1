package models

import (
	"github.com/Unknwon/com"
)

type MpArticle struct {
	ID       int64
	Title    string `form:"title" valid:"Required"`
	MpId     int64  `valid:"Required"`
	Content  string `valid:"Required"`
	Author   string
	Cover    string
	Datetime int64
	Digest   string
}

//批量添加文章
func AddMpArticleList(articleList *[]MpArticle) error {
	//使用原生sql添加数据
	var sql string = "INSERT INTO `wx`.`wx_article` (`id`, `title`, `mp_id`,`content`, `author`, `cover`, `datetime`, `digest`) VALUES"
	for _, article := range *articleList {
		sql += "(" + com.ToStr(article.ID, 10) + ",'"
		sql += article.Title + "',"
		sql += com.ToStr(article.ID, 10) + ",'"
		sql += article.Content + "','"
		sql += article.Author + "','"
		sql += article.Cover + "',"
		sql += com.ToStr(article.Datetime) + ",'"
		sql += article.Digest + "')"
	}

	db.Exec(sql)

	return nil
}
