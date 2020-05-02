package logedit

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type level int

const (
	UNKNOW level = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type log struct {
	Level level
}

func NewLoger(level string) log {
	lv := levelParse(level)
	return log{
		Level: lv,
	}
}
func levelParse(lv string) level {
	lv = strings.ToUpper(lv)
	switch lv {
	case "DEBUG":
		return DEBUG
	case "TRACE":
		return TRACE
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return UNKNOW
	}
}
func unLevelParse(Level level) string {
	switch Level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func (l *log) Debug(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if l.Level <= DEBUG {
		l.TerminalLog(logContent, DEBUG)
	}
	return nil
}
func (l *log) Trace(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if l.Level <= TRACE {
		l.TerminalLog(logContent, TRACE)
	}
	return nil
}
func (l *log) Warning(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if l.Level <= WARNING {
		l.TerminalLog(logContent, WARNING)
	}
	return nil
}
func (l *log) Info(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if l.Level <= INFO {
		l.TerminalLog(logContent, INFO)
	}
	return nil
}
func (l *log) Error(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if l.Level <= ERROR {
		l.TerminalLog(logContent, ERROR)
	}
	return nil
}
func (l *log) FATAL(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if l.Level <= FATAL {
		l.TerminalLog(logContent, FATAL)
	}
	return nil
}
func (l *log) TerminalLog(logContent string, lv level) {
	timeNow := currentTime()
	level := unLevelParse(lv)
	funcName, fileName, line, err := getLogInfo(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(os.Stdout, "[%s][%s][%s:%d:%s]:%s\n", timeNow, level, fileName, line, funcName, logContent)
}
func currentTime() string {
	return time.Now().Format("2006/01/02 15:04:05")
}
func getLogInfo(x int) (funcName, file string, line int, err error) {
	pc, file, line, ok := runtime.Caller(x)
	if !ok {
		return "", "", -1, errors.New("获取数据失败")
	}
	FuncPoint := runtime.FuncForPC(pc)
	funcName = FuncPoint.Name()
	funcName = strings.Split(funcName, ".")[1]
	return funcName, file, line, nil
}
