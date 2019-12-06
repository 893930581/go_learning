package mylogger

import "os"

type ConsoleLogger struct {
	level logLevel
}

func NewConsoleLogger(loglev string) (ConsoleLogger) {
	if lv,err := ConvertStrLevel(loglev); err!=nil{
		panic(err)
	}else{
		return ConsoleLogger{lv}
	}
}


func (l ConsoleLogger) Debug(format string,a ...interface{}){
	if l.level <= DEBUG {
		log(os.Stdout,DEBUG,format,a...)
	}
}

func (l ConsoleLogger) Trace(format string,a ...interface{}){
	if l.level <= TRACE {
		log(os.Stdout,TRACE,format,a...)
	}
}

func (l ConsoleLogger) Info(format string,a ...interface{}){
	if l.level <= INFO {
		log(os.Stdout,INFO,format,a...)
	}
}

func (l ConsoleLogger) Warning(format string,a ...interface{}){
	if l.level <= WARNING {
		log(os.Stdout,WARNING,format,a...)
	}
}

func (l ConsoleLogger) Error(format string,a ...interface{}){
	if l.level <= ERROR {
		log(os.Stdout,ERROR,format,a...)
	}
}

func (l ConsoleLogger) Fatal(format string,a ...interface{}){
	if l.level <= FATAL {
		log(os.Stdout,FATAL,format,a...)
	}
}
