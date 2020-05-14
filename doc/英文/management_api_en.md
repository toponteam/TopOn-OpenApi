# Change Log

| version | date  | notes                     |
| :--------: | ------------- | -------------------------------- |
| v 1.0    | 2019/7/17 | supports apps and placements |
| v 2.0    | 2019/11/4 | supports waterfall and segments |
| v 2.1    | 2020/3/16 | supports network and adsources |
| v 2.2    | 2020/5/14 | segments function adjustment |


## 1. Introduction

In order to improve the monetization efficiency of publishers, TopOn provides API with developers' background-related operations, such as creating app and placements, changing Waterfall priorities and so on. This document is the detailed instruction of API. If you need any assistance, please feel free to reach us. Thank you!

## 2. Authentication acquisition

After the account has been successfully registered, the developer backend management API permission has been automatically activated. Log in to the developer backend account management page to view publisher_key.

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

## 4. App API

### 4.1 Batch create and update apps

#### 4.1.1 Request URL

<https://openapi.toponad.com/v1/deal_app>

#### 4.1.2 Request method 

POST

#### 4.1.3 Request params

| params            | type   | required | notes                                                        |
| ----------------- | ------ | -------- | ------------------------------------------------------------ |
| count             | Int    | Y        | Quantity of created apps                                     |
| apps.app_id       | String | N        | app id                                                       |
| apps.app_name     | String | Y        | app name                                                     |
| apps.platform     | Int    | Y        | platform 1 or 2  (1:android，2:iOS)                          |
| apps.market_url   | String | N        | Need to be in compliance with requirements of app store links |
| apps.screen_orientation | Int    | Y        | 1:portrait <br>2:landscape <br>3:both                  |
| apps.package_name | String | N        | Need to be in compliance with requirements of app package name.  com.xxx |
| apps.category     | String | N        | category.See:Appendix1_App category and sub category enum.App that is not in the store at the time of creation must be passed. |
| apps.sub_category | String | N        | sub category.See:Appendix1_App category and sub category enum.App that is not in the store at the time of creation must be passed. |
| apps.coppa       | String | N        | Whether to comply with COPPA protocol. Default: no < br > 1: no, 2: yes |
| apps.ccpa       | String | N        | Whether to comply with CCPA protocol. Default: no < br > 1: no, 2: yes  |

 

#### 4.1.4 Return data 

| fields   | type   | required | notes                               |
| -------- | ------ | -------- | ----------------------------------- |
| app_id   | String | Y        | app id                              |
| app_name | String | Y        | app name                            |
| errors   | String | N        | error messages                      |
| platform | Int    | Y        | platform 1 or 2  (1:android，2:iOS) |
| screen_orientation | Int    | Y        | 1:portrait <br>2:landscape <br>3:both                  |
| apps.coppa       | String | N        | Whether to comply with COPPA protocol. Default: no < br > 1: no, 2: yes |
| apps.ccpa       | String | N        | Whether to comply with CCPA protocol. Default: no < br > 1: no, 2: yes  |

 

#### 4.1.5 Sample

request sample：
```
{
    "count": 1,
    "apps": [
        {
            "app_name": "oddman",
            "platform": 1,
            "screen_orientation":1,
            "market_url": "https://play.google.com/store/apps/details?id=com.hunantv.imgo.activity.inter"
        }
    ]
}
```


return sample：
```
[
    {
        "app_name": "oddman",
        "app_id": "",
        "platform": 1,
        "screen_orientation": 1,
        "errors": "repeat app name error"
    }
]
```

### 4.2 Get app list

#### 4.2.1 Request URL

<https://openapi.toponad.com/v1/apps>

#### 4.2.2 Request method

POST

#### 4.2.3 Request params

| params  | type        | required | notes                  |
| ------- | ----------- | -------- | ---------------------- |
| app_ids | Array[String] | N        | ["abc", "acc"]   （Cannot be used with the following parameters）      |
| start   | Int         | N        | Default 0              |
| limit   | Int         | N        | Default 100 , [0, 100] |

 

#### 4.2.4 Return data

| fields       | type   | required | notes                               |
| ------------ | ------ | -------- | ----------------------------------- |
| app_id       | String | Y        | app id                              |
| app_name     | String | Y        | app name                            |
| platform     | Int    | Y        | platform 1 or 2  (1:android，2:iOS) |
| market_url   | String | N        | -                                   |
| screen_orientation | Int  | Y    | 1:portrait <br>2:landscape <br>3:both   |
| package_name | String | N        | -                                   |
| category     | String | N        | -                                   |
| sub-category | String | N        | -                                   |
| apps.coppa       | String | N        | Whether to comply with COPPA protocol. Default: no < br > 1: no, 2: yes |
| apps.ccpa       | String | N        | Whether to comply with CCPA protocol. Default: no < br > 1: no, 2: yes  |

 

#### 4.2.5 Sample

request sample：
```
{
	"limit":1
}
```


return sample：
```
[
    {
        "app_name": "uparputest",
        "app_id": "a5bc9921f7fdb4",
        "platform": 2,
        "market_url": "https://itunes.apple.com/cn/app/%E7%A5%9E%E5%9B%9E%E9%81%BF/id1435756371?mt=8",
        "category": "Game",
        "sub_category": "Action",
        "screen_orientation": 3
    }
]
```

### 4.3 Batch delete apps

#### 4.3.1 Request URL

<https://openapi.toponad.com/v1/del_apps>

#### 4.3.2 Request method

POST

#### 4.3.3 Request params

| params  | type        | required | notes                  |
| ------- | ----------- | -------- | ---------------------- |
| app_ids | Array[String] | Y        | ["abc", "acc"]         |

#### 4.3.4 Return data

None. If it is an error, errors will be returned.

#### 4.3.5 Sample

request sample：
```
{
	"app_ids": ["a1bu2thutsq3mn"]
}
```


return sample：

None. If it is an error, errors will be returned.

## 5. Placement API

### 5.1 Batch create and update placements

#### 5.1.1 Request URL

<https://openapi.toponad.com/v1/deal_placement>

#### 5.1.2 Request method

POST

#### 5.1.3 Request params

| params                                | type   | required | notes                                                        |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| count                                 | Int    | Y        | Quantity of created placements                               |
| app_id                                | String | Y        | APP ID of created placements                                 |
| placements.placement_name             | String | Y        | placement name. max length 30                                |
| placements.adformat                   | String | Y        | native,banner,rewarded_video,interstitial,splash             |
| placements.template                   | Int    | N        | Configurations for native ads: 0：standard<br/>1：Native Banner<br/>2：Native Splash |
| placements.template.cdt               | Int    | N        | template is Native Splash：countdown time, default 5s        |
| placements.template.ski_swt           | Int    | N        | template is Native Splash：it can skipped or not, it could be skipped by default.<br/>0：No<br/>1：Yes |
| placements.template.aut_swt           | Int    | N        | template is Native Splash：it can be auto closed or not, it could be auto closed by default.<br/>0：No<br/>1：Yes |
| placements.template.auto_refresh_time | Int    | N        | template is Native Banner：it can be auto refreshed or not, it could not be auto refreshed by default<br/>-1 no auto refresh<br/>0-n auto refresh time (s) |
| remark                                | String | N        | remarks information                                          |
| status                                | Int    | N        | placement's status                                           |
 

