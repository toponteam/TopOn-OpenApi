package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	body1 := ""
	request("https://openapi.toponad.com/v1/waterfall/units?placement_id=xxxxxx", body1, "GET")

}

func request(demoUrl string, body string, httpMethod string) {
	//your publisherKey
	publisherKey := ""
	//request method
	//httpMethod := "POST"
	contentType := "application/json"
	publisherTimestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headers := map[string]string{
		"X-Up-Timestamp": publisherTimestamp,
		"X-Up-Key":       publisherKey,
	}
	//queryPath
	urlParsed, err := url.Parse(demoUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	//resource
	resource := urlParsed.Path

	//body
	h := md5.New()
	h.Write([]byte(body))
	contentMD5 := hex.EncodeToString(h.Sum(nil))
	contentMD5 = strings.ToUpper(contentMD5)

	publisherSignature := signature(httpMethod, contentMD5, contentType, headerJoin(headers), resource)

	request, err := http.NewRequest(httpMethod, demoUrl, bytes.NewReader([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return
	}
	client := &http.Client{}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-Up-Key", publisherKey)
	request.Header.Set("X-Up-Signature", publisherSignature)
	request.Header.Set("X-Up-Timestamp", publisherTimestamp)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return
	}

	//return
	fmt.Println(string(content))
}

func headerJoin(headers map[string]string) string {
	headerKeys := []string{
		"X-Up-Timestamp",
		"X-Up-Key",
	}
	sort.Strings(headerKeys)
	ret := make([]string, 0)
	for _, k := range headerKeys {
		v := headers[k]
		ret = append(ret, k+":"+v)
	}
	return strings.Join(ret, "\n")
}

func signature(httpMethod, contentMD5, contentType, headerString, resource string) string {
	stringSection := []string{
		httpMethod,
		contentMD5,
		contentType,
		headerString,
		resource,
	}
	stringToSign := strings.Join(stringSection, "\n")

	h := md5.New()
	h.Write([]byte(stringToSign))
	resultMD5 := hex.EncodeToString(h.Sum(nil))
	fmt.Println(stringToSign)
	return strings.ToUpper(resultMD5)
}
