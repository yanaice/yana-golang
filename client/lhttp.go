package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WrapHttpClient struct {
	Client *http.Client
}

func WrapHttpClientInit(c *http.Client) *WrapHttpClient {
	return &WrapHttpClient{
		Client: c,
	}
}

func (c *WrapHttpClient) Do(req *http.Request) (*http.Response, error) {

	// DO SOMETHING WITH REQUEST BODY
	var reqBodyBytes []byte
	if req.Body != nil {
		reqBodyBytes, _ = ioutil.ReadAll(req.Body)
		newBody := ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes)) // have to create new buffer
		req.Body = newBody
	}
	var reqBodyJson interface{}
	json.Unmarshal(reqBodyBytes, &reqBodyJson)

	// MAKING REQUEST
	resp, err := c.Client.Do(req)

	// DO SOMETHING WITH RESPONSE BODY
	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	newRespBody := ioutil.NopCloser(bytes.NewBuffer(respBodyBytes)) // have to create new buffer
	resp.Body = newRespBody

	var respBodyJson interface{}
	json.Unmarshal(respBodyBytes, &respBodyJson)

	// log.WithFields(log.Fields{
	// 	"## HTTPCLIENT ##": time.Now().Format("2006-01-02 15:04:05"),
	// 	"url":              req.URL.Path,
	// 	"reqHeader":        req.Header,
	// 	"reqBody":          reqBodyJson,
	// 	"respStatus":       resp.StatusCode,
	// 	"respBody":         respBodyJson,
	// 	"z":                "## END ##",
	// }).Info()

	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}
