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
              "Mail": "xxxxxxx@minieye.cc"
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
              "Mail": "xxxxxxx@minieye.cc"
          },
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
    	"name":"一等奖", //奖项名称不可重复
    	"sum":10 //奖项数量
  }
  ```

##### 获取奖项列表

路径: `/api/prize/list`

HTTP: `GET`

##### 修改一个奖项的数量

路径: `/api/prize/update`

HTTP: `POST`

参数类型: `json`

请求示例:

```json
{
  	"name":"一等奖", //要修改奖项的名称
  	"sum":10 //修改后的数量
}
```