#### Innotrik Linphone网络接口调试问题

**简要描述:**

`目前获取token以及获取二维码的接口测试没有问题,未调通的接口主要有以下几个,下面做简要描述:`

___

**接口1:创建会议**

**请求URL:**

- `https://ludiqiao.com/conferences`

**请求方式: POST**

**参数:**

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|subject |是  |string |设备名称,没有填写设备ID   |
|deviceId     |否  |string | 设备ID,用处不详    |

- 注: 请求header中添加DeviceToken

**问题与描述:**

`能返回数据但是没有进会议室`

___

**接口2:查找会议室**

**请求URL:**

- `https://ludiqiao.com/conferences/search/findByAccessCode?accessCode=111222333`

**请求方式: GET**

**参数:**

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|accessCode |是  |string |用户输入的accessCode   |

- 注:请求header中添加DeviceToken

**问题与描述**

`这个接口没有问题`

___

**接口3:加入会议**

**请求URL:**

- `https://ludiqiao.com/conferences/{conferenceId}/participants`

**请求方式: POST**

**参数:**

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|deviceId |是  |string |deviceId   |

**问题与描述**

`deviceId在文档中给的参数名是deviceId,但参数给的是 “device-token”,   -----扫面电视机二维码获得的scene中携带. 
传deviceId或者device-token都是返回401`