#### 5.1.4 Return data

| fields                                | type   | required | notes                                                        |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| app_id                                | String | Y        | APP ID                                                       |
| placement_name                        | String | Y        | placement name                                               |
| placement_id                          | String | Y        | placement ID                                                 |
| adformat                              | String | Y        | Native, banner, rewarded video, interstitial, splash         |
| placements.template                   | Int    | N        | Configurations for native ads: 0：standard 1：Native Banner 2：Native Splash |
| placements.template.cdt               | Int    | N        | template is Native Splash：countdown time, default 5s        |
| placements.template.ski_swt           | Int    | N        | Template is Native Splash：it can skipped or not, it could be skipped by default.<br/>0：No<br/>1：Yes |
| placements.template.aut_swt           | Int    | N        | Template is Native Splash：it can be auto closed or not, it could be auto closed by default.<br/>0：No<br/>1：Yes |
| placements.template.auto_refresh_time | Int    | N        | template is Native Banner：it can be auto refreshed or not, it could not be auto refreshed by default<br/>-1 no auto refresh<br/>0-n auto refresh time (s) |
| remark                                | String | N        | remarks information                                          |
| status                                | Int    | N        | placement's status                                           |
 

#### 5.1.5 Sample

request sample：
```
{
    "count": 1,
    "app_id": "a5c41a9ed1679c",
    "placements": [
        {
            "placement_name": "6",
            "adformat": "native",
            "remark": "remark",
            "template":2,
            "template_extra":{
            	"cdt":55,
            	"ski_swt":1,
            	"aut_swt":1
            }
            
        }
        
    ]
}
```


return sample：
```
[
    {
        "app_name": "test1",
        "app_id": "a5c41a9ed1679c",
        "platform": 2,
        "placement_id": "b1bv57tielnlts",
        "placement_name": "6",
        "adformat": "native",
        "remark": "remark",
        "template": 2,
        "template_extra": {
            "cdt": 55,
            "ski_swt": 1,
            "aut_swt": 1
        }
    }
]
```

### 5.2 Get placement list

#### 5.2.1 Request URL

<https://openapi.toponad.com/v1/placements>

#### 5.2.2 Request method 

POST

#### 5.2.3 Request params

| params        | type        | required | notes              |
| ------------- | ----------- | -------- | ------------------ |
| app_ids       | Array[String] | N        | eg: ["abc", "acc"] |
| placement_ids | Array[String] | N        | eg: ["abc", "acc"] |
| start         | Int         | N        | Default 0   (Not required when both app and ad slot are specified)       |
| limit         | Int         | N        | Default 100  (Not required when both app and ad slot are specified)           |

 

#### 5.2.4 Return data

| fields         | type   | required | notes                     |
| -------------- | ------ | -------- | ------------------------- |
| app_id         | String | Y        | app id                    |
| app_name       | String | Y        | app name                  |
| platform       | Int    | Y        | 1 or 2  (1:android,2:IOS) |
| placement_id   | String | N        | -                         |
| placement_name | String | N        | -                         |
| adformat       | String | N        | -                         |
| placements.template                   | Int    | N        | Configurations for native ads: 0：standard 1：Native Banner 2：Native Splash |
| placements.template.cdt               | Int    | N        | template is Native Splash：countdown time, default 5s        |
| placements.template.ski_swt           | Int    | N        | Template is Native Splash：it can skipped or not, it could be skipped by default.<br/>0：No<br/>1：Yes |
| placements.template.aut_swt           | Int    | N        | Template is Native Splash：it can be auto closed or not, it could be auto closed by default.<br/>0：No<br/>1：Yes |
| placements.template.auto_refresh_time | Int    | N        | template is Native Banner：it can be auto refreshed or not, it could not be auto refreshed by default<br/>-1 no auto refresh<br/>0-n auto refresh time (s) |
| remark                                | String | N        | remarks information                                          |
| status                                | Int    | N        | placement's status                                           |
 

#### 5.2.5 Sample

request sample：
```
{
	"placement_ids":["b5bc9bc2951216"]
}
```


return sample：
```
[
    {
        "app_name": "topontest",
        "app_id": "a5bc9921f7fdb4",
        "platform": 2,
        "placement_id": "b5bc9bc2951216",
        "placement_name": "topontest_rewardvideo",
        "adformat": "rewarded_video"
    }
]
```

### 5.3 Batch delete placemens

#### 5.3.1 Request URL

<https://openapi.toponad.com/v1/del_placements>

#### 5.3.2 Request method 

POST

#### 5.3.3 Request params

| params        | type        | required | notes              |
| ------------- | ----------- | -------- | ------------------ |
| placement_ids | Array[String] | Y        | eg: ["abc", "acc"] |

#### 5.3.4 Return data

| fields         | type   | required | notes                     |
| -------------- | ------ | -------- | ------------------------- |
| msg | String | N        | -         |
 

#### 5.3.5 Sample
request sample：
```
{
	"placement_ids":["b5bc9bc2951216"]
}
```


return sample：
```
{
    "msg": "suc"
}
```

## 6. Segment API

### 6.1 Batch create and update segments

#### 6.1.1 Request URL

<https://openapi.toponad.com/v2/deal_segment>

#### 6.1.2 Request method 

POST

#### 6.1.3 Request params

| params        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| count         | Int    | Y        | segment number                                               |
| app_id                  | String    | Y        | app_id                                                    |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group  |
| segments      | Array  | Y        | -                                                            |
| segments.name          | String | Y        | segment name （The newly added segment priority is ranked in front of the existing group）                                               |
| segments.segment_id    | String | N        | must reture segment id when updating segment                 |
| segments.rules         | Array  | Y        | segment rules                                                |
| segments.rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| segments.rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| segments.rules.content | string | Y        | See:Appendix2_segment rule enum |

#### 6.1.4 Return data

|   fields      | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| count         | Int    | Y        | segment number                                               |
| app_id                  | String    | Y        | app_id                                                    |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group  |
| segments      | Array  | Y        | -                                                            |
| segments.name          | String | Y        | segment name （The newly added segment priority is ranked in front of the existing group）                                               |
| segments.errors    | String | N        | segment error message                 |
| segments.segment_id    | String | N        | must reture segment id when updating segment                 |
| segments.rules         | Array  | Y        | segment rules                                                |
| segments.rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| segments.rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| segments.rules.content | string | Y        | See:Appendix2_segment rule enum |



#### 6.1.5 Sample

request sample：

```
{
    "count": 3,
    "app_id":"a5bc9921f7fdb4",
    "placement_id":"b5bc9bbfb0f913",
    "is_abtest":0,
    "segments": [
        {
            "name": "999",
            "segment_id": "c1c3femr2h7smb",
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                }
            ]
        },
        {
            "name": "5555",
            "segment_id": "c5ea52b0e79baf",
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                }
            ]
        },
        {
            "name": "2123123434"
        }
    ]
}
```

 

return sample：

```
{
    "count": 3,
    "placement_id": "b5ebbb200f10af",
    "app_id": "a5e68b165154d5",
    "segments": [
        {
            "name": "999",
            "segment_id": "c1c3kadvqpuffb",
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                }
            ]
        },
        {
            "name": "5555",
            "segment_id": "c1c3kadvr7dkfu",
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                }
            ]
        },
        {
            "name": "2123123434",
            "segment_id": "",
            "errors": "segment rule length must 1"
        }
    ]
}
```

### 6.2 Get segment list

#### 6.2.1 Request URL

<https://openapi.toponad.com/v2/waterfall/get_segment>

#### 6.2.2 Request method 

POST

#### 6.2.3 Request params

| params  | type | required | notes                                                     |
| ----------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id | String | Y        | placement_id                         |
| app_id                  | String    | Y        | app_id                                                    |
| is_abtest             | Int    | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group         |

 

#### 6.2.4 Return data

| fields        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| priority      | Int    | Y        | Priority parameter                                           |
| name          | String | Y        | Segment name                                                 |
| segment_id    | String | Y        | Segment ID                                                   |
| parallel_request_number    | Int | Y        | Number of parallel requests                             |
| auto_load    | Int | Y        | Default 0: off, only 0 or positive integer < br/ > for Banner, automatic refresh time can be set, and greater than 0 means automatic refresh time < br/ > for RV and plug-in screen, only the switch status of automatic request is controlled, and non-zero means on |
| day_cap    | Int | Y        | Default -1: indicates off                            |
| hour_cap    | Int | Y        | Default -1: indicates off                            |
| priority    | Int | Y        | Default -1: indicates off                             |
| rules         | Array  | Y        | Segment rules                                                |
| rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| rules.content | string | Y        | See:Appendix2_segment rule enum|



#### 6.2.5 Sample

request sample：

```
https://openapi.toponad.com/v2/waterfall/get_segment?placement_id=b5bc9bbfb0f913&app_id=a5bc9921f7fdb4&is_abtest=1
```

return sample：

```
[
    {
        "name": "segment1",
        "segment_id": "c1c3eo1tahts80",
        "parallel_request_number": 1,
        "auto_load": 0,
        "day_cap": -1,
        "hour_cap": -1,
        "pacing": -1,
        "priority": 1,
        "rules": [
            {
                "type": 3,
                "rule": 0,
                "content": [
                    "4g",
                    "3g",
                    "2g"
                ]
            },
            {
                "type": 17,
                "rule": 0,
                "content": [
                    "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                    "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                    "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                ]
            },
            {
                "type": 10,
                "rule": 0,
                "content": [
                    "iphone"
                ]
            },
            {
                "type": 9,
                "rule": 7,
                "content": "5.0.0"
            }
        ]
    },
    {
        "name": "Default Segment",
        "segment_id": "",
        "parallel_request_number": 1,
        "auto_load": 0,
        "day_cap": 0,
        "hour_cap": 0,
        "pacing": 0,
        "priority": 2
    }
]
]
```

### 6.3 Batch delete segments

#### 6.3.1 Request URL

<https://openapi.toponad.com/v1/waterfall/del_segment>

#### 6.3.2 Request method 

POST

#### 6.3.3 Request params

| params  | type | required | notes                         |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | Multiple segment is an array |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    |N       | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group    |

 

#### 6.3.4 Return data

| params          | type   | required | notes                                                          |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | Y        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group            |
| segments               | Array  | Y        | -                                                             |
| segments.name          | String | Y        | Segment name                                                  |
| segments.priority      | Int | Y        | Segment priority                                                  |
| segments.segment_id    | String | N        | Segment ID                              |
| segments.rules         | Array  | Y        | Segment rules                                                 |
| segments.rules.type    | Int    | Y        | Segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| segments.rules.rule    | Int    | Y        | Segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value |
| segments.rules.content | string | Y        | See:Appendix2_segment rule enum           |


#### 6.3.5 Sample

request sample：

```
{
    "placement_id": "111111",
    "is_abtest": 1,
    "segment_ids": [
        "22222"
    ]
}
```

return sample：

```
{
    "placement_id": "b5bc9bbfb0f913",
    "is_abtest": 0,
    "segments": [
        {
            "priority": 1,
            "name": "segment1",
            "segment_id": "c1c3eo1tahts80",
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                }
            ]
        },
        {
            "priority": 2,
            "name": "Default Segment",
            "segment_id": ""
        }
    ]
}
```

### 6.4 Set Segment Priority

#### 6.4.1 Request URL

<https://openapi.toponad.com/v2/waterfall/set_segment_rank>

#### 6.4.2 Request method 

POST

#### 6.4.3 Request params

| params  | type | required | notes                         |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | Multiple segment is an array |
| placement_id | int32 | Y        | placement_id |
| is_abtest | int32 | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group |
| app_id | int32 | Y        | app_id |


#### 6.4.4 Return data

| fields        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | Y        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group   |
| priority      | Int    | Y        | Priority parameter                                           |
| segments               | Array  | Y        | -                                                             |
| segments.name         | String | Y        | Segment name                                                 |
| segments.priority      | Int | Y        | Segment priority                                                 |
| segments.segment_id    | String | Y        | Segment ID                                                   |
| segments.parallel_request_number    | Int | Y        | Number of parallel requests                             |
| segments.auto_load    | Int | Y        | Default 0: off, only 0 or positive integer < br/ > for Banner, automatic refresh time can be set, and greater than 0 means automatic refresh time < br/ > for RV and plug-in screen, only the switch status of automatic request is controlled, and non-zero means on |
| segments.day_cap    | Int | Y        | Default -1: indicates off                            |
| segments.hour_cap    | Int | Y        | Default -1: indicates off                            |
| segments.priority    | Int | Y        | Default -1: indicates off                             |
| segments.rules         | Array  | Y        | Segment rules                                                |
| segments.rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| segments.rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| segments.rules.content | string | Y        | See:Appendix2_segment rule enum|


#### 6.4.5 Sample

request sample：

```
{
    "app_id":"a5bc9921f7fdb4",
    "placement_id":"b5bc9bbfb0f913",
    "is_abtest": 1,
    "segment_ids": [
    	"c1c3eo129ou5v9",
    	"c1c3eo1tahts80",
        "c5ea52b0e79baf"
    ]
}
```

return sample：

```
{
    "placement_id": "b5bc9bbfb0f913",
    "is_abtest": 1,
    "segments": [
        {
            "name": "segment1",
            "segment_id": "c1c3eo129ou5v9",
            "parallel_request_number": 0,
            "auto_load": 0,
            "day_cap": 0,
            "hour_cap": 0,
            "pacing": 0,
            "priority": 0,
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                },
                {
                    "type": 0,
                    "rule": 0,
                    "content": []
                }
            ]
        },
        {
            "name": "segment2",
            "segment_id": "c1c3eo1tahts80",
            "parallel_request_number": 0,
            "auto_load": 0,
            "day_cap": 0,
            "hour_cap": 0,
            "pacing": 0,
            "priority": 1,
            "rules": [
                {
                    "type": 3,
                    "rule": 0,
                    "content": [
                        "4g",
                        "3g",
                        "2g"
                    ]
                },
                {
                    "type": 17,
                    "rule": 0,
                    "content": [
                        "591B0524-9BC6-4AFC-BE75-7DDD4937DBE1",
                        "DA973F33-9A9D-4B47-82FB-4C6B9B19E09D",
                        "C093B2E8-849B-45AE-B11A-E862B1EE1025"
                    ]
                },
                {
                    "type": 10,
                    "rule": 0,
                    "content": [
                        "iphone"
                    ]
                },
                {
                    "type": 9,
                    "rule": 7,
                    "content": "5.0.0"
                },
                {
                    "type": 0,
                    "rule": 0,
                    "content": []
                }
            ]
        }
    ]
}
```


