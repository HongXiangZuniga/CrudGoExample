run:
	docker-compose up -d
	go run cmd/server/api.go

install:
	go mod download

build:
	go build cmd/server/api.go


docker-build:
	docker build \
	-f build/docker/Dockerfile \
	-t  portafolio:local .

docker-run:
	docker-compose up -d
	docker run --rm -it -p 3000:3000 \
	--env-file ./.env \
	portafolio:local