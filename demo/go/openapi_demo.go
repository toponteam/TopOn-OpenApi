package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	PublisherKey = "publisher key"
	ContentType  = "application/json"
)

func main() {
	reqUrl := `https://openapi.toponad.com/v1/apps`
	reqContent := `` // request content

	respData, err := dealFunc(http.MethodPost, reqUrl, reqContent)
	if err != nil {
		panic(err)
	}
	fmt.Println(respData)
}

// request and response content
func dealFunc(reqMethod, reqUrl, reqContent string) (string, error) {
	urlInfo, err := url.Parse(reqUrl)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(reqMethod, reqUrl, bytes.NewReader([]byte(reqContent)))
	if err != nil {
		return "", err
	}

	h := md5.New()
	h.Write([]byte(reqContent))
	contentMD5 := hex.EncodeToString(h.Sum(nil))
	contentMD5 = strings.ToUpper(contentMD5)
	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	headerString := fmt.Sprintf("X-Up-Key:%s\nX-Up-Timestamp:%s", PublisherKey, timestamp)

	sign := signatureFunc(reqMethod, contentMD5, ContentType, headerString, urlInfo.Path)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("X-Up-Key", PublisherKey)
	req.Header.Set("X-Up-Timestamp", timestamp)
	req.Header.Set("X-Up-Signature", sign)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respByte), nil
}

func signatureFunc(httpMethod, contentMD5, contentType, headerString, urlPath string) string {
	stringSection := []string{httpMethod, contentMD5, contentType, headerString, urlPath}
	stringToSign := strings.Join(stringSection, "\n")
	h := md5.New()
	h.Write([]byte(stringToSign))
	resultMD5 := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(resultMD5)
}