### 6.5 Set attributes of segments in Waterfall

#### 6.5.1 Request URL

<https://openapi.toponad.com/v2/waterfall/set_segment>

#### 6.5.2 Request method 

POST

#### 6.5.3 Request params

| params  | type | required | notes                         |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | segment_ids |
| placement_id | String | Y        | placement_id |
| is_abtest | int32 | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group |
| app_id | String | Y        | app_id |
| segments               | Array  | Y        | -                                                             |
| segments.segment_id    | String | N        | Segment id                              |
| segments.parallel_request_number    | Int | Y        | parallel request number                            |
| segments.auto_load    | Int | Y        | Default 0: off, only 0 or positive integer < br/ > for Banner, automatic refresh time can be set, and greater than 0 means automatic refresh time < br/ > for RV and plug-in screen, only the switch status of automatic request is controlled, and non-zero means on.|
| segments.day_cap    | Int | Y        | Default -1 ：off                            |
| segments.hour_cap    | Int | Y        | Default -1 ：off                             |
| segments.pacing    | Int | Y        | Default -1 ：off                            |


#### 6.5.4 Return data

| fields        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | Y        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group               |
| segments               | Array  | Y        | -                                                             |
| segments.segment_id    | String | N        | Segment id                              |
| segments.parallel_request_number    | Int | Y        | parallel request number                            |
| segments.auto_load    | Int | Y        | Default 0: off, only 0 or positive integer < br/ > for Banner, automatic refresh time can be set, and greater than 0 means automatic refresh time < br/ > for RV and plug-in screen, only the switch status of automatic request is controlled, and non-zero means on.|
| segments.day_cap    | Int | Y        | Default -1 ：off                            |
| segments.hour_cap    | Int | Y        | Default -1 ：off                             |
| segments.pacing    | Int | Y        | Default -1 ：off                            |


#### 6.5.5 Sample

request sample：

```
{
    "app_id": "a5e68b165154d5",
    "placement_id": "b5ebbb200f10af",
    "is_abtest": 0,
    "segments": [
        {
            "segment_id": "c1c3kadvqpuffb",
            "auto_load": 3,
            "day_cap": 1,
            "hour_cap": 6,
            "pacing": 7
        },
        {
            "segment_id": "c5ebbb2823ada1",
            "auto_load": 7,
            "day_cap": 3,
            "hour_cap": 4,
            "parallel_request_number":24,
            "pacing": 5
        }
    ]
}
```

return sample：

```
{
    "app_id": "a5e68b165154d5",
    "placement_id": "b5ebbb200f10af",
    "is_abtest": 0,
    "segments": [
        {
            "segment_id": "c1c3kadvqpuffb",
            "parallel_request_number": 1,
            "auto_load": 3,
            "day_cap": 1,
            "hour_cap": 6,
            "pacing": 7
        },
        {
            "segment_id": "c5ebbb2823ada1",
            "parallel_request_number": 24,
            "auto_load": 7,
            "day_cap": 3,
            "hour_cap": 4,
            "pacing": 5
        }
    ]
}
```


## 7. Waterfall API

### 7.1 Get waterfall's adsources

#### 7.1.1 Request URL

<https://openapi.toponad.com/v1/waterfall/units>

#### 7.1.2 Request method 

GET

#### 7.1.3 Request params

| params   | type | required | notes         |
| ------------ | ------ | -------- | --------------- |
| placement_id | String | Y        | placement ID |
| segment_id   | String | Y        | Segment ID      |
| is_abtest             | Int    | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group     |

#### 7.1.4 Return data

| fields                              | type    | required | notes                                                        |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement ID                                                 |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group  |
| ad_source_list                      | Array   | Y        | adsource list in used                                        |
| ad_source_list.ad_source_id         | Int     | N        | adsource ID                                                  |
| ad_source_list.ecpm                 | float64 | N        | eCPM                                                         |
| ad_source_list.auto_ecpm            | float64 | N        | auto eCPM                                                  |
| ad_source_list.header_bidding_witch | Int     | N        | if support Header Bidding<br />1：not support，<br />2：support |
| ad_source_list.auto_switch          | Int     | N        | 1：not open auto eCPM sort switch，<br />2：open auto eCPM sort switch |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：close                                           |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：close                                           |
| ad_source_list.pacing               | Int     | N        | Default -1 ：close                                           |
| free_ad_source_list                 | Array   | N        | adsource list not in used                                    |
| offer_list                          | Array   | N        | my offer list in used                                        |
| offer_list.offer_id                 | String  | N        | offer id                                                     |
| offer_list.offer_name               | String  | N        | offer name                                                   |

#### 7.1.5 Sample

request sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1"
}
```

return sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1",
    "ad_source_list": [
        {
            "priority": 1,
            "ad_source_id": "ad_source_id1",
            "ecpm": "ecpm1",
            "header_bidding_witch": 0,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        },
        {
            "priority": 2,
            "ad_source_id": "ad_source_id2",
            "ecpm": "ecpm2",
            "header_bidding_witch": 0,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        }
    ]
}
```

### 7.2 Set waterfall's adsources

#### 7.2.1 Request URL

<https://openapi.toponad.com/v1/waterfall/set_units>

#### 7.2.2 Request method 

POST

#### 7.2.3 Request params

| params                              | type    | required | notes                                                        |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement ID                                                 |
| is_abtest             | Int    | N        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group   |
| segment_id                          | String  | Y        | segment ID                                                   |
| parallel_request_number             | Int     | Y        | parallel request number                                      |
| offer_switch                        | Int     | N        | my offer switch                                              |
| unbind_adsource_list                | Array   | N        | unbind the adsource and send only the adsource id            |
| ad_source_list                      | Array   | Y        | adsources need to binding                                    |
| ad_source_list.ad_source_id         | Int     | Y        | adsource ID                                                  |
| ad_source_list.ecpm                 | float64 | Y        | eCPM                                                         |
| ad_source_list.header_bidding_witch | Int     | N        | if support Header Bidding<br />1：not support，<br />2：support |
| ad_source_list.auto_switch          | Int     | Y        | 1：not open auto eCPM sort switch，<br />2：open auto eCPM sort switch |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：close                                           |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：close                                           |
| ad_source_list.pacing               | Int     | N        | Default -1 ：close                                           |

#### 7.2.4 Return data

