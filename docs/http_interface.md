### 年会抽奖接口文档

##### 获取所有员工

  路径: `/api/user/list`

  HTTP: `GET`

  返回示例:

  ```json
  {
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
              "Contract": "深圳佑驾创新科技有限公司", //合同公司
              "Mail": "xxxxxxx@minieye.cc",
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
              "Contract": "深圳佑驾创新科技有限公司",
              "Mail": "xxxxxxx@minieye.cc",
            "IsLucky": false //是否中奖过
          }
      ]
  }
  ```

##### 新增员工

  路径: `/api/user/add`

  HTTP: `POST`

  参数类型: `json`

  请求示例:

  ```json
  {
      "name":"张三",
      "phone":"110",
      "type":"实习", //只有 `实习` 和 `全职` 两种类型
      "number":"111",//工号
      "contract":"深圳佑驾创新科技有限公司",
      "mail":"xxx@minieye.cc"
  }
  ```

  返回示例:

  ```json
  {
      "msg": "ok",
      "status": true
  }
  ```

##### 添加一个奖项

路径: `/api/prize/add`

HTTP: `POST`

参数类型: `json`

请求示例:

  ```json
  {
    	"level":"一等奖", //奖项等级不可重复
    	"name":"神秘大礼包", //奖品名称
    	"sum":10,//奖项数量
    	"image_base64":"xxxxxxxx" //图片base64数据
  }
  ```

请求成功返回示例:

```json
{
  	"Status":true,
  	"Msg":"ok",
  	"Prize":{
      	"ID": 1,
        "CreatedAt": "2021-12-21T17:22:27.732045+08:00",
        "UpdatedAt": "2021-12-21T17:22:27.732045+08:00",
        "DeletedAt": null,
        "Level": "一等奖", // 奖项级别
        "Name": "神秘大礼包", // 奖品名称
        "Sum": 10,    // 奖项数量
        "ImageUrl": "http://xxxxxxx/api/images/isdwkkskw.jpg" //图片url
    }
}
```

##### 获取奖项列表

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
          "ImageUrl": "http://xxxxxxx/api/images/isdwkkskw.jpg" //图片url
        },
      	.....
    ]
}
```

##### 修改一个奖项的数量

路径: `/api/prize/update`

HTTP: `POST`

参数类型: `json`

请求示例:

```json
{
  	"level":"一等奖", //要修改奖项的级别
  	"sum":10 //修改后的数量
}
```

##### 删除一个奖项

路径: `/api/prize/delete`

HTTP: `DELETE`

参数: `id` 奖项的id

请求示例:

```http
http://localhost:8080/api/prize/delete?id=1&count=10
```

##### 获取中奖名单列表

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
        "PrizeLevel": "一等奖"
      }
    ]
}
```

##### 下载中奖名单表格文件

路径: `/api/lucky/file`

HTTP: `GET`

##### 抽奖接口

路径: `/api/lucky/random`

HTTP: `GET`

参数: `id`  奖项的id

参数: `count` 抽奖数量

请求示例:

```http
http://localhost:8080/api/lucky/random?id=1&count=10
```

##### 发送一条中奖通知短信

路径: `/api/sms/send`

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

