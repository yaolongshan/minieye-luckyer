### 🥇 年会抽奖接口文档

##### ⛴ 用户登录及校验

路径: `/api/login`

HTTP: `POST`

请求示例:

```json
{
    "unm":"admin", // 账号
    "pwd":"Minieye2022" // 密码
}
```

##### 🗽 判断用户是否登录

路径: `/api/islogin`

HTTP: `GET`

##### 🚗 获取所有员工

  路径: `/api/user/list`

  HTTP: `GET`

  返回示例:

  ```json
  {
    "Status": true,
    "Count": 430, //员工数量
    "Users": [
      {
        "ID": 1,
        "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
        "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
        "DeletedAt": null,
        "Name": "XXX",
        "Phone": "18811111111",
        "Type": "实习", //工作类型：实习、全职
        "Number": "00888", //工号
        "IsLucky": false //是否中奖过
      },
      {
        "ID": 2,
        "CreatedAt": "2021-12-21T17:22:27.733072+08:00",
        "UpdatedAt": "2021-12-21T17:22:27.733072+08:00",
        "DeletedAt": null,
        "Name": "XXX",
        "Phone": "13611111111",
        "Type": "全职",
        "Number": "00666",
        "IsLucky": false //是否中奖过
      }
    ]
  }
  ```

##### 🚕 新增员工

  路径: `/api/user/add`

  HTTP: `POST`

  参数类型: `json`

  请求示例:

  ```json
  {
    "name": "张三",
    "phone": "110",
    "type": "实习", //只有 `实习` 和 `全职` 两种类型
    "number": "111" //工号
  }
  ```

  返回示例:

  ```json
  {
    "Status": true,
    "Msg": "ok"
  }
  ```

##### 🚙 添加一个奖项

路径: `/api/prize/add`

HTTP: `POST`

参数类型: `json`

请求示例:

  ```json
  {
    "level": "一等奖", //奖项等级不可重复
    "name": "神秘大礼包", //奖品名称
    "sum": 10,//奖项数量
    "draw_number": 2, //设置这个奖项每次抽奖的数量
    "image_base64": "xxxxxxxx" //图片base64数据
  }
  ```

请求成功返回示例:

```json
{
  "Status": true,
  "Msg": "ok",
  "Prize": {
    "ID": 1,
    "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
    "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
    "DeletedAt": null,
    "Level": "一等奖", // 奖项级别
    "Name": "神秘大礼包", // 奖品名称
    "Sum": 10,    // 奖项数量
    "AlreadyUsed": 0, // 已抽数量
    "DrawNumber": 2, // 每次抽奖的数量
    "ImageUrl": "/api/images/isdwkkskw.jpg" //图片url
  }
}
```

##### 🚌 获取奖项列表

路径: `/api/prize/list`

HTTP: `GET`

返回示例:

```json
{
  "Status": true,
  "Count": 5,
  "Prizes": [
    {
      "ID": 1,
      "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
      "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
      "DeletedAt": null,
      "Level": "一等奖", // 奖项级别
      "Name": "神秘大礼包", // 奖品名称
      "Sum": 10,   // 奖项数量
      "AlreadyUsed": 10, // 已抽数量
      "DrawNumber": 2, // 每次抽奖的数量
      "ImageUrl": "/api/images/isdwkkskw.jpg" //图片url
    }
  ]
}
```

##### 🚎 修改一个奖项的数量

路径: `/api/prize/update`

HTTP: `POST`

参数类型: `json`

请求示例:

```json
{
  "id":1, //要修改奖项的id
  "sum":10 //修改后的数量
}
```

##### ⛱ 修改一个奖项每次抽奖的数量

路径: `/api/prize/change`

HTTP: `POST`

参数类型: `json`

请求示例:

```json
{
  "id":1, //要修改奖项的id
  "number":10 //修改为每次抽奖抽取多少个
}
```

##### 🚐 删除一个奖项

路径: `/api/prize/delete`

HTTP: `DELETE`

参数: `id` 奖项的id

请求示例:

```http
http://localhost:8080/api/prize/delete?id=1
```

##### 🚒 获取中奖名单列表

路径: `/api/lucky/list`

HTTP: `GET`

返回示例:

