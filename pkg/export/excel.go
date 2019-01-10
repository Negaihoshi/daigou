package export

import "github.com/negaihoshi/daigou/pkg/setting"

const EXT = ".xlsx"

func GetExcelFullUrl(name string) string {
	return setting.AppConfig.App.PrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppConfig.App.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.AppConfig.App.RuntimeRootPath + GetExcelPath()
}
