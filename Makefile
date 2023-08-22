create:
	protoc --proto_path=protofiles/ --go_out=server/ protofiles/answer.proto
	protoc --proto_path=protofiles/ --go-grpc_out=server/ protofiles/answer.proto
#	protoc --proto_path=protofiles/ --grpc-gateway_out .  \
#	--grpc-gateway_opt logtostderr=true \
#	--grpc-gateway_opt paths=source_relative \
#	protofiles/answer.proto
	protoc --proto_path=protofiles \
	--js_out=import_style=commonjs,binary:app/src/protos \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:app/src/protos \
	protofiles/answer.proto
clean:
	rm -rf server/protogen/*.go
	rm -rf app/src/protos/*.js

build:
	docker compose -f docker-compose.yml build

run:
	docker compose up -d

stop:
	docker compose down

logs:
	docker compose logs -f
