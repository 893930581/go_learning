package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	s := strings.NewReader("aaa")
	b,_ := ioutil.ReadAll(s)
	fmt.Printf("%s",b)

	f,_ := os.Open("1.txt")
	b,_ = ioutil.ReadAll(f)
	fmt.Printf("%s",b)
}