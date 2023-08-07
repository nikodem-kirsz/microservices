include .env
export

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh trainings internal/trainings/ports ports
