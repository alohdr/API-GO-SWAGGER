all: validate build

validate:
	swagger validate ./api/swagger.yaml

build:
	swagger generate server -f ./swagger.yaml -A blog-api -t ./api


