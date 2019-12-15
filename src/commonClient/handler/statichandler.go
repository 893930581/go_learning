package handler

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}


const prefix string = "/static/"

func StaticFileHandler(res http.ResponseWriter,req *http.Request) error {
		if strings.Index(req.URL.Path,prefix) != 0 {
			return userError(fmt.Sprintf("Path %v must start with %v",req.URL.Path,prefix))
		}

		filePath := req.URL.Path
		//获取所取路径
		appPath, err := os.Executable()
		// 获得程序路径(application)
		wantedFile := filepath.Join(filepath.Dir(appPath),filePath)
		//拼接为文件路径
		fmt.Println("user visit \n","->",wantedFile)

		fr, err := os.Open(wantedFile)
		if err != nil {
			return err
		}

		defer fr.Close()
		fileContent, err := ioutil.ReadAll(fr)
		if err != nil {
			return err
		}

		SetContentTypeFromExtension(res,path.Ext(filePath))
		//根据文件类型设置请求头。

		_,err = res.Write(fileContent)
		if err != nil {
			return err
		}
		return nil
}

func SetContentTypeFromExtension(w http.ResponseWriter,extension string){
		mimeType := mime.TypeByExtension(extension)
		if mimeType != "" {
			w.Header().Set("Content-Type",mimeType)
		}
}