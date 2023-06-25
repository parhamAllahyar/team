proto:
	cd ./api/proto && protoc -I ./ ./customer.proto  --go_out=./  --go-grpc_out=./

run:
	cd ./cmd/main && go run .
build:	
	cd ./cmd/main && go build  main.go
docker-compose:
		sh env.sh && docker compose up

compose:
	docker compose build --no-cache && docker compose up

new:
	docker-compose rm -f && docker-compose pull && docker-compose up --build

rec:
	docker-compose up --build --force-recreate	


AA:
	docker compose down && docker compose build --no-cache && sudo docker compose up 

