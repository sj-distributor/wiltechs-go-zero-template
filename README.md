## wiltechs-go-zero-template

### 环境准备
* 运行环境 https://github.com/nivin-studio/gonivinck
* 安装protoc & protoc-gen-go https://go-zero.dev/cn/docs/prepare/protoc-install
* Goctl工具 https://go-zero.dev/cn/docs/goctl/goctl
* 安装golang-migrate客户端
    ```
    Mac
    brew install golang-migrate
    
    Windows
    scoop install migrate
    ```

### 开始
克隆项目
```
git clone -o seed -b main --single-branch https://github.com/sj-distributor/wiltechs-go-zero-template.git
```

安装依赖包
```
go mod tidy
```

项目配置
```
cp ./service/app/api/etc/app.yaml.example ./service/app/api/etc/app.yaml 
```
```
cp ./service/app/rpc/etc/app.yaml.example ./service/app/rpc/etc/app.yaml 
```

启动服务
```
make run_rpc & make run_api
```

### 目录结构
```
wiltechs-go-zero-template  //工程名称
├── Makefile
├── README.md
├── common  //通用库
│   └── jwtx
│     
├── goctl  
│   └── template  //goctl模版文件
│    
├── go.mod
├── go.sum
└── service  //服务存放目录
    └── app  //app服务
        ├── api
        ├── migration
        ├── model
        └── rpc
```

### 注意事项
1. 数据库迁移 https://github.com/golang-migrate/migrate
   * 生成sql文件
    ```
    migrate create -ext sql -dir ./service/app/migration/sql -seq create_table_name_table
    ```
   * 编写好sql文件之后运行迁移
    ```
    migrate -path ./service/app/migration/sql -database "mysql://username:password@tcp(localhost:3306)/database" up
    ```
   * 连接数据库生成或更新model
    ```
    goctl model mysql datasource -url="username:password@tcp(localhost:3306)/database" -table="table_name"  -dir="./service/app/model" --home="./goctl/template" -c -style go_zero
    ```
2. 错误处理
   ```go
   // rpc返回错误
   if in.Id != "1" {
       return nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQES_PARAM_ERROR), "GetUser req: %+v", in)
   }
   ```
   这里使用 go 默认的 errors 的包 errors.Wrapf，第一个参数是给前端的友好提示，第二个参数是写入日志的内容。可在xerr包中配置常用错误码和提示信息

   ```go
    // api返回错误
   user, err := l.svcCtx.AppRpc.GetUser(l.ctx, &appclient.IdRequest{ Id: req.Id })
   if err != nil {
       return nil, errors.Wrapf(err, "GetUser req: %+v", req)
   }
   ```
   第一个参数直接拿rpc返回的err（如果是单体应用则传入给前端的友好提示），第二个参数是写入日志的内容

   ```go
   if err != nil {
       logx.WithContext(l.ctx).Errorf("User is exist : %+v", err)
   }
   ```
   如果不需要 return ，只需要日志记录一下错误可以使用logx.WithContext