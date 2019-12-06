package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FIleLogger struct {
	level         logLevel
	filepath      string
	filename      string
	maxFIleSize	  int64
	errFIleObj    *os.File
	commonFileObj *os.File

}

func NewFIleLogger(loglev,fp,fn string, maxSize int64) *FIleLogger {
	lv,err := ConvertStrLevel(loglev)
	if err != nil {
		panic(err)
	}
	f1 := &FIleLogger{
		level:         lv,
		filepath:      fp,
		filename:      fn,
		commonFileObj: nil,
		errFIleObj:    nil,
		maxFIleSize:   maxSize,
	}
	initErr := f1.initFile()
	if initErr != nil {
		panic(initErr)
	}
	return f1
}

func (f *FIleLogger) splitFile (file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n",err)
		return nil, err
	}else if(fileInfo.Size() < f.maxFIleSize){
		return nil, err
	}
	logName := path.Join(f.filepath,fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s",logName,nowStr)

	file.Close()
	os.Rename(logName,newLogName)

	fileObj, err := os.OpenFile(logName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n",err)
		return nil, err
	}
	return fileObj, nil
}


func (f *FIleLogger) initFile() error {
	fullPath := path.Join(f.filepath,f.filename)
	fileObj, err := os.OpenFile(fullPath,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error:%v\n",err)
		return err
	}
	errorFileObj,err := os.OpenFile(fullPath+".err",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error:%v\n",err)
		return err
	}

	f.commonFileObj = fileObj
	f.errFIleObj = errorFileObj
	return nil
}

func (f FIleLogger) Debug(format string,a ...interface{}){
	if f.level <= DEBUG {
		log(f.commonFileObj,DEBUG,format,a...)
	}
}

func (f FIleLogger) Trace(format string,a ...interface{}){
	if f.level <= TRACE {
		log(f.commonFileObj,TRACE,format,a...)
	}
}

func (f FIleLogger) Info(format string,a ...interface{}){
	if f.level <= INFO {
		log(f.commonFileObj,INFO,format,a...)
	}
}

func (f FIleLogger) Warning(format string,a ...interface{}){
	if f.level <= WARNING {
		log(f.commonFileObj,WARNING,format,a...)
	}
}

func (f FIleLogger) Error(format string,a ...interface{}){
	if f.level <= ERROR {
		log(f.commonFileObj,ERROR,format,a...)
		log(f.errFIleObj,ERROR,format,a...)
	}
}

func (f FIleLogger) Fatal(format string,a ...interface{}){
	if f.level <= FATAL {
		log(f.commonFileObj,FATAL,format,a...)
		log(f.errFIleObj,ERROR,format,a...)
	}
}

func (f FIleLogger) CloseFIle (){
	f.commonFileObj.Close()
	f.errFIleObj.Close()
}