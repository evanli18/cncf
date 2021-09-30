export tag=v1.0
root:
	export ROOT=github.com/evanli18/cncf
	export GO111MODULE="auto"

build:
	echo "building httpserver binary"
	mkdir -p output
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/httpserver ./main.go

release: build
	echo "building httpserver container"
	docker build -t evanchn/httpserver:${tag} .

push: release
	echo "pushing evanchn/httpserver"
	docker push evanchn/httpserver:${tag}