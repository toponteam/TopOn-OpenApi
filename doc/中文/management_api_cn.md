
# 修订历史

| 文档版本 | 发布时间      | 修订说明                          |
| :--------: | ----------- | -------------------------------- |
| v 1.0    | 2019年7月17日 | 支持批量创建和查询应用、广告位    |
| v 2.0    | 2019年11月4日 | 支持Waterfall、流量分组等相关配置 |
| v 2.1    | 2020年3月16日 | 新增广告平台、广告源管理接口      |
| v 2.2    | 2020年5月14日 | 流量分组接口功能调整      |


## 1. 关于文档

为提高合作伙伴的变现效率，TopOn平台专门提供了对接开发者后台相关操作的API接口，如创建应用和广告位、调整Waterfall优先级等。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

## 2. 权限申请

账号注册成功后，已自动开通开发者后台管理API权限，请登录开发者后台的账号管理页面查看publisher_key

## 3. 接口校验

### 3.1 接口请求流程说明

- 请求端根据 API 请求内容（包括 HTTP Header 和 Body）生成签名字符串。

- 请求端使用MD5对第一步生成的签名字符串进行签名，形成该 API 请求的数字签名。

- 请求端把 API 请求内容和数字签名一同发送给服务端。

- 服务端在接到请求后会重复如上的第一、二步工作，并在服务端计算出该请求期望的数字签名。

- 服务端用期望的数字签名和请求端发送过来的数字签名做比对，如果完全一致则认为该请求通过安全验证，否则直接拒绝该请求。

### 3.2 Header公共请求参数

| 参数           | 说明                                                         | 样例                                       |
| -------------- | ------------------------------------------------------------ | ------------------------------------------ |
| X-Up-Key       | publisher_key                                                | X-Up-Key: i8XNjC4b8KVok4uw5RftR38Wgp2BFwql |
| X-Up-Timestamp | API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为15分钟。 |  -                                          |
| X-Up-Signature | 签名字符串                                                   |                                            |-

### 3.3 签名字段

