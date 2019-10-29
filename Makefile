# GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# DOCKER_BUILD=$(shell pwd)/.docker_build
# DOCKER_CMD=$(DOCKER_BUILD)/go-webserver

# $(DOCKER_CMD): clean
# 	mkdir -p $(DOCKER_BUILD)
# 	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

# clean:
# 	rm -rf $(DOCKER_BUILD)

# heroku: $(DOCKER_CMD)
# 	heroku container:push web
run: bin/go-webserver
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/go-webserver: main.go
	go build -o bin/go-webserver main.go

clean:
	rm -rf bin