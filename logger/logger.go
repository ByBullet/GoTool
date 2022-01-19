package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	1. 日志级别
	2. 日志输出方向
	3. 选择需要输出的日志级别
	4. 不同级别的日志输出到不同文件
*/

//日志级别
const (
	LOGGER_ERROR = iota
	LOGGER_WARNING
	LOGGER_INFO
	LOGGER_DEBUG
)

//日志输出方式
const (
	LOGGER_CONSOLE = iota + 4
	LOGGER_FILE
)

/*
	日志配置信息
*/
type Config struct {
	// 日志级别 对于LOGGER_DEBUG，LOGGER_INFO，LOGGER_WARNING，LOGGER_ERROR
	LoggerLevel int
	//日志输出文件夹
	OutDir string
	//日志输出方式 LOGGER_CONSOLE,LOGGER_FILE
	OutType int
}

var loggerLevel = LOGGER_DEBUG

var loggerList = []*log.Logger{
	log.New(os.Stdout, "error> ", log.LstdFlags|log.Lshortfile),
	log.New(os.Stdout, "warn> ", log.LstdFlags|log.Lshortfile),
	log.New(os.Stdout, "info> ", log.LstdFlags|log.Lshortfile),
	log.New(os.Stdout, "debug> ", log.LstdFlags|log.Lshortfile),
}

var logFileNames = []string{"error.log", "warn.log", "info.log", "debug.log"}

/*
	获取项目所在根路径
	通过闭包缓存路径
*/
func getCurrentPath() (dir string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	return
}

/*
	初始化配置
*/
func SetConfig(config Config) {
	loggerLevel = config.LoggerLevel
	if config.OutType == LOGGER_FILE {
		for i := 0; i <= config.LoggerLevel; i++ {
			f, err := os.Create(fmt.Sprintf("%s/%s/%s", getCurrentPath(), config.OutDir, logFileNames[i]))
			if err != nil {
				loggerList[LOGGER_ERROR].Println(err)
				continue
			}
			loggerList[i].SetOutput(f)
		}
	}
}

func println(lv int, v ...interface{}) {
	if lv <= loggerLevel {
		loggerList[lv].Output(3, fmt.Sprintln(v...))
	}
}

func printf(lv int, format string, v ...interface{}) {
	if lv <= loggerLevel {
		loggerList[lv].Output(3, fmt.Sprintf(format, v...))
	}
}

//println
func Error(v ...interface{}) {
	println(LOGGER_ERROR, v...)
}
func Warn(v ...interface{}) {
	println(LOGGER_WARNING, v...)
}
func Info(v ...interface{}) {
	println(LOGGER_INFO, v...)
}
func Debug(v ...interface{}) {
	println(LOGGER_DEBUG, v...)
}

//format
func ErrorFormat(format string, v ...interface{}) {
	printf(LOGGER_ERROR, format, v...)
}
func WarnFormat(format string, v ...interface{}) {
	printf(LOGGER_WARNING, format, v...)
}
func InfoFormat(format string, v ...interface{}) {
	printf(LOGGER_INFO, format, v...)
}
func DebugFormat(format string, v ...interface{}) {
	printf(LOGGER_DEBUG, format, v...)
}
