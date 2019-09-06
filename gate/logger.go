package main

import (
	"app/common/serverConf"
	"fmt"
	"github.com/phachon/go-logger"
	"os"
)

var logger = go_logger.NewLogger()

func init() {
	logLevel := logger.LoggerLevel(serverConf.GetServerConf().LogLevel)

	logger.Detach("console")

	// 配置 console adapter
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}
	// 添加输出到命令行
	logger.Attach("console", logLevel, consoleConfig)

	logDir := "./log"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Printf("create dir fail, %v\n", err)
		panic("")
	}

	filePath := logDir + "/" + serverName + ".log"
	errFilePath := logDir + "/" + serverName + "_error.log"

	// 配置 file adapter
	fileConfig := &go_logger.FileConfig{
		Filename: filePath, // 所有日志输出文件名，不存在自动创建
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): errFilePath, // Error 级别日志被写入 error .log 文件
		},
		MaxSize:    1024 * 1024, // 文件最大值（KB），默认值0不限
		MaxLine:    100000,      // 文件最大行数，默认 0 不限制
		DateSlice:  "d",         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,       // 写入文件的数据是否 json 格式化
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}

	// 添加输出到文件
	logger.Attach("file", logLevel, fileConfig)
}
