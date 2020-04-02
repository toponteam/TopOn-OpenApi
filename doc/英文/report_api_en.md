# Change Log

| version | date  | notes    |
| :--------: | :------------ | -------------------- |
| v 1.0    | 2019/7/17 | supports full report |
| v 2.0    | 2019/8/30 | supports LTV & retention report |
| v 2.1    | 2020/3/17 | update full report metrics |

## 1. Introduction

In order to improve the monetization efficiency of publishers, TopOn provides API for data report, which can query comprehensive report, LTV & retention report and other data. This document is the detailed instruction of API. If you need any assistance, please feel free to reach us. Thank you!

## 2. Authentication acquisition

Before using the batch creation API of TopOn, publishers shall apply  for publisher_key that can identify the request from the publisher. For more details to apply the authority, please consult with the business manager contacted you.

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
| Headers      | Headers except X-Up-Signature              | X-Up-Timestamp:1562813567000 X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2 |
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

## 4. Full report

### 4.1 Request URL

<https://openapi.toponad.com/v2/fullreport>

### 4.2 Request Method

POST

### 4.3 Request params

| params       | type                | required | notes                                | sample                                     |
| ------------ | ------------------- | -------- | ----------------------------------- | ------------------------------------------ |
| startdate    | Int                 | Y        | start date，format：YYYYmmdd                                 | 20190501        |
| enddate      | Int                 | Y        | end date，format：YYYYmmdd                                   | 20190506          |
| app_id_list          | Array[String] | N        | app id                               | ['xxxxx']                                  |
| placement_id_list    | Array[String] | N        | placement id                         | ['xxxxx']                                  |
| time_zone            | String        | N        | report time zone                     | UTC-8,UTC+8,UTC+0                          |
| network_firm_id_list | Array[int32]  | N        | network firm id list                 |  ['xxxxx']                                          |
| adsource_id_list     | Array[int32]  | N        | adsource id list                     | [121]                                      |
| area_list            | Array[String] | N        | area list                            | ['xxxxx']                                  |
| placement_id | String              | N        | Placement ID(single)                     | xxxxx       |
| group_by     | Array               | N        | maximum three：date（default），app，placement，adformat，area，network，adsource，network firm id | ["app","placement","area"]                 |
| metric       | Array               | N        | return metrics. default（default values），all，dau，arpu，request，fillrate，impression，click，ctr，ecpm，revenue，request_api，fillrate_api，impression_api，click_api，ctr_api，ecpm_api | ["dau","arpu","request","click","ctr_api"] |
| start        | Int                 | N        | offset                                                       | 0                    |
| limit        | Int                 | N        | limit row number. default 1000.  [1,1000]                    | 1000                 |

 

- default return metrics：

dau，arpu，request，fillrate，impression，click，ecpm，revenue，impression_api，click_api，ecpm_api

 

### 4.4 Return data

| fileds           | type   | required | notes                                                        |
| ---------------- | ------ | -------- | ------------------------------------------------------------ |
| count            | Int    | Y        | count of the return rows                                     |
| date             | String | Y        | date，format：YYYYmmdd. Return if in param 'group_by'        |
| app.id           | String | Y        | APP ID                                                       |
| app.name         | String | N        | APP name                                                     |
| app.platform     | String | N        | APP platform                                                 |
| placement.id     | String | N        | Placement ID                                                 |
| placement.name   | String | N        | Placement name                                               |
| adformat         | String | N        | rewarded_video/interstitial/banner/native/splash.        Return if in param 'group_by' |
| area             | String | N        | country code.Return if in param 'group_by'                   |
| network_firm_id  | String | N        | network firm id.Return if in param 'group_by' |
| network_firm     | String | N        | network firm name.Return if in param 'group_by' |
| network          | String | N        | account id.Return if in param 'group_by' |
| network_name     | String | N        | account name.Return if in param 'group_by' |
| adsource.network | String | N        | adsource network name                                        |
| adsource.token   | String | N        | adsource token.adsource's appid,slotid and so on.Return if in param 'group_by' |
| time_zone        | String | N        | UTC+8、UTC+0、UTC-8                                  |
| currency         | String | N        | currency |
| new_users        | String | N        | new users                                                     |
| new_user_rate    | String | N        | new user rate                                                  |
| day2_retention   | String | N        | day2 retention                                                     |
| deu              | String | N        | deu                                                          |
| engaged_rate     | String | N        | engaged rate                                                       |
| imp_dau          | String | N        | imp/dau                                                    |
| imp_deu          | String | N        | imp/deu                                                    |
| impression_rate  | String | N        | impression rate                                                       |
| dau              | String | N        | Return if in param 'group_by'                                |
| arpu             | String | N        | need dau                                                     |
| request          | String | N        | request numbers                                              |
| fillrate         | String | N        | fillrate                                                     |
| impression       | String | N        | impression numbers                                           |
| click            | String | N        | click numbers                                                |
| ctr              | String | N        | ctr                                                          |
| ecpm             | String | N        | ecpm                                                         |
| revenue          | String | N        | revenue                                                      |
| request_api      | String | N        | network data:request numbers                                 |
| fillrate_api     | String | N        | network data:fillrate                                        |
| impression_api   | String | N        | network data:impression numbers                              |
| click_api        | String | N        | network data:click numbers                                   |
| ctr_api          | String | N        | network data:ctr                                             |
| ecpm_api         | String | N        | network data:ecpm                                            |

