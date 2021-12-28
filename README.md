### 🏆年会抽奖系统

#### 🐛项目选型
- go
  
- gin
  
- sqlite
  
- gorm

#### 🔧如何使用

- 拉取项目
  
  `git clone https://git.minieye.tech/yaolongshan/minieye-luckyer.git`

- 下载go依赖

  `go mod download`
  
- 修改本地配置文件`local_conf.json`

  RootPath: 项目根目录
  
  AccessPath: 服务访问的URL，IP或者域名

  ```json
  {
    "RootPath": "/Users/yaolongshan/go/src/code/minieye-luckyer",
    "AccessPath": "http://localhost:8080"
  }
  ```
  
#### 🐒运行

  `go build`生成可执行文件，运行后会在项目根目录生成一个数据库文件`data.db`