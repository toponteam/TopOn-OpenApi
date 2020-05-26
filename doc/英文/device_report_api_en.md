# Change Log

| version | date  | notes        |
| :-------: | ------------- | -------------------- |
| v 1.0    | 2019/8/29 | supports device report |
| v 1.1    | 2020/3/17 | supports currency and timezone |
| v 1.2    | 2020/5/16 | supports abtest and segment dimension |


## 1. Introduction

In order to improve the monetization efficiency of publishers, TopOn provides the API for device dimension data reporting, which can understand the monetization of each device in detail and realize fine operation. This document is the detailed instruction of API. If you need any assistance, please feel free to reach us. Thank you!

## 2. Authentication acquisition

Before using the device reporting API of TopOn, publishers shall apply that can identify the request from the publisher. For more details to apply the authority, please consult with the business manager contacted you.

## 3. Authentication check

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
| X-Up-Signature | signature string                                             |                                            |-


### 3.3 Params to create signature

| params       | notes                                      | sample                                                       |
| ------------ | ------------------------------------------ | ------------------------------------------------------------ |
| Content-MD5  | MD5 from HTTP Body string（upper letters） | 875264590688CA6171F6228AF5BBB3D2                             |
| Content-Type | type of HTTP Body                          | application/json                                             |
| Headers      | Headers except X-Up-Signature              | X-Up-Timestamp: 1562813567000X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2\n |
| HTTPMethod   | HTTP method(upper letters)                 | PUT、GET、POST                                               |
| Resource Path     | strings from HTTP path    | /v1/fullreport                          |


### 3.4 Create signature

Create signature string：
```
     SignString = HTTPMethod + "\n" 
                        \+ Content-MD5 + "\n" 
                        \+ Content-Type + "\n"  
                        \+ Headers + "\n"
                        \+ Resource 
```
If HTTP body is empty：
    
```
    SignString = HTTPMethod + "\n" 
                        \+ "\n" 
                        \+ "\n" 
                        \+ Headers + "\n"
                        \+ Resource 
```
Resource:
```
    URL Path and query params       
```
Headers：
```
    // X-Up-Key + X-Up-Timestamp (sort by first letter)
    // except X-Up-Signature 
    Headers = Key1 + ":" + Value1 + '\n' + Key2 + ":" + Value2   


​    
    Sign = MD5(SignString)

```

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

## 4. Device report

### 4.1 Request URL

<https://openapi.toponad.com/v1/devicereport>

### 4.2 Request method

GET

### 4.3 Request 

| params   | type | required | notes                                                     | sample                                  |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| day    | Int    | Y        | start date, format：YYYYmmdd                   | 20190501,Earliest date is the day before yesterday |
| app_id       | String | Y        | APP ID(single)                        | xxxxx                                                     |
| timezone | Int | N | Time Zone | -8 or 8 or 0, default 8 |

- Your device reporting data will create in the date which open authentication. If the data is not obtained at the time, please try again the next day <br/>

- the data update time point of each time zone : <br/>
1. UTC + 8: delayed 2 days, update at 5 am (UTC + 8) <br/>
2. UTC + 0: delayed by 2 days, updated at 10 am (UTC + 8) <br/>
3. UTC - 8: delayed by 3 days, updated at 0 am (UTC + 8) <br/>

### 4.4 Return data

API will return the download url, you can get data from the url. <br>
https://topon-openapi.s3.amazonaws.com/topon_report_device/dt%3D2019-07-10/publisher_id%3D22/app_id%3Da5d147334b3685/000000_0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIA35FGARBHLHHS7TWB%2F20190828%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20190828T095315Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=6aaf947f9b2cf02f3acb49d64a3daf719cb0b57a3d5221b0121a006e58b04b10 <br>

The data file is CSV, explode by ',' .

Fields detail:

| fields | type | notes                                                     |
| ---------------- | ------  | ------------------------------------------------------------ |
| placement_id            | String      | Placement ID                                          |
| placement_name             | String      | Placement name    |
| placement_format          | String     | adformat 0: native,1: rewarded_video,2: banner,3: interstitial,4: splash                    |
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
| is_abtest             | String       | control group or test group <br> 0: control group or the A / B test is not activated, 1: test group                      |
| traffic_group_id             | String      | control or test group id |
| segment_id             | String       | segment id                                          |
| segment_name             | String      | segment name                                         |
## 5. Notices

Please control the frequency of requests:

•  1000 per hour for single user

•  10000 per day for single user
