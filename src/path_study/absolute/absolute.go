package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	execpath, _ := os.Executable() // 获得程序路径
	configfile := filepath.Join(filepath.Dir(execpath),"./../../1.txt")
	fmt.Println(execpath)
	fmt.Println(configfile)
	dir := filepath.Join(filepath.Dir(execpath))
	base := filepath.Join(filepath.Base(execpath))
	fmt.Println(dir)
	fmt.Println(base)
	ppath := "c:/asd/../../dsad"
	fmt.Println(filepath.Clean(ppath))
}
