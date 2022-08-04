all: validate build

validate:
	swagger validate ./api/swagger.yaml

build:
	swagger generate server -f ./api/swagger.yaml -A blog-api
