# TopOn平台报表数据查询API对接文档

## 修订历史

| 文档版本 | 发布时间      | 修订说明             |
| :--------: | ------------- | -------------------- |
| v 1.0    | 2019年7月17日 | 支持查询综合报表数据 |
| v 2.0    | 2019年8月30日 | 支持LTV报表数据 |

## 目录

[1. 关于文档](#关于文档)</br>
[2. 申请开通权限](#申请开通权限)</br>
[3. 接口校验](#接口校验)</br>
[4. 查询综合报表数据](#查询综合报表数据)</br>
[5. 查询LTV报表数据](#查询LTV报表数据)</br>
[6. 注意事项](#注意事项)</br>
[7. 附录1：go语言示例代码](#附录1：go语言示例代码)</br>

<h2 id='关于文档'>1. 关于文档</h2>

为提高合作伙伴的变现效率，TopOn平台专门提供了报表查询的API接口。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

<h2 id='申请开通权限'>2. 申请开通权限</h2>

在使用TopOn平台的批量创建 API 前，合作伙伴需向TopOn申请 publisher_key，用于识别来自合作伙伴的请求，申请方法请咨询与您对接的商务经理。

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

   如请求包含查询字符串（QueryString），则在 Resource 字符串尾部添加 ? 和查询字符串

   QueryString是 URL 中请求参数按字典序排序后的字符串，其中参数名和值之间用 = 相隔组成字符串，并对参数名-值对按照字典序升序排序，然后以 & 符号连接构成字符串。

    Key1 + "=" + Value1 + "&" + Key2 + "=" + Value2        

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

 
<h2 id='查询综合报表数据'>4. 查询综合报表数据</h2>

### 4.1 请求URL

<https://openapi.toponad.com/v1/fullreport>

### 4.2 请求方式

POST

### 4.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                       |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| startdate    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501                                   |
| enddate      | Int    | Y        | 结束日期，格式：YYYYmmdd                                     | 20190506                                   |
| app_id       | String | N        | Up开发者后台的App ID，单选                                   | xxxxx                                      |
| placement_id | String | N        | Up开发者后台的Placement ID，单选                             | xxxxx                                      |
| group_by     | Array  | N        | 可选，最多选三个：date（默认值），app，placement，adformat，area，network，adsource | ["app","placement","area"]                 |
| metric       | Array  | N        | 可选，当同时选了all和其他指标时即返回全部指标：default（默认值），all，dau，arpu，request，fillrate，impression，click，ctr，ecpm，revenue，request_api，fillrate_api，impression_api，click_api，ctr_api，ecpm_api | ["dau","arpu","request","click","ctr_api"] |
| start        | Int    | N        | 偏移数，代表从第几条数据开始，默认为0                        | 0                                          |
| limit        | Int    | N        | 每次拉取数据的最大条数，默认是1000，可选[1,1000]             | 1000                                       |

 

- 默认返回的指标：

dau，arpu，request，fillrate，impression，click，ecpm，revenue，impression_api，click_api，ecpm_api

 

### 4.4 返回参数

| 字段             | 类型   | 是否必传 | 备注                                                         |
| ---------------- | ------ | -------- | ------------------------------------------------------------ |
| count            | Int    | Y        | 总条数                                                       |
| date             | String | Y        | 日期，格式：YYYYmmdd。group_by有选才有返回                   |
| app.id           | String | Y        | Up开发者后台的App ID                                         |
| app.name         | String | N        | App名称                                                      |
| app.platform     | String | N        | App的系统平台                                                |
| placement.id     | String | N        | Up开发者后台的Placement ID                                   |
| placement.name   | String | N        | Placement名称                                                |
| adformat         | String | N        | rewarded_video/interstitial/banner/native/splash。group_by有选才有返回 |
| area             | String | N        | 国家码。group_by有选才有返回                                 |
| network          | String | N        | facebook/admob/toutiao/gdt/baidu/mintegral……。group_by有选才有返回 |
| adsource.network | String | N        | 三方unit的network名称                                        |
| adsource.token   | String | N        | 三方unit的token信息，请求广告的appid，slotid等。group_by有选才有返回 |
| dau              | String | N        | 根据group_by条件才有返回                                     |
| arpu             | String | N        | 有dau才有该项返回                                            |
| request          | String | N        | 请求数                                                       |
| fillrate         | String | N        | 填充率                                                       |
| impression       | String | N        | 展示数                                                       |
| click            | String | N        | 点击数                                                       |
| ctr              | String | N        | 点击率                                                       |
| ecpm             | String | N        | ecpm                                                         |
| revenue          | String | N        | 收益                                                         |
| request_api      | String | N        | 三方广告平台的请求数                                         |
| fillrate_api     | String | N        | 三方广告平台的填充率                                         |
| impression_api   | String | N        | 三方广告平台的展示数                                         |
| ecpm_api         | String | N        | 三方广告平台的点击数                                         |
| click_api        | String | N        | 三方广告平台的点击率                                         |
| ecpm_api         | String | N        | 三方广告平台的ecpm                                           |

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

<h2 id='查询LTV报表数据'>5. 查询LTV报表数据</h2>

### 5.1 请求URL

<https://openapi.toponad.com/v1/ltvreport>

### 5.2 请求方式

POST
### 5.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                       |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| startdate    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501                                   |
| enddate      | Int    | Y        | 结束日期，格式：YYYYmmdd                                     | 20190506                                   |
| app_id      | String    | N        | Up开发者后台的App ID，单选                                     | a5c41a9ed1679c                                   |
| metric      | array    | N        | 可选，默认值：[“ltv_day_1”、”ltv_day_7”、”retention_day_2”、”retention_day_7”][“all”] 表示所有指标  | [“ltv_day_1”， “retention_day_2”]                                   |
| order_by      | array    | N        | 可选，默认值：[“date_time”, “desc”, “revenue”, “desc”, “dau”, “desc”, “new_user”, “desc”, “app_id”, “desc”]  |["date_time", “asc”, “app_id”, “desc”]                             |
| group_by    | array    | N        | 可选，默认值：["app_id”, "date_time", "area", "channel"]                                     | ["area", "channel"]                                   |
| start    | Int    | N        |     偏移数，代表从第几条数据开始，默认为0                                 |                                    0|
| limit    | Int    | N        | 每次拉取数据的最大条数，默认是1000，可选[1,1000]                                   | 1000                                 |

### 5.4 返回参数

| 字段             | 类型    | 备注                                                         |
| ---------------- | ------ | ------------------------------------------------------------ |
| count            | Int           | 总条数                                                       |
| records             | array       | 记录                   |

**records元素结构如下：**

| 字段名           | 类型   | 备注                     |     |
| ---------------- | ------ | ------------------------ | --- |
| date             | string | 默认返回                 |     |
| app.id           | string    | 默认返回                 |     |
| app.name         | string | 默认返回                 |     |
| new_user         | string | 默认返回                 |     |
| dau              | string | 默认返回                 |     |
| revenue          | string | group by channel时不返回 |     |
| arpu             | string | 跟随revenue指标          |     |
| ltv\_day\_1        | string | 默认返回                 |     |
| ltv\_day\_2        | string |                          |     |
| ltv\_day\_3        | string |                          |     |
| ltv\_day\_4        | string |                          |     |
| ltv\_day\_5        | string |                          |     |
| ltv\_day\_6        | string |                          |     |
| ltv\_day\_7        | string | 默认返回                 |     |
| ltv\_day\_14       | string |                          |     |
| ltv\_day\_30       | string |                          |     |
| ltv\_day\_60       | string |                          |     |
| retention\_day\_2  | string | 默认返回                 |     |
| retention\_day\_3  | string |                          |     |
| retention\_day\_4  | string |                          |     |
| retention\_day\_5  | string |                          |     |
| retention\_day\_6  | string |                          |     |
| retention\_day\_7  | string |                          |     |
| retention\_day\_14 | string | 默认返回                 |     |
| retention\_day\_30 | string |                          |     |
| retention\_day\_60 | string |                          |     |

> 备注
> 1. 只能查询今天往前推2天的数据
> 2. ltv\_day\_N和retention\_day\_N指标若返回值为“－”，表示这一天该指标不存在，请开发者注意区分

### 5.5 样例

``` javascript
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

<h2 id='注意事项'>6. 注意事项</h2>

为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵

守。

• 每小时最多请求 1000 次

• 每天请求 10000 次

<h2 id='附录1：go语言示例代码'>7. 附录1：go语言示例代码</h2>

• java、php、python等语言示例代码请参考demo目录

```
package main

import (

​	"bytes"

​	"crypto/md5"

​	"encoding/hex"

​	"fmt"

​	"io/ioutil"

​	"net/http"

​	"net/url"

​	"sort"

​	"strconv"

​	"strings"

​	"time"

)

 

func main() {

​	//openapi的地址

​	demoUrl := "请求URL"

​	//提交的body数据

​	body := "{}"

​	//您申请的publisherKey

​	publisherKey := "请填写您的publisherKey"

​	//请求方式

​	httpMethod := "POST"

​	contentType := "application/json"

​	publisherTimestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

​	headers := map[string]string{

​		"X-Up-Timestamp": publisherTimestamp,

​		"X-Up-Key":       publisherKey,

​	}

​	//处理queryPath

​	urlParsed, err := url.Parse(demoUrl)

​	if err != nil {

​		fmt.Println(err)

​		return

​	}

​	//处理resource

​	resource := urlParsed.Path

​	m, err := url.ParseQuery(urlParsed.RawQuery)

​	if err != nil {

​		fmt.Println(err)

​		return

​	}

​	queryString := m.Encode()

​	if queryString != "" {

​		resource += "?" + queryString

​	}

 

​	//处理body

​	h := md5.New()

​	h.Write([]byte(body))

​	contentMD5 := hex.EncodeToString(h.Sum(nil))

​	contentMD5 = strings.ToUpper(contentMD5)

 

​	publisherSignature := signature(httpMethod, contentMD5, contentType, headerJoin(headers), resource)

 

​	request, err := http.NewRequest(httpMethod, demoUrl, bytes.NewReader([]byte(body)))

​	if err != nil {

​		fmt.Println("Fatal error", err.Error())

​		return

​	}

​	client := &http.Client{}

​	request.Header.Set("Content-Type", contentType)

​	request.Header.Set("X-Up-Key", publisherKey)

​	request.Header.Set("X-Up-Signature", publisherSignature)

​	request.Header.Set("X-Up-Timestamp", publisherTimestamp)

​	resp, err := client.Do(request)

​	defer resp.Body.Close()

​	content, err := ioutil.ReadAll(resp.Body)

​	if err != nil {

​		fmt.Println("Fatal error", err.Error())

​		return

​	}

 

​	//返回数据

​	fmt.Println(string(content))

 

}

 

func headerJoin(headers map[string]string) string {

​	headerKeys := []string{

​		"X-Up-Timestamp",

​		"X-Up-Key",

​	}

​	sort.Strings(headerKeys)

​	ret := make([]string, 0)

​	for _, k := range headerKeys {

​		v := headers[k]

​		ret = append(ret, k+":"+v)

​	}

​	return strings.Join(ret, "\n")

}

 

func signature(httpMethod, contentMD5, contentType, headerString, resource string) string {

​	stringSection := []string{

​		httpMethod,

​		contentMD5,

​		contentType,

​		headerString,

​		resource,

​	}

​	stringToSign := strings.Join(stringSection, "\n")

 

​	h := md5.New()

​	h.Write([]byte(stringToSign))

​	resultMD5 := hex.EncodeToString(h.Sum(nil))

​	fmt.Println(stringToSign)

​	return strings.ToUpper(resultMD5)

}
```
