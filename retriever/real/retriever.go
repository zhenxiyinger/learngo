package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) Post(url string) string {
	panic("implement me")
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	response, err := httputil.DumpResponse(resp, true)
	err = resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(response)
}
