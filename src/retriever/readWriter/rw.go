package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile (filename string){
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContent(file)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main(){
	s := `adbas
			dasda
				dasdas`
	printFileContent(strings.NewReader(s))
}