| 字段         | 说明                                                   | 样例                                                         |
| ------------ | ------------------------------------------------------ | ------------------------------------------------------------ |
| Content-MD5  | HTTP 请求中 Body 部分的 MD5 值（必须为大写字符串）     | 875264590688CA6171F6228AF5BBB3D2                             |
| Content-Type | HTTP 请求中 Body 部分的类型                            | application/json                                             |
| Headers      | 除X-Up-Signature的其它header                           | X-Up-Timestamp: 1562813567000X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2\n |
| HTTPMethod   | HTTP 请求的方法名称，全部大写                          | PUT、GET、POST 等                                            |
| Resource.Path     | 由 HTTP 请求资源构造的字符串(如果有querystring不要加上） | /v1/fullreport                           |

### 3.4 签名方式

参与签名计算的字符串：
```
     SignString = HTTPMethod + "\n" 
                        \+ Content-MD5 + "\n" 
                        \+ Content-Type + "\n"  
                        \+ Headers + "\n"
                        \+ Resource 
```
如果无body，如下：
```
    SignString = HTTPMethod + "\n" 
                        \+ "\n" 
                        \+ "\n" 
                        \+ Headers + "\n"
                        \+ Resource 
````
Resource:
```
    URL的Path      
```
Headers：
```
    X-Up-Key + X-Up-Timestamp 按字典序升序
    
    X-Up-Signature不参与签名计算
    
    Key1 + ":" + Value1 + '\n' + Key2 + ":" + Value2   
        
    Sign = MD5(HTTPMethod + Content-MD5+ Content-Type + Header + Resource)
```
服务端会比对计算Sign和X-Up-Signature

### 3.5 Http状态码和业务状态码

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

## 4. 应用管理

### 4.1 批量创建和修改应用

#### 4.1.1 请求URL

<https://openapi.toponad.com/v1/deal_app>

#### 4.1.2 请求方式 

POST

#### 4.1.3 请求参数

| 字段                    | 类型   | 是否必传 | 备注                                                        |
| ----------------------- | ------ | -------- | ----------------------------------------------------------- |
| count                   | Int    | Y        | 创建应用的数量                                              |
| apps.app_id             | String | N        | 开发者后台的应用ID                  |
| apps.app_name           | String | Y        | 应用名称                                                    |
| apps.platform           | Int    | Y        | 1或者2  (1:安卓平台，2是iOS平台)                            |
| apps.market_url         | String | N        | 需符合商店链接规范                                          |
| apps.screen_orientation | Int    | Y        | 1：竖屏 <br />2：横屏 <br />3：所有                         |
| apps.package_name       | String | N        | 需符合包名规范，示例：com.xxx，创建时必传                   |
| apps.category           | String | N        | 一级分类，需符合附录1规范，创建时未上架的应用必传 |
| apps.sub_category       | String | N        | 二级分类，需符合附录1规范，创建时未上架的应用必传 |
| apps.coppa       | String | N        | 是否遵守COPPA协议，默认：否<br>1：否，2：是 |
| apps.ccpa       | String | N        | 是否遵守CCPA协议，默认：否<br>1：否，2：是  |
 

#### 4.1.4 返回参数

| 字段               | 类型   | 是否必传 | 备注                             |
| ------------------ | ------ | -------- | -------------------------------- |
| app_id             | String | Y        | 开发者后台的应用ID               |
| app_name           | String | Y        | 应用名称                         |
| errors             | String | N        | 错误信息（错误时返回）           |
| platform           | Int    | Y        | 1或者2  (1:安卓平台，2是iOS平台) |
| screen_orientation | Int    | Y        | 1：竖屏<br />2：横屏<br />3：所有  |
| coppa       | String | N        | 是否遵守COPPA协议，默认：否<br>1：否，2：是 |
| ccpa       | String | N        | 是否遵守CCPA协议，默认：否<br>1：否，2：是  |

 

#### 4.1.5 样例

请求样例：
```
{
    "count": 1,
    "apps": [
        {
            "app_name": "oddman华为",
            "platform": 1,
            "screen_orientation":1,
            "market_url": "https://play.google.com/store/apps/details?id=com.hunantv.imgo.activity.inter"
        }
    ]
}
```


返回样例：
```
[
    {
        "app_name": "oddman华为",
        "app_id": "",
        "platform": 1,
        "screen_orientation": 1,
        "errors": "repeat app name error"
    }
]
```

### 4.2 获取应用列表

#### 4.2.1 请求URL

<https://openapi.toponad.com/v1/apps>

#### 4.2.2 请求方式 

POST

#### 4.2.3 请求参数

| 字段    | 类型   | 是否必传 | 备注                           |
| ------- | ------ | -------- | ------------------------------ |
| app_ids | Array[String] | N  | 多个应用ID是数组 （和下面两参数不能一起使用） |
| start   | Int    | N        | Default 0                      |
| limit   | Int    | N        | Default 100 最大一次性获取100  |

 

#### 4.2.4 返回参数

| 字段                    | 类型   | 是否必传 | 备注                                |
| ----------------------- | ------ | -------- | ----------------------------------- |
| app_id                  | String | Y        | 开发者后台的应用ID                  |
| app_name                | String | Y        | 应用名称                            |
| platform                | Int    | Y        | 1或者2  (1:安卓平台，2是iOS平台)    |
| market_url              | String | N        | -                                   |
| screen_orientation      | Int    | Y        | 1：竖屏 <br />2：横屏<br />3：所有 |
| package_name            | String | N        | -                                   |
| category                | String | N        | -                                   |
| sub-category            | String | N        | -                                   |
| coppa       | String | N        | 是否遵守COPPA协议，默认：否<br>1：否，2：是 |
| ccpa       | String | N        | 是否遵守CCPA协议，默认：否<br>1：否，2：是  |


#### 4.2.5 样例

请求样例：
```
{
	"limit":1
}
```


返回样例：
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

### 4.3 批量删除应用

#### 4.3.1 请求URL

<https://openapi.toponad.com/v1/del_apps>

#### 4.3.2 请求方式 

POST

#### 4.3.3 请求参数

| 字段    | 类型   | 是否必传 | 备注                           |
| ------- | ------ | -------- | ------------------------------ |
| app_ids | Array[String] | Y        | 多个应用ID是数组 |


#### 4.3.4 返回参数

 无，如果是错误会返回errors

#### 4.3.5 样例

请求样例：
```
{
	"app_ids": ["a1bu2thutsq3mn"]
}
```

返回样例：

返回状态码或者错误码


## 5. 广告位管理

### 5.1 批量创建和修改广告位

#### 5.1.1 请求URL

<https://openapi.toponad.com/v1/deal_placement>

#### 5.1.2 请求方式

POST

#### 5.1.3 请求参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- |---------------------------------------------- |
| count                                 | Int    | Y        | 创建的广告位数量                                             |
| app_id                                | String | Y        | 创建广告位的应用ID                                           |
| placements.placement_name             | String | Y        | 广告位名称，30个汉字或字符以内                               |
| placements.adformat                   | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | N        | 针对于native广告才有相关配置。<br />0：标准<br />1：原生Banner<br />2：原生开屏 |
| placements.template_extra.cdt               | Int    | N        | template为原生开屏时：倒计时时间，默认5秒                    |
| placements.template_extra.ski_swt           | Int    | N        | template为原生开屏时：是否可跳过，默认可跳过<br />0：表示No<br />1：表示Yes    |
| placements.template_extra.aut_swt           | Int    | N        | template为原生开屏时：是否自动关闭，默认自动关闭<br />0：表示No<br />1：表示Yes  |
| placements.template_extra.auto_refresh_time | Int    | N        | template为原生Banner时：是否自动刷新，默认不启动<br />-1表示不启动<br />0-n表示刷新时间  |
| remark                                 | String    | N        | 备注                                             |
| status                                 | Int   | N        | 广告位状态                                             |


#### 5.1.4 返回参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| app_id                                | String | Y        | 开发者后台的应用ID                                           |
| placement_name                        | String | Y        | 广告位名称                                                   |
| placement_id                          | String | Y        | 开发者后台的广告位ID                                         |
| adformat                              | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | N        | 针对于native广告才有相关配置。<br />0：标准<br />1：原生Banner<br />2：原生开屏 |
| placements.template_extra.cdt               | Int    | N        | template为原生开屏时：倒计时时间                             |
| placements.template_extra.ski_swt           | Int    | N        | template为原生开屏时：是否可调过                             |
| placements.template_extra.aut_swt           | Int    | N        | template为原生开屏时：是否自动关闭                           |
| placements.template_extra.auto_refresh_time | Int    | N        | template为原生Banner时：是否自动刷新                         |
| remark                                 | String    | N        | 备注                                             |
| status                                 | Int   | N        | 广告位状态                                             |

#### 5.1.5 样例

请求样例：
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


返回样例：
```
[
    {
        "app_name": "我要翘课",
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

### 5.2 获取广告位列表

#### 5.2.1 请求URL

<https://openapi.toponad.com/v1/placements>

#### 5.2.2 请求方式 

POST

#### 5.2.3 请求参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| app_ids       | Array[String] | N        | 多个应用ID是数组                               |
| placement_ids | Array[String] | N        | 多个广告位ID是数组 默认可以为空                 |
| start         | Int    | N        | Default 0。当应用和广告位都指定时不需要填写                   |
| limit         | Int    | N        | Default 100 最大一次性获取100。当应用和广告位都指定时不需要填写 |

 

#### 5.2.4 返回参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| app_id                                | String | Y        | 开发者后台的应用ID                                           |
| app_name                              | String | Y        | 应用名称                              |
| platform                              | Int    | Y        | 1 or 2  (1:android，2IOS)              |
| placement_name                        | String | Y        | 广告位名称                                                   |
| placement_id                          | String | Y        | 开发者后台的广告位ID                                         |
| adformat                              | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | N        | 针对于native广告才有相关配置。<br />0：标准<br />1：原生Banner<br />2：原生开屏 |
| placements.template_extra.cdt               | Int    | N        | template为原生开屏时：倒计时时间                             |
| placements.template_extra.ski_swt           | Int    | N        | template为原生开屏时：是否可调过                             |
| placements.template_extra.aut_swt           | Int    | N        | template为原生开屏时：是否自动关闭                           |
| placements.template_extra.auto_refresh_time | Int    | N        | template为原生Banner时：是否自动刷新                         |
| remark                                 | String    | N        | 备注                                             |
| status                                 | Int   | N        | 广告位状态                                             |


#### 5.2.5 样例

请求样例：
```
{
	"placement_ids":["b5bc9bc2951216"]
}
```


返回样例：
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

### 5.3 批量删除广告位

#### 5.3.1 请求URL

<https://openapi.toponad.com/v1/del_placements>

#### 5.3.2 请求方式 

POST

#### 5.3.3 请求参数

| 字段          | 类型  | 是否必传 | 备注                                         |
| ------------- | ----- | -------- | -------------------------------------------- |
| placement_ids | Array | Y        | 默认传Array，多个广告位ID是数组 |

 

#### 5.3.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| msg | String | N        | 默认返回String         |


#### 5.3.5 样例

请求样例：
```
{
	"placement_ids":["b5bc9bc2951216"]
}
```


返回样例：
```
{
    "msg": "suc"
}
```


## 6. 广告位维度的流量分组管理

### 6.1 批量创建和修改流量分组

#### 6.1.1 请求URL

<https://openapi.toponad.com/v2/deal_segment>

#### 6.1.2 请求方式 

POST

#### 6.1.3 请求参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| count                  | Int    | Y        | 请求条数                                                      |
| app_id                  | String    | Y        | app_id                                                    |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |
| segments               | Array  | Y        | -                                                             |
| segments.name          | String | Y        | Segment名称 (默认新增的segment优先级排在现有分组前面)                                                 |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见附录2规范           |

#### 6.1.4 返回参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| count                  | Int    | Y        | 请求条数                                                      |
| app_id                  | String    | Y        | app_id                                                    |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组               |
| segments               | Array  | Y        | -                                                             |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.name          | String | Y        | Segment名称 (默认新增的segment优先级排在现有分组前面)                                                 |
| segments.errors    | String | N        | Segment处理异常的错误                             |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见附录2规范           |       |



#### 6.1.5 样例

请求样例：

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

 

返回样例：

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

### 6.2 获取已启用的流量分组列表

#### 6.2.1 请求URL

<https://openapi.toponad.com/v2/waterfall/get_segment>

#### 6.2.2 请求方式 

GET

#### 6.2.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                              |
| ------------ | ------ | -------- | --------------------------------- |
| placement_id | String | Y        | 广告位ID                          |
| app_id                  | String    | Y        | app_id                                                    |
| is_abtest             | Int    | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |

#### 6.2.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| priority      | Int    | Y        | 优先级参数                                                   |
| name          | String | Y        | Segment名称                                               |
| segment_id    | String | Y        | Segment ID                                                   |
| parallel_request_number    | Int | Y        | 并发请求数                             |
| auto_load    | Int | Y        | Default 0：表示关，只能传0或正整数<br/>对于Banner，可以设置自动刷新时间，大于0表示自动刷新时间<br/>对于RV和插屏，仅控制自动请求的开关状态，非0表示开 |
| day_cap    | Int | Y        |  -1 ：表示关                            |
| hour_cap    | Int | Y        |  -1 ：表示关                             |
| pacing    | Int | Y        |  -1 ：表示关                             |
| rules         | Array  | Y        | Segment的规则                                                |
| rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则详见附录2规范                       |

#### 6.2.5 样例

请求样例：

```
https://openapi.toponad.com/v2/waterfall/get_segment?placement_id=b5bc9bbfb0f913&app_id=a5bc9921f7fdb4&is_abtest=1
```

返回样例：

```
[
    {
        "name": "解决为",
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

### 6.3 批量删除流量分组

#### 6.3.1 请求URL

<https://openapi.toponad.com/v1/waterfall/del_segment>

#### 6.3.2 请求方式 

POST

#### 6.3.3 请求参数

| 字段        | 类型   | 是否必传 | 备注                            |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | 默认传Array，多个segment是数组 |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    |N       | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |

 

#### 6.3.4 返回参数
| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | Y        | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |
| segments               | Array  | Y        | -                                                             |
| segments.name          | String | Y        | Segment名称                                                  |
| segments.priority      | Int | Y        | 优先级排序                                                  |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见附录2规范           |


#### 6.3.5 样例

请求样例：

```
{
    "placement_id": "111111",
    "is_abtest": 1,
    "segment_ids": [
        "22222"
    ]
}
```

返回样例：

```
{
    "placement_id": "b5bc9bbfb0f913",
    "is_abtest": 0,
    "segments": [
        {
            "priority": 1,
            "name": "解决为",
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


### 6.4 调整流量分组优先级

#### 6.4.1 请求URL

<https://openapi.toponad.com/v2/waterfall/set_segment_rank>

#### 6.4.2 请求方式 

POST

#### 6.4.3 请求参数

| 字段        | 类型   | 是否必传 | 备注                            |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | 默认传Array，多个segment是数组 |
| placement_id | String | Y        | placement_id |
| is_abtest | int32 | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组 |
| app_id | String | Y        | app_id |

 

#### 6.4.4 返回参数
| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | Y        | 是否是测试组，默认：0<br/>0：默认组，1：测试组               |
| segments               | Array  | Y        | -                                                             |
| segments.name          | String | Y        | Segment名称                                                  |
| segments.priority      | Int | Y        | 优先级排序                                                  |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.parallel_request_number    | Int | Y        | 并发请求数                             |
| segments.auto_load    | Int | Y        | Default 0：表示关，只能传0或正整数<br/>对于Banner，可以设置自动刷新时间，大于0表示自动刷新时间<br/>对于RV和插屏，仅控制自动请求的开关状态，非0表示开|
| segments.day_cap    | Int | Y        | -1 ：表示关                            |
| segments.hour_cap    | Int | Y        |  -1 ：表示关                             |
| segments.pacing    | Int | Y        |  -1 ：表示关                             |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见附录2规范           |




#### 6.4.5 样例

请求样例：

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

返回样例：
```
{
    "placement_id": "b5bc9bbfb0f913",
    "is_abtest": 1,
    "segments": [
        {
            "name": "解决",
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
            "name": "解决为",
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

### 6.5 批量修改流量分组在Waterfall的属性

#### 6.5.1 请求URL

<https://openapi.toponad.com/v2/waterfall/set_segment>

#### 6.5.2 请求方式 

POST

#### 6.5.3 请求参数

| 字段        | 类型   | 是否必传 | 备注                            |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | 默认传Array，多个segment是数组 |
| placement_id | String | Y        | placement_id |
| is_abtest | int32 | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组 |
| app_id | String | Y        | app_id |
| segments               | Array  | Y        | -                                                             |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.parallel_request_number    | Int | Y        | 并发请求数                             |
| segments.auto_load    | Int | Y        | Default 0：表示关，只能传0或正整数<br/>对于Banner，可以设置自动刷新时间，大于0表示自动刷新时间<br/>对于RV和插屏，仅控制自动请求的开关状态，非0表示开|
| segments.day_cap    | Int | Y        |  -1 ：表示关                            |
| segments.hour_cap    | Int | Y        |  -1 ：表示关                             |
| segments.pacing    | Int | Y        |  -1 ：表示关                             |

 

#### 6.5.4 返回参数
| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id            | String    | Y        | placement_id                                        |
| is_abtest             | Int    | Y        | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |
| segments               | Array  | Y        | -                                                             |
| segments.name          | String | Y        | Segment名称                                                  |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.parallel_request_number    | Int | Y        | 并发请求数                             |
| segments.auto_load    | Int | Y        | Default 0：表示关，只能传0或正整数<br/>对于Banner，可以设置自动刷新时间，大于0表示自动刷新时间<br/>对于RV和插屏，仅控制自动请求的开关状态，非0表示开|
| segments.day_cap    | Int | Y        |  -1 ：表示关                            |
| segments.hour_cap    | Int | Y        |  -1 ：表示关                             |
| segments.pacing    | Int | Y        |  -1 ：表示关                             |



#### 6.5.5 样例

请求样例：

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

返回样例：
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


## 7. 聚合管理基本操作

### 7.1 查询Waterfall的广告源列表

#### 7.1.1 请求URL

<https://openapi.toponad.com/v1/waterfall/units>

#### 7.1.2 请求方式 

GET

#### 7.1.3 请求参数

| 字段         | 类型   | 是否必传 | 备注            |
| ------------ | ------ | -------- | --------------- |
| placement_id | String | Y        | 广告位ID        |
| segment_id   | String | N        | Segment ID,默认是default segment      |
| is_abtest             | Int    | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |

#### 7.1.4 返回参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | 广告位ID                                                     |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组               |
| ad_source_list                      | Array   | Y        | 如果为空，则当前没有启用广告源                               |
| ad_source_list.ad_source_id         | Int     | N        | 广告源ID                                                     |
| ad_source_list.ecpm                 | float64 | N        | eCPM价格                                                     |
| ad_source_list.auto_ecpm            | float64 | N        | 自动eCPM价格                                                     |
| ad_source_list.header_bidding_switch | Int     | N        | 是否支持Header Bidding，广告源创建时已确定<br />1：表示不支持，<br />2：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 1：表示不开启自动优化，<br />2：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | Y        |  -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | Y        |  -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | Y        |  -1 ：表示关                                          |
| free_ad_source_list                 | Array   | N        | 未使用adsource_list（其他参数参照ad_source_list）            |
| offer_list                          | Array   | N        | 正在使用的交叉推广列表                                        |
| offer_list.offer_id                 | String  | N        | offer_id                                                     |
| offer_list.offer_name               | String  | N        | offer名称                                                    |

#### 7.1.5 样例

请求样例：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1"
}
```

返回样例：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1,
    "segment_id": "segment_id1",
    "ad_source_list": [
        {
            "ad_source_id": "ad_source_id1",
            "ecpm": "ecpm1",
            "header_bidding_switch": 1,
            "day_cap": -1,
            "hour_cap": -1,
            "pacing": -1
        },
        {
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

### 7.2 批量修改广告源在Waterfall的属性

#### 7.2.1 请求URL

<https://openapi.toponad.com/v1/waterfall/set_units>

#### 7.2.2 请求方式 

POST

#### 7.2.3 请求参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | 广告位ID                                                     |
| is_abtest             | Int    | N        | 是否是测试组，默认：0<br/>0：默认组，1：测试组                |
| segment_id                          | String  | Y        | Segment ID                                                   |
| parallel_request_number             | Int     | Y        | 并行请求数据                                                 |
| offer_switch                        | Int     | N        | 交叉推广开关                                                    |
| unbind_adsource_list                | Array   | N        | 取消绑定的广告源，只传广告源ID                                    |
| ad_source_list                      | Array   | Y        | 要绑定的广告源配置信息                                       |
| ad_source_list.ad_source_id         | Int     | Y        | 广告源ID                                                     |
| ad_source_list.ecpm                 | float64 | Y        | eCPM价格                                                     |
| ad_source_list.header_bidding_switch | Int     | N        | 是否支持Header Bidding，广告源创建时已确定<br />1：表示不支持，<br />2：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 1：表示不开启自动优化，<br />2：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | Y        |  -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | Y        |  -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | Y        |  -1 ：表示关                                          |

#### 7.2.4 返回参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | 广告位ID                                                     |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        |是否是测试组，默认：0<br/>0：默认组，1：测试组               |
| parallel_request_number             | Int     | Y        | 并行请求数据                                                 |
| offer_switch                        | Int     | N        | offer开关<br />1：关<br />2：开                              |
| unbind_adsource_list                | Array   | N        | 取消绑定的adsource                                           |
| ad_source_list                      | Array   | Y        | 要绑定的广告源配置信息                                       |
| ad_source_list.ad_source_id         | Int     | Y        | 广告源ID                                                     |
| ad_source_list.ecpm                 | float64 | Y        | eCPM                                                         |
| ad_source_list.header_bidding_switch | Int     | N        | 是否支持Header Bidding，广告源创建时已确定<br />1：表示不支持，<br />2：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 1：表示不开启自动优化，<br />2：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | Y        |  -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | Y        |  -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | Y        |  -1 ：表示关                                          |

#### 7.2.5 样例

请求样例：

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

返回样例：

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



## 8. 广告平台管理

### 8.1 创建和修改广告平台Publisher、App维度参数

#### 8.1.1 请求URL

<https://openapi.toponad.com/v1/set_networks>

#### 8.1.2 请求方式

POST

#### 8.1.3 请求参数

| 字段                              | 类型   | 是否必传 | 备注                |
| --------------------------------- | ------ | -------- | ------------------- |
| network_name                      | String | N        | 广告平台账号名称，开通多账号时必传        |
| nw_firm_id                           | Int    | Y        | 广告平台ID              |
| network_id                        | Int    | N        | 广告平台账号ID          |
| is_open_report                    | Int    | N        | 是否开通Report API  |
| auth_content                      | Object | N        | 广告平台Publisher维度参数    |
| network_app_info                  | Array  | N        | -     |
| network_app_info.app_id           | String | N        | TopOn的应用ID              |
| network_app_info.app_auth_content | Object | N        | 广告平台App维度参数，详见附录3规范 |


#### 8.1.4 返回参数

| 字段                              | 类型   | 是否必传 | 备注                            |
| --------------------------------- | ------ | -------- | ------------------------------- |
| network_name                      | String | Y        | 广告平台账号名称                    |
| nw_firm_id                        | Int    | Y        | 广告平台ID                          |
| network_id                        | Int    | N        | 广告平台账号ID                      |
| is_open_report                    | Int    | N        | 是否开通Report API              |
| auth_content                      | Object | N        | 广告平台Publisher维度参数    |
| network_app_info                  | Array  | N        | -                |
| network_app_info.app_id           | String | N        | TopOn的应用ID                         |
| network_app_info.app_auth_content | Object | N        | 广告平台App维度参数 |


#### 8.1.5 样例

请求样例：

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


返回样例：

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

### 8.2 获取广告平台Publisher、App维度参数

#### 8.2.1 请求URL

<https://openapi.toponad.com/v1/networks>

#### 8.2.2 请求方式 

POST

#### 8.2.3 请求参数

无 

#### 8.2.4 返回参数

| 字段                              | 类型   | 是否必传 | 备注                |
| --------------------------------- | ------ | -------- | ------------------- |
| network_name                      | String | N        | 广告平台账号名称，开通多账号时必传        |
| nw_firm_id                        | Int    | Y        | 广告平台ID              |
| network_id                        | Int    | N        | 广告平台账号ID          |
| is_open_report                    | Int    | N        | 是否开通Report API  |
| auth_content                      | Object | N        | 广告平台Publisher维度参数    |
| network_app_info                  | Array  | N        | -     |
| network_app_info.app_id           | String | N        | TopOn的应用ID              |
| network_app_info.app_auth_content | Object | N        | 广告平台App维度参数 |

 

#### 8.2.5 样例


返回样例：

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

## 9. 广告源管理

### 9.1 批量创建和修改广告源

#### 9.1.1 请求URL

<https://openapi.toponad.com/v1/set_units>

#### 9.1.2 请求方式 

POST

#### 9.1.3 请求参数

| 字段                | 类型   | 是否必传 | 备注                             |
| ------------------- | ------ | -------- | -------------------------------- |
| count               | Int32  | Y        | 广告源总数                             |
| units               | Array  | Y        | -                        |
| units.network_id    | Int    | Y        | 广告平台账号ID                       |
| units.adsource_id   | Int    | N        | 广告源ID，修改时必传|
| units.adsource_name | String | Y        | 广告源名称                 |
| units.adsource_token | Object | Y        | 广告平台Unit维度参数，详见附录3规范 |
| units.placement_id  | String | Y        | TopOn的广告位ID                     |
| units.default_ecpm  | String | Y        | 广告源默认价格                             |
| units.header_bidding_switch  | String | Y        | 1：表示不支持<br>2：表示支持 |

#### 9.1.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                             |
| ------------- | ------ | -------- | -------------------------------- |
| network_id    | Int    | N        | 广告平台账号ID                       |
| adsource_id   | Int    | N        | 广告源ID               |
| adsource_name | String | Y        | 广告源名称                 |
| adsource_token | Object | Y        | 广告平台Unit维度参数 |
| placement_id  | String | Y        | TopOn的广告位ID                    |
| default_ecpm  | String | Y        | 广告源默认价格                       |

#### 9.1.5 样例

请求样例：
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

返回样例：

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


### 9.2 获取广告源列表

#### 9.2.1 请求URL

<https://openapi.toponad.com/v1/units>

#### 9.2.2 请求方式

POST

#### 9.2.3 请求参数

| 字段             | 类型          | 是否必传 | 备注                                             |
| ---------------- | ------------- | -------- | ------------------------------------------------ |
| network_firm_ids | Array[int32]  | N        | 支持传入多个广告平台ID                           |
| app_ids          | Array[String] | N        | 支持传入多个应用ID                               |
| placement_ids    | Array[String] | N        | 支持传入多个广告位ID                             |
| adsource_ids     | Array[int32]  | N        | 支持传入多个广告源ID                             |
| start            | int32         | N        | 默认值：0 (和上面参数不能一起使用)               |
| limit            | int32         | N        | 默认值：100，最大一次性获取100 (和上面参数不能一起使用)    |
| metrics          | Array[String] | N        | 从ad_source_list中指定返回的字段，不传则全部返回 |


#### 9.2.4 返回参数

| 字段                                   | 类型   | 是否必传 | 备注                |
| -------------------------------------- | ------ | -------- | ------------------- |
| network_id                             | String | N        | 广告平台账号ID        |
| network_name                           | String | N        | 广告平台账号名称        |
| nw_firm_id                             | Int    | N        | 广告平台ID              |
| adsource_id                            | Int    | N        | 广告源ID         |
| adsource_name                          | Int    | N        | 广告源名称       |
| adsource_token                          | Object | N        | 广告源配置参数 |
| app_id                                 | String | N        | TopOn的应用ID     |
| app_name                               | String | N        | TopOn的应用名称   |
| platform                               | Int    | N        | 平台 |
| placement_id                           | String | N        | TopOn的广告位ID              |
| placement_name                         | Object | N        | TopOn的广告位名称 |
| placement_format                       | String | N        |   广告位广告形式                  |
| waterfall_list                         |  Array |   N      |    当前正在使用该广告源的waterfall信息                 |
| waterfall_list.ecpm                    |   String     |      N    |  waterfall关联的ecpm                   |
| waterfall_list.auto_ecpm               |   String     |     N     |   waterfall自动优化的ecpm                    |
| waterfall_list.header_bidding_switch    |   Int     |     N     |     是否支持headerbidding                |
| waterfall_list.auto_switch             |    Int    |      N    |  是否开启了自动优化                   |
| waterfall_list.day_cap                 |   Int     |      N    |  daycap                   |
| waterfall_list.hour_cap                |  Int      |     N     |   hour cap                  |
| waterfall_list.pacing                  |   Int     |    N      |    pacing                 |
| waterfall_list.segment_name            |  String  |   N       |   segment名称                  |
| waterfall_list.segment_id              |  String      |   N       |   关联的segment_id                  |
| waterfall_list.priority                |   Int     |     N     |  关联的segment优先级排序                   |
| waterfall_list.parallel_request_number |   Int     |     N     |    关联的segment的并发请求数                 |
| waterfall_list.is_abtest |   Int     |     N     |    是否是测试组，默认：0<br/>0：默认组，1：测试组                 |


#### 9.2.5 样例

请求样例：

```
{
	"adsource_ids":[19683]
}
```


返回样例：

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



### 9.3 批量删除广告源

#### 9.3.1 请求URL

<https://openapi.toponad.com/v1/del_units>

#### 9.3.2 请求方式 

POST

#### 9.3.3 请求参数

| 字段         | 类型         | 是否必传 | 备注            |
| ------------ | ------------ | -------- | --------------- |
| adsource_ids | Array[Int32] | Y        | adsource_id列表 |

#### 9.3.4 返回参数

| 字段 | 类型   | 是否必传 | 备注     |
| ---- | ------ | -------- | -------- |
| msg  | String | N        | 处理信息 |

#### 9.3.5 样例

请求样例：

```
{
	"adsource_ids":[19683]
}
```


返回样例：

```
{
    "msg": "suc"
}
```


## 10. 注意事项

为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵守。

• 每小时最多请求 1000 次

• 每天请求 10000 次


## 附录1：应用一级和二级分类列表

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

## 附录2：流量分组规则数据格式

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

## 附录3：广告平台详细参数

所有的Key和Value数据类型均为String

| 广告平台ID | 广告平台名称 | auth_content | app_auth_content | 广告样式 | adsource_token |  key-value对应值  |
| --------- | ----------- | ------------ | ---------------- | ------- | -------------  | ---------------- |
| 1         | Facebook    | - | app_id<br>app_token | native<br>rewarded_video<br>interstitial | unit_id | app_id：AppID <br> app_token：AccessToken <br> unit_id：PlacementID |
| 1         | Facebook    | - | app_id<br>app_token | bannner | unit_id<br>size | size枚举值：320x50,320x90,320x250 |
| 2         | Admob       | account_id<br>oauth_key | app_id | native<br>rewarded_video<br>interstitial | unit_id | account_id：PublisherID <br/> oauth_key：AccessToken <br/> app_id：AppID <br/> unit_id：UnitID |
| 2         | Admob       | account_id<br>oauth_key | app_id | bannner | unit_id<br>size | size枚举值：320x50,320x100,320x250,468x60,728x90 |
| 3         | Inmobi      | username<br>password<br>apikey<br>app_id | - | native<br>rewarded_video<br>interstitial | unit_id |    username：EmailID </br> app_id：Account ID </br> password：Password </br> apikey：API Key </br> unit_id：Placement ID |
| 3         | Inmobi      | username<br>password<br>apikey<br>app_id | - | bannner | unit_id<br>size | size枚举值：320x50 |
| 4         | Flurry      | token | sdk_key | native<br>rewarded_video<br>interstitial | ad_space | token：Token </br> sdk_key：API Key </br> ad_space：AD Unit Name |
| 4         | Flurry      | token | sdk_key | banner | ad_space<br>size | size枚举值：320x50 |
| 5         | Applovin    | sdkkey<br>apikey | - | native | - | sdkkey：SDK Key </br> apikey：Report Key  |
| 5         | Applovin    | sdkkey<br>apikey | -  | rewarded_video<br>interstitial | zone_id | zone_id：Zone ID |
| 5         | Applovin    | sdkkey<br>apikey | -  | banner | zone_id<br>size | size枚举值：320x50,300x250  |
| 6         | Mintegral   | skey<br>secret<br>appkey | app_id | native<br>rewarded_video  | unit_id | appkey：App Key </br> skey：Skey </br> secret：Secret </br> appid：AppID </br> unitid：UnitID |
| 6         | Mintegral   | skey<br>secret<br>appkey | app_id | bannner | unit_id<br>size | size枚举值：320x50,300x250,320x90,smart  |
| 6         | Mintegral   | skey<br>secret<br>appkey | app_id | interstitial | unit_id<br>is_video | is_video枚举值：0,1 |
| 7         | Mopub       | repkey<br>apikey | - | native<br>rewarded_video<br>interstitial | unit_id | repkey：Inventory Report ID </br> apikey：API Key </br> unitid：Unit ID  |
| 7         | Mopub       | repkey<br>apikey | -  | bannner | unit_id<br>size | size枚举值：320x50,300x250,728x90  |
| 8         | 腾讯广告     | agid<br>publisher_id<br>app_key<br>qq | app_id | native | unit_id<br>unit_version<br>unit_type | qq：账号ID </br> agid：AGID </br> publisher_id：App ID </br> app_key：App Key </br> app_id：媒体ID </br> unit_id：UnitID</br>unit_version枚举值：1,2</br>unit_type枚举值：1,2 |
| 8         | 腾讯广告     | agid<br>publisher_id<br>app_key<br>qq | app_id | rewarded_video,splash | unit_id | - |
| 8         | 腾讯广告     | agid<br>publisher_id<br>app_key<br>qq | app_id | bannner| unit_id<br>unit_version<br>size | unit_version枚举值：2</br> size枚举值：320x50 |
| 8         | 腾讯广告     | agid<br>publisher_id<br>app_key<br>qq | app_id | interstitial | unit_id<br>unit_version<br>video_muted<br>video_autoplay<br>video_duration<br>is_fullscreen | video_duration_switch：videoDuration</br>unit_version枚举值：2</br> video_muted枚举值：0,1 </br>video_autoplay枚举值：0,1</br> video_duration：时长可选</br>is_fullscreen枚举值：0，1 |
| 9         | Chartboost  | user_id<br>user_signature | app_id<br>app_signature | rewarded_video<br>interstitial | location | user_id：UserID </br> user_signature：UserSignature </br> app_id：UserAppID </br> app_signature：AppSignature </br> location：Location |
| 10        | Tapjoy      | apikey | sdk_key | rewarded_video<br>interstitial | placement_name | apikey：APIKey </br> sdk_key：SDKKey </br> placement_name：PlacementName |
| 11        | Ironsource  | username<br>secret_key | app_key | rewarded_video<br>interstitial | instance_id |   username：Username </br> secret_key：Secret Key </br> app_key：App Key </br> instance_id：Instance ID |
| 12        | UnityAds    | apikey | game_id | rewarded_video<br>interstitial | placement_id | apikey：API Key </br> organization_core_id：Organization core ID </br> game_id：Game ID </br> placement_id：Placement ID |
| 13        | Vungle      | apikey | app_id | rewarded_video<br>interstitial | placement_id | apikey：Reporting API Key </br> app_id：App ID </br> placement_id：PlacementID |
| 14        | AdColony    | user_credentials | app_id | rewarded_video<br>interstitial | zone_id | user_credentials：Read-Only API key </br> app_id：App ID </br> zone_id：Zone ID |
| 15        | 穿山甲       | user_id<br>secure_key | app_id | native | slot_id<br>is_video<br>layout_type<br>media_size | user_id：UserID </br> secure_key：Secure Key </br> app_id：AppID </br> slot_id：SlotID </br> is_video枚举值：0,1,2,3 <br> layout_type枚举值：0,1 </br> media_size枚举值（layout_type = 1必填）：1,2 |
| 15        | 穿山甲       | user_id<br>secure_key | app_id | rewarded_video | slot_id<br>personalized_template | personalized_template枚举值：0,1 |
| 15        | 穿山甲       | user_id<br>secure_key | app_id | banner | slot_id<br>layout_type<br>size | layout_type枚举值：1 </br> size枚举值：640x100,600x90,600x150,600x500,600x400,600x300,600x260,690x388 |
| 15        | 穿山甲       | user_id<br>secure_key | app_id | interstitial | slot_id<br>is_video<br>layout_type<br>size<br>personalized_template | is_video为0时，以下两个参数必填<br>layout_type枚举值：1 <br> size枚举值：1:1,3:2,2:3 </br> is_video为1时，以下参数必填<br>personalized_template枚举值：0,1 |
| 15        | 穿山甲       | user_id<br>secure_key | app_id | splash | slot_id<br>personalized_template | personalized_template枚举值：0,1 |
| 16        | 聚量传媒     | - | - | rewarded_video<br>interstitial | app_id | app_id：App ID |
| 16        | 聚量传媒     | - | - | banner | app_id<br>size | size枚举值：320x50,480x75,640x100,960x150,728x90 |
| 17        | OneWay      | access_key | publisher_id | rewarded_video<br>interstitial | slot_id | access_key：Access Key </br> publisher_id：Publisher ID </br> slot_id：Placement ID |
| 18        | MobPower    | publisher_id<br>api_key  | app_id | native<br>rewarded_video<br>interstitial | placement_id | api_key：API Key </br> publisher_id：Publisher ID </br> app_id：App ID </br> placement_id：Placement ID |
| 18        | MobPower    | publisher_id<br>api_key  | app_id | banner | placement_id<br>size | size枚举值：320x50 |
| 19        | 金山云       | - | media_id | rewarded_video | slot_id | media_id：Media ID </br> slot_id：Slot ID |
| 21        | AppNext     | email<br>password<br>key  | - | native<br>rewarded_video<br>interstitial | placement_id | email：Email </br> password：Password </br> key：Key </br> placement_id：Placement ID |
| 21        | AppNext     | email<br>password<br>key  | - | banner | placement_id<br>size | size枚举值：320x50,320x100,300x250 |
| 22        | Baidu       | access_key | app_id | native<br>rewarded_video<br>interstitial<br>splash | ad_place_id | access_key：Access Key </br> app_id：AppID </br> ad_place_id：ADPlaceID |
| 22        | Baidu       | access_key | app_id | banner | ad_place_id<br>size | size枚举值：375x56,200x30,375x250,200x133,375x160,200x85,375x187,200x100 |
| 23        | Nend        | api_key | - | naitve | spot_id<br>api_key<br>is_video | api_key：APIKey </br> spot_id：spotID </br> is_video枚举值：0,1 |
| 23        | Nend        | api_key | - | rewarded_video | spot_id<br>api_key | - |
| 23        | Nend        | api_key | - | banner | spot_id<br>api_key<br>size | size枚举值：320x50,320x100,300x100,300x250,728x90 |
| 23        | Nend        | api_key | - | interstitial | spot_id<br>api_key<br>is_video | is_video枚举值：0,1,2 |
| 24        | Maio        | api_id<br>api_key | media_id | rewarded_video<br>interstitial | zone_id | api_id：API ID </br> api_key：API Key </br> media_id：Media ID </br> zone_id：Zone ID |
| 25        | StartAPP    | partner_id<br>token  | app_id | rewarded_video<br>interstitial | ad_tag | partner_id：Partner ID </br> token：Token </br> app_id：APP ID </br> ad_tag：AD Tag |
| 26        | SuperAwesome | - | property_id | rewarded_video | placement_id | property_id：Property ID </br> placement_id：Placement ID |
| 28        | 快手        | access_key<br>security_key | app_id<br>app_name | native | position_id<br>layout_type<br>video_sound<br>is_video<br>unit_type | access_key：Access Key </br> security_key：Security Key </br> app_id：AppID </br> app_name：AppName </br> position_id：PosID </br> unit_type枚举值：0,1<br>unit_type为1时，以下三个参数必填<br>layout_type枚举值：0<br>is_video枚举值：0,1<br>video_sound枚举值：0,1 |
| 28        | 快手        | access_key<br>security_key | app_id<br>app_name | rewarded_video<br>interstitial | orientation | orientation枚举值：1,2 |
| 29        | Sigmob      | public_key<br>secret_key  | app_id<br>app_key | rewarded_video<br>interstitial<br>splash | placement_id |    public_key：Public Key </br> secret_key：Secret Key </br> app_id：AppID </br> app_key：App Key </br> placement_id：PlacementID |
| 36        | Ogury       | api_key<br>api_secret | key | rewarded_video<br>interstitial | unit_id | api_key：API KEY </br> api_secret：API SECRET </br> key：KEY </br> unit_id：AD Unit ID |
