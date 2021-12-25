export tag=v1.0
root:
	export GO111MODULE="auto"

build-user:
	echo "building user binary"
	mkdir -p dockerfiles/user/output
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dockerfiles/user/output/user ./cmd/user/main.go

release-user: build-user
	echo "building user container"
	cd dockerfiles/user && \
	docker build -t evanchn/user:${tag} .

push-user: release-user
	echo "pushing evanchn/user"
	cd dockerfiles/user && \
	docker push evanchn/user:${tag}

build-auth:
	echo "building auth binary"
	mkdir -p dockerfiles/auth/output
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dockerfiles/auth/output/auth ./cmd/auth/main.go

release-auth: build-auth
	echo "building auth container"
	cd dockerfiles/auth && \
	docker build -t evanchn/auth:${tag} .

push-auth: release-auth
	echo "pushing evanchn/auth"
	cd dockerfiles/auth && \
	docker push evanchn/auth:${tag}

build-order:
	echo "building order binary"
	mkdir -p dockerfiles/order/output
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dockerfiles/order/output/order ./cmd/order/main.go

release-order: build-order
	echo "building order container"
	cd dockerfiles/order && \
	docker build -t evanchn/order:${tag} .

push-order: release-order
	echo "pushing evanchn/order"
	cd dockerfiles/order && \
	docker push evanchn/order:${tag}

build-bff:
	echo "building bff binary"
	mkdir -p dockerfiles/bff/output
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dockerfiles/bff/output/bff ./cmd/bff/main.go

release-bff: build-bff
	echo "building bff container"
	cd dockerfiles/bff && \
	docker build -t evanchn/bff:${tag} .

push-bff: release-bff
	echo "pushing evanchn/bff"
	cd dockerfiles/bff && \
	docker push evanchn/bff:${tag}
