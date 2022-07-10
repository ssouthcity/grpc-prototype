all: protoc build

build:
	go build -o bin ./...

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/pb/*.proto

clean:
	rm bin/*
	rm -rf proto/pb/*.pb.go
