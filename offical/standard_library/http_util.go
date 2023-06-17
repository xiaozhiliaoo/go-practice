package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(headers map[string]string, url string) (data string, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, respErr := client.Do(req)
	return processResp(respErr, resp)
}

func Post(url string, params map[string]string) (data string, err error) {
	payload, _ := json.Marshal(params)
	resp, respErr := http.Post(url, "application/json", bytes.NewBuffer(payload))
	return processResp(respErr, resp)
}

func processResp(respErr error, resp *http.Response) (data string, err error) {
	if respErr != nil {
		return "", respErr
	}

	if resp == nil {
		return "", respErr
	}

	if resp.StatusCode != http.StatusOK {
		statusErr := fmt.Errorf("processResp rsp code error, code: %d", resp.StatusCode)
		return "", statusErr
	}

	defer resp.Body.Close()
	body, bodyErr := io.ReadAll(resp.Body)
	if err != nil {
		return "", bodyErr
	}
	return string(body), nil
}

func main() {
	data, err := Get(nil, "https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
