package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestGetCurl(t *testing.T) {

	secretId := "XXXX"
	secretKey := "XXXX"
	host := "account.tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"
	version := "2018-12-25"
	action := "GetAuthInfoByUin"
	var timestamp = time.Now().Unix()
	// step 1: build canonical request string
	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\n" + "host:" + host + "\n"
	signedHeaders := "content-type;host"
	// 请求参数，必须和文档一样，
	payload := `{"XXX":"XXXX","XXXX":[{"XXX":"XXXX","XXX":"XXXXX"}]}`
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	fmt.Println(canonicalRequest)

	// step 2: build string to sign
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)
	fmt.Println(string2sign)

	// step 3: sign string
	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	fmt.Println(signature)

	// step 4: build authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature)
	fmt.Println(authorization)

	fmt.Println("---------------CURL CMD----------------------")
	curl := fmt.Sprintf(`curl -X POST https://%s\
 -H "Authorization: %s"\
 -H "Content-Type: application/json; charset=utf-8"\
 -H "Host: %s" -H "X-TC-Action: %s"\
 -H "X-TC-Timestamp: %d"\
 -H "X-TC-Version: %s"\
 -H "X-TC-Region: %s"\
 -d '%s'`, host, authorization, host, action, timestamp, version, region, payload)
	fmt.Println(curl)
	fmt.Println("---------------DO REQUEST----------------------")

	request, _ := http.NewRequest("POST", "https://tione.tencentcloudapi.com", bytes.NewBuffer([]byte(payload)))

	request.Header.Add("Authorization", authorization)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Host", "https://XXXX.tencentcloudapi.com")
	request.Header.Add("X-TC-Timestamp", strconv.FormatInt(timestamp, 10))
	request.Header.Add("X-TC-Version", version)
	request.Header.Add("X-TC-Action", action)
	request.Header.Add("X-TC-Region", region)

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Printf("error: %+v\n", err)
		assert.Fail(t, "fail")
	}
	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("resp body:%s", string(respBody))
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
