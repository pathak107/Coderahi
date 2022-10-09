run-api:
	cd cmd && go run rest_api.go

swagger-doc:
	cd cmd 
	swag init --output ../docs
