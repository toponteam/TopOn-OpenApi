# TopOn平台DeviceReportAPI对接文档



修订历史

 

| 文档版本 | 发布时间      | 修订说明             |
| -------- | ------------- | -------------------- |
| v 1.0    | 2019年8月29日 | 支持Device Report数据查询 |

 

## 1. 关于文档

为提高合作伙伴的变现效率，TopOn平台专门提供了报表查询的API接口。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

## 2. 申请开通权限

在使用TopOn平台的批量创建 API 前，合作伙伴需向TopOn申请 publisher_key，用于识别来自合作伙伴的请求，申请方法请咨询与您对接的商务经理。

## 3. 接口相关

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

​     SignString = HTTPMethod + "\n" 

​                        \+ Content-MD5 + "\n" 

​                        \+ Content-Type + "\n" 

​                        \+ Headers + "\n"

​                        \+ Resource 

​    如果无body，如下：

​     SignString = HTTPMethod + "\n" 

​                         \+ "\n" 

​                         \+ "\n" 

​                        \+ Headers + "\n"

​                        \+ Resource 

Resource:

​    如请求包含查询字符串（QueryString），则在 Resource 字符串尾部添加 ? 和查询字符串

   QueryString是 URL 中请求参数按字典序排序后的字符串，其中参数名和值之间用 = 相隔组成字符串，并对参数名-值对按照字典序升序排序，然后以 & 符号连接构成字符串。

​    Key1 + "=" + Value1 + "&" + Key2 + "=" + Value2        

Headers：

​     X-Up-Key + X-Up-Timestamp 按字典序升序

​     X-Up-Signature不参与签名计算

​    Key1 + ":" + Value1 + '\n' + Key2 + ":" + Value2           Sign = MD5(HTTPMethod + Content-MD5+ Content-Type + Header + Resource)

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

 

### 3.6. DeviceReportAPI说明

#### 3.6.1. 请求URL

<https://openapi.toponad.com/v1/devicereport>

#### 3.6.2. 请求方式

GET

#### 3.6.3. 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                       |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| day    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501，仅支持2天前的日期                                   |
| app_id       | String | N        | Up开发者后台的App ID，单选                                   | xxxxx                                                                            |

注：数据支持的时间从TopOn运营开通权限后开始支持生成(仅支持2天前的日期)，返回的内容以文件形式提供

#### 3.6.4. 返回参数

此接口返回报表数据的下载链接，开发者获取此链接后，可直接在浏览器中下载报表数据。链接格式如下：<br>
https://topon-openapi.s3.amazonaws.com/topon_report_device/dt%3D2019-07-10/publisher_id%3D22/app_id%3Da5d147334b3685/000000_0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIA35FGARBHLHHS7TWB%2F20190828%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20190828T095315Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=6aaf947f9b2cf02f3acb49d64a3daf719cb0b57a3d5221b0121a006e58b04b10 <br>

文件以csv形式提供，用","作为分隔符，具体参数内容如下：

| 字段             | 类型   | 备注                                                         |
| ---------------- | ------  | ------------------------------------------------------------ |
| placement_id            | String      | 广告位id                                                      |
| placement_name             | String      | 广告位名称                  |
| placement_format          | String     | 格式 0: native,1: rewarded_video,2: banner,3: interstitial,4: splash                                        |
| unit_id         | String      | TopOn平台生成的第三方adsource id                                                      |
| unit_network     | String       | TopOn平台生成的第三方adsource对应的network名称                                                |
| unit_token     | String       | TopOn存储的第三方adsource的具体信息                                   |
| android_id   | String     | 设备id，androidid                                                |
| gaid         | String      | google的广告设备id |
| idfa             | String        | ios的设备id                                |
| area          | String       | 国家|
| impression | String       | 展示数                                        |
| click   | String      | 点击数 |
| revenue              | decimal(18,6)       | 收益，货币单位同开发者后台配置一致                                     |
| ecpm             | decimal(18,6)       | 千次展示收益，货币单位同开发者后台配置一致                                          |


## 4. 注意事项

为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵

守。

• 每小时最多请求 1000 次

• 每天请求 10000 次

## 5. 附录1：go语言示例代码

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
