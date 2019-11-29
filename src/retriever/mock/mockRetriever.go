package mock

import "fmt"

type Http struct {
	Contents string
}

func (r *Http) Get(url string) string {
	return r.Contents
}

func (r *Http) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Http) String() string {
	return fmt.Sprintf("Http:{Contents=%s}",r.Contents)
}