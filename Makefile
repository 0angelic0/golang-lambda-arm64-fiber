SHELL := /usr/bin/env bash
ROOT := ${CURDIR}

.PHONY: help
help:
	@echo 'Usage:'
	@echo '    make codegen               	Generate the http handler code from openapi spec.'
	@echo '    make umlgen               	Generate the uml class diagram from code.'
	@echo '    make lint                  	Run the linters.'
	@echo '    make test                  	Run the tests.'
	@echo '    make run                  	Run the program.'
	@echo '    make build                  	Build the programs.'
	@echo '    make deploy-dev              Deploy the programs into dev environment.'
	@echo

.PHONY: codegen
codegen:
	redocly bundle api/user/openapi.yml --output api/user/openapi_bundled.yml
	oapi-codegen --old-config-style -templates tools/oapi-codegen-templates/ -generate types -o pkg/user/http/user_types.gen.go -package http api/user/openapi_bundled.yml
	oapi-codegen --old-config-style -templates tools/oapi-codegen-templates/ -generate server -o pkg/user/http/user_apis.gen.go -package http api/user/openapi_bundled.yml
	sed -i '' "s%github.com/gofiber/fiber%github.com/gofiber/fiber/v2%g" pkg/user/http/user_apis.gen.go
	rm -f api/user/openapi_bundled.yml
	redocly bundle api/role/openapi.yml --output api/role/openapi_bundled.yml
	oapi-codegen --old-config-style -templates tools/oapi-codegen-templates/ -generate types -o pkg/role/http/role_types.gen.go -package http api/role/openapi_bundled.yml
	oapi-codegen --old-config-style -templates tools/oapi-codegen-templates/ -generate server -o pkg/role/http/role_apis.gen.go -package http api/role/openapi_bundled.yml
	sed -i '' "s%github.com/gofiber/fiber%github.com/gofiber/fiber/v2%g" pkg/role/http/role_apis.gen.go
	rm -f api/role/openapi_bundled.yml

.PHONY: umlgen
umlgen:
	goplantuml -recursive -title demo . > diagram.puml
# add this line after @startuml
# !pragma layout smetana

.PHONY: runuser
runuser:
	IS_OFFLINE=true go run ./services/user/.

.PHONY: runrole
runrole:
	IS_OFFLINE=true go run ./services/role/.

.PHONY: build
build:
	export GO111MODULE=on
	go mod tidy
	GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -tags lambda.norpc -o bin/user/bootstrap services/user/main.go
	ditto -c -k --sequesterRsrc bin/user/bootstrap bin/user/bootstrap_user.zip
	GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -tags lambda.norpc -o bin/role/bootstrap services/role/main.go
	ditto -c -k --sequesterRsrc bin/role/bootstrap bin/role/bootstrap_role.zip

.PHONY: deploy-dev
deploy-dev:
	serverless deploy --stage dev --aws-profile mfa