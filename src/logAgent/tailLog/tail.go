package tailLog

import (
	"fmt"
	"github.com/hpCloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:		true,
		Follow:		true,
		Location:	&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:	false,
		Poll:		true,
	}
	tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v",err)
		return
	}
	return
}

func ReadLog() <-chan *tail.Line{
	return tailObj.Lines
}