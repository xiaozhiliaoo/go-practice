package main

import (
	"fmt"
	"github.com/avast/retry-go"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	var body []byte
	var retryCount uint64 = 0
	req, _ := http.NewRequest("GET", "https://www.baidu.com", nil)
	client := &http.Client{Timeout: 5 * time.Second}

	retryableFunc := func() error {
		retryCount += 1
		fmt.Printf("retry count:%d\n", retryCount)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("error1:%+v\n", err)
			return err
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error2:%+v\n", err)
			return err
		}
		return nil
	}
	err := retry.Do(retryableFunc, retry.Attempts(5), retry.Delay(1*time.Second))

	if err != nil {
		fmt.Printf("retry fails:%+v", err)
	}

	fmt.Printf("retry success:%+v", string(body))
}
