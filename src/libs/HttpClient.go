package libs

import (
	"log"
	"net/http"
)

type IHttpClient interface {
	Get(url string) (*http.Response, error)
}

type HttpClient struct {
}


func (c HttpClient) Get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error in request")
	}

	return client.Do(req)
}