include .env

IMAGE_VERSION=$(DICT_VERSION)
IMAGE_NAME=lapierre/{{.Name}}

CONNECTION_STRING=postgres://app:qwedsazxc@localhost:5432/app?sslmode=disable
DATASOURCE_USER=app
DICT_DATASOURCE_PASSWORD=qwedsazxc

modelgen:
	genna model -c $(CONNECTION_STRING) -o model/model.go -k -g 9

build:
	cd cmd/{{.Name}} && CGO_ENABLED=0 go build -a -installsuffix cgo -o {{.Name}} .

docker: build
	cd cmd/{{.Name}} && docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) .
	docker tag $(IMAGE_NAME):$(IMAGE_VERSION) $(IMAGE_NAME):latest

push:
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)
	docker push $(IMAGE_NAME):latest

run:
	cd cmd/{{.Name}} && go build -o /tmp/___go_build_main_go main.go
	DICT_SHOW_SQL=true DICT_DATASOURCE_USER=$(DATASOURCE_USER) DICT_DATASOURCE_PASSWORD=$(DICT_DATASOURCE_PASSWORD) /tmp/___go_build_main_go