| fields                              | type    | required | notes                                                        |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement ID                                                 |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | Whether it is a test group, default: 0 <br/> 0: control group, 1: test group  |
| parallel_request_number             | Int     | Y        | parallel request number                                      |
| offer_switch                        | Int     | N        | my offer switch                                              |
| unbind_adsource_list                | Array   | N        | unbind the adsource and send only the adsource id            |
| ad_source_list                      | Array   | Y        | adsources need to binding                                    |
| ad_source_list.ad_source_id         | Int     | Y        | adsource ID                                                  |
| ad_source_list.ecpm                 | float64 | Y        | eCPM                                                         |
| ad_source_list.header_bidding_witch | Int     | N        | if support Header Bidding<br />1：not support，<br />2：support |
| ad_source_list.auto_switch          | Int     | Y        | 1：not open auto eCPM sort switch，<br />2：open auto eCPM sort switch |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：close                                           |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：close                                           |
| ad_source_list.pacing               | Int     | N        | Default -1 ：close                                           |

#### 7.2.5 Sample

request sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1",
    "ad_source_list": [
        {
            "auto_switch": 1,
            "ad_source_id": "ad_source_id1",
            "ecpm": "ecpm1",
            "header_bidding_switch": 1,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        },
        {
            "auto_switch": 2,
            "ad_source_id": "ad_source_id2",
            "ecpm": "ecpm2",
            "header_bidding_switch": 1,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        }
    ]
}
```

return sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1",
    "ad_source_list": [
        {
            "priority": 1,
            "ad_source_id": "ad_source_id1",
            "ecpm": "ecpm1",
            "header_bidding_switch": 1,
            "auto_switch": 1,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        },
        {
            "priority": 2,
            "ad_source_id": "ad_source_id2",
            "ecpm": "ecpm2",
            "header_bidding_switch": 1,
            "auto_switch": 1,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        }
    ]
}
```

## 8. Network API

### 8.1 Create and update network publisher and app parameters

#### 8.1.1 Request URL

<https://openapi.toponad.com/v1/set_networks>

#### 8.1.2 Request method

POST

#### 8.1.3 Request params

| params                            | type   | required | notes                                                        |
| --------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| network_name                      | String | N        | the account name of network, which must be passed when opening multiple accounts     |
| nw_firm_id                        | Int    | Y        | network id              |
| network_id                        | Int    | N        | account id          |
| is_open_report                    | Int    | N        | whether to activate Report API  |
| auth_content                      | Object | N        | network publisher params    |
| network_app_info                  | Array  | N        | -     |
| network_app_info.app_id           | String | N        | TopOn app id              |
| network_app_info.app_auth_content | Object | N        | network app params,See:Appendix3_Detailed parameters of network |
 

#### 8.1.4 Return data

| fields                            | type   | required | notes                                                        |
| --------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| network_name                      | String | Y        | account name                   |
| nw_firm_id                        | Int    | Y        | network id                          |
| network_id                        | Int    | N        | account id                     |
| is_open_report                    | Int    | N        | whether to activate Report API             |
| auth_content                      | Object | N        | network publisher params    |
| network_app_info                  | Array  | N        | -                |
| network_app_info.app_id           | String | N        | TopOn app id                         |
| network_app_info.app_auth_content | Object | N        | network app params |
 

#### 8.1.5 Sample

request sample：
```
   {
        "network_name": "Default",
        "nw_firm_id": 2,
        "network_id": 226,
        "is_open_report": 2,
        "auth_content": {
            "account_id": "pub-1310074477383748",
            "oauth_key": "1/CW8VoZRbc5UCscXs3ddTXzwT9LQ71uFUMSE6iEwcRlk"
        },
        "network_app_info": [
            {
                "app_id": "a5bc9921f7fdb4",
                "app_auth_content": {
                    "app_id": "ca-app-pub-1310074477383748~6985182465"
                }
            }
        ]
    }
```


return sample：
```
{
    "network_name": "Default",
    "nw_firm_id": 2,
    "network_id": 226,
    "is_open_report": 2,
    "auth_content": {
        "account_id": "pub-1310074477383748",
        "oauth_key": "1/CW8VoZRbc5UCscXs3ddTXzwT9LQ71uFUMSE6iEwcRlk"
    },
    "network_app_info": [
        {
            "app_id": "a5bc9921f7fdb4",
            "app_auth_content": {
                "app_id": "ca-app-pub-1310074477383748~6985182465"
            }
        }
    ]
}
```

### 8.2 Get network publisher and app parameters

#### 8.2.1 Request URL

<https://openapi.toponad.com/v1/networks>

#### 8.2.2 Request method 

POST

#### 8.2.3 Request params

None

#### 8.2.4 Return data

| fields         | type   | required | notes                     |
| -------------- | ------ | -------- | ------------------------- |
| network_name                      | String | Y        | account name                   |
| nw_firm_id                        | Int    | Y        | network id                          |
| network_id                        | Int    | N        | account id                     |
| is_open_report                    | Int    | N        | whether to activate Report API             |
| auth_content                      | Object | N        | network publisher params    |
| network_app_info                  | Array  | N        | -                |
| network_app_info.app_id           | String | N        | TopOn app id                         |
| network_app_info.app_auth_content | Object | N        | network app params |


#### 8.2.5 Sample

return sample：
```
[
    {
        "network_name": "Default",
        "nw_firm_id": 1,
        "network_id": 307,
        "is_open_report": 2,
        "network_app_info": [
            {
                "app_id": "appid1",
                "app_auth_content": {
                    "app_id": "24234234",
                    "app_token": "1"
                }
            }
        ]
    },
    {
        "network_name": "24523423",
        "nw_firm_id": 1,
        "network_id": 1418,
        "is_open_report": 2,
        "network_app_info": [
            {
                "app_id": "appid2",
                "app_auth_content": {
                    "app_id": "232323",
                    "app_token": "1"
                }
            }
        ]
    }
]
```

## 9. Adsource API

### 9.1 Batch create and update adsource

#### 9.1.1 Request URL

<https://openapi.toponad.com/v1/set_units>

#### 9.1.2 Request method 

POST

#### 9.1.3 Request params

| params            | type   | required | notes                                                        |
| ----------------- | ------ | -------- | ------------------------------------------------------------ |
| count               | Int32  | Y        | adsource number                             |
| units               | Array  | Y        | -                         |
| units.network_id    | Int    | Y        | account id                       |
| units.adsource_id   | Int    | N        | adsource id,it must be passed on when it is updated|
| units.adsource_name | String | Y        | adsource name                 |
| units.adsource_token | Object | Y        | network unit params.See:Appendix3_Detailed parameters of network |
| units.placement_id  | String | Y        | TopOn placemtne id                     |
| units.default_ecpm  | String | Y        | adsource default ecpm                             |
| units.header_bidding_switch  | String | Y        | 1：on,2：off |

 

#### 9.1.4 Return data 

| fields   | type   | required | notes                               |
| -------- | ------ | -------- | ----------------------------------- |
| network_id    | Int    | N        | account id                       |
| adsource_id   | Int    | N        | adsource id               |
| adsource_name | String | Y        | adsource name                  |
| adsource_token | Object | Y        | network unit params |
| placement_id  | String | Y        | TopOn placemtne id                    |
| default_ecpm  | String | Y        | adsource default ecpm                       |

 

#### 9.1.5 Sample

