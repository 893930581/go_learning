package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}

func openIt(filename string){
	f,err := os.Open(filename)
	if err != nil {

	}
	defer f.Close()
	if err != nil {
		panic(err)
	}
	b := make([]byte,10)
	f.Read(b)
	n,err := f.Read(b)
	fmt.Println(n,err)
	fmt.Println(fmt.Sprintf("%s",b))
}

func openItAt(filename string,off int64){
	f,err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	b := make([]byte,10)
	f.ReadAt(b,off)
	fmt.Println(fmt.Sprintf("%s",b))
}

func writeIt(b []byte,filename string){
	f,err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	n,err := f.Write(b)
	fmt.Println(n,err)
}

func writeItBufio(filename string){
	f,_ := os.Create(filename)
	fmt.Printf("%T",f)
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	for i:=0;i<10;i++ {
		fmt.Fprintln(writer,i)
	}
}

func tryDefer2(){
	for i:=0;i<10;i++ {
		defer fmt.Println(i)
	}
}

func main() {
	b := []byte("jsioadj")
	filePath := "C:\\GOPKG\\src\\go_learning\\src\\eorrorhandle\\test.txt"
	writeIt(b,filePath)
	writeItBufio(filePath)
	openIt(filePath)
	openItAt(filePath,3)
	tryDefer()
	fmt.Fprint(os.Stdout,"吃饭真相")
	//var a string = "1234"
	tryDefer2()
}
