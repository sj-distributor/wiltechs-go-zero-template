## wiltechs-go-zero-template

### 环境准备
* 运行环境 https://github.com/nivin-studio/gonivinck
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
* 目录结构树和说明

### 注意事项
1. 数据库版本管理
   * 参考 https://github.com/golang-migrate/migrate
   * 生成sql文件
    ```
    migrate create -ext sql -dir ./service/app/migration/sql -seq create_table_name_table
    ```
   * 编写好sql文件之后运行迁移
    ```
    migrate -path ./service/foundation/migration/sql -database "mysql://username:password@tcp(localhost:3306)/database" up
    ```
