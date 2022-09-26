proto.build:
	protoc -I=./proto --go_out=./api/gen --go_opt=paths=source_relative proto/*.proto
