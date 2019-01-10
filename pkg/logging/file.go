package logging

import (
	"fmt"
	"time"

	"github.com/negaihoshi/daigou/pkg/setting"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppConfig.App.RuntimeRootPath, setting.AppConfig.App.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppConfig.App.LogSaveName,
		time.Now().Format(setting.AppConfig.App.TimeFormat),
		setting.AppConfig.App.LogFileExt,
	)
}

//func openLogFile(fileName, filePath string) (*os.File, error) {
//	dir, err := os.Getwd()
//	if err != nil {
//		return nil, fmt.Errorf("os.Getwd err: %v", err)
//	}
//
//	src := dir + "/" + filePath
//	perm := file.CheckPermission(src)
//	if perm == true {
//		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
//	}
//
//	err = file.IsNotExistMkDir(src)
//	if err != nil {
//		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
//	}
//
//	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
//	}
//
//	return f, nil
//}
