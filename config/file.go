package config

import "os"

const filePathENV = "CONFIG_FILE_PATH"
const defaultFilePath = "./config.yaml"

// getFilePath 获得配置文件路径
func getFilePath() string {
	if configFilePath := os.Getenv(filePathENV); configFilePath != "" {
		return configFilePath
	}
	return defaultFilePath
}