```json
{
  "Status": true,
  "Count": 20,
  "LuckyList": [
    {
      "ID": 1,
      "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
      "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
      "DeletedAt": null,
      "UserID": 225,
      "Name": "张三",
      "Number": "00001",
      "Phone": "18088888888",
      "PrizeLevel": "特等奖",
      "Content": "美的洗衣机一台"
    }
  ]
}
```

##### 🏎 下载中奖名单表格文件

路径: `/api/lucky/file`

HTTP: `GET`

##### 🚄 获取阳光普照奖列表

路径: `/api/lucky/notlist`

HTTP: `GET`

返回示例:

```json
{
  "Status": true,
  "Count": 20,
  "LuckyList": [
    {
      "ID": 1,
      "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
      "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
      "DeletedAt": null,
      "UserID": 225,
      "Name": "张三",
      "Number": "00001",
      "Phone": "18088888888",
      "PrizeLevel": "阳光普照奖",
      "Content": "京东卡/沃尔玛购物卡"
    }
  ]
}
```

##### 🚔 下载阳光普照奖表格文件

路径: `/api/lucky/notfile`

HTTP: `GET`

##### 🚜 抽奖接口

路径: `/api/lucky/random`

HTTP: `GET`

参数: `id`  奖项的id

请求示例:

```http
http://localhost:8080/api/lucky/random?id=1
```

返回示例:

```json
{
  "Status": true,
  "Count": 10, //返回结果数量
  "Participants": 100, //本次抽奖参与人数
  "PrizeRemaining": 10, //奖项剩余数量
  "Results": [
    {
      "Name": "张三",
      "Phone": "18000000000",
      "Number": "00001",
    }
  ]
}
```

##### 🛴 发送一条中奖通知  短信&钉钉消息

短信: `/api/sms/send`

钉钉消息: `/api/ding/send`

HTTP: `POST`

参数类型: `json` 

请求示例:

```json
{
  "name":"张三",
  "phone":"18088888888",
  "content":"特等奖，美的冰箱一台" //内容可以是这种格式，奖项级别+奖品内容，也可以只有奖项级别
}
```

返回示例:

```json
{
  "Status": true,
  "Msg": "发送成功"
}
```

##### ✈️ 获取所有祝福语

路径: `/api/greeting/list`

HTTP: `GET`

返回示例:

```json
{
  "Status": true,
  "Count": 20,
  "Greetings": [
    {
      "ID": 1,
      "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
      "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
      "DeletedAt": null,
      "Name": "张三",
      "Number": "00999",
      "Phone": "18088888888",
      "Greeting": "祝福语内容xxxxxxxx",
      "IsLucky": false
    }
  ]
}
```

##### 🛩 添加一条祝福语

路径: `/api/greeting/add`

HTTP: `POST`

参数类型: `json`

请求示例:

```json
{
  "name": "张三",
  "number": "00999",
  "phone": "18088888888",
  "greeting": "祝福语内容xxxxx"
}
```

##### 🚀 祝福语抽奖接口

路径:  `/api/greeting/random`

HTTP:  `GET`

参数:  `count` 抽取祝福语的数量

请求示例:

```http
http://192.168.17.115:8080/api/greetings/random?count=10
```

返回示例:

```json
{
  "Status": true,
  "Count": 10, // 返回结果数量
  "Participants": 300, // 参与抽奖的祝福语数量
  "Results": [
    {
      "Name": "张三",
      "Number": "00999",
      "Greeting": "这是一条祝福语"
    }
  ]
}
```

##### 🚤 获取中奖的祝福语

路径:  `/api/greeting/luckylist`

HTTP:  `GET`

返回示例:

```json
{
  "Status": true,
  "Count": 20,
  "Greetings": [
    {
      "ID": 1,
      "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
      "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
      "DeletedAt": null,
      "Name": "张三",
      "Number": "00999",
      "Greeting": "祝福语内容xxxxxxxx",
      "IsLucky": true
    }
  ]
}
```

##### 🛰 下载中奖的祝福语表格文件

路径:  `/api/greeting/file`

HTTP:  `GET`

##### 📞 初始化数据库，用于重置表中数据

路径:  `/api/db/init`

HTTP:  `GET`
