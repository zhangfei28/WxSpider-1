package router

import (
	"WxSpider/api/v1"

	"github.com/gin-gonic/gin"
)

func apiRouter(r *gin.Engine) {
	//爬虫爬取接口
	apiv1 := r.Group("/api/v1/spider")
	apiv1.Use()
	{
		apiv1.POST("/post", v1.SpiderArticle)
		apiv1.GET("/", v1.SpiderIndex)
	}
}
