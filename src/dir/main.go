package main

import (
	"fmt"
	"io/ioutil"
)

func showALl(dirPath string){
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
	}
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			fmt.Print(fileInfo.Name(),"->")
			showALl(dirPath+"/"+fileInfo.Name())
		}else{
			fmt.Println(fileInfo.Name())
		}
	}
}

func main(){
	showALl("./test")
}
