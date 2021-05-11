SHELL = /bin/bash
.PHONY: proto build version

VERSION_SHORT   ?= $(shell git rev-parse --short HEAD)

define generate-yml
	sed 's/{{ NAMESPACE_KUBERNETES }}/${NAMESPACE_KUBERNETES}/g; s/{{ UPDATED_AT }}/$(shell date)/g; s/{{ ACCOUNT }}/${ACCOUNT}/g; s/{{ ENV_MIOXXO }}/${ENV_MIOXXO}/g; s/{{ AWS_DEFAULT_REGION }}/${AWS_DEFAULT_REGION}/g; s/{{ SERVERADDRESS }}/${SERVERADDRESS}/g' ./kubernetes.tmpl > ./kubernetes.yaml
endef

build-ci:
	docker build -f ./Dockerfile -t app .
	docker tag $(repository)-api ${ACCOUNT}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/$(repository)-api:$(VERSION_SHORT)
	docker push ${ACCOUNT}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/$(repository)-api:$(VERSION_SHORT)
	docker tag $(repository)-api ${ACCOUNT}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/$(repository)-api:latest
	docker push ${ACCOUNT}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/$(repository)-api:latest

deploy-app:
	$(generate-yml)
	kubectl apply --namespace=${NAMESPACE_KUBERNETES} -f ./kubernetes.yaml

unit-tests:
	go test ./...

init-integration-test:
	go run ./main.go

cli-integration-test:
	go test integration/*_test.go