### 4.5 Sample

```
{
​    "startdate": 20190706,
​    "enddate": 201907010,
​     "limit":120,  
​    "group_by":["adsource"],
​    "metric":["all"],
​    "start":0,
​    "app_id":"a5c41a9ed1679c",
​    "placement_id":""
}
```


return data sample：
```
{
​	"count": 64,
​	"records": [{
​		"adsource": {
​			"network": "TouTiao",
​			"token": "{\"app_id\":\"5008225\",\"slot_id\":\"908225577\",\"is_video\":\"1\"}"
​		},
​		"revenue": "12995.80"
​	}]
}
```

## 5. LTV & retention report

### 5.1 Request URL

<https://openapi.toponad.com/v2/ltvreport>

### 5.2 Request method

POST
### 5.3 Request params

| params  | type | required | notes                                                    | 样例                                       |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| startdate    | Int    | Y        | start date, format：YYYYmmdd                   | 20190501                                   |
| enddate      | Int    | Y        | end date, format：YYYYmmdd                     | 20190506                                   |
| area_list | Array[String] | N | area list| ["xxx"] |
| appid_list | String    | N        | app id                                    | a5c41a9ed1679c                                   |
| time_zone | String | Y | timezone | UTC+8、UTC+0、UTC-8 |
| metric      | array    | N        | default：[“ltv_day_1”、”ltv_day_7”、”retention_day_2”、”retention_day_7”][“all”] all: all metrics | [“ltv_day_1”， “retention_day_2”]                                   |                         |
| group_by    | array    | N        | defaults：["app_id”, "date_time", "area"]                             | ["area"]                                   |
| start    | Int    | N        |     offset                           |                                    0|
| limit    | Int    | N        | limit row number. default 1000.  [1,1000] | 1000                                 |

### 5.4 Return data

| fields     | type | notes                                                    |
| ---------------- | ------ | ------------------------------------------------------------ |
| count            | Int           | count of the row numbers                              |
| records             | array       |   -                 |

**records:**

| fields | type | notes                |
| ---------------- | ------ | ------------------------ |
| date             | string | default return  |
| app.id           | string    | default return   |
| app.name         | string | default return   |
| new_user         | string | default return   |
| dau              | string | default return   |
| revenue          | string | don't return if group by channel |
| arpu             | string | with revenue     |
| ltv\_day\_1        | string | default return   |
| ltv\_day\_2        | string | -                         |
| ltv\_day\_3        | string | -                         |
| ltv\_day\_4        | string | -                         |
| ltv\_day\_5        | string | -                         |
| ltv\_day\_6        | string | -                         |
| ltv\_day\_7        | string | default return   |
| ltv\_day\_14       | string | -                         |
| ltv\_day\_30       | string | -                         |
| ltv\_day\_60       | string | -                         |
| retention\_day\_2  | string | default return   |
| retention\_day\_3  | string | -                         |
| retention\_day\_4  | string | -                         |
| retention\_day\_5  | string | -                         |
| retention\_day\_6  | string | -                         |
| retention\_day\_7  | string | default return   |
| retention\_day\_14 | string | -                         |
| retention\_day\_30 | string | -                         |
| retention\_day\_60 | string | -                         |
| time_zone | string | - |
| arpu | string | - |
| currency | string | - |

> notes:
> 1. Earliest date is the day before yesterday
> 2. ltv\_day\_N and retention\_day\_N reutrn '-', means the metrics are not exist.

### 5.5 Sample

``` 
{
    "count": 5,
    "records": [
        {
            "date": "20190823",
            "app": {
                "id": "122",
                "name": "abcde",
                "platform": "2"
            },
            "new_user": "15202",
            "dau": "0",
            "revenue": "5880.77",
            "ltv_day_1": "0.2334",
            "ltv_day_7": "-",
            "retention_day_2": "0.269",
            "retention_day_7": "-"
        }
    ]
}

```

## 6. Notices
Please control the frequency of requests:

•  1000 per hour

•  10000 per day
