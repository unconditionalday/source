package netx

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) Download(src string) ([]byte, error) {
	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
