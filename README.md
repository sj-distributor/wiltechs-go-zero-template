## wiltechs-go-zero-template

### 环境准备
* 参考 https://github.com/nivin-studio/gonivinck

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

### 注意事项
