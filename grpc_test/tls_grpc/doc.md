# 1.pb生成文件
```
protoc --go_out=plugins=grpc:./ .\product.proto (旧)
protoc  --go_out=./service --go-grpc_out=./service  pbfile\product.proto （新）
    > $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
grpc:生成存放的目录 proto文件
```

