APP = skeleton-svc
CURRENT_DIR = $(shell pwd)

.PHONY: test
test: 
	go test ./... -coverprofile=coverage.out --cover

.PHONY: build
build: 
	env GOOS=linux GOARCH=amd64 go build

.PHONY: report
report:
	goreportcard-cli -v > goreportcard.txt

.PHONY: init
init:
	go mod init

.PHONY: install
install:
	go mod tidy

.PHONY: mod
mod: init install

.PHONY: serve
serve:
	./$(APP)

.PHONY: run
run:
	go run main.go

.PHONY: build-docker
build-docker:
	go mod vendor
	docker buildx build --platform=linux/amd64 -t $(APP):latest .
	rm -rf vendor

.PHONY: reporter
reporter:
	go get -u github.com/360EntSecGroup-Skylar/goreporter
	goreporter -p $(CURRENT_DIR) -r $(CURRENT_DIR) -f html

.PHONY: goose-create
DIR ?= $(shell bash -c 'read -p "Directory : " dir; echo $$dir')
NAME ?= $(shell bash -c 'read -p "Name Migration : " name; echo $$name')
FORMAT ?= $(shell bash -c 'read -p "Format migration [sql/go] : " format; echo $$format')
goose-create:
	goose -dir=$$(DIR) create $$(NAME) $$(FORMAT)

.PHONY: mock
TARGET_DIR ?= $(shell bash -c 'read -p "Target Dir: " target_dir; echo $$target_dir')
PACKAGE ?= $(shell bash -c 'read -p "Package: " package; echo $$package')
OUTPUT ?= $(shell bash -c 'read -p "Output: " output; echo $$output')
mock:
	cd $(TARGET_DIR)
	mockery --name $(PACKAGE) --output $(OUTPUT)

.PHONY: push-image
push-image:
	docker tag $(APP) 12.12.12.5:5000/$(APP):latest
	docker image push 12.12.12.5:5000/$(APP):latest

.PHONY: push-docker-staging
push-docker-staging: build-docker push-image

# BEGIN __INCLUDE_TEMPLATE__
.PHONY: new-project
DESTINATION ?= $(shell bash -c 'read -p "Destination: " destination; echo $$destination')
new-project:
	@go get github.com/rantav/go-archetype  
	@go install github.com/rantav/go-archetype
	@go-archetype transform --transformations=transformations.yml --source=. --destination=$(DESTINATION)
# END __INCLUDE_TEMPLATE__