request sample：
```
{
    "count": 2,
    "units": [
        {
            "network_id": 307,
            "adsource_name": "5234",
            "adsource_token": {
                "unit_id": "fasfasf",
                "is_video": "0",
                "personalized_template": "1",
                "size": "320x50",
                "layout_type": "1"
            },
            "placement_id": "b5bc993ab0966a",
            "default_ecpm": "69"
        },
        {
            "network_id": 225,
            "adsource_name": "5234",
            "adsource_id": 19759,
            "adsource_token": {
                "slot_id": "fasfasf",
                "is_video": "0",
                "personalized_template": "1",
                "size": "640x100",
                "layout_type": "1"
            },
            "placement_id": "b5bc993ab0966a",
            "default_ecpm": "69"
        }
    ]
}
```


return sample：
```
[
    {
        "network_id": 307,
        "adsource_id": 19743,
        "adsource_name": "23423423423",
        "adsource_token": {
            "size": "sdsd",
            "unit_id": "xcxc"
        },
        "placement_id": "12312312",
        "default_ecpm": "",
        "errors": "adsource_id error"
    },
    {
        "network_id": 307,
        "adsource_name": "asfdasdasd",
        "adsource_token": {
            "size": "asfasd",
            "unit_id": "asdasdafsdddd"
        },
        "placement_id": "123123123",
        "default_ecpm": "",
        "errors": "ad_source_name repeated"
    }
]
```

### 9.2 Get adsource list

#### 9.2.1 Request URL

<https://openapi.toponad.com/v1/units>

#### 9.2.2 Request method

POST

#### 9.2.3 Request params

| params  | type        | required | notes                  |
| ------- | ----------- | -------- | ---------------------- |
| network_firm_ids | Array[int32]  | N        | multiple values are supported        |
| app_ids          | Array[String] | N        | multiple values are supported        |
| placement_ids    | Array[String] | N        | multiple values are supported        |
| adsource_ids     | Array[int32]  | N        | multiple values are supported        |
| start            | int32         | N        | default value: 0 (cannot be used with the above parameters)        |
| limit            | int32         | N        | default value: 100, and the maximum is 100 at a time  (cannot be used with the above parameters)               |
| metrics          | Array[String] | N        | specify the returned fields from the ad_source_list. If you do not pass, all of them will be returned |


#### 9.2.4 Return data

| fields       | type   | required | notes                               |
| ------------ | ------ | -------- | ----------------------------------- |
| network_id                             | String | N        | account id        |
| network_name                           | String | N        | account name        |
| nw_firm_id                             | Int    | N        | network firm id              |
| adsource_id                            | Int    | N        | adsource id        |
| adsource_name                          | Int    | N        | adsource name       |
| adsource_token                          | Object | N        | adsource params |
| app_id                                 | String | N        | TopOn app id     |
| app_name                               | String | N        | TopOn app name   |
| platform                               | Int    | N        | platform |
| placement_id                           | String | N        | TopOn placement id              |
| placement_name                         | Object | N        | TopOn placement name |
| placement_format                       | String | N        | adformat                  |
| waterfall_list                         |  Array |   N      | the waterfall of adsource is being used               |
| waterfall_list.ecpm                    |   String     |      N    | adsource ecpm                   |
| waterfall_list.auto_ecpm               |   String     |     N     | adsource auto ecpm                   |
| waterfall_list.header_bidding_switch    |   Int     |     N     |  header bidding switch                |
| waterfall_list.auto_switch             |    Int    |      N    |  auto ecpm switch                   |
| waterfall_list.day_cap                 |   Int     |      N    |  daycap                   |
| waterfall_list.hour_cap                |  Int      |     N     |   hour cap                  |
| waterfall_list.pacing                  |   Int     |    N      |    pacing                 |
| waterfall_list.segment_name            |  String  |   N       |   segment name                  |
| waterfall_list.segment_id              |  String      |   N       |  segment_id                  |
| waterfall_list.priority                |   Int     |     N     |  segment priority                  |
| waterfall_list.parallel_request_number |   Int     |     N     |  parallel request number                |
| waterfall_list.is_abtest |   Int     |     N     |    Whether it is a test group, default: 0 <br/> 0: control group, 1: test group                |

 

#### 9.2.5 Sample

request sample：
```
{
	"adsource_ids":[19683]
}
```


return sample：
```
[
    {
        "nw_firm_id": 12,
        "network_name": "Default",
        "adsource_id": 19683,
        "adsource_name": "Unity Ads_int_2",
        "adsource_token": {
            "game_id": "234234",
            "placement_id": "23434"
        },
        "app_id": "232323",
        "app_name": "234234",
        "platform": 2,
        "placement_id": "234234234234",
        "placement_name": "234234234",
        "placement_format": "3",
        "waterfall_list": [
            {
                "ecpm": "1",
                "auto_ecpm": "",
                "header_bidding_switch": 1,
                "auto_switch": 1,
                "day_cap": 0,
                "hour_cap": 0,
                "pacing": 0,
                "name": "日韩",
                "segment_id": "2324234",
                "priority": 3,
                "parallel_request_number": 2
            },
            {
                "ecpm": "2",
                "auto_ecpm": "",
                "header_bidding_switch": 1,
                "auto_switch": 1,
                "day_cap": -1,
                "hour_cap": -1,
                "pacing": -1,
                "name": "ipad",
                "segment_id": "23423423423",
                "priority": 2,
                "parallel_request_number": 2
            }
        ]
    }
]
```

### 9.3 Batch delete adsource

#### 9.3.1 Request URL

<https://openapi.toponad.com/v1/del_units>

#### 9.3.2 Request method

POST

#### 9.3.3 Request params

| params  | type        | required | notes                  |
| ------- | ----------- | -------- | ---------------------- |
| adsource_ids | Array[Int32] | Y        | adsource_id |

#### 9.3.4 Return data

| params  | type        | required | notes                  |
| ------- | ----------- | -------- | ---------------------- |
| msg  | String | N        | result |

#### 9.3.5 Sample

request sample：

```
{
	"adsource_ids":[19683]
}
```


return sample：

```
{
    "msg": "suc"
}
```

## 10. Notices

Please control the frequency of requests:

•  1000 per hour

•  10000 per day

## Appendix1：APP category and sub category enum

| Platform | ategory | Sub Category            |
| -------- | ------- | ----------------------- |
| Android  | App     | Daydream                |
| Android  | App     | Android Wear            |
| Android  | App     | Art & Design            |
| Android  | App     | Auto & Vehicles         |
| Android  | App     | Beauty                  |
| Android  | App     | Books & Reference       |
| Android  | App     | Business                |
| Android  | App     | Comics                  |
| Android  | App     | Communication           |
| Android  | App     | Dating                  |
| Android  | App     | Education               |
| Android  | App     | Entertainment           |
| Android  | App     | Events                  |
| Android  | App     | Finance                 |
| Android  | App     | Food & Drink            |
| Android  | App     | Health & Fitness        |
| Android  | App     | House & Home            |
| Android  | App     | Libraries & Demo        |
| Android  | App     | Lifestyle               |
| Android  | App     | Maps & Navigation       |
| Android  | App     | Medical                 |
| Android  | App     | Music & Audio           |
| Android  | App     | News & Magazines        |
| Android  | App     | Parenting               |
| Android  | App     | Personalisation         |
| Android  | App     | Photography             |
| Android  | App     | Productivity            |
| Android  | App     | Shopping                |
| Android  | App     | Social                  |
| Android  | App     | Sports                  |
| Android  | App     | Tools                   |
| Android  | App     | Travel & Local          |
| Android  | App     | Video Players & Editors |
| Android  | App     | Weather                 |
| Android  | Game    | Action                  |
| Android  | Game    | Adventure               |
| Android  | Game    | Arcade                  |
| Android  | Game    | Board                   |
| Android  | Game    | Card                    |
| Android  | Game    | Casino                  |
| Android  | Game    | Casual                  |
| Android  | Game    | Educational             |
| Android  | Game    | Music                   |
| Android  | Game    | Puzzle                  |
| Android  | Game    | Racing                  |
| Android  | Game    | Role Playing            |
| Android  | Game    | Simulation              |
| Android  | Game    | Sports                  |
| Android  | Game    | Strategy                |
| Android  | Game    | Trivia                  |
| Android  | Game    | Word                    |
| Android  | Family  | Ages 5 & Under          |
| Android  | Family  | Ages 6-8                |
| Android  | Family  | Ages 9 & Over           |
| Android  | Family  | Action & Adventure      |
| Android  | Family  | Brain Games             |
| Android  | Family  | Creativity              |
| Android  | Family  | Education               |
| Android  | Family  | Music and video         |
| Android  | Family  | Pretend play            |
| iOS      | Game    | Action                  |
| iOS      | Game    | Adventure               |
| iOS      | Game    | Arcade                  |
| iOS      | Game    | Board                   |
| iOS      | Game    | Card                    |
| iOS      | Game    | Casino                  |
| iOS      | Game    | Dice                    |
| iOS      | Game    | Educational             |
| iOS      | Game    | Family                  |
| iOS      | Game    | Music                   |
| iOS      | Game    | Puzzle                  |
| iOS      | Game    | Racing                  |
| iOS      | Game    | Role Playing            |
| iOS      | Game    | Simulation              |
| iOS      | Game    | Sports                  |
| iOS      | Game    | Strategy                |
| iOS      | Game    | Trivia                  |
| iOS      | Game    | Word                    |
| iOS      | App     | Books                   |
| iOS      | App     | Business                |
| iOS      | App     | Catalogs                |
| iOS      | App     | Education               |
| iOS      | App     | Entertainment           |
| iOS      | App     | Finance                 |
| iOS      | App     | Food & Drink            |
| iOS      | App     | Health & Fitness        |
| iOS      | App     | Lifestyle               |
| iOS      | App     | Magazines & Newspapers  |
| iOS      | App     | Medical                 |
| iOS      | App     | Music                   |
| iOS      | App     | Navigation              |
| iOS      | App     | News                    |
| iOS      | App     | Photo & Video           |
| iOS      | App     | Productivity            |
| iOS      | App     | Reference               |
| iOS      | App     | Shopping                |
| iOS      | App     | Social Networking       |
| iOS      | App     | Sports                  |
| iOS      | App     | Stickers                |
| iOS      | App     | Travel                  |
| iOS      | App     | Utilities               |
| iOS      | App     | Weather                 |

## Appendix2：segment rule enum

| rule | type                           | sample                                     |
| :--- | :----------------------------- | :----------------------------------------- |
| 0    | include（set）                 | one dimension JSON ["CN", "US"]            |
| 1    | exclude（set）                 | one dimension JSON [1,2,3]                 |
| 2    | Greater than or equal（value） | int or float 124                           |
| 3    | Less than or equal（value）    | int or float 222.36                        |
| 4    | in interval（interval）        | two dimension JSON [[122,456],[888,12322]] |
| 5    | not in interval（interval）    | two dimension JSON [[122,456],[888,12322]] |
| 6    | custom rule（custom）          | bb=1&c!=3&p=3                              |
| 7    | Greater than（value）          | int,float or string 124                    |
| 8    | Less than（value）             | int,float or string 222.36                 |

## Appendix3：Detailed parameters of network

All Key and Value data types are String

| network firm id | network firm name | auth_content | app_auth_content | adformat | adsource_token |  key-value  |
| --------- | ----------- | ------------ | ---------------- | ------- | -------------  | ---------------- |
| 1         | Facebook    | - | app_id<br>app_token | native<br>rewarded_video<br>interstitial | unit_id | app_id：AppID <br> app_token：AccessToken <br> unit_id：PlacementID |
| 1         | Facebook    | - | app_id<br>app_token | bannner | unit_id<br>size | size：320x50,320x90,320x250 |
| 2         | Admob       | account_id<br>oauth_key | app_id | native<br>rewarded_video<br>interstitial | unit_id | account_id：PublisherID <br/> oauth_key：AccessToken <br/> app_id：AppID <br/> unit_id：UnitID |          
| 2         | Admob       | account_id<br>oauth_key | app_id | bannner | unit_id<br>size | size：320x50,320x100,320x250,468x60,728x90 |
| 3         | Inmobi      | username<br>password<br>apikey<br>app_id | - | native<br>rewarded_video<br>interstitial | unit_id |    username：EmailID </br> app_id：Account ID </br> password：Password </br> apikey：API Key </br> unit_id：Placement ID | 
| 3         | Inmobi      | username<br>password<br>apikey<br>app_id | - | bannner | unit_id<br>size | size：320x50 |  
| 4         | Flurry      | token | sdk_key | native<br>rewarded_video<br>interstitial | ad_space | token：Token </br> sdk_key：API Key </br> ad_space：AD Unit Name |  
| 4         | Flurry      | token | sdk_key | banner | ad_space<br>size | size：320x50 |  
| 5         | Applovin    | sdkkey<br>apikey | - | native | - | sdkkey：SDK Key </br> apikey：Report Key  | 
| 5         | Applovin    | sdkkey<br>apikey | -  | rewarded_video<br>interstitial | zone_id | zone_id：Zone ID |  
| 5         | Applovin    | sdkkey<br>apikey | -  | banner | zone_id<br>size | size：320x50,300x250  | 
| 6         | Mintegral   | skey<br>secret<br>appkey | app_id | native<br>rewarded_video  | unit_id | appkey：App Key </br> skey：Skey </br> secret：Secret </br> appid：AppID </br> unitid：UnitID | 
| 6         | Mintegral   | skey<br>secret<br>appkey | app_id | bannner | unit_id<br>size | size：320x50,300x250,320x90,smart  | 
| 6         | Mintegral   | skey<br>secret<br>appkey | app_id | interstitial | unit_id<br>is_video | is_video：0,1 | 
| 7         | Mopub       | repkey<br>apikey | - | native<br>rewarded_video<br>interstitial | unit_id | repkey：Inventory Report ID </br> apikey：API Key </br> unitid：Unit ID  |
| 7         | Mopub       | repkey<br>apikey | -  | bannner | unit_id<br>size | size：320x50,300x250,728x90  | 
| 8         | Tencent Ads     | agid<br>publisher_id<br>app_key<br>qq | app_id | native | unit_id<br>unit_version<br>unit_type | qq：MemberID </br> agid：AGID </br> publisher_id：App ID </br> app_key：App Key </br> app_id：Media ID </br> unit_id：UnitID</br>unit_version：1,2</br>unit_type：1,2 |
| 8         | Tencent Ads     | agid<br>publisher_id<br>app_key<br>qq | app_id | rewarded_video,splash | unit_id | - |
| 8         | Tencent Ads     | agid<br>publisher_id<br>app_key<br>qq | app_id | bannner| unit_id<br>unit_version<br>size | unit_version：2</br> size：320x50 | 
| 8         | Tencent Ads     | agid<br>publisher_id<br>app_key<br>qq | app_id | interstitial | unit_id<br>unit_version<br>video_muted<br>video_autoplay<br>video_duration<br>is_fullscreen | video_duration_switch：videoDuration</br>unit_version：2</br> video_muted：0,1 </br>video_autoplay：0,1</br> video_duration：optional</br>is_fullscreen：0，1 |
| 9         | Chartboost  | user_id<br>user_signature | app_id<br>app_signature | rewarded_video<br>interstitial | location | user_id：UserID </br> user_signature：UserSignature </br> app_id：UserAppID </br> app_signature：AppSignature </br> location：Location |  
| 10        | Tapjoy      | apikey | sdk_key | rewarded_video<br>interstitial | placement_name | apikey：APIKey </br> sdk_key：SDKKey </br> placement_name：PlacementName |  
| 11        | Ironsource  | username<br>secret_key | app_key | rewarded_video<br>interstitial | instance_id |   username：Username </br> secret_key：Secret Key </br> app_key：App Key </br> instance_id：Instance ID |  
| 12        | UnityAds    | apikey | game_id | rewarded_video<br>interstitial | placement_id | apikey：API Key </br> organization_core_id：Organization core ID </br> game_id：Game ID </br> placement_id：Placement ID |  
| 13        | Vungle      | apikey | app_id | rewarded_video<br>interstitial | placement_id | apikey：Reporting API Key </br> app_id：App ID </br> placement_id：PlacementID |  
| 14        | AdColony    | user_credentials | app_id | rewarded_video<br>interstitial | zone_id | user_credentials：Read-Only API key </br> app_id：App ID </br> zone_id：Zone ID |  
| 15        | Pangle(Tiktok Ads)       | user_id<br>secure_key | app_id | native | slot_id<br>is_video<br>layout_type<br>media_size | user_id：UserID </br> secure_key：Secure Key </br> app_id：AppID </br> slot_id：SlotID </br> is_video：0,1,2,3 <br> layout_type：0,1 </br> media_size（when layout_type = 1 required）：1,2 |  
| 15        | Pangle(Tiktok Ads)       | user_id<br>secure_key | app_id | rewarded_video | slot_id<br>personalized_template | personalized_template：0,1 |  
| 15        | Pangle(Tiktok Ads)      | user_id<br>secure_key | app_id | banner | slot_id<br>layout_type<br>size | layout_type：1 </br> size：640x100,600x90,600x150,600x500,600x400,600x300,600x260,690x388 |  
| 15        | Pangle(Tiktok Ads)       | user_id<br>secure_key | app_id | interstitial | slot_id<br>is_video<br>layout_type<br>size<br>personalized_template | when is_video=0 required<br>layout_type：1 <br> size：1:1,3:2,2:3 </br> when is_video=1 required<br>personalized_template：0,1 | 
| 15        | Pangle(Tiktok Ads)       | user_id<br>secure_key | app_id | splash | slot_id<br>personalized_template | personalized_template：0,1 | 
| 16        | Joomob     | - | - | rewarded_video<br>interstitial | app_id | app_id：App ID |  
| 16        | Joomob     | - | - | banner | app_id<br>size | size：320x50,480x75,640x100,960x150,728x90 |
| 17        | OneWay      | access_key | publisher_id | rewarded_video<br>interstitial | slot_id | access_key：Access Key </br> publisher_id：Publisher ID </br> slot_id：Placement ID |  
| 18        | MobPower    | publisher_id<br>api_key  | app_id | native<br>rewarded_video<br>interstitial | placement_id | api_key：API Key </br> publisher_id：Publisher ID </br> app_id：App ID </br> placement_id：Placement ID |  
| 18        | MobPower    | publisher_id<br>api_key  | app_id | banner | placement_id<br>size | size：320x50 |  
| 19        | Kingsoft       | - | media_id | rewarded_video | slot_id | media_id：Media ID </br> slot_id：Slot ID |  
| 21        | AppNext     | email<br>password<br>key  | - | native<br>rewarded_video<br>interstitial | placement_id | email：Email </br> password：Password </br> key：Key </br> placement_id：Placement ID |  
| 21        | AppNext     | email<br>password<br>key  | - | banner | placement_id<br>size | size：320x50,320x100,300x250 | 
| 22        | Baidu       | access_key | app_id | native<br>rewarded_video<br>interstitial<br>splash | ad_place_id | access_key：Access Key </br> app_id：AppID </br> ad_place_id：ADPlaceID |  
| 22        | Baidu       | access_key | app_id | banner | ad_place_id<br>size | size：375x56,200x30,375x250,200x133,375x160,200x85,375x187,200x100 | 
| 23        | Nend        | api_key | - | naitve | spot_id<br>api_key<br>is_video | api_key：APIKey </br> spot_id：spotID </br> is_video：0,1 |  
| 23        | Nend        | api_key | - | rewarded_video | spot_id<br>api_key | - | 
| 23        | Nend        | api_key | - | banner | spot_id<br>api_key<br>size | size：320x50,320x100,300x100,300x250,728x90 | 
| 23        | Nend        | api_key | - | interstitial | spot_id<br>api_key<br>is_video | is_video：0,1,2 | 
| 24        | Maio        | api_id<br>api_key | media_id | rewarded_video<br>interstitial | zone_id | api_id：API ID </br> api_key：API Key </br> media_id：Media ID </br> zone_id：Zone ID |  
| 25        | StartAPP    | partner_id<br>token  | app_id | rewarded_video<br>interstitial | ad_tag | partner_id：Partner ID </br> token：Token </br> app_id：APP ID </br> ad_tag：AD Tag |  
| 26        | SuperAwesome | - | property_id | rewarded_video | placement_id | property_id：Property ID </br> placement_id：Placement ID |  
| 28        | Kuaishou Ads        | access_key<br>security_key | app_id<br>app_name | native | position_id<br>layout_type<br>video_sound<br>is_video<br>unit_type | access_key：Access Key </br> security_key：Security Key </br> app_id：AppID </br> app_name：AppName </br> position_id：PosID </br> unit_type：0,1<br>when unit_type=1 required<br>layout_type：0<br>is_video：0,1<br>video_sound：0,1 |  
| 28        | Kuaishou Ads        | access_key<br>security_key | app_id<br>app_name | rewarded_video<br>interstitial | orientation | orientation：1,2 | 
| 29        | Sigmob      | public_key<br>secret_key  | app_id<br>app_key | rewarded_video<br>interstitial<br>splash | placement_id |    public_key：Public Key </br> secret_key：Secret Key </br> app_id：AppID </br> app_key：App Key </br> placement_id：PlacementID |  
| 36        | Ogury       | api_key<br>api_secret | key | rewarded_video<br>interstitial | unit_id | api_key：API KEY </br> api_secret：API SECRET </br> key：KEY </br> unit_id：AD Unit ID |  
