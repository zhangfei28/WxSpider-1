package logging

import (
	"WxSpider/pkg/file"
	"WxSpider/pkg/setting"
	"fmt"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppConfig.RuntimeRootPath+setting.AppConfig.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppConfig.LogSaveName,
		time.Now().Format(setting.AppConfig.TimeFormat),
		setting.AppConfig.LogFileExt,
	)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err :%v", err)
	}

	src := dir + "/" + filePath
	perm := file.HasPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Touch(src + fileName)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
