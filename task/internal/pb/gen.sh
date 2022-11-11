protoc --go_out=.  --go-grpc_out=.  ./*.proto -I ./

# https://github.com/favadi/protoc-go-inject-tag


# 1. go support for protobuf: 
#    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# 2. go install github.com/favadi/protoc-go-inject-tag 

# 注入自定义标签
protoc-go-inject-tag -input="*.pb.go"