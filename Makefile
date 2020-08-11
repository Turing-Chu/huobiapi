build:
	@swag init -g api.go -d gin-swagger
	@openapi-generator generate -i docs/swagger.yaml -g go  --package-name huobiapi
fmt: ## go fmt
	@go fmt ./...
    @CGO_ENABLED=0 go vet ${PKG_LIST}
    @CGO_ENABLED=0 swag init -g api.go -d gin-swagger/