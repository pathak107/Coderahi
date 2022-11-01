run-api:
	cd cmd && go run rest_api.go

run-admin-panel:
	cd admin_panel && npm start

swagger-doc:
	swag init --output ../docs

build-admin-panel:
	cd admin_panel && npm run build

build-api:
	cd cmd && go build rest_api.go

local-build: build-admin-panel build-api
	rm -rf build/*
	mv cmd/rest_api build
	mv admin_panel/build/* build/
	cp .env build
	cd build && echo "\nGIN_MODE=release" >> .env

docker-build:

local-run:
	cd build && ./rest_api