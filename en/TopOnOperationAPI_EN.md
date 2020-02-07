# TopOn Opration API

## Change Log


| version | date  | notes                     |
| :--------: | ------------- | -------------------------------- |
| v 1.0    | 2019/7/17 | supports create and query apps and placements |
| v 2.0    | 2019/11/4 | supports operate waterfall and segments |


## Contents

[1. Introduction](#Introduction)</br> 
[2. Authentication acquisition](#Authentication_acquisition)</br> 
[3. Authentication check](#Authentication_check)</br> 
[4. APP API](#APP_API)</br> 
- [4.1 Batch create APPs](#Batch_create_APPs)</br>  
- [4.2 Get APP list](#Get_APP_list)</br>

[5. Placement API](#Placement_API)</br>
- [5.1 Batch create placements](#Batch_create_placements)</br>  
- [5.2 Get placement list](#Get_placement_list)</br>  

[6. Segment API](#Segment_API)</br>
- [6.1 Create and update segments](#Create_and_update_segments)</br>
- [6.2 Get segment list](#Get_segment_list)</br>
- [6.3 Batch delete segments](#Batch_delete_segments)</br>

[7. Waterfall API](#Waterfall_API)</br>
- [7.1 Get placement's segment list](#Get_placements_segment_list)</br>  
- [7.2 Set priorities or create segments for placements](#Set_priorities_or_create_segments_for_placements)</br>
- [7.3 Batch delete placement's segments](#Batch_delete_placements_segments)</br>
- [7.4 Get waterfall's adsources](#Get_waterfalls_adsources)</br>  
- [7.5 Set waterfall's adsources](#Set_waterfalls_adsources)</br>

[8. Notices](#Notices)</br>
[9. Appendix1：golang demo](#Appendix1：golang_demo)</br>
[10. Appendix2：APP category and sub category enum](#Appendix2：APP_category_and_sub_category_enum)</br>
[11. Appendix3：segment rule enum](#Appendix3：segment_rule_enum)

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
| Headers      | Headers except X-Up-Signature              | X-Up-Timestamp:1562813567000 X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2 |
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

<h2 id='APP_API'>4. APP API</h2>

<h3 id='Batch_create_APPs'>4.1 Batch create APPs</h3>

#### 4.1.1 Request URL

<https://openapi.toponad.com/v1/create_app>

#### 4.1.2 Request method 

POST

#### 4.1.3 Request params

| params            | type   | required | notes                                                        |
| ----------------- | ------ | -------- | ------------------------------------------------------------ |
| count             | Int    | Y        | Quantity of created APPs                                     |
| apps.app_name     | String | Y        | APP name                                                     |
| apps.platform     | Int    | Y        | platform 1 or 2  (1:android，2:iOS)                          |
| apps.market_url   | String | N        | Need to be in compliance with requirements of app store links |
| apps.package_name | String | N        | Need to be in compliance with requirements of APP package name.  com.xxx |
| apps.category     | String | N        | category.[Appendix2：APP category and sub category enum](#Appendix2：APP_category_and_sub_category_enum) |
| apps.sub_category | String | N        | sub category.[Appendix2：APP category and sub category enum](#Appendix2：APP_category_and_sub_category_enum) |

 

#### 4.1.4 Return data 

| fields   | type   | required | notes                               |
| -------- | ------ | -------- | ----------------------------------- |
| app_id   | String | Y        | APP ID                              |
| app_name | String | Y        | APP name                            |
| errors   | String | N        | error messages                      |
| platform | Int    | Y        | platform 1 or 2  (1:android，2:iOS) |

 

#### 4.1.5 Sample

request sample：
```
{
    "count": 1,
    "apps": [
        {
            "app_name": "111",
            "platform": 1,
            "market_url": ""
        }
    ]
}
```


return sample：
```
[
    {
        "app_name": "111",
        "errors": "app package name is required"  
    }
]
```

<h3 id='Get_APP_list'>4.2 Get APP list</h3>

#### 4.2.1 Request URL

<https://openapi.toponad.com/v1/apps>

#### 4.2.2 Request method

POST

#### 4.2.3 Request params

| params  | type        | required | notes                  |
| ------- | ----------- | -------- | ---------------------- |
| app_ids | string List | N        | ["abc", "acc"]         |
| start   | Int         | N        | Default 0              |
| limit   | Int         | N        | Default 100 , [0, 100] |

 

#### 4.2.4 Return data

| fields       | type   | required | notes                               |
| ------------ | ------ | -------- | ----------------------------------- |
| app_id       | String | Y        | APP ID                              |
| app_name     | String | Y        | APP name                            |
| platform     | Int    | Y        | platform 1 or 2  (1:android，2:iOS) |
| market_url   | String | N        | -                                   |
| package_name | String | N        | -                                   |
| category     | String | N        | -                                   |
| sub-category | String | N        | -                                   |

 

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
        "app_name": "topontest",
        "app_id": "a5bc9921f7fdb4",
        "platform": 2,
        "market_url": "https://itunes.apple.com/cn/app/%E7%A5%9E%E5%9B%9E%E9%81%BF/id1435756371?mt=8",
        "category": "Game",
        "sub_category": "Action"
    }
]
```

<h2 id='Placement_API'>5. Placement API</h2>

<h3 id='Batch_create_placements'>5.1 Batch create placements</h3>

#### 5.1.1 Request URL

<https://openapi.toponad.com/v1/create_placement>

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
| placements.template.ski_swt           | Int    | N        | Template is Native Splash：it can skipped or not, it could be skipped by default.<br/>0：No<br/>1：Yes |
| placements.template.aut_swt           | Int    | N        | Template is Native Splash：it can be auto closed or not, it could be auto closed by default.<br/>0：No<br/>1：Yes |
| placements.template.auto_refresh_time | Int    | N        | template is Native Banner：it can be auto refreshed or not, it could not be auto refreshed by default<br/>-1 no auto refresh<br/>0-n auto refresh time (s) |

 

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

 

#### 5.1.5 Sample

request sample：
```
{
    "count": 1, 
    "app_id": "a5bc9921f7fdb4",
    "placements": [
        {
            "placement_name": "xxx",
            "adformat": "native"
        }
    ]
}
```


return sample：
```
[
    {
        "app_name": "",
        "app_id": "a5bc9921f7fdb4",
        "platform": 0,
        "placement_name": "xxx", 
        "adformat": "native"   
    }
]
```

<h3 id='Get_placement_list'>5.2 Get placement list</h3>

#### 5.2.1 Request URL

<https://openapi.toponad.com/v1/placements>

#### 5.2.2 Request method 

POST

#### 5.2.3 Request params

| params        | type        | required | notes              |
| ------------- | ----------- | -------- | ------------------ |
| app_ids       | string List | N        | eg: ["abc", "acc"] |
| placement_ids | string List | N        | eg: ["abc", "acc"] |
| start         | Int         | N        | Default 0          |
| limit         | Int         | N        | Default 100        |

 

#### 5.2.4 Return data

| fields         | type   | required | notes                     |
| -------------- | ------ | -------- | ------------------------- |
| app_id         | String | Y        | APP ID                    |
| app_name       | String | Y        | app name                  |
| platform       | Int    | Y        | 1 or 2  (1:android，2IOS) |
| placements     | String | Y        | -                         |
| placement_id   | String | N        | -                         |
| placement_name | String | N        | -                         |
| adformat       | String | N        | -                         |

 

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

<h2 id='Segment_API'>6. Segment API</h2>

<h3 id='Create_and_update_segments'>6.1 Create and update segments</h3>

#### 6.1.1 Request URL

<https://openapi.toponad.com/v1/deal_segment>

#### 6.1.2 Request method 

POST

#### 6.1.3 Request params

| params        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | Segment name                                                 |
| segment_id    | String | N        | Must reture Segment ID when updating Segment                 |
| rules         | Array  | Y        | Segment rules                                                |
| rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| rules.content | string | Y        | [Appendix3：segment rule enum](#Appendix3：segment_rule_enum) |

#### 6.1.4 Return data

| fields        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | Segment name                                                 |
| segment_id    | String | Y        | Segment ID                                                   |
| rules         | Array  | Y        | Segment rules                                                |
| rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| rules.content | string | Y        | [Appendix3：segment rule enum](#Appendix3：segment_rule_enum) |



#### 6.1.5 Sample

request sample：

```
{
    "name": "segment1",
    "rules": [
        {
            "type": 1,
            "rule": 1,
            "content": "sdsd"
        }
    ]
}
```

 

return sample：

```
{
    "name": "segment1",
    "segment_id": "asasdsdsd",
    "rules": [
        {
            "type": 1,
            "rule": 1,
            "content": "sdsd"
        }
    ]
}
```

<h3 id='Get_segment_list'>6.2 Get segment list</h3>

#### 6.2.1 Request URL

<https://openapi.toponad.com/v1/segment_list>

#### 6.2.2 Request method 

POST

#### 6.2.3 Request params

| params  | type | required | notes                                                     |
| ----------- | ------ | -------- | ------------------------------------------------------------ |
| segment_ids | string List | N        | ["uuid1","uuid2"]          |
| start       | Int    | N        | Default 0                   |
| limit       | Int    | N        | Default 100 |

 

#### 6.2.4 Return data

| fields        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | Segment name                                                 |
| segment_id    | String | Y        | Segment ID                                                   |
| rules         | Array  | Y        | Segment rules                                                |
| rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| rules.content | string | Y        | [Appendix3：segment rule enum](#Appendix3：segment_rule_enum) |



#### 6.2.5 Sample

request sample：

```
{
   "segment_ids":["uuid1","uuid2"]
}
```

return sample：

```
[
    {
        "name": "segment1",
        "segment_id": "asasdsdsd",
        "rules": [
            {
                "type": 1,
                "rule": 1,
                "content": "sdsd"
            }
        ]
    },
    {
        "name": "segment2",
        "segment_id": "uuid2",
        "rules": [
            {
                "type": 1,
                "rule": 1,
                "content": "sdsd"
            }
        ]
    }
]
```

<h3 id='Batch_delete_segments'>6.3 Batch delete segments</h3>

#### 6.3.1 Request URL

<https://openapi.toponad.com/v1/del_segment>

#### 6.3.2 Request method 

POST

#### 6.3.3 Request params

| params  | type | required | notes                         |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | string List | Y        | ["uuid1","uuid2"] |

 

#### 6.3.4 Return data

It will return HTTP code 200 when success Otherwise, it will return segments data. It could not be deleted if the segment has been used in the waterfall setting and all the segment list in this request will failed to be deleted.

#### 6.3.5 Sample

request sample：

```
{
   "segment_ids":["uuid1","uuid2"]
}
```

return sample：

HTTP code 200

<h2 id='Waterfall_API'>7. Waterfall API</h2>

<h3 id='Get_placements_segment_list'>7.1 Get placement's segment list</h3>

#### 7.1.1 Request URL

<https://openapi.toponad.com/v1/waterfall/segment>

#### 7.1.2 Request method 

GET

#### 7.1.3 Request params

| params   | type | required | notes                           |
| ------------ | ------ | -------- | --------------------------------- |
| placement_id | String | Y        | placement ID                |
| is_abtest    | Int    | Y        | 0 : control group, or not activate ab test<br/>1 test group |

#### 7.1.4 Return data

| fields        | type   | required | notes                                                        |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| priority      | Int    | Y        | priority                                                     |
| name          | String | Y        | Segment name                                                 |
| segment_id    | String | Y        | Segment ID                                                   |
| rules         | Array  | Y        | Segment rules                                                |
| rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| rules.content | string | Y        | [Appendix3：segment rule enum](#Appendix3：segment_rule_enum) |

#### 7.1.5 Sample

request sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1
}
```

return sample：

```
[
    {
        "name": "segment1",
        "segment_id": "segment_id1",
        "priority": 1,
        "rules": [
            {
                "type": 1,
                "rule": 1,
                "content": "sdsd"
            }
        ]
    },
    {
        "name": "segment2",
        "segment_id": "segment_id2",
        "priority": 2,
        "rules": [
            {
                "type": 1,
                "rule": 1,
                "content": "sdsd"
            }
        ]
    }
]
```

<h3 id='Set_priorities_or_create_segments_for_placements'>7.2 Set priorities or create segments for placements</h3>

#### 7.2.1 Request URL

<https://openapi.toponad.com/v1/waterfall/set_segment>

#### 7.2.2 Request method 

POST

#### 7.2.3 Request params

|                 | type | required | notes                      |
| ------------------- | ------ | -------- | ---------------------------- |
| placement_id        | String | Y        | placement ID           |
| is_abtest           | Int    | Y        | 0 : control group, or not activate ab test<br/>1 test group |
| segments            | Array  | Y        | Segment priority List |
| segments.priority   | Int    | Y        | Segment priority     |
| segments.segment_id | String | Y        | Segment ID                   |

#### 7.2.4 Return data

| fields                 | type   | required | notes                                                        |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id           | String | Y        | placement ID                                                 |
| is_abtest              | Int    | Y        | 0 : control group, or not activate ab test<br/>1 test group  |
| segments.priority      | Int    | Y        | priority                                                     |
| segments.name          | String | Y        | Segment name                                                 |
| segments.segment_id    | String | Y        | Segment ID                                                   |
| segments.rules         | Array  | Y        | Segment rules                                                |
| segments.rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| segments.rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| segments.rules.content | string | Y        | [Appendix3：segment rule enum](#Appendix3：segment_rule_enum) |

#### 7.2.5 Sample

request sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segments": [
        {
            "priority": 1,
            "segment_id": "segment_id1"
        },
        {
            "priority": 2,
            "segment_id": "segment_id2"
        }
    ]
}
```

return sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segments": [
        {
            "priority": 1,
            "segment_id": "segment_id1",
            "name": "name1",
            "rules": [
                {
                    "type": 1,
                    "rule": 1,
                    "content": "sdsd"
                }
            ]
        },
        {
            "priority": 2,
            "segment_id": "segment_id2",
            "name": "name2",
            "rules": [
                {
                    "type": 1,
                    "rule": 1,
                    "content": "sdsd"
                }
            ]
        }
    ]
}
```

<h3 id='Batch_delete_placements_segments'>7.3 Batch delete placement's segments</h3>

#### 7.3.1 Request URL

<https://openapi.toponad.com/v1/waterfall/del_segment>

#### 7.3.2 Request method 

POST

#### 7.3.3 Request params

| params   | type | required | notes                      |
| ------------ | ------ | -------- | ---------------------------- |
| placement_id | String | Y        | placement ID           |
| is_abtest    | Int    | Y        | 0 : control group, or not activate ab test<br/>1 test group |
| segment_ids  | string List | Y        | delete Segment List |

#### 7.3.4 Return data

|                        | type   | required | notes                                                        |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id           | String | Y        | placement ID                                                 |
| is_abtest              | Int    | Y        | 0 : control group, or not activate ab test<br/>1 test group  |
| segments.priority      | Int    | Y        | priority                                                     |
| segments.name          | String | Y        | Segment name                                                 |
| segments.segment_id    | String | Y        | Segment ID                                                   |
| segments.rules         | Array  | Y        | Segment rules                                                |
| segments.rules.type    | Int    | Y        | segment rule type.Default 0 <br />0 country code（set）<br/>1 time（interval）<br/>2 weekday（set）<br/>3 network_type（set）<br/>4 hour/1225/2203（interval）<br/>5 custom rule（custom）<br/>8 app version （set）<br/>9 sdk version （set）<br/>10 device_type （set）<br/>11 device brand（set）<br/>12 os version （set）<br/>16 timezone (value)<br/>17 Device ID （set）<br/>19 city code （set） |
| segments.rules.rule    | Int    | Y        | segment rule action.Default 0<br />0 include（set）<br/>1 exclude（set）<br/>2 Greater than or equal（value）<br/>3 Less than or equal（value）<br/>4 in interval（interval）<br/>5 not in interval（interval）<br/>6 custom rule（custom）<br/>7 Greater than（value）<br/>8 Less than（value） |
| segments.rules.content | string | Y        | [Appendix3：segment rule enum](#Appendix3：segment_rule_enum) |

#### 7.3.5 Sample

request sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_ids": [
        "segment_id1",
        "segment_id2"
    ]
}
```

return sample：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segments": [
        {
            "priority": 1,
            "segment_id": "segment_id1",
            "name": "name1",
            "rules": [
                {
                    "type": 1,
                    "rule": 1,
                    "content": "sdsd"
                }
            ]
        },
        {
            "priority": 2,
            "segment_id": "segment_id2",
            "name": "name2",
            "rules": [
                {
                    "type": 1,
                    "rule": 1,
                    "content": "sdsd"
                }
            ]
        }
    ]
}
```

<h3 id='Get_waterfalls_adsources'>7.4 Get waterfall's adsources</h3>

#### 7.4.1 Request URL

<https://openapi.toponad.com/v1/waterfall/units>

#### 7.4.2 Request method 

GET

#### 7.4.3 Request params

| params   | type | required | notes         |
| ------------ | ------ | -------- | --------------- |
| placement_id | String | Y        | placement ID |
| segment_id   | String | Y        | Segment ID      |
| is_abtest    | Int    | Y        | 0 : control group, or not activate ab test<br/>1 test group |

#### 7.4.4 Return data

| fields                              | type    | required | notes                                                        |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement ID                                                 |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 : control group, or not activate ab test<br/>1 test group  |
| ad_source_list                      | Array   | Y        | empty means has no adsource                                  |
| ad_source_list.ad_source_id         | Int     | N        | adsource ID                                                  |
| ad_source_list.ecpm                 | float64 | N        | eCPM                                                         |
| ad_source_list.pirority             | Int     | N        | adsource pirority                                            |
| ad_source_list.header_bidding_witch | Int     | N        | if support Header Bidding<br />0：not support，<br />1：support |
| ad_source_list.auto_switch          | Int     | N        | 0：not open auto eCPM sort switch，<br />1：open auto eCPM sort switch |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：close                                           |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：close                                           |
| ad_source_list.pacing               | Int     | N        | Default -1 ：close                                           |

#### 7.4.5 Sample

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

<h3 id='Set_waterfalls_adsources'>7.5 Set waterfall's adsources</h3>

#### 7.5.1 Request URL

<https://openapi.toponad.com/v1/waterfall/set_units>

#### 7.5.2 Request method 

POST

#### 7.5.3 Request params

| params                              | type    | required | notes                                                        |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement ID                                                 |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 : control group, or not activate ab test<br/>1 test group  |
| ad_source_list                      | Array   | Y        | adsources need to binding                                    |
| ad_source_list.ad_source_id         | Int     | Y        | adsource ID                                                  |
| ad_source_list.ecpm                 | float64 | Y        | eCPM                                                         |
| ad_source_list.header_bidding_witch | Int     | N        | if support Header Bidding<br />0：not support，<br />1：support |
| ad_source_list.auto_switch          | Int     | Y        | 0：not open auto eCPM sort switch，<br />1：open auto eCPM sort switch |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：close                                           |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：close                                           |
| ad_source_list.pacing               | Int     | N        | Default -1 ：close                                           |

#### 7.5.4 Return data

| fields                              | type    | required | notes                                                        |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement ID                                                 |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 : control group, or not activate ab test<br/>1 test group  |
| ad_source_list                      | Array   | Y        | adsources need to binding                                    |
| ad_source_list.ad_source_id         | Int     | Y        | adsource ID                                                  |
| ad_source_list.ecpm                 | float64 | Y        | eCPM                                                         |
| ad_source_list.pirority             | Int     | N        | adsource pirority                                            |
| ad_source_list.header_bidding_witch | Int     | N        | if support Header Bidding<br />0：not support，<br />1：support |
| ad_source_list.auto_switch          | Int     | Y        | 0：not open auto eCPM sort switch，<br />1：open auto eCPM sort switch |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：close                                           |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：close                                           |
| ad_source_list.pacing               | Int     | N        | Default -1 ：close                                           |

#### 7.5.5 Sample

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
            "header_bidding_witch": 0,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        },
        {
            "auto_switch": 2,
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
            "auto_switch": 0,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        },
        {
            "priority": 2,
            "ad_source_id": "ad_source_id2",
            "ecpm": "ecpm2",
            "header_bidding_witch": 0,
            "auto_switch": 0,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        }
    ]
}
```

<h2 id='Notices'>8. Notices</h2>

Please control the frequency of requests:

•  1000 per hour

•  10000 per day

<h2 id='Appendix1：golang_demo'>9. Appendix1：golang demo</h2>

 Java,PHP,Python demos are in the Git path /demo

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
	
	demoUrl := "Request URL"
	
	//request body
	
	body := "{}"
	
	//your publisherKey
	
	publisherKey := "your publisherKey"
	
	//Request method
	
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

	//return data
	
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
<h2 id='Appendix2：APP_category_and_sub_category_enum'>10. Appendix2：APP category and sub category enum</h2>

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

<h2 id='Appendix3：segment_rule_enum'>11. Appendix3：segment rule enum</h2>

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
