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
	PublisherKey = "your publisher key"
	ContentType  = "application/json"
)

func main() {
	// POST example:
	reqUrl := `https://openapi.toponad.com/v1/apps`
	reqBody := `{"limit":1}`
	respData, err := doRequest(http.MethodPost, reqUrl, reqBody)
	if err != nil {
		panic(err)
	}
	fmt.Println("post response: ", respData)

	// GET example:
	reqUrl = `https://openapi.toponad.com/v1/waterfall/units?placement_id='xxx'`
	respData, err = doRequest(http.MethodGet, reqUrl, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("get response: ", respData)
}

// doRequest requests the given reqUrl and returns the response.
func doRequest(reqMethod, reqUrl, reqBody string) (string, error) {
	urlInfo, err := url.Parse(reqUrl)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(reqMethod, reqUrl, bytes.NewReader([]byte(reqBody)))
	if err != nil {
		return "", err
	}

	// create the MD5 value of the request body.
	h := md5.New()
	h.Write([]byte(reqBody))
	contentMD5 := hex.EncodeToString(h.Sum(nil))
	contentMD5 = strings.ToUpper(contentMD5)
	// create the content of common headers except the header 'X-Up-Signature'.
	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	headerString := fmt.Sprintf("X-Up-Key:%s\nX-Up-Timestamp:%s", PublisherKey, timestamp)

	// create the final signature.
	sign := genSignature(reqMethod, contentMD5, ContentType, headerString, urlInfo.Path)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("X-Up-Key", PublisherKey)
	req.Header.Set("X-Up-Timestamp", timestamp)
	req.Header.Set("X-Up-Signature", sign)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respData), nil
}

// genSignature returns the final signature
func genSignature(httpMethod, contentMD5, contentType, headerString, relativePath string) string {
	stringSection := []string{httpMethod, contentMD5, contentType, headerString, relativePath}
	stringToSign := strings.Join(stringSection, "\n")
	h := md5.New()
	h.Write([]byte(stringToSign))
	resultMD5 := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(resultMD5)
}
