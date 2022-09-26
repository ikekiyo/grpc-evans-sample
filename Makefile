proto.build:
	protoc -I=./proto --go_out=./api --go_opt=paths=source_relative proto/*.proto
