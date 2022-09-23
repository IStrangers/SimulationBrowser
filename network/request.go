package network

import (
	"browser"
	"io"
	"log"
	"net/http"
	"network/structs"
	"strings"
)

var httpClient = &http.Client{}
var defaultHeaders = map[string]string{
	"User-Agent": browser.WebBrowserName + "/" + browser.WebBrowserVersion,
}

func SendGetRequest(url string) (*structs.RequestResult, error) {
	return SendRequest("GET", url, nil)
}

func SendPostRequest(url string, body string) (*structs.RequestResult, error) {
	return SendRequest("POST", url, strings.NewReader(body))
}

func SendRequest(method string, url string, body io.Reader) (*structs.RequestResult, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalln(err)
	}
	AddRequestHeaders(request, defaultHeaders)
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	requestResult := &structs.RequestResult{
		Body:        res,
		ContentType: response.Header.Get("Content-Type"),
		StatusCode:  response.StatusCode,
	}
	return requestResult, nil
}

func AddRequestHeader(request *http.Request, key string, value string) {
	request.Header.Set(key, value)
}

func AddRequestHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		AddRequestHeader(request, key, value)
	}
}