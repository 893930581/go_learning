package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)
type Retriever struct{
	UsertAgent string
	TimeOut time.Duration
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	return result[0]
	resp.Body.Close()

	if err !=nil {
		panic(err)
	}
}