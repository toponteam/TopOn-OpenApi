# 修订历史

| 文档版本 | 发布时间      | 修订说明             |
| :-------: | ------------- | -------------------- |
| v 1.0    | 2019年8月29日 | 支持设备维度数据查询 |
| v 1.1    | 2020年3月17日 | 支持币种、时区 |


## 1. 关于文档
为提高合作伙伴的变现效率，TopOn平台专门提供了设备维度数据报告的API接口，可详细了解每个设备的变现情况，实现精细化运营。该文档详细描述了API的使用方法，如需要帮助，请及时与我们联系，谢谢！

## 2. 申请开通权限
在使用设备层级数据报告API前，合作伙伴需向TopOn申请开通权限，具体请咨询与您对接的商务经理。

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
| X-Up-Timestamp | API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为15分钟。 | -                                           |
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

 

## 4. 设备维度数据报告

### 4.1 请求URL

<https://openapi.toponad.com/v1/devicereport>

### 4.2 请求方式

GET

### 4.3 请求参数

| 字段         | 类型   | 是否必传 | 备注                                                         | 样例                                       |
| ------------ | ------ | -------- | ------------------------------------------------------------ | ------------------------------------------ |
| day    | Int    | Y        | 开始日期，格式：YYYYmmdd                                     | 20190501，仅支持2天前的日期                                   |
| app_id       | String | N        | 开发者后台的应用ID，单选                                   | xxxxx                                                                            |
| timezone | Int | N | 时区支持 | -8,8,0三个时区 |

注：数据支持的时间从TopOn运营开通权限后开始支持生成(仅支持2天前的日期)，返回的内容以文件形式提供

### 4.4 返回参数

此接口返回报表数据的下载链接，开发者获取此链接后，可直接在浏览器中下载报表数据。链接格式如下：<br>
https://topon-openapi.s3.amazonaws.com/topon_report_device/dt%3D2019-07-10/publisher_id%3D22/app_id%3Da5d147334b3685/000000_0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIA35FGARBHLHHS7TWB%2F20190828%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20190828T095315Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&X-Amz-Signature=6aaf947f9b2cf02f3acb49d64a3daf719cb0b57a3d5221b0121a006e58b04b10 <br>

文件以csv形式提供，用","作为分隔符，具体参数内容如下：

| 字段             | 类型   | 备注                                                         |
| ---------------- | ------  | ------------------------------------------------------------ |
| placement_id            | String      | 广告位ID                                                      |
| placement_name             | String      | 广告位名称                  |
| placement_format          | String     | 格式 0: native,1: rewarded_video,2: banner,3: interstitial,4: splash                                        |
| unit_id         | String      | TopOn平台生成的广告源id                                                      |
| unit_network     | String       | 广告源所属的广告平台名称                                                |
| unit_token     | String       | 广告源三方ID信息                                   |
| android_id   | String     | 设备ID，androidid                                                |
| gaid         | String      | Google的广告设备ID |
| idfa             | String        | iOS的设备ID                                |
| area          | String       | 国家|
| impression | String       | 展示数                                        |
| click   | String      | 点击数 |
| revenue              | decimal(18,6)       | 收益，货币单位同开发者后台配置一致                                     |
| ecpm             | decimal(18,6)       | 千次展示收益，货币单位同开发者后台配置一致                                          |
| is_abtest             | String       | 对照组或测试组  	<br/>0 表示对照组或未开通A/B测试 <br/> 1 表示测试组             |
| traffic_group_id             | String      | 对照组或测试组id |
| segment_id             | String       | 流量分组ID                                          |
| segment_name             | String      | 流量分组名称                                          |
## 5. 注意事项

为防止频繁请求造成服务器故障，特对请求的频率进行控制，策略如下，请各位合作伙伴遵

守。

• 每小时最多请求 1000 次

• 每天请求 10000 次
