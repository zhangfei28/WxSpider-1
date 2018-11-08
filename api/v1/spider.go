//爬取主要接口
package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"

	"WxSpider/models"
	"WxSpider/pkg/app"
	"WxSpider/pkg/e"
	"WxSpider/pkg/logging"
	"WxSpider/pkg/splider"
	"WxSpider/pkg/validate"
)

//获取微信内容
type MpData struct {
	ID         int64  `json:"id"`
	Datetime   int64  `json:"datetime"`
	Digest     string `json:"digest"`
	ContentUrl string `json:"content_url"`
	Cover      string `json:"cover"`
}

func SpiderIndex(c *gin.Context) {
	c.String(http.StatusOK, "It work")
}

//公众号内容采集静态页面
func SpiderArticle(c *gin.Context) {
	var mpData []MpData
	var mp models.Mp
	var mpArticleList []models.MpArticle

	appG := app.Gin{c}
	status_code := e.SUCCESS
	mpArticle_chan := make(chan models.MpArticle)
	delayTime := time.Second * 1 //超时几秒，表示任务结束

	//获取公众号信息
	mp.Biz = c.PostForm("biz")
	mpid, err := mp.GetMpIdByBiz()
	if err != nil {
		status_code = e.INVALID_MP
		logging.Warn(e.GetMsg(status_code))
		appG.Response(http.StatusBadRequest, status_code, nil)
		return
	}

	//从请求中获取相应参数
	data := c.PostForm("mpdata")
	json.NewDecoder(strings.NewReader(data)).Decode(&mpData)

	if len(mpData) <= 0 {
		status_code = e.INVALID_PARAMS
		logging.Warn(e.GetMsg(status_code))
		appG.Response(http.StatusBadRequest, status_code, nil)
		return
	}

	for _, articleData := range mpData {
		go getMpArticle(articleData, mpid, mpArticle_chan)
	}

L:
	for {
		select {
		case article := <-mpArticle_chan:
			mpArticleList = append(mpArticleList, article)
		case <-time.After(delayTime):
			logging.Warn("timeout...")
			break L
		}
	}

	if len(mpArticleList) > 0 {
		logging.Info("执行成功，添加数据" + com.ToStr(len(mpArticleList)) + "条")
		models.AddMpArticleList(&mpArticleList)
	}

	appG.Response(http.StatusBadRequest, status_code, "执行成功，新增"+com.ToStr(len(mpArticleList))+"条数据")
}

//根据链接获取公众号的内容
func getMpArticle(data MpData, mpid int64, mpArticle_chan chan models.MpArticle) {
	//	err := checkDataFormat(&data)
	//	if err != nil {
	//		logging.Warn(err)
	//		return
	//	}
	log.Printf("%+v", data)
	article, err := splider.GetArticle(data.ContentUrl)
	if err != nil {
		logging.Warn("getMpArticle err:", err)
		return
	}

	var mpArticle models.MpArticle
	mpArticle.ID = data.ID
	mpArticle.Datetime = data.Datetime
	mpArticle.Digest = data.Digest
	mpArticle.Cover = data.Cover
	mpArticle.Content = article["content"]
	mpArticle.Title = article["title"]
	mpArticle.Author = article["author"]
	mpArticle.MpId = mpid
	log.Println(mpArticle)
	mpArticle_chan <- mpArticle
}

//采集微信公众号信息
func GetMpInfo(c *gin.Context) {
	var mp models.Mp
	appG := app.Gin{c}

	err := c.ShouldBind(&mp)
	if err != nil {
		logging.Warn("getMpInfo error:", err)
	}

	errs := validate.Check(&mp)
	if errs != nil {
		appG.Response(http.StatusBadRequest, 400, errs)
		return
	}

	if isExist := mp.IsExist(); isExist == true {
		appG.Response(http.StatusBadRequest, 400, "mp already exist")
	}

	if err := mp.AddMp(); err != true {
		appG.Response(http.StatusBadRequest, 400, "add mp error")
		return
	}

	appG.Response(http.StatusOK, 200, nil)
}
