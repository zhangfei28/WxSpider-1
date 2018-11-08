package logging

import (
	"WxSpider/pkg/setting"
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetLogFilePath(t *testing.T) {
	setting.InitSetUp("../../conf/app.ini")
	log.Printf(getLogFilePath())
	Convey("两个文件的路径应该是相同的", t, func() {
		So(getLogFilePath(), ShouldEqual, setting.AppConfig.RuntimeRootPath+setting.AppConfig.LogSavePath)
	})
}

func TestGetLogFileName(t *testing.T) {
	setting.InitSetUp("../../conf/app.ini")
	log.Printf("%s", getLogFileName())
	Convey("能够获取到日志的文件名称", t, func() {
		So(getLogFileName(), ShouldEqual, fmt.Sprintf("%s%s.%s",
			setting.AppConfig.LogSaveName,
			time.Now().Format(setting.AppConfig.TimeFormat),
			setting.AppConfig.LogFileExt,
		))
	})
}
