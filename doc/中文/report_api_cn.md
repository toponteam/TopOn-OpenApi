# 修订历史

| 文档版本 | 发布时间      | 修订说明             |
| :--------: | ------------- | -------------------- |
| v 1.0    | 2019年7月17日 | 支持综合报表数据查询 |
| v 2.0    | 2019年8月30日 | 支持用户价值&留存报表数据查询 |
| v 2.1    | 2020年3月17日 | 综合报表支持新增用户、渗透率等指标查询 |
| v 2.2    | 2020年5月15日 | 支持返回广告源ID，LTV和留存支持返回1至60天数据 |


## 1. 关于文档
为提高合作伙伴的变现效率，TopOn平台专门提供了数据报表查询的API接口，可查询综合报表、LTV&留存报表等数据。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

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
| X-Up-Timestamp | API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为15分钟。 |   -                                         |
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
```
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


## 4. 综合报表

### 4.1 请求URL

<https://openapi.toponad.com/v2/fullreport>

### 4.2 请求方式

POST

### 4.3 请求参数

| 字段                 | 类型          | 是否必传 | 备注                                                         | 样例                                       |
| -------------------- | ------------- | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| startdate            | Int           | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501                                   |
| enddate              | Int           | Y        | 结束日期，格式：YYYYmmdd                                     | 20190506                                   |
| app_id_list          | Array[String] | N        | 开发者后台的应用ID，多选                                     | ['xxxxx']                                  |
| placement_id_list    | Array[String] | N        | 开发者后台的广告位ID，多选                                   | ['xxxxx']                                  |
| time_zone            | String        | N        | 时区                                                         | UTC-8,UTC+8,UTC+0                          |
| network_firm_id_list | Array[int32]  | N        | 广告平台ID列表                                                   | ['xxxxx']         |
| adsource_id_list     | Array[int32]  | N        | 广告源ID列表                                              | [121]                   |
| area_list            | Array[String] | N        | 国家列表                                                     | ['xxxxx']           |
| group_by             | Array         | N        | 可选，最多选三个：date（默认值），app，placement，adformat，area，network，adsource，network_firm_id | ["app","placement","area"]<br>network为广告平台账号层级，network_firm_id为广告平台层级  |
| metric               | Array         | N        | 可选，当同时选了all和其他指标时即返回全部指标：default（默认值），all，dau，arpu，request，fillrate，impression，click，ctr，ecpm，revenue，request_api，fillrate_api，impression_api，click_api，ctr_api，ecpm_api | ["dau","arpu","request","click","ctr_api"] |
| start                | Int           | N        | 偏移数，代表从第几条数据开始，默认为0                        | 0                                          |
| limit                | Int           | N        | 每次拉取数据的最大条数，默认是1000，可选[1,1000]             | 1000                                       |

 

- 默认返回的指标：

dau，arpu，request，fillrate，impression，click，ecpm，revenue，impression_api，click_api，ecpm_api

 

### 4.4 返回参数

| 字段             | 类型   | 是否必传 | 备注                                                         |
| ---------------- | ------ | -------- | ------------------------------------------------------------ |
| count            | Int    | Y        | 总条数                                                       |
| date             | String | Y        | 日期，格式：YYYYmmdd。group_by有选才有返回                   |
| app.id           | String | Y        | 开发者后台的应用ID                                           |
| app.name         | String | N        | 应用名称                                                     |
| app.platform     | String | N        | 应用的系统平台                                               |
| placement.id     | String | N        | 开发者后台的广告位ID                                         |
| placement.name   | String | N        | 广告位名称                                                   |
| adformat         | String | N        | rewarded_video/interstitial/banner/native/splash。group_by有选才有返回 |
| area             | String | N        | 国家码。group_by有选才有返回                                 |
| network_firm_id  | String | N        | 广告平台ID。group_by有选network_firm_id才有返回 |
| network_firm     | String | N        | 广告平台名称。group_by有选network_firm_id才有返回 |
| network          | String | N        | 广告平台账号ID。group_by有选network才有返回 |
| network_name     | String | N        | 广告平台账号名称。group_by有选network才有返回 |
| adsource.network | String | N        | 广告源所属的广告平台名称                                     |
| adsource.token   | String | N        | 广告源的三方ID信息，请求广告的appid，slotid等。group_by有选才有返回 |
| time_zone        | String | N        | 枚举值：UTC+8、UTC+0、UTC-8                                  |
| currency         | String | N        | 开发者账号币种，该字段与revenue字段组成的收益需与开发者后台报表的收益一致 |
| new_users        | String | N        | 新增用户                                                     |
| new_user_rate    | String | N        | 新增用户占比                                                 |
| day2_retention   | String | N        | 次日留存                                                     |
| deu              | String | N        | DEU                                                          |
| engaged_rate     | String | N        | 渗透率                                                       |
| imp_dau          | String | N        | 展示 / DAU                                                   |
| imp_deu          | String | N        | 展示 / DEU                                                   |
| impression_rate  | String | N        | 展示率                                                       |
| dau              | String | N        | 根据group_by条件才有返回                                     |
| arpu             | String | N        | 有dau才有该项返回                                            |
| request          | String | N        | 请求数                                                       |
| fillrate         | String | N        | 填充率                                                       |
| impression       | String | N        | 展示数                                                       |
| click            | String | N        | 点击数                                                       |
| ctr              | String | N        | 点击率                                                       |
| ecpm             | String | N        | eCPM                                                         |
| revenue          | String | N        | 收益                                                         |
| request_api      | String | N        | 三方广告平台的请求数                                         |
| fillrate_api     | String | N        | 三方广告平台的填充率                                         |
| impression_api   | String | N        | 三方广告平台的展示数                                         |
| ecpm_api         | String | N        | 三方广告平台的点击数                                         |
| click_api        | String | N        | 三方广告平台的点击率                                         |
| ecpm_api         | String | N        | 三方广告平台的eCPM                                           |
| adsource.adsource_id      | String | N        | 广告源ID，group_by adsource时返回                            |
| adsource.adsource_name    | String | N        | 广告源名称，group_by adsource时返回                          |

### 4.5 样例

请求样例：
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


返回样例：
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

## 5. 用户价值&留存报表

### 5.1 请求URL

<https://openapi.toponad.com/v2/ltvreport>

### 5.2 请求方式

POST
### 5.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ---------------------------------- |
| startdate    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501                            |
| enddate      | Int    | Y        | 结束日期，格式：YYYYmmdd                                     | 20190506                            |
| area_list | Array[String] | N | 国家列表 | ["xxx"] |
| appid_list | Array[String]    | N        | 开发者后台的应用ID                                     | a5c41a9ed1679c                    |
| time_zone | String | N | 时区 | 枚举值：UTC+8、UTC+0、UTC-8 |
| metric      | array    | N        | 可选，默认值：[“ltv_day_1”、”ltv_day_7”、”retention_day_2”、”retention_day_7”][“all”] 表示所有指标  | [“ltv_day_1”， “retention_day_2”]                                   |                        |
| group_by    | array    | N        | 可选，默认值：["app_id”, "date_time", "area"]                                     | ["area"]                                   |
| start    | Int    | N        |     偏移数，代表从第几条数据开始，默认为0                                 |                               0|
| limit    | Int    | N        | 每次拉取数据的最大条数，默认是1000，可选[1,1000]                                   | 1000               |

### 5.4 返回参数

| 字段             | 类型    | 备注                                                         |
| ---------------- | ------ | ------------------------------------------------------------ |
| count            | Int           | 总条数                                                       |
| records             | array       | 记录                   |

**records元素结构如下：**

| 字段名           | 类型   | 备注                     |
| ---------------- | ------ | ------------------------ |
| time_zone        | string | 枚举值：UTC+8、UTC+0、UTC-8  |
| date             | string | 默认返回                 |
| channel          | string | group_by channel时返回   |
| area             | string | group_by area时返回      |
| app.id           | string | 默认返回                 |
| app.name         | string | 默认返回                 |
| new_user         | string | 默认返回                 |
| dau              | string | 默认返回                 |
| currency         | string | 开发者账号币种           |
| revenue          | string | group_by channel时不返回 |
| arpu             | string | 跟随revenue指标          |
| ltv\_day\_1        | string | 默认返回                 |
| ltv\_day\_2        | string | -                         |
| ltv\_day\_3        | string | -                         |
| ltv\_day\_4        | string | -                         |
| ltv\_day\_5        | string | -                         |
| ltv\_day\_6        | string | -                         |
| ltv\_day\_7        | string | 默认返回                 |
| ltv\_day\_14       | string | -                        |
| ltv\_day\_30       | string | -                         |
| ltv\_day\_60       | string | -                         |
| retention\_day\_2  | string | 默认返回                 |
| retention\_day\_3  | string | -                         |
| retention\_day\_4  | string | -                         |
| retention\_day\_5  | string | -                         |
| retention\_day\_6  | string | -                         |
| retention\_day\_7  | string | 默认返回                 |
| retention\_day\_14 | string | -                         |
| retention\_day\_30 | string | -                        |
| retention\_day\_60 | string | -                         |
| time_zone | string | - | 
| arpu | string | - |
| currency | string | - |

> 备注
> 1. 只能查询今天往前推2天的数据
> 2. ltv\_day\_N和retention\_day\_N指标若返回值为“－”，表示这一天该指标不存在，请开发者注意区分

### 5.5 样例

```
{
    "count": 5,
    "records": [
        {
            "date": "20190823",
            "app": {
                "id": "122",
                "name": "我要翘课",
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

## 6. 用户价值1-60天报表

### 6.1 请求URL

<https://openapi.toponad.com/v3/ltvreport>

### 6.2 请求方式

POST
### 6.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ---------------------------------- |
| start_date    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501                            |
| end_date      | Int    | Y        | 结束日期，格式：YYYYmmdd                                     | 20190506                            |
| appid_list | string[] | N | 开发者后台的应用ID列表 | ["xxx"] |
| area_list | string[]    | N        |     国家短码列表        | ["xxxxxx","ddddd"]                   |不传默认publisher 下全部app
| currency | string | Y | 币种：USD |默认按照用户币种信息 |
| time_zone | String | Y | 时区 | 枚举值：UTC+8、UTC+0、UTC-8 |不传默认UTC+8                    |                        |
| group_by    | array    | Y        | 可选，默认值：["app_id”, "date_time", "area"]    | ["area"]        |
| start    | Int    | Y        |     偏移数，代表从第几条数据开始，默认为0                                 |                               0|
| limit    | Int    | Y        | 每次拉取数据的最大条数，默认是1000，可选[1,1000]                                   |         不传默认1000，最大1000       |
| metric    | string[]    | Y        | 维度：["ltv_day_11","ltv_day_12","ltv_day_13"]                 | 只传一个["all"] 代表全部。 其中ltv_day_xx 代表完整的收益，不再是比例               |
| group_by    | string[]    | Y        | group by 维度：["date","app","area"]                                   | date,app 固定存在，一直会有               |


### 6.4 返回参数

| 字段             | 类型    | 备注                                                         |
| ---------------- | ------ | ------------------------------------------------------------ |
| records             | array       | 记录                   |
| count            | Int           | 总条数                                                       |
| time_zone | string | - | 
| currency | string | - |

**records元素结构如下：**

| 字段名           | 类型   | 备注                     |
| ---------------- | ------ | ------------------------ |
| app        | string | app维度信息  |
| app.app_id             | string | app app_id                 |
| app.name             | string | app name                 |
| app.platform          | int32 | app 平台  |
| date             | int32| 日期     |
| ltv_day_xx           | float64 | ltv_day_数字  数字（1-60）                |



> 备注
> 1. 只能查询今天往前推2天的数据

### 6.5 样例

```
{
    "records": [
        {
            "app": {
                "app_id": 67,
                "name": "分手回避",
                "platform": 2
            },
            "date": 20200424,
            "ltv_day_1": 19.563598,
            "ltv_day_10": 0,
            "ltv_day_11": 0,
            "ltv_day_12": 0,
            "ltv_day_13": 0,
            "ltv_day_14": 0,
            "ltv_day_15": 0,
            "ltv_day_16": 0,
            "ltv_day_17": 0,
            "ltv_day_18": 0,
            "ltv_day_19": 0,
            "ltv_day_2": 0,
            "ltv_day_20": 0,
            "ltv_day_21": 0,
            "ltv_day_22": 0,
            "ltv_day_23": 0,
            "ltv_day_24": 0,
            "ltv_day_25": 0,
            "ltv_day_26": 0,
            "ltv_day_27": 0,
            "ltv_day_28": 0,
            "ltv_day_29": 0,
            "ltv_day_3": 0,
            "ltv_day_30": 0,
            "ltv_day_31": 0,
            "ltv_day_32": 0,
            "ltv_day_33": 0,
            "ltv_day_34": 0,
            "ltv_day_35": 0,
            "ltv_day_36": 0,
            "ltv_day_37": 0,
            "ltv_day_38": 0,
            "ltv_day_39": 0,
            "ltv_day_4": 0,
            "ltv_day_40": 0,
            "ltv_day_41": 0,
            "ltv_day_42": 0,
            "ltv_day_43": 0,
            "ltv_day_44": 0,
            "ltv_day_45": 0,
            "ltv_day_46": 0,
            "ltv_day_47": 0,
            "ltv_day_48": 0,
            "ltv_day_49": 0,
            "ltv_day_5": 0,
            "ltv_day_50": 0,
            "ltv_day_51": 0,
            "ltv_day_52": 0,
            "ltv_day_53": 0,
            "ltv_day_54": 0,
            "ltv_day_55": 0,
            "ltv_day_56": 0,
            "ltv_day_57": 0,
            "ltv_day_58": 0,
            "ltv_day_59": 0,
            "ltv_day_6": 0,
            "ltv_day_60": 0,
            "ltv_day_7": 0,
            "ltv_day_8": 0,
            "ltv_day_9": 0
        }
    ],
    "time_zone": "",
    "count": 128
}

```

## 7. 用户留存2-60天报表

### 7.1 请求URL

<https://openapi.toponad.com/v3/retentionreport>

### 7.2 请求方式

POST
### 7.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ---------------------------------- |
| start_date    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501                            |
| end_date      | Int    | Y        | 结束日期，格式：YYYYmmdd                                     | 20190506                            |
| appid_list | string[] | N | 开发者后台的应用ID列表 | ["xxx"] |
| area_list | string[]    | N        |     国家短码列表        | ["xxxxxx","ddddd"]                   |不传默认publisher 下全部app
| currency | string | Y | 币种：USD |不传按用户自己配置来 |
| time_zone | String | Y | 时区 | 枚举值：UTC+8、UTC+0、UTC-8 |不传默认UTC+8                    |                        |
| group_by    | array    | Y        | 可选，默认值：["app_id”, "date_time", "area"]    | ["area"]        |
| start    | Int    | Y        |     偏移数，代表从第几条数据开始，默认为0                                 |                               0|
| limit    | Int    | Y        | 每次拉取数据的最大条数，默认是1000，可选[1,1000]                                   |         不传默认1000，最大1000       |
| metric    | string[]    | Y        | 维度：["retention_day_42","retention_day_43","retention_day_46"]                 | 只传一个["all"] 代表全部。 其中retention_day_46 代表第46天的留存人数，不再是比例             |
| group_by    | string[]    | Y        | group by 维度：["date","app","area"]                                   | date,app 固定存在，一直会有               |


### 7.4 返回参数

| 字段             | 类型    | 备注                                                         |
| ---------------- | ------ | ------------------------------------------------------------ |
| records             | array       | 记录                   |
| count            | Int           | 总条数                                                       |
| time_zone | string | - | 
| currency | string | - |

**records元素结构如下：**

| 字段名           | 类型   | 备注                     |
| ---------------- | ------ | ------------------------ |
| app        | string | app维度信息  |
| app.app_id             | string | app app_id                 |
| app.name             | string | app name                 |
| app.platform          | int32 | app 平台  |
| date             | int32| 日期     |
| retention_day_xx          | int32 | retention_day_数字 （数字2-60）                |



> 备注
> 1. 只能查询今天往前推2天的数据

### 7.5 样例

```
{
    "records": [
        {
            "app": {
                "app_id": 388,
                "name": "一群小辣鸡",
                "platform": 1
            },
            "date": 20200424,
            "retention_day_10": 32,
            "retention_day_11": 20,
            "retention_day_12": 21,
            "retention_day_13": 19,
            "retention_day_14": 0,
            "retention_day_15": 0,
            "retention_day_16": 0,
            "retention_day_17": 0,
            "retention_day_18": 0,
            "retention_day_19": 0,
            "retention_day_2": 297,
            "retention_day_20": 0,
            "retention_day_21": 0,
            "retention_day_22": 0,
            "retention_day_23": 0,
            "retention_day_24": 0,
            "retention_day_25": 0,
            "retention_day_26": 0,
            "retention_day_27": 0,
            "retention_day_28": 0,
            "retention_day_29": 0,
            "retention_day_3": 169,
            "retention_day_30": 0,
            "retention_day_31": 0,
            "retention_day_32": 0,
            "retention_day_33": 0,
            "retention_day_34": 0,
            "retention_day_35": 0,
            "retention_day_36": 0,
            "retention_day_37": 0,
            "retention_day_38": 0,
            "retention_day_39": 0,
            "retention_day_4": 104,
            "retention_day_40": 0,
            "retention_day_41": 0,
            "retention_day_42": 0,
            "retention_day_43": 0,
            "retention_day_44": 0,
            "retention_day_45": 0,
            "retention_day_46": 0,
            "retention_day_47": 0,
            "retention_day_48": 0,
            "retention_day_49": 0,
            "retention_day_5": 75,
            "retention_day_50": 0,
            "retention_day_51": 0,
            "retention_day_52": 0,
            "retention_day_53": 0,
            "retention_day_54": 0,
            "retention_day_55": 0,
            "retention_day_56": 0,
            "retention_day_57": 0,
            "retention_day_58": 0,
            "retention_day_59": 0,
            "retention_day_6": 50,
            "retention_day_60": 0,
            "retention_day_7": 45,
            "retention_day_8": 35,
            "retention_day_9": 32
        }
    ],
    "time_zone": "",
    "count": 128
}

```

## 8. 注意事项
为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵

守。

• 单个用户每小时最多请求 1000 次

• 单个用户每天请求 10000 次
