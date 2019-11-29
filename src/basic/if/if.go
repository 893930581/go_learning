package main

import "fmt"
import "io/ioutil"

const filename = "abc.txt"

func readfile() {
	if contents, err := ioutil.ReadFile(filename); contents != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//fmt.Println(err)//报错,if里定义的变量在外部没有作用域
}


func ifif() {
	if a := 1; a > 0 {
		fmt.Println("允许")
	}
}

func main() {

	readfile()
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
