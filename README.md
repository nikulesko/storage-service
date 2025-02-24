protoc --go_out=crypto --go_opt=paths=source_relative --go-grpc_out=crypto --go-grpc_opt=paths=source_relative crypto.proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u gorm.io/gen
go get -u gorm.io/driver/postgres