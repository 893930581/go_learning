package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func WriteTo(writer io.Writer,con []byte) error {
	n, err := writer.Write(con)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	return err
}

func writeBufio(writer io.Writer,con []byte) {
	bw := bufio.NewWriter(writer)
	defer bw.Flush()
	fmt.Fprintf(bw,string(con))
}

func writeFileBufio(filename string,content []uint8){
	frw, err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(frw)
	defer bw.Flush()
	fmt.Fprintf(bw,string(content))
}

func ReadFrom(reader io.Reader,num int) ([]byte,error){
	con := make([]byte,num)
	n, err := reader.Read(con)
	fmt.Println("读取的字节数:",n)
	if err != nil && err!=io.EOF {
		panic(err)
	}
	return con,err
}

func readBufio(reader io.Reader) ([]byte,error) {
	br := bufio.NewReader(reader)
	line, err := br.Peek(100)
	fmt.Println(line)
	return line, err
}

func ReadFromFile(filename string, num int) ([]byte,error) {
	con := make([]byte,num)
	fr, err := os.Open(filename)
	if err != nil {
		return nil,err
	}
	_, err = fr.Read(con)
	if err != nil {
		return nil,err
	}
	return con,err
}

func main(){
	resp, err := http.Get("http://127.0.0.1:8888/getData")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//body, err := ReadFrom(resp.Body,20)
	fmt.Printf("%T -> %v",body,body)
	fmt.Println(string(body))
	fw,_ := os.Create("./nodejs.txt")
	WriteTo(fw,body)
	fr,_ := os.Open("./nodejs.txt")
	fc,_ := readBufio(fr)
	fmt.Printf("%s",fc)
	//writeFileBufio("./nodejs.txt",body)
	//err = ioutil.WriteFile("./nodejs.txt",body,0666)
	//if err != nil {
	//	panic(err)
	//}
}