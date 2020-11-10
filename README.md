## go_micro_demo for micro v2

### 1、安装protocol相关工具
- protoc cmd工具下载地址：https://github.com/protocolbuffers/protobuf/releases

- protoc-gen-go
```
go get github.com/golang/protobuf/protoc-gen-go
```

- protoc-gen-micro（v2.9.3）
```
go get github.com/micro/micro/v2/cmd/protoc-gen-micro
```

- 安装 go-micro v2（v2.9.1）
```
go get github.com/micro/go-micro/v2
```

### proto文件编译看api文件夹内 .proto 文件下的注释