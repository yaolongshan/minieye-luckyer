### 🥇 年会抽奖系统

#### 🐛 项目选型
- go
  
- gin
  
- sqlite
  
- gorm

#### 🔧 如何使用

- 拉取项目
  
  `git clone https://git.minieye.tech/yaolongshan/minieye-luckyer.git`

- 下载go依赖

  `go mod download`
  
- 在项目根目录下编辑本地配置文件`local_conf.json`

  ```json
  {
    "Port": 8080, // 端口
    "RootPath": "/Users/yaolongshan/go/src/code/minieye-luckyer", // 项目根目录
    "SMS": { // 阿里云短信相关配置
      "AccessKeyId": "XXXXXXXX",
      "AccessKeySecret": "XXXXXXXXXXX",
      "SignName": "短信签名",
      "TemplateCode": "XXXXXXXXX"
    },
    "DingDing": { // 钉钉消息通知相关配置
      "AppKey": "xxxxxxxxxx",
      "AppSecret": "xxxxxxxxxxxxxxxxxxxxxx"
    }
  }
  ```
  
#### 🐒 运行

  `go build`生成可执行文件，运行后会在项目根目录生成一个数据库文件`data.db`

#### 🎧 工具

`tools/reader.go`文件是辅助工具类

在项目根目录下将`tools/reader.go`编译可执行文件

`go build tools/reader.go`

使用

`./reader -h`