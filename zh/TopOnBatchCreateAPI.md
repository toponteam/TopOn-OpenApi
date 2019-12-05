# TopOn平台批量创建API对接文档

修订历史


| 文档版本 | 发布时间      | 修订说明                         |
| -------- | ------------- | -------------------------------- |
| v 1.0    | 2019年7月17日 | 支持批量创建和查询应用、广告位   |
| v 2.0    | 2019年11月4日 | 支持waterfall 相关配置（草稿版） |

 

 

## 1. 关于文档

为提高合作伙伴的变现效率，TopOn平台专门提供了批量创建应用和广告位，以及查询应用和广告位状态的API接口。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

## 2. 申请开通权限

在使用TopOn平台的批量创建 API 前，合作伙伴需向TopOn申请 publisher_key，用于识别来自合作伙伴的请求，申请方法请咨询与您对接的商务经理。

##3. 接口相关

### 3.1. 接口请求流程说明

l 请求端根据 API 请求内容（包括 HTTP Header 和 Body）生成签名字符串。

l 请求端使用MD5对第一步生成的签名字符串进行签名，形成该 API 请求的数字签名。

l 请求端把 API 请求内容和数字签名一同发送给服务端。

l 服务端在接到请求后会重复如上的第一、二步工作，并在服务端计算出该请求期望的数字签名。

l 服务端用期望的数字签名和请求端发送过来的数字签名做比对，如果完全一致则认为该请求通过安全验证，否则直接拒绝该请求。

### 3.2. Header公共请求参数

| 参数           | 说明                                                         | 样例                                       |
| -------------- | ------------------------------------------------------------ | ------------------------------------------ |
| X-Up-Key       | publisher_key                                                | X-Up-Key: i8XNjC4b8KVok4uw5RftR38Wgp2BFwql |
| X-Up-Timestamp | API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为15分钟。 |                                            |
| X-Up-Signature | 签名字符串                                                   |                                            |

 

### 3.3. 签名字段

