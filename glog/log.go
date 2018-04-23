package glog

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	level int
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	depth int
}

var Glogger *Logger
var (
	loglevel       = 2
	defaultLogName = "./api_log/api.log"
	defaultLogSize = 100
)

func InitLogger() {
	initLoggerInstance()
}

func initLogConfig() *lumberjack.Logger {
	name := viper.GetString("log.logName")
	if name == "" {
		name = defaultLogName
	}
	size := viper.GetInt("log.logSize")
	if size == 0 {
		size = defaultLogSize
	}
	l := &lumberjack.Logger{
		Filename:  name,
		MaxSize:   size, // megabytes
		LocalTime: true,
		//MaxBackups: 3,
		//MaxAge:     28,   // days
		//Compress: true, // disabled by default
	}
	return l
}

func initLoggerInstance() {
	if Glogger == nil {
		f := initLogConfig()
		Glogger = new(Logger)
		flag := log.Ldate | log.Ltime | log.Lshortfile
		Glogger.err = log.New(f, "[ERROR]\t", flag)
		Glogger.warn = log.New(f, "[WARN]\t", flag)
		Glogger.info = log.New(f, "[INFO]\t", flag)
		Glogger.debug = log.New(f, "[DEBUG]\t", flag)
	}
}

func INFO(v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	//Glogger.info.Println(v)
	Glogger.info.Output(loglevel, fmt.Sprintln(v))
}

func INFOF(format string, v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	Glogger.info.Output(loglevel, fmt.Sprintf(format, v))
}

func DEBUG(v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	//Glogger.debug.Println(v)
	Glogger.debug.Output(loglevel, fmt.Sprintln(v))
}

func DEBUGF(format string, v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	Glogger.debug.Output(loglevel, fmt.Sprintf(format, v))
}

func ERROR(v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	Glogger.err.Output(loglevel, fmt.Sprintln(v))
}
func ERRORF(format string, v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	Glogger.err.Output(loglevel, fmt.Sprintf(format, v))
}

func WARN(v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	Glogger.warn.Output(loglevel, fmt.Sprintln(v))
}

func WARNF(format string, v ...interface{}) {
	if Glogger == nil {
		initLoggerInstance()
	}
	Glogger.warn.Output(loglevel, fmt.Sprintf(format, v))
}
