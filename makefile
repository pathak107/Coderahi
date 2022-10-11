run-api:
	cd cmd && go run rest_api.go

swagger-doc:
	swag init --output ../docs
