package filelisten

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter,request *http.Request){
	path := request.URL.Path[len("/list/"):]
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	writer.Write(all)
}