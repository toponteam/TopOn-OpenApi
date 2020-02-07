# TopOn Device Reporting API

## Change Log

| version | date  | notes        |
| :-------: | ------------- | -------------------- |
| v 1.0    | 2019/8/29 | supports device reporting |


## Contents

[1. Introduction](#Introduction)</br>
[2. Authentication acquisition](#Authentication_acquisition)</br>
[3. Authentication check](#Authentication_check)</br>
[4. Device report](#Device_report)</br>
[5. Notices](#Notices)</br>
[6. Appendix1：golang demo](#Appendix1：golang_demo)</br>

<h2 id='Introduction'>1. Introduction</h2>

In order to improve the monetization efficiency of publishers, TopOn provides the reporting API. This document is the detailed instruction of API. If you need any assistance, please feel free to reach us. Thank you!

<h2 id='Authentication_acquisition'>2. Authentication acquisition</h2>

Before using the batch creation API of TopOn, publishers shall apply  for publisher_key that can identify the request from the publisher. For more details to apply the authority, please consult with the business manager contacted you.

<h2 id='Authentication_check'>3. Authentication check</h2>

### 3.1 The process description of API request

- The client generates a key based on the content of the API request, including the HTTP headers and bodies.
- The client uses MD5 to sign on the key that generated in the first step.
- The client sends the API request content along with the signed key to the server.
- After receiving the request, the server repeats the above first and second steps and calculates the expected signature at the server.
- The server compares the expected signature with the signed key that sent by the client.If they are entirely consistent with eachother, the request can pass the security verification.Otherwise, it will be rejected.

### 3.2 Header general request params

| params         | notes                                                        | sample                                     |
| -------------- | ------------------------------------------------------------ | ------------------------------------------ |
| X-Up-Key       | publisher_key                                                | X-Up-Key: i8XNjC4b8KVok4uw5RftR38Wgp2BFwql |
| X-Up-Timestamp | Unix timestamp(ms), the millisecond from 1970/1/1. Valid duration is 15 minutes. | 1562813567000                              |
| X-Up-Signature | signature string                                             |                                            |


### 3.3 Params to create signature

| params       | notes                                      | sample                                                       |
| ------------ | ------------------------------------------ | ------------------------------------------------------------ |
| Content-MD5  | MD5 from HTTP Body string（upper letters） | 875264590688CA6171F6228AF5BBB3D2                             |
| Content-Type | type of HTTP Body                          | application/json                                             |
| Headers      | Headers except X-Up-Signature              | X-Up-Timestamp:1562813567000X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2\n |
| HTTPMethod   | HTTP method(upper letters)                 | PUT、GET、POST                                               |
| Resource     | strings from HTTP path and query params    | /v1/fullreport?key1=val1&key2=val2                           |


### 3.4 Create signature

Create signature string：

     SignString = HTTPMethod + "\n" 
                        \+ Content-MD5 + "\n" 
                        \+ Content-Type + "\n"  
                        \+ Headers + "\n"
                        \+ Resource 

If HTTP body is empty：
    

    SignString = HTTPMethod + "\n" 
                        \+ "\n" 
                        \+ "\n" 
                        \+ Headers + "\n"
                        \+ Resource 

Resource:

    URL Path and query params       

Headers：

    // X-Up-Key + X-Up-Timestamp (sort by first letter)
    // except X-Up-Signature 
    Headers = Key1 + ":" + Value1 + '\n' + Key2 + ":" + Value2   
    
    
    Sign = MD5(SignString)



Server will create sign and campare the sign with X-Up-Signature

 

### 3.5 Response HTTP code

| HTTP code | response message         | notes                       |
| --------- | ------------------------ | --------------------------- |
| 200       | -                        | success                     |
| 500       | -                        | general exception           |
| 600       | StatusHeaderParamError   | request Header params error |
| 601       | StatusSign               | Sign error                  |
| 602       | StatusParam              | params error                |
| 603       | StatusPublisherRestrict  | no authentication           |
| 604       | StatusAppLengthError     | App creation error          |
| 605       | StatusRpcParamError      | base Server error           |
| 606       | StatusRequestRepeatError | duplicated requests         |

<h2 id='Device_report'>4. Device report</h2>

### 4.1 Request URL

<https://openapi.toponad.com/v1/devicereport>

### 4.2 Request method

GET

### 4.3 Request 

| params   | type | required | notes                                                     | sample                                  |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| day    | Int    | Y        | start date, format：YYYYmmdd                   | 20190501,Earliest date is the day before yesterday |
| app_id       | String | N        | APP ID(single)                        | xxxxx                                                                            |

notes: Your device reporting data will create in the date which open authentication 

### 4.4 Return data

API will return the download url, you can get data from the url. <br>
https://topon-openapi.s3.amazonaws.com/topon_report_device/dt%3D2019-07-10/publisher_id%3D22/app_id%3Da5d147334b3685/000000_0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIA35FGARBHLHHS7TWB%2F20190828%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20190828T095315Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=6aaf947f9b2cf02f3acb49d64a3daf719cb0b57a3d5221b0121a006e58b04b10 <br>

The data file is CSV, explode by ',' .

Fields detail:

| fields | type | notes                                                     |
| ---------------- | ------  | ------------------------------------------------------------ |
| placement_id            | String      | Placement ID                                          |
| placement_name             | String      | Placement name    |
| placement_format          | String     | adformat 0: native,1: rewarded_video,2: banner,3: interstitial,4: splash                            |
| unit_id         | String      | TopOn's adsource id                                                  |
| unit_network     | String       | TopOn's network name                                        |
| unit_token     | String       | TopOn's adsource token                  |
| android_id   | String     | androidid                                          |
| gaid         | String      | google advertising id |
| idfa             | String        | ios device id                        |
| area          | String       | country code |
| impression | String       | impression number                       |
| click   | String      | click number |
| revenue              | decimal(18,6)       | revenue                              |
| ecpm             | decimal(18,6)       | ecpm                                      |

<h2 id='Notices'>5. Notices</h2>

Please control the frequency of requests:

•  1000 per hour

•  10000 per day

<h2 id='Appendix1：golang_demo'>6. Appendix1：golang_demo</h2>

• Java,PHP,Python demos are in the Git path /demo

```
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

  //openapi url
	
	demoUrl := "request URL"
	
	//request body
	
	body := "{}"
	
	//your publisherKey
	
	publisherKey := "your publisherKey"
	
	//request method
	
	httpMethod := "POST"
	
	contentType := "application/json"
	
	publisherTimestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	
	headers := map[string]string{
	
		"X-Up-Timestamp": publisherTimestamp,
	
		"X-Up-Key":       publisherKey,
	
	}
	
	//handle queryPath
	
	urlParsed, err := url.Parse(demoUrl)
	
	if err != nil {
	
		fmt.Println(err)
	
		return
	
	}
	
	//handle resource
	
	resource := urlParsed.Path
	
	_, err = url.ParseQuery(urlParsed.RawQuery)
	
	if err != nil {
	
		fmt.Println(err)
	
		return
	
	}

	//handle body
	
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
	
	defer resp.Body.Close()
	
	content, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
	
		fmt.Println("Fatal error", err.Error())
	
		return
	
	}

	//return reporting data
	
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
```
