package testproto

//go:generate sh -c "protoc ./*.proto --go_out=paths=source_relative,plugins=grpc:./"
