# TopOn开发者后台操作API对接文档

## 修订历史


| 文档版本 | 发布时间      | 修订说明                         |
| :--------: | ------------- | -------------------------------- |
| v 1.0    | 2019年7月17日 | 支持批量创建和查询应用、广告位   |
| v 2.0    | 2019年11月4日 | 支持Waterfall、流量分组等相关配置 |


## 目录

[1. 关于文档](#关于文档)</br>
[2. 申请开通权限](#申请开通权限)</br>
[3. 接口校验](#接口校验)</br>
[4. 应用管理](#应用管理)</br> 
- [4.1 批量创建应用](#批量创建应用)</br>  
- [4.2 获取应用列表](#获取应用列表)</br>
- 4.3批量删除应用</br>

[5. 广告位管理](#广告位管理)</br>
- [5.1 批量创建广告位](#批量创建广告位)</br>  
- [5.2 获取广告位列表](#获取广告位列表)</br>

[6. 流量分组管理](#流量分组管理)</br>
- [6.1 创建和修改流量分组](#创建和修改流量分组)</br>  
- [6.2 获取流量分组列表](#获取流量分组列表)</br>
- [6.3 批量删除流量分组](#批量删除流量分组)</br>

[7. 聚合管理基本操作](#聚合管理基本操作)</br>

- [7.1 查询广告位已启用的流量分组列表](#查询广告位已启用的流量分组列表)</br>  
- [7.2 为广告位启用新流量分组或调整流量分组优先级](#为广告位启用新流量分组或调整流量分组优先级)</br>
- [7.3 为广告位批量移除流量分组](#为广告位批量移除流量分组)</br>
- [7.4 查询Waterfall已启用的广告源列表](#查询Waterfall已启用的广告源列表)</br>  
- [7.5 为Waterfall批量启用广告源](#为Waterfall批量启用广告源)</br>

[8. 注意事项](#注意事项)</br>
[9. 附录1：go语言示例代码](#附录1：go语言示例代码)</br>
[10. 附录2：应用一级和二级分类列表](#附录2：应用一级和二级分类列表)</br>
[11. 附录3：流量分组规则数据格式](#附录3：流量分组规则数据格式)

<h2 id='关于文档'>1. 关于文档</h2>
为提高合作伙伴的变现效率，TopOn平台专门提供了批量创建及查询API接口。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

<h2 id='申请开通权限'>2. 申请开通权限</h2>
在使用TopOn平台的批量创建API前，合作伙伴需向TopOn申请publisher_key，用于识别来自合作伙伴的请求，申请方法请咨询与您对接的商务经理。

<h2 id='接口校验'>3. 接口校验</h2>
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
| X-Up-Timestamp | API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为15分钟。 |                                            |
| X-Up-Signature | 签名字符串                                                   |                                            |

 

### 3.3 签名字段

| 字段         | 说明                                                   | 样例                                                         |
| ------------ | ------------------------------------------------------ | ------------------------------------------------------------ |
| Content-MD5  | HTTP 请求中 Body 部分的 MD5 值（必须为大写字符串）     | 875264590688CA6171F6228AF5BBB3D2                             |
| Content-Type | HTTP 请求中 Body 部分的类型                            | application/json                                             |
| Headers      | 除X-Up-Signature的其它header                           | X-Up-Timestamp:1562813567000X-Up-Key:aac6880633f102bce2174ec9d99322f55e69a8a2\n |
| HTTPMethod   | HTTP 请求的方法名称，全部大写                          | PUT、GET、POST 等                                            |
| Resource     | 由 HTTP 请求资源构造的字符串(如果有querystring要加上） | /v1/fullreport?key1=val1&key2=val2                           |

 

### 3.4 签名方式

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

    URL的Path      

Headers：

    X-Up-Key + X-Up-Timestamp 按字典序升序
    
    X-Up-Signature不参与签名计算
    
    Key1 + ":" + Value1 + '\n' + Key2 + ":" + Value2   
        
    Sign = MD5(HTTPMethod + Content-MD5+ Content-Type + Header + Resource)

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

<h2 id='应用管理'>4. 应用管理</h2>
<h3 id='批量创建应用'>4.1 批量创建和修改应用</h3>
#### 4.1.1 请求URL

<https://openapi.toponad.com/v1/deal_app>

#### 4.1.2 请求方式 

POST

#### 4.1.3 请求参数

| 字段                    | 类型   | 是否必传 | 备注                                                        |
| ----------------------- | ------ | -------- | ----------------------------------------------------------- |
| count                   | Int    | Y        | 创建应用的数量                                              |
| apps.app_name           | String | Y        | 应用名称                                                    |
| apps.platform           | Int    | Y        | 1或者2  (1:安卓平台，2是iOS平台)                            |
| apps.market_url         | String | N        | 需符合商店链接规范                                          |
| apps.screen_orientation | Int    | Y        | 1：竖屏 <br />2：横屏,<br />3：所有                         |
| apps.package_name       | String | N        | 需符合包名规范，示例：com.xxx                               |
| apps.category           | String | N        | 一级分类，需符合[附录2规范](#附录2：应用一级和二级分类列表) |
| apps.sub_category       | String | N        | 二级分类，需符合[附录2规范](#附录2：应用一级和二级分类列表) |
| apps.screen_orientation | Int    | Y        | 1:竖屏<br />2:横屏<br />3：所有                             |

 

#### 4.1.4 返回参数

| 字段               | 类型   | 是否必传 | 备注                             |
| ------------------ | ------ | -------- | -------------------------------- |
| app_id             | String | Y        | 开发者后台的应用ID               |
| app_name           | String | Y        | 应用名称                         |
| errors             | String | N        | 错误信息（错误时返回）           |
| platform           | Int    | Y        | 1或者2  (1:安卓平台，2是iOS平台) |
| screen_orientation | Int    | Y        | 1:竖屏<br />2:横屏<br />3：所有  |

 

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

<h3 id='获取应用列表'>4.2 获取应用列表</h3>
#### 4.2.1 请求URL

<https://openapi.toponad.com/v1/apps>

#### 4.2.2 请求方式 

POST

#### 4.2.3 请求参数

| 字段    | 类型   | 是否必传 | 备注                           |
| ------- | ------ | -------- | ------------------------------ |
| app_ids | String | N        | 默认传Object，多个应用ID是数组 |
| start   | Int    | N        | Default 0                      |
| limit   | Int    | N        | Default 100 最大一次性获取100  |

 

#### 4.2.4 返回参数

| 字段                    | 类型   | 是否必传 | 备注                                |
| ----------------------- | ------ | -------- | ----------------------------------- |
| app_id                  | String | Y        | 开发者后台的应用ID                  |
| app_name                | String | Y        | 应用名称                            |
| platform                | Int    | Y        | 1或者2  (1:安卓平台，2是iOS平台)    |
| market_url              | String | N        | -                                   |
| apps.screen_orientation | Int    | Y        | 1：竖屏 <br />2：横屏,<br />3：所有 |
| package_name            | String | N        | -                                   |
| category                | String | N        | -                                   |
| sub-category            | String | N        | -                                   |

 

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

<h3 id='获取应用列表'>4.3 批量删除应用</h3>
#### 4.3.1 请求URL

<https://openapi.toponad.com/v1/del_apps>

#### 4.3.2 请求方式 

POST

#### 4.3.3 请求参数

| 字段    | 类型   | 是否必传 | 备注                           |
| ------- | ------ | -------- | ------------------------------ |
| app_ids | String | N        | 默认传Object，多个应用ID是数组 |

 

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


<h2 id='广告位管理'>5. 广告位管理</h2>
<h3 id='批量创建广告位'>5.1 批量创建广告位</h3>
#### 5.1.1 请求URL

<https://openapi.toponad.com/v1/deal_placement>

#### 5.1.2 请求方式

POST

#### 5.1.3 请求参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| count                                 | Int    | Y        | 创建的广告位数量                                             |
| app_id                                | String | Y        | 创建广告位的应用ID                                           |
| placements.placement_name             | String | Y        | 广告位名称，30个汉字或字符以内                               |
| placements.adformat                   | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | N        | 针对于native广告才有相关配置。<br />0：标准<br />1：原生Banner<br />2：原生开屏 |
| placements.template.cdt               | Int    | N        | template为原生开屏时：倒计时时间，默认5秒                    |
| placements.template.ski_swt           | Int    | N        | template为原生开屏时：是否可跳过，默认可跳过<br />0：表示No<br />1：表示Yes    |
| placements.template.aut_swt           | Int    | N        | template为原生开屏时：是否自动关闭，默认自动关闭<br />0：表示No<br />1：表示Yes  |
| placements.template.auto_refresh_time | Int    | N        | template为原生Banner时：是否自动刷新，默认不启动<br />-1表示不启动<br />0-n表示刷新时间  |

 

#### 5.1.4 返回参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| app_id                                | String | Y        | 开发者后台的应用ID                                           |
| placement_name                        | String | Y        | 广告位名称                                                   |
| placement_id                          | String | Y        | 开发者后台的广告位ID                                         |
| adformat                              | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | N        | 针对于native广告才有相关配置。<br />0：标准<br />1：原生Banner<br />2：原生开屏 |
| placements.template.cdt               | Int    | N        | template为原生开屏时：倒计时时间                             |
| placements.template.ski_swt           | Int    | N        | template为原生开屏时：是否可调过                             |
| placements.template.aut_swt           | Int    | N        | template为原生开屏时：是否自动关闭                           |
| placements.template.auto_refresh_time | Int    | N        | template为原生Banner时：是否自动刷新                         |

 

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

<h3 id='获取广告位列表'>5.2 获取广告位列表</h3>
#### 5.2.1 请求URL

<https://openapi.toponad.com/v1/placements>

#### 5.2.2 请求方式 

POST

#### 5.2.3 请求参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| app_ids       | Object | N        | 默认传Object，多个应用ID是数组                               |
| placement_ids | Object | N        | 默认传Object，多个广告位ID是数组 默认可以为空                 |
| start         | Int    | N        | Default 0。当应用和广告位都指定时不需要填写                   |
| limit         | Int    | N        | Default 100 最大一次性获取100。当应用和广告位都指定时不需要填写 |

 

#### 5.2.4 返回参数

| 字段                                  | 类型   | 是否必传 | 备注                                                         |
| ------------------------------------- | ------ | -------- | ------------------------------------------------------------ |
| app_id                                | String | Y        | 开发者后台的应用ID                                           |
| placement_name                        | String | Y        | 广告位名称                                                   |
| placement_id                          | String | Y        | 开发者后台的广告位ID                                         |
| adformat                              | String | Y        | native、banner、rewarded_video、interstitial、splash （单选） |
| placements.template                   | Int    | N        | 针对于native广告才有相关配置。<br />0：标准<br />1：原生Banner<br />2：原生开屏 |
| placements.template.cdt               | Int    | N        | template为原生开屏时：倒计时时间                             |
| placements.template.ski_swt           | Int    | N        | template为原生开屏时：是否可调过                             |
| placements.template.aut_swt           | Int    | N        | template为原生开屏时：是否自动关闭                           |
| placements.template.auto_refresh_time | Int    | N        | template为原生Banner时：是否自动刷新                         |

 

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

<h3 id='获取广告位列表'>5.3 删除广告位</h3>
#### 5.3.1 请求URL

<https://openapi.toponad.com/v1/del_placements>

#### 5.3.2 请求方式 

POST

#### 5.3.3 请求参数

| 字段          | 类型  | 是否必传 | 备注                                         |
| ------------- | ----- | -------- | -------------------------------------------- |
| placement_ids | Array | Y        | 默认传Array，多个广告位ID是数组 默认可以为空 |

 

#### 5.2.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| msg | String | N        | 默认返回String         |


#### 5.2.5 样例

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


<h2 id='流量分组管理'>6. 流量分组管理</h2>
<h3 id='创建和修改流量分组'>6.1 创建和修改流量分组</h3>
#### 6.1.1 请求URL

<https://openapi.toponad.com/v1/deal_segment>

#### 6.1.2 请求方式 

POST

#### 6.1.3 请求参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| count                  | Int    | Y        | 请求条数                                                     |
| segments               | Array  | Y        | 请求的大的segment array                                      |
| segments.name          | String | Y        | Segment名称                                                  |
| segments.segment_id    | String | N        | Segment修改的时候必传Segment ID                              |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时/1225/2203（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见[附录3规范](#附录3：流量分组规则数据格式)            |

#### 6.1.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | Segment名称                                              |
| segment_id    | String | Y        | Segment ID                                                   |
| rules         | Array  | Y        | Segment的规则                                                |
| rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则详见[附录3规范](#附录3：流量分组规则数据格式)            |



#### 6.1.5 样例

请求样例：

```
{
    "count": 2,
    "segments": [
        {
            "name": "2123123",
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
[
    {
        "name": "2123123",
        "segment_id": "c1boq7f7apetou",
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
```

<h3 id='获取流量分组列表'>6.2 获取流量分组列表</h3>
#### 6.2.1 请求URL

<https://openapi.toponad.com/v1/segment_list>

#### 6.2.2 请求方式 

POST

#### 6.2.3 请求参数

| 字段        | 类型   | 是否必传 | 备注                                                         |
| ----------- | ------ | -------- | ------------------------------------------------------------ |
| segment_ids | Array | N        | 默认传Array，多个Segment ID是数组                           |
| start       | Int    | N        | Default 0。当Segment ID都指定时不需要填写                   |
| limit       | Int    | N        | Default 100 最大一次性获取100。当Segment ID都指定时不需要填写 |

 

#### 6.2.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| name          | String | Y        | Segment名称                                                |
| segment_id    | String | Y        | Segment ID                                                   |
| rules         | Array  | Y        | Segment的规则                                                |
| rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则详见[附录3规范](#附录3：流量分组规则数据格式)     |



#### 6.2.5 样例

请求样例：

```
{
   "segment_ids":["uuid1","uuid2"]
}
```

返回样例：

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

<h3 id='批量删除流量分组'>6.3 批量删除流量分组</h3>
#### 6.3.1 请求URL

<https://openapi.toponad.com/v1/del_segment>

#### 6.3.2 请求方式 

POST

#### 6.3.3 请求参数

| 字段        | 类型   | 是否必传 | 备注                            |
| ----------- | ------ | -------- | ------------------------------- |
| segment_ids | Array | Y        | 默认传Array，多个segment是数组 |

 

#### 6.3.4 返回参数

成功只返回状态码200，失败则返回数据。如果其中一个Segment正在Waterfall中使用，则不允许删除，本次请求的Segment列表都会删除失败

#### 6.3.5 样例

请求样例：

```
{
   "segment_ids":["uuid1","uuid2"]
}
```

返回样例：

成功只返回状态码200，失败则返回数据

<h2 id='聚合管理基本操作'>7. 聚合管理基本操作</h2>
<h3 id='查询广告位已启用的流量分组列表'>7.1 查询广告位已启用的流量分组列表</h3>
#### 7.1.1 请求URL

<https://openapi.toponad.com/v1/waterfall/segment>

#### 7.1.2 请求方式 

GET

#### 7.1.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                              |
| ------------ | ------ | -------- | --------------------------------- |
| placement_id | String | Y        | 广告位ID                          |
| is_abtest    | Int    | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组 |

#### 7.1.4 返回参数

| 字段          | 类型   | 是否必传 | 备注                                                         |
| ------------- | ------ | -------- | ------------------------------------------------------------ |
| priority      | Int    | Y        | 优先级参数                                                   |
| name          | String | Y        | Segment名称                                               |
| segment_id    | String | Y        | Segment ID                                                   |
| rules         | Array  | Y        | Segment的规则                                                |
| rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| rules.content | string | Y        | 规则详见[附录3规范](#附录3：流量分组规则数据格式)                       |

#### 7.1.5 样例

请求样例：

```
{
    "placement_id": "placementid1",
    "is_abtest": 1
}
```

返回样例：

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

<h3 id='为广告位启用新流量分组或调整流量分组优先级'>7.2 为广告位启用新流量分组或调整流量分组优先级</h3>
#### 7.2.1 请求URL

<https://openapi.toponad.com/v1/waterfall/set_segment>

#### 7.2.2 请求方式 

POST

#### 7.2.3 请求参数

| 字段                | 类型   | 是否必传 | 备注                         |
| ------------------- | ------ | -------- | ---------------------------- |
| placement_id        | String | Y        | 广告位ID                     |
| is_abtest           | Int    | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组 |
| segments            | Array  | Y        | Segment排序的完整列表           |
| segments.priority   | Int    | Y        | Segment优先级               |
| segments.segment_id | String | Y        | Segment ID                   |

#### 7.2.4 返回参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id           | String | Y        | 广告位ID                                                     |
| is_abtest              | Int    | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组                |
| segments.priority      | Int    | Y        | 优先级参数                                                   |
| segments.name          | String | Y        | Segment的名字                                                |
| segments.segment_id    | String | Y        | Segment ID                                                   |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见[附录3规范](#附录3：流量分组规则数据格式)                       |

#### 7.2.5 样例

请求样例：

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

返回样例：

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

<h3 id='为广告位批量移除流量分组'>7.3 为广告位批量移除流量分组</h3>
#### 7.3.1 请求URL

<https://openapi.toponad.com/v1/waterfall/del_segment>

#### 7.3.2 请求方式 

POST

#### 7.3.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                         |
| ------------ | ------ | -------- | ---------------------------- |
| placement_id | String | Y        | 广告位ID                     |
| is_abtest    | Int    | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组 |
| segment_ids  | Array  | Y        | 要移除的Segment列表          |

#### 7.3.4 返回参数

| 字段                   | 类型   | 是否必传 | 备注                                                         |
| ---------------------- | ------ | -------- | ------------------------------------------------------------ |
| placement_id           | String | Y        | 广告位ID                                                     |
| is_abtest              | Int    | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组                |
| segments.priority      | Int    | Y        | 优先级参数                                                   |
| segments.name          | String | Y        | Segment的名字                                                |
| segments.segment_id    | String | Y        | Segment ID                                                   |
| segments.rules         | Array  | Y        | Segment的规则                                                |
| segments.rules.type    | Int    | Y        | Default 0 <br />下面是各种数字的对应的值。<br />0 地区（集合）<br/>1 时间（区间）<br/>2 天（星期）（集合）<br/>3 网络（集合）<br/>4 小时（区间）<br/>5 自定义规则（custom）<br/>8 app version （集合）<br/>9 sdk version （集合）<br/>10 device_type （集合）<br/>11 device brand（集合）<br/>12 os version （集合）<br/>16 timezone (值，特殊处理)<br/>17 Device ID （集合）<br/>19 城市 （集合） |
| segments.rules.rule    | Int    | Y        | Default 0<br />下面是各种数字对应的值<br />0 包含（集合）<br/>1 不包含（集合）<br/>2 大于等于（值）<br/>3 小于等于（值）<br/>4 区间内（区间）<br/>5 区间外（区间）<br/>6 自定义规则（custom）<br/>7 大于（值）<br/>8 小于（值） |
| segments.rules.content | string | Y        | 规则详见[附录3规范](#附录3：流量分组规则数据格式)            |

#### 7.3.5 样例

请求样例：

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

返回样例：

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

<h3 id='查询Waterfall已启用的广告源列表'>7.4 查询Waterfall已启用的广告源列表</h3>
#### 7.4.1 请求URL

<https://openapi.toponad.com/v1/waterfall/units>

#### 7.4.2 请求方式 

GET

#### 7.4.3 请求参数

| 字段         | 类型   | 是否必传 | 备注            |
| ------------ | ------ | -------- | --------------- |
| placement_id | String | Y        | 广告位ID        |
| segment_id   | String | Y        | Segment ID      |
| is_abtest    | Int    | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组 |

#### 7.4.4 返回参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | 广告位ID                                                     |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组                |
| ad_source_list                      | Array   | Y        | 如果为空，则当前没有启用广告源                                |
| ad_source_list.ad_source_id         | Int     | N        | 广告源ID                                                     |
| ad_source_list.ecpm                 | float64 | N        | eCPM价格                                                     |
| ad_source_list.pirority             | Int     | N        | 广告源优先级                                                 |
| ad_source_list.header_bidding_witch | Int     | N        | 是否支持Header Bidding，广告源创建时已确定<br />0：表示不支持，<br />1：表示支持 |
| ad_source_list.auto_switch          | Int     | N        | 0：表示不开启自动优化，<br />1：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | N        | Default -1 ：表示关                                          |

#### 7.4.5 样例

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

<h3 id='为Waterfall批量启用广告源'>7.5 为Waterfall批量启用广告源</h3>
#### 7.5.1 请求URL

<https://openapi.toponad.com/v1/waterfall/set_units>

#### 7.5.2 请求方式 

POST

#### 7.5.3 请求参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | 广告位ID                                                     |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组               |
| parallel_request_number             | Int     | Y        | 并行请求数据                                                 |
| offer_switch                        | Int     | N        | offer开关                                                    |
| unbind_adsource_list                | Array   | N        | 取消绑定的adsource                                           |
| ad_source_list                      | Array   | Y        | 要绑定的广告源配置信息                                       |
| ad_source_list.ad_source_id         | Int     | Y        | 广告源ID                                                     |
| ad_source_list.ecpm                 | float64 | Y        | eCPM价格                                                     |
| ad_source_list.header_bidding_witch | Int     | N        | 是否支持Header Bidding，广告源创建时已确定<br />0：表示不支持，<br />1：表示支持 |
| ad_source_list.auto_switch          | Int     | Y        | 0：表示不开启自动优化，<br />1：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | N        | Default -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | N        | Default -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | N        | Default -1 ：表示关                                          |

#### 7.5.4 返回参数

| 字段                                | 类型    | 是否必传 | 备注                                                         |
| ----------------------------------- | ------- | -------- | ------------------------------------------------------------ |
| placement_id                        | String  | Y        | 广告位ID                                                     |
| segment_id                          | String  | Y        | Segment ID                                                   |
| is_abtest                           | Int     | Y        | 0 表示对照组或未开通A/B测试 <br />1 表示测试组               |
| parallel_request_number             | Int     | Y        | 并行请求数据                                                 |
| offer_switch                        | Int     | N        | offer开关                                                    |
| unbind_adsource_list                | Array   | N        | 取消绑定的adsource                                           |
| ad_source_list                      | Array   | Y        | 要绑定的广告源配置信息                                       |
| ad_source_list.ad_source_id         | Int     | Y        | 广告源ID                                                     |
| ad_source_list.ecpm                 | float64 | Y        | eCPM                                                         |
| ad_source_list.pirority             | Int     | N        | 广告源优先级                                                 |
| ad_source_list.header_bidding_witch | Int     | N        | 是否支持Header Bidding，广告源创建时已确定<br />0：表示不支持，<br />1：表示支持 |
| ad_source_list.auto_switch          | Int     | Y        | 0：表示不开启自动优化，<br />1：表示开启自动优化             |
| ad_source_list.day_cap              | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.hour_cap             | Int     | N        | default -1 ：表示关                                          |
| ad_source_list.pacing               | Int     | N        | default -1 ：表示关                                          |

#### 7.5.5 样例

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



<h2 id='厂商参数配置'>8. 厂商参数配置</h2>
<h3 id='单个创建厂商数据配置'>8.1 单个创建厂商数据配置</h3>

#### 8.1.1 请求URL

<https://openapi.toponad.com/v1/set_networks>

#### 8.1.2 请求方式

POST

#### 8.1.3 请求参数

| 字段                              | 类型   | 是否必传 | 备注                |
| --------------------------------- | ------ | -------- | ------------------- |
| network_name                      | String | Y        | 厂商账号名称        |
| firm_id                           | Int    | Y        | 厂商Id              |
| network_id                        | Int    | N        | 厂商账号id          |
| is_open_report                    | Int    | N        | 是否开通report_api  |
| auth_content                      | Object | N        | 厂商维度配置参数    |
| network_app_info                  | Array  | N        | 厂商app维度数据     |
| network_app_info.app_id           | String | N        | app_id              |
| network_app_info.app_auth_content | Object | N        | 厂商对应app维度参数 |

 

#### 8.1.4 返回参数

| 字段                              | 类型   | 是否必传 | 备注                |
| --------------------------------- | ------ | -------- | ------------------- |
| network_name                      | String | Y        | 厂商账号名称        |
| nw_firm_id                        | Int    | Y        | 厂商Id              |
| network_id                        | Int    | N        | 厂商账号id          |
| is_open_report                    | Int    | N        | 是否开通report_api  |
| auth_content                      | Object | N        | 厂商维度配置参数    |
| network_app_info                  | Array  | N        | 厂商app维度数据     |
| network_app_info.app_id           | String | N        | app_id              |
| network_app_info.app_auth_content | Object | N        | 厂商对应app维度参数 |

****

#### 8.1.5 样例

请求样例：

```
   {
        "network_name": "Default",
        "firm_id": 2,
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
    "firm_id": 2,
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

<h3 id='获取广告位列表'>5.2 获取厂商和厂商app维度信息列表</h3>

#### 9.2.1 请求URL

<https://openapi.toponad.com/v1/networks>

#### 9.2.2 请求方式 

POST

#### 9.2.3 请求参数

无 

#### 9.2.4 返回参数

| 字段                              | 类型   | 是否必传 | 备注                |
| --------------------------------- | ------ | -------- | ------------------- |
| network_name                      | String | Y        | 厂商账号名称        |
| firm_id                           | Int    | Y        | 厂商Id              |
| network_id                        | Int    | N        | 厂商账号id          |
| is_open_report                    | Int    | N        | 是否开通report_api  |
| auth_content                      | Object | N        | 厂商维度配置参数    |
| network_app_info                  | Array  | N        | 厂商app维度数据     |
| network_app_info.app_id           | String | N        | app_id              |
| network_app_info.app_auth_content | Object | N        | 厂商对应app维度参数 |

 

#### 9.2.5 样例


返回样例：

```
[
    {
        "network_name": "Default",
        "firm_id": 1,
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
        "firm_id": 1,
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

<h2 id='AdSource 配置管理'>10. AdSource 配置管理</h2>
<h3 id='获取adSource列表'>10.1 获取adSource列表</h3>

#### 10.1.1 请求URL

<https://openapi.toponad.com/v1/units

#### 10.1.2 请求方式

POST

#### 10.1.3 请求参数

| 字段             | 类型          | 是否必传 | 备注                                             |
| ---------------- | ------------- | -------- | ------------------------------------------------ |
| network_firm_ids | Array[int32]  | N        | 支持传入多个广告平台ID                           |
| app_ids          | Array[String] | N        | 支持传入多个应用ID                               |
| placement_ids    | Array[String] | N        | 支持传入多个广告位ID                             |
| adsource_ids     | Array[int32]  | N        | 支持传入多个广告源ID                             |
| start            | int32         | N        | 默认值：0 (和上面参数不能一起使用)               |
| limit            | int32         | N        | 默认值：100，最大一次性获取100                   |
| metrics          | Array[String] | N        | 从ad_source_list中指定返回的字段，不传则全部返回 |

 

#### 10.1.4 返回参数

| 字段                                   | 类型   | 是否必传 | 备注                |
| -------------------------------------- | ------ | -------- | ------------------- |
| network_name                           | String | N        | 厂商账号名称        |
| nw_firm_id                             | Int    | N        | 厂商Id              |
| adsource_id                            | Int    | N        | adsource_id         |
| adsource_name                          | Int    | N        | adsource_name       |
| adsouce_token                          | Object | N        | adsouce维度配置参数 |
| app_id                                 | String | N        | 厂商app维度数据     |
| app_name                               | String | N        | app_name            |
| platform                               | Int    | N        |                     |
| placement_id                           | String | N        | app_id              |
| placement_name                         | Object | N        | 厂商对应app维度参数 |
| placement_format                       |        |          |                     |
| waterfall_list                         |        |          |                     |
| waterfall_list.ecpm                    |        |          |                     |
| waterfall_list.auto_ecpm               |        |          |                     |
| waterfall_list.header_bidding_witch    |        |          |                     |
| waterfall_list.auto_switch             |        |          |                     |
| waterfall_list.day_cap                 |        |          |                     |
| waterfall_list.hour_cap                |        |          |                     |
| waterfall_list.pacing                  |        |          |                     |
| waterfall_list.name                    |        |          |                     |
| waterfall_list.segment_id              |        |          |                     |
| waterfall_list.priority                |        |          |                     |
| waterfall_list.parallel_request_number |        |          |                     |

****

#### 10.1.5 样例

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
        "adsouce_token": {
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
                "header_bidding_witch": 0,
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
                "header_bidding_witch": 0,
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



<h3 id='批量删除ad_source'>10.3 批量删除ad_source</h3>

#### 10.2.1 请求URL

<https://openapi.toponad.com/v1/del_units>

#### 10.2.2 请求方式 

POST

#### 10.2.3 请求参数

| 字段         | 类型         | 是否必传 | 备注            |
| ------------ | ------------ | -------- | --------------- |
| adsource_ids | Array[Int32] | Y        | adsource_id列表 |

#### 10.2.4 返回参数

| 字段 | 类型   | 是否必传 | 备注     |
| ---- | ------ | -------- | -------- |
| msg  | String | N        | 处理信息 |

#### 10.2.5 样例


返回样例：

```
{
    "msg": "suc"
}
```

<h3 id='新增ad_source'>10.2 批量新增和修改ad_source</h3>

#### 10.2.1 请求URL

<https://openapi.toponad.com/v1/set_units>

#### 10.2.2 请求方式 

POST

#### 10.2.3 请求参数

| 字段                | 类型   | 是否必传 | 备注               |
| ------------------- | ------ | -------- | ------------------ |
| count               | Int32  | Y        | 总数               |
| units               | Array  | Y        | unit总数           |
| units.network_id    | Int    | N        | 厂商账号id         |
| units.adsource_id   | Int    | N        | 是否开通report_api |
| units.adsource_name | String | Y        | 厂商维度配置参数   |
| units.adsouce_token | Object | Y        | 厂商app维度数据map |
| units.placement_id  | String | Y        | placement_id       |
| units.default_ecpm  | String | Y        | ecpm               |

#### 10.2.4 返回参数

| 字段          | 类型   | 是否必传 | 备注               |
| ------------- | ------ | -------- | ------------------ |
| network_id    | Int    | N        | 厂商账号id         |
| adsource_id   | Int    | N        | 是否开通report_api |
| adsource_name | String | Y        | 厂商维度配置参数   |
| adsouce_token | Object | Y        | 厂商app维度数据map |
| placement_id  | String | Y        | placement_id       |
| default_ecpm  | String | Y        | ecpm               |

#### 10.2.5 样例


返回样例：

```
[
    {
        "network_id": 307,
        "adsource_id": 19743,
        "adsource_name": "23423423423",
        "adsouce_token": {
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
        "adsouce_token": {
            "size": "asfasd",
            "unit_id": "asdasdafsdddd"
        },
        "placement_id": "123123123",
        "default_ecpm": "",
        "errors": "ad_source_name repeated"
    }
]
```





<h2 id='注意事项'>8. 注意事项</h2>

为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵守。

• 每小时最多请求 1000 次

• 每天请求 10000 次

<h2 id='附录1：go语言示例代码'>9. 附录1：go语言示例代码</h2>
• java、php、python等语言示例代码请参考demo目录

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
	
	_, err = url.ParseQuery(urlParsed.RawQuery)
	
	if err != nil {
	
		fmt.Println(err)
	
		return
	
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
```
<h2 id='附录2：应用一级和二级分类列表'>10. 附录2：应用一级和二级分类列表</h2>
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

<h2 id='附录3：流量分组规则数据格式'>11. 附录3：流量分组规则数据格式</h2>
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
