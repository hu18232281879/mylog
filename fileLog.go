package logedit

import (
	"fmt"
	"os"
	"time"
)

type fileLoger struct {
	Level       level
	filePath    string
	fileName    string
	errFileName string
	fileSzie    int64
}

//NewFileLoger... 构造函数,返回日志指针
func NewFileLoger(level string, filePath, fileName string, fileSize int64) *fileLoger {
	lv := levelParse(level)
	return &fileLoger{
		Level:       lv,
		filePath:    filePath,
		fileName:    fileName,
		errFileName: fileName + ".err",
		fileSzie:    fileSize,
	}
}
func (f *fileLoger) Debug(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if f.Level <= DEBUG {
		f.WriteLogToFile(logContent, DEBUG)
	}
	return nil
}

func (f *fileLoger) Trace(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if f.Level <= TRACE {
		f.WriteLogToFile(logContent, DEBUG)
	}
	return nil
}
func (f *fileLoger) Warning(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if f.Level <= WARNING {
		f.WriteLogToFile(logContent, WARNING)
	}
	return nil
}
func (f *fileLoger) Info(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if f.Level <= INFO {
		f.WriteLogToFile(logContent, INFO)
	}
	return nil
}
func (f *fileLoger) Error(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if f.Level <= ERROR {
		f.WriteLogToFile(logContent, ERROR)
	}
	return nil
}
func (f *fileLoger) FATAL(format string, a ...interface{}) error {
	logContent := fmt.Sprintf(format, a...)
	if f.Level <= FATAL {
		f.WriteLogToFile(logContent, FATAL)
	}
	return nil
}
func (f *fileLoger) WriteLogToFile(logContent string, lv level) {
	file, err := os.OpenFile(f.filePath+f.fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if lv >= ERROR {
		errfile, err := os.OpenFile(f.filePath+f.errFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("错误日志文件打开失败", err)
			return
		}
		defer errfile.Close()

		ok := f.getfileStatus(errfile)
		if ok {
			newerrFilePor := f.cuttingFile(errfile, ".err")
			f.writeLogOrErrorLog(newerrFilePor, lv, logContent)
		} else {
			f.writeLogOrErrorLog(errfile, lv, logContent)
		}

	}
	ok := f.getfileStatus(file)
	if ok {
		newFilepor := f.cuttingFile(file, "")
		f.writeLogOrErrorLog(newFilepor, lv, logContent)
	} else {
		f.writeLogOrErrorLog(file, lv, logContent)
	}

}
func (f *fileLoger) writeLogOrErrorLog(filePointer *os.File, lv level, logContent string) {
	timeNow := currentTime()
	level := unLevelParse(lv)
	funcName, fileName, line, err := getLogInfo(4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(filePointer, "[%s][%s][%s:%d:%s]:%s\n", timeNow, level, fileName, line, funcName, logContent)
	filePointer.Close()
}

func (f *fileLoger) getfileStatus(filePointer *os.File) (status bool) {
	fileInfo, err := filePointer.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败")
		return
	}
	if fileInfo.Size() >= f.fileSzie {
		return true
	}
	return false

}
func (f *fileLoger) cuttingFile(filepor *os.File, iserr string) (newfilepor *os.File) {
	filepor.Close()
	timeformat := time.Now().Format("20060102150405000")
	os.Rename(f.filePath+f.fileName+iserr, f.filePath+f.fileName+iserr+timeformat)
	newfilepor, err := os.OpenFile(f.filePath+f.fileName+iserr, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	return newfilepor
}
