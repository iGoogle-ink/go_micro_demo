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
或
go get github.com/micro/protoc-gen-micro/v2
```

- 安装 go-micro v2（v2.9.1）

```
go get github.com/micro/go-micro/v2
```

### proto文件编译看api文件夹内 .proto 文件下的注释

```
# cd 到 .proto 文件夹下，执行
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto（官方）
protoc --proto_path=. --micro_out=. --go_out=. greeter.proto
```