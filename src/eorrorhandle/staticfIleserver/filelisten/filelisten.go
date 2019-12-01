package filelisten

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix string = "/list/"

func HandleFileList(writer http.ResponseWriter,request *http.Request) error{
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError(
			fmt.Sprintf("path %s must start "+
				"with %s",
				request.URL.Path, prefix))
	}

	path := request.URL.Path[len(prefix):]
	fmt.Println(path)
	fmt.Println("用户访问go_learning/static/"+path)
	file, err := os.Open("go_learning/static/"+path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}