| 字段         | 说明                                                   | 样例                                                         |
| ------------ | ------------------------------------------------------ | ------------------------------------------------------------ |
| Content-MD5  | HTTP 请求中 Body 部分的 MD5 值（必须为大写字符串）     | 875264590688CA6171F6228AF5BBB3D2                             |
| Content-Type | HTTP 请求中 Body 部分的类型                            | application/json                                             |
| Headers      | 除X-Up-Signature的其它header                           | X-Up-Timestamp:1562813567000X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2\n |
| HTTPMethod   | HTTP 请求的方法名称，全部大写                          | PUT、GET、POST 等                                            |
| Resource     | 由 HTTP 请求资源构造的字符串(如果有querystring要加上） | /v1/fullreport?key1=val1&key2=val2                           |

 

### 3.4. 签名方式

 参与签名计算的字符串：

     SignString = HTTPMethod + "\n" 
    
                        \+ Content-MD5 + "\n" 
    
                        \+ Content-Type + "\n" 
    
                        \+ Headers + "\n"
    
                        \+ Resource 
    
    如果无body，如下：
    
     SignString = HTTPMethod + "\n" 
    
                         \+ "\n" 
    
                         \+ "\n" 
    
                        \+ Headers + "\n"
    
                        \+ Resource 

Resource:

    如请求包含查询字符串（QueryString），则在 Resource 字符串尾部添加 ? 和查询字符串

   QueryString是 URL 中请求参数按字典序排序后的字符串，其中参数名和值之间用 = 相隔组成字符串，并对参数名-值对按照字典序升序排序，然后以 & 符号连接构成字符串。

    Key1 + "=" + Value1 + "&" + Key2 + "=" + Value2        

Headers：

     X-Up-Key + X-Up-Timestamp 按字典序升序
    
     X-Up-Signature不参与签名计算
    
    Key1 + ":" + Value1 + '\n' + Key2 + ":" + Value2           Sign = MD5(HTTPMethod + Content-MD5+ Content-Type + Header + Resource)

服务端会比对计算Sign和X-Up-Signature

 

### 3.5. Http状态码和业务状态码

| 状态码 | 返回信息                 | 含义               |
| ------ | ------------------------ | ------------------ |
| 200    | -                        | 成功               |
| 500    | -                        | 通用异常           |
| 600    | StatusHeaderParamError   | Header请求参数异常 |
| 601    | StatusSign               | Sign异常           |
| 602    | StatusParam              | 参数错误           |
| 603    | StatusPublisherRestrict  | 用户未开通权限     |
| 604    | StatusAppLengthError     | App创建错误        |
| 605    | StatusRpcParamError      | 中间服务异常       |
| 606    | StatusRequestRepeatError | 重复请求           |

 

### 3.6. 批量创建应用

#### 3.6.1. 请求URL

<https://openapi.toponad.com/v1/create_app>

### 3.6.2. 请求方式 

POST

### 3.6.3. 请求参数

| 字段              | 类型   | 是否必传 | 备注                             |
| ----------------- | ------ | -------- | -------------------------------- |
| count             | Int    | Y        | 创建app的数量                    |
| apps.app_name     | String | Y        | App名称                          |
| apps.platform     | Int    | Y        | 1或者2  (1:安卓平台，2是ios平台) |
| apps.market_url   | String | N        | 需符合Market URL规范             |
| apps.package_name | String | N        | 需符合包名规范                   |
| apps.category     | String | N        | 需符合附录规范                   |
| apps.sub_category | String | N        | 需符合附录规范                   |

 

### 3.6.4. 返回参数

| 字段     | 类型   | 是否必传 | 备注                             |
| -------- | ------ | -------- | -------------------------------- |
| app_id   | String | Y        | Up开发者后台的App ID             |
| app_name | String | Y        | App名称                          |
| errors   | String | N        | 错误信息（错误时返回）           |
| platform | Int    | Y        | 1或者2  (1:安卓平台，2是ios平台) |

 

### 3.6.5. 样例

请求样例：

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

 

返回样例：

[

    {
    
        "app_name": "111",
    
        "errors": "app package name is required"
    
    }

]

 

## 3.7. 批量读取应用列表

### 3.7.1. 请求URL

<https://openapi.toponad.com/v1/apps>

### 3.7.2. 请求方式 

POST

### 3.7.3. 请求参数

| 字段    | 类型   | 是否必传 | 备注                           |
| ------- | ------ | -------- | ------------------------------ |
| app_ids | String | N        | 默认传object，多个app_id是数组 |
| start   | Int    | N        | Default 0                      |
| limit   | Int    | N        | Default 100 最大一次性获取100  |

 

### 3.7.4. 返回参数

| 字段         | 类型   | 是否必传 | 备注                             |
| ------------ | ------ | -------- | -------------------------------- |
| app_id       | String | Y        | Up开发者后台的App ID             |
| app_name     | String | Y        | App名称                          |
| platform     | Int    | Y        | 1或者2  (1:安卓平台，2是ios平台) |
| market_url   | String | N        | -                                |
| package_name | String | N        | -                                |
| category     | String | N        | -                                |
| sub-category | String | N        | -                                |

 

### 3.7.5. 样例

请求样例：

{

	"limit":1

}

 

返回样例：

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

 

## 3.8. 批量创建广告位

### 3.8.1. 请求URL

<https://openapi.toponad.com/v1/create_placement>

### 3.8.2. 请求方式

POST

### 3.8.3. 请求参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| count                                 | Int    | Y        | 本次创建的Placement数量                                      |
| app_id                                | String | Y        | 创建广告位的应用id                                           |
| placements. placement_name            | String | Y        | 广告位名称，30个汉字或字符以内                               |
| placements. adformat                  | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | No       | 针对于native广告才有相关配置。<br />0：标准<br />1：原生banner<br />2：原生开屏 |
| placements.template.cdt               | Int    | No       | template为原生开屏时：倒计时时间                             |
| placements.template.ski_swt           | Int    | No       | template为原生开屏时：是否可调过                             |
| placements.template.aut_swt           | Int    | No       | template为原生开屏时：是否自动关闭                           |
| placements.template.auto_refresh_time | Int    | No       | -1表示不启动<br />0-n表示刷新时间                            |

 

### 3.8.4. 返回参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
|                                       |        |          |                                                              |
| app_id                                | String | Y        | Up开发者后台的App ID                                         |
| placement_name                        | String | Y        | Placement名称                                                |
| placement_id                          | String | Y        | Up开发者后台的Placement ID                                   |
| adformat                              | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | No       | 针对于native广告才有相关配置。<br />0：标准<br />1：原生banner<br />2：原生开屏 |
| placements.template.cdt               | Int    | No       | template为原生开屏时：倒计时时间                             |
| placements.template.ski_swt           | Int    | No       | template为原生开屏时：是否可调过<br />0：表示No<br />1：表示Yes |
| placements.template.aut_swt           | Int    | No       | template为原生开屏时：是否自动关闭<br />0：表示No<br />1：表示Yes |
| placements.template.auto_refresh_time | Int    | No       | -1表示不启动<br />0-n表示刷新时间                            |

 

### 3.8.5. 样例

请求样例：

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

 

返回样例：

[

    {
    
        "app_name": "",
    
        "app_id": "a5bc9921f7fdb4",
    
        "platform": 0,
    
        "placement_name": "xxx",
    
        "adformat": "native"
    
    }

]

 

## 3.9. 批量读取广告位列表

### 3.9.1. 请求URL

<https://openapi.toponad.com/v1/placements>

### 3.9.2. 请求方式 

POST

### 3.9.3. 请求参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| app_ids       | Object | N        | 默认传object，多个app_id是数组                               |
| placement_ids | Object | N        | 默认传object，多个placement_id是数组 默认可以为空            |
| start         | Int    | N        | Default 0。当App和Placement都指定时不需要填写                |
| limit         | Int    | N        | Default 100 最大一次性获取100。当App和Placement都指定时不需要填写 |

 

### 3.9.4. 返回参数

| 字段           | 类型   | 是否必传 | 备注                             |
| -------------- | ------ | -------- | -------------------------------- |
| app_id         | String | Y        | Up开发者后台的App ID             |
| app_name       | String | Y        | App名称                          |
| platform       | Int    | Y        | 1或者2  (1:安卓平台，2是ios平台) |
| placements     | String | Y        | -                                |
| placement_id   | String | N        | -                                |
| placement_name | String | N        | -                                |
| adformat       | String | N        | -                                |

 

### 3.9.5. 样例

请求样例：

{

	"placement_ids":["b5bc9bc2951216"]

}

 

返回样例：

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

 

## 3.10. 添加和修改segment

### 3.10.1. 请求URL

<https://openapi.toponad.com/v1/deal_segment>

### 3.10.2. 请求方式 

POST

### 3.10.3. 请求参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | segment的名字                                                |
| segment_id    | String | N        | segment修改的时候必传segment_id                              |
| rules         | Array  | Y        | segment的规则                                                |
| rules.type    | Int    | Y        | default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则string，详见：rule_content数据格式                       |

 

### 3.10.4. 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | segment的名字                                                |
| segment_id    | String | Y        | segment_id                                                   |
| rules         | Array  | Y        | segment的规则                                                |
| rules.type    | Int    | Y        | default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则string，详见：rule_content数据格式                       |



### 3.10.5. 样例

请求样例：

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

 

返回样例：

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



## 3.11. segment获取列表

### 3.11.1. 请求URL

<https://openapi.toponad.com/v1/segment_list>

### 3.11.2. 请求方式 

POST

### 3.11.3. 请求参数

| 字段        | 类型   | 是否必传 | 备注                                                         |
| ----------- | ------ | -------- | ------------------------------------------------------------ |
| segment_ids | Object | N        | 默认传object，多个segment_id是数组                           |
| start       | Int    | N        | Default 0。当segment_ids都指定时不需要填写                   |
| limit       | Int    | N        | Default 100 最大一次性获取100。当segment_ids都指定时不需要填写 |

 

### 3.11.4. 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | segment的名字                                                |
| segment_id    | String | Y        | segment_id                                                   |
| rules         | Array  | Y        | segment的规则                                                |
| rules.type    | Int    | Y        | default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则string，详见：rule_content数据格式                       |



### 3.11.5. 样例

请求样例：

{
   "segment_ids":{"uuid1","uuid2"}
}

 

返回样例：

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



## 3.12. 删除segment

### 3.12.1. 请求URL

<https://openapi.toponad.com/v1/del_segment>

### 3.12.2. 请求方式 

POST

### 3.12.3. 请求参数

| 字段        | 类型   | 是否必传 | 备注                            |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Object | Y        | 默认传object，多个segment是数组 |

 

### 3.12.4. 返回参数

成功无返回参数，失败则又返回数据

 

### 3.12.5. 样例

请求样例：

{
   "segment_ids":{"uuid1","uuid2"}
}

 

返回样例：成功无返回，失败有返回



## 3.13. 查询placement下面绑定的segment

### 3.13.1. 请求URL

<https://openapi.toponad.com/v1/waterfall/segment>

### 3.13.2. 请求方式 

POST

### 3.13.3. 请求参数

| 字段         | 类型   | 是否必传 | 备注                              |
| ------------ | ------ | -------- | --------------------------------- |
| placement_id | String | Y        | Placement Id String               |
| is_abtest    | Int    | Y        | 0 表示默认 <br />1 表示abtest分组 |

 

### 3.13.4. 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| priority      | Int    | Y        | 优先级参数                                                   |
| name          | String | Y        | segment的名字                                                |
| segment_id    | String | Y        | segment_id                                                   |
| rules         | Array  | Y        | segment的规则                                                |
| rules.type    | Int    | Y        | default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则string，详见：rule_content数据格式                       |

 

### 3.13.5. 样例

请求样例：

{
    "placement_id": "placementid1",
    "is_abtest": 1
}

 

返回样例：

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



## 3.14. 设置placement下面segment

### 3.14.1. 请求URL

<https://openapi.toponad.com/v1/waterfall/set_segment>

### 3.14.2. 请求方式 

POST

### 3.14.3. 请求参数

| 字段                | 类型   | 是否必传 | 备注                         |
| ------------------- | ------ | -------- | ---------------------------- |
| placement_id        | String | Y        | Placement Id String          |
| is_abtest           | Int    | Y        | 0 表示默认  1 表示abtest分组 |
| segments            | Array  | Y        | segment 排序的列表           |
| segments.priority   | Int    | Y        | segment 优先级               |
| segments.segment_id | String | Y        | segment id                   |

 

### 3.14.4. 返回参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id           | String | Y        | Placement Id String                                          |
| is_abtest              | Int    | Y        | 0 表示默认  <br />1 表示abtest分组                           |
| segments.priority      | Int    | Y        | 优先级参数                                                   |
| segments.name          | String | Y        | segment的名字                                                |
| segments.segment_id    | String | Y        | segment_id                                                   |
| segments.rules         | Array  | Y        | segment的规则                                                |
| segments.rules.type    | Int    | Y        | default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则string，详见：rule_content数据格式                       |

 

### 3.14.5. 样例

请求样例：

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

 

返回样例：

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
                    "content": ""
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
                    "content": ""
                }
            ]
        }
    ]
}



## 3.15. 删除placement下面segment

### 3.15.1. 请求URL

<https://openapi.toponad.com/v1/waterfall/del_segment>

### 3.15.2. 请求方式 

POST

### 3.15.3. 请求参数

| 字段         | 类型   | 是否必传 | 备注                         |
| ------------ | ------ | -------- | ---------------------------- |
| placement_id | String | Y        | Placement Id String          |
| is_abtest    | Int    | Y        | 0 表示默认  1 表示abtest分组 |
| segment_ids  | Array  | Y        | segment的删除的列表          |

 

### 3.15.4. 返回参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id           | String | Y        | Placement Id String                                          |
| is_abtest              | Int    | Y        | 0 表示默认   <br />1 表示abtest分组                          |
| segments.priority      | Int    | Y        | 优先级参数                                                   |
| segments.name          | String | Y        | segment的名字                                                |
| segments.segment_id    | String | Y        | segment_id                                                   |
| segments.rules         | Array  | Y        | segment的规则                                                |
| segments.rules.type    | Int    | Y        | default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则string                                                   |

 

### 3.15.5. 样例

请求样例：

{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_ids": [
        "segment_id1",
        "segment_id2"
    ]
}

 

返回样例：

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
                    "content": ""
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
                    "content": ""
                }
            ]
        }
    ]
}



## 3.16. waterfall的unit的查询

### 3.16.1. 请求URL

<https://openapi.toponad.com/v1/waterfall/units>

### 3.16.2. 请求方式 

GET

### 3.16.3. 请求参数

| 字段         | 类型   | 是否必传 | 备注            |
| ------------ | ------ | -------- | --------------- |
| placement_id | String | Y        | placement id    |
| segment_id   | String | Y        | segment id      |
| is_abtest    | Int    | Y        | 默认是0 1表示是 |

 

### 3.16.4. 返回参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement id                                                 |
| segment_id                          | String  | Y        | segment id                                                   |
| is_abtest                           | Int     | Y        | 0 表示默认   <br />1 表示abtest分组                          |
| ad_source_list                      | Array   | Y        | 如果这里填入为空的情况下就是查询， 如果不填入为空则是设置waterfall的设置 |
| ad_source_list.ad_source_id         | Int     | N        | adsource_id                                                  |
| ad_source_list.ecpm                 | float64 | N        | ecpm                                                         |
| ad_source_list.pirority             | Int     | N        | adsource优先级                                               |
| ad_source_list.header_bidding_witch | Int     | N        | adsource是否支持header bidding<br />0：表示默认，<br />1：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 0：表示不开启自动优化，<br />1：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | N        | default -1 ：表示关                                          |

 

### 3.16.5. 样例

请求样例：

{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1"
}

 

返回样例：

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



## 3.17. waterfall的unit设置

### 3.17.1. 请求URL（只能对未绑定的重新绑定关系，不能重复绑定）

<https://openapi.toponad.com/v1/waterfall/set_units>

### 3.17.2. 请求方式 

POST

### 3.17.3. 请求参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement id                                                 |
| segment_id                          | String  | Y        | segment id                                                   |
| is_abtest                           | Int     | Y        | 0 表示默认   <br />1 表示abtest分组                          |
| ad_source_list                      | Array   | N        | 如果这里填入为空的情况下就是查询，<br />如果不填入为空则是设置waterfall的设置 |
| ad_source_list.ad_source_id         | Int     | N        | adsource_id                                                  |
| ad_source_list.ecpm                 | float64 | N        | ecpm                                                         |
| ad_source_list.header_bidding_witch | Int     | N        | adsource是否支持header bidding<br />0：表示默认，<br />1：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 0：表示不开启自动优化，<br />1：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | N        | default -1 ：表示关                                          |

 

### 3.17.4. 返回参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | placement id                                                 |
| segment_id                          | String  | Y        | segment id                                                   |
| is_abtest                           | Int     | Y        | 0 表示默认   <br />1 表示abtest分组                          |
| ad_source_list                      | Array   | Y        | 如果这里填入为空的情况下就是查询， 如果不填入为空则是设置waterfall的设置 |
| ad_source_list.ad_source_id         | Int     | N        | adsource_id                                                  |
| ad_source_list.ecpm                 | float64 | N        | ecpm                                                         |
| ad_source_list.pirority             | Int     | N        | adsource优先级                                               |
| ad_source_list.header_bidding_witch | Int     | N        | adsource是否支持header bidding<br />0：表示默认，<br />1：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 0：表示不开启自动优化，<br />1：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | N        | default -1 ：表示关                                          |

 

### 3.17.5. 样例

请求样例：

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

 

返回样例：

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



# **4.** **注意事项**

为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵

守。

• 每小时最多请求 1000 次

• 每天请求 10000 次

# **5.** **附录1：go语言示例代码**

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

	//openapi的地址
	
	demoUrl := "请求URL"
	
	//提交的body数据
	
	body := "{}"
	
	//您申请的publisherKey
	
	publisherKey := "请填写您的publisherKey"
	
	//请求方式
	
	httpMethod := "POST"
	
	contentType := "application/json"
	
	publisherTimestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	
	headers := map[string]string{
	
		"X-Up-Timestamp": publisherTimestamp,
	
		"X-Up-Key":       publisherKey,
	
	}
	
	//处理queryPath
	
	urlParsed, err := url.Parse(demoUrl)
	
	if err != nil {
	
		fmt.Println(err)
	
		return
	
	}
	
	//处理resource
	
	resource := urlParsed.Path
	
	m, err := url.ParseQuery(urlParsed.RawQuery)
	
	if err != nil {
	
		fmt.Println(err)
	
		return
	
	}
	
	queryString := m.Encode()
	
	if queryString != "" {
	
		resource += "?" + queryString
	
	}

 










	//处理body
	
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

 










	//返回数据
	
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

 

# **6.** **附录****2****：应用一级和二级分类列表**

| 应用    | 一级分类 | 二级分类                |
| ------- | -------- | ----------------------- |
| Android | App      | Daydream                |
| Android | App      | Android Wear            |
| Android | App      | Art & Design            |
| Android | App      | Auto & Vehicles         |
| Android | App      | Beauty                  |
| Android | App      | Books & Reference       |
| Android | App      | Business                |
| Android | App      | Comics                  |
| Android | App      | Communication           |
| Android | App      | Dating                  |
| Android | App      | Education               |
| Android | App      | Entertainment           |
| Android | App      | Events                  |
| Android | App      | Finance                 |
| Android | App      | Food & Drink            |
| Android | App      | Health & Fitness        |
| Android | App      | House & Home            |
| Android | App      | Libraries & Demo        |
| Android | App      | Lifestyle               |
| Android | App      | Maps & Navigation       |
| Android | App      | Medical                 |
| Android | App      | Music & Audio           |
| Android | App      | News & Magazines        |
| Android | App      | Parenting               |
| Android | App      | Personalisation         |
| Android | App      | Photography             |
| Android | App      | Productivity            |
| Android | App      | Shopping                |
| Android | App      | Social                  |
| Android | App      | Sports                  |
| Android | App      | Tools                   |
| Android | App      | Travel & Local          |
| Android | App      | Video Players & Editors |
| Android | App      | Weather                 |
| Android | Game     | Action                  |
| Android | Game     | Adventure               |
| Android | Game     | Arcade                  |
| Android | Game     | Board                   |
| Android | Game     | Card                    |
| Android | Game     | Casino                  |
| Android | Game     | Casual                  |
| Android | Game     | Educational             |
| Android | Game     | Music                   |
| Android | Game     | Puzzle                  |
| Android | Game     | Racing                  |
| Android | Game     | Role Playing            |
| Android | Game     | Simulation              |
| Android | Game     | Sports                  |
| Android | Game     | Strategy                |
| Android | Game     | Trivia                  |
| Android | Game     | Word                    |
| Android | Family   | Ages 5 & Under          |
| Android | Family   | Ages 6-8                |
| Android | Family   | Ages 9 & Over           |
| Android | Family   | Action & Adventure      |
| Android | Family   | Brain Games             |
| Android | Family   | Creativity              |
| Android | Family   | Education               |
| Android | Family   | Music and video         |
| Android | Family   | Pretend play            |
| iOS     | Game     | Action                  |
| iOS     | Game     | Adventure               |
| iOS     | Game     | Arcade                  |
| iOS     | Game     | Board                   |
| iOS     | Game     | Card                    |
| iOS     | Game     | Casino                  |
| iOS     | Game     | Dice                    |
| iOS     | Game     | Educational             |
| iOS     | Game     | Family                  |
| iOS     | Game     | Music                   |
| iOS     | Game     | Puzzle                  |
| iOS     | Game     | Racing                  |
| iOS     | Game     | Role Playing            |
| iOS     | Game     | Simulation              |
| iOS     | Game     | Sports                  |
| iOS     | Game     | Strategy                |
| iOS     | Game     | Trivia                  |
| iOS     | Game     | Word                    |
| iOS     | App      | Books                   |
| iOS     | App      | Business                |
| iOS     | App      | Catalogs                |
| iOS     | App      | Education               |
| iOS     | App      | Entertainment           |
| iOS     | App      | Finance                 |
| iOS     | App      | Food & Drink            |
| iOS     | App      | Health & Fitness        |
| iOS     | App      | Lifestyle               |
| iOS     | App      | Magazines & Newspapers  |
| iOS     | App      | Medical                 |
| iOS     | App      | Music                   |
| iOS     | App      | Navigation              |
| iOS     | App      | News                    |
| iOS     | App      | Photo & Video           |
| iOS     | App      | Productivity            |
| iOS     | App      | Reference               |
| iOS     | App      | Shopping                |
| iOS     | App      | Social Networking       |
| iOS     | App      | Sports                  |
| iOS     | App      | Stickers                |
| iOS     | App      | Travel                  |
| iOS     | App      | Utilities               |
| iOS     | App      | Weather                 |

 

rule_content 数据格式

| rule | 描述                 | 示例                                 |
| :--- | :------------------- | :----------------------------------- |
| 0    | 包含（集合）         | 一维数组JSON ["CN", "US"]            |
| 1    | 不包含（集合）       | 一维数组JSON [1,2,3]                 |
| 2    | 大于等于（值）       | 整形或浮点 124                       |
| 3    | 小于等于（值）       | 整形或浮点 222.36                    |
| 4    | 区间内（区间）       | 二维数组JSON [[122,456],[888,12322]] |
| 5    | 区间外（区间）       | 二维数组JSON [[122,456],[888,12322]] |
| 6    | 自定义规则（custom） | bb=1&c!=3&p=3                        |
| 7    | 大于（值）           | 整形、浮点或字符串 124               |
| 8    | 小于（值）           | 整形、浮点或字符串 222.36            |