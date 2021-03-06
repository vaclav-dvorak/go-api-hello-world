check-swagger:
	which swagger || (go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	go mod vendor && swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml --no-open

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile
