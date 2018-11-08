package splider

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//通过链接获取微信文章
func GetArticle(url string) (map[string]string, error) {
	body, reponse_status := HttpGet(url, "local")
	if reponse_status != 200 {
		return nil, errors.New("响应错误")
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	article := make(map[string]string)

	article["title"] = strings.TrimSpace(doc.Find("title").Text())
	article["author"] = strings.Replace(doc.Find("#meta_content .rich_media_meta_text").Text(), " ", "", -1)
	article["content"], err = doc.Find("#js_content").Html()
	if err != nil {
		return nil, errors.New("获取文章内容失败")
	}

	return article, nil
}
