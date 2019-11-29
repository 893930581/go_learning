package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	fmt.Println(n,err)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

type lion struct {
	age int
	name string
}

func (l lion) String() string {
	return fmt.Sprint("狮子名叫",l.name,"，今年",l.age,"岁了。")
}

func main(){
	s := strings.NewReader("123")
	data, err := ReadFrom(s,1)
	data2, err := ReadFrom(s,2)
	data3, err := ReadFrom(s,3)

	fmt.Println(data,err)
	fmt.Println(data2,err)
	fmt.Println(data3,err)
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)

	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)

	fmt.Printf("%%")
	fmt.Println()
	fmt.Println("单纯的输出百分号")

	fmt.Printf("%t",true)
	fmt.Println()
	fmt.Println("单纯的输出真假值（true/false）")

	fmt.Printf("%v","string")
	fmt.Println()
	fmt.Println("原格式输出")
	fmt.Printf("%T","string")
	fmt.Println()
	fmt.Println("打印数据类型")

	fmt.Printf("%d",12)
	fmt.Println()
	fmt.Println("十进制方法表示，只能传递int")
	fmt.Printf("%b",12)
	fmt.Println()
	fmt.Println("二进制方法表示，只能传递int")

	fmt.Printf("%q","string")
	fmt.Println()
	fmt.Println("双引号围绕的字符串，由Go语法安全地转义")

	fmt.Printf("%s", []byte("Go语言中文网"))
	fmt.Println()
	fmt.Println("输出字符串表示（string类型或[]byte)")

	fmt.Printf("%c", 0x4E2D)
	fmt.Println()
	fmt.Println("相应Unicode码点所表示的字符")

	xb := lion{name:"xinba",age:7}
	fmt.Print(xb)
	fmt.Printf("%T",fmt.Sprint("狮子名叫","，今年","岁了。"))
}