OWNER := fabianlee
PROJECT := docker-golang-hello-world-web
VERSION := 1.0.0
OPV := $(OWNER)/$(PROJECT):$(VERSION)
WEBPORT := 8080:8080

BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
# unique id from last git commit
MY_GITREF := $(shell git rev-parse --short HEAD)

## builds docker image
docker-build:
	echo $(MY_COMMIT)
	docker build --build-arg MY_VERSION=$(VERSION) --build-arg MY_BUILDTIME=$(BUILD_TIME) -f Dockerfile -t $(OPV) .

## cleans docker image
clean:
	sudo docker image rm $(OPV) | true

## runs container in foreground, testing a couple of override values
docker-test-fg:
	sudo docker run -it -p $(WEBPORT) -e APP_CONTEXT=/hello/ -e MY_NODE_NAME=node1 --rm $(OPV)

## runs container in foreground, override entrypoint to use use shell
docker-test-cli:
	sudo docker run -it --rm --entrypoint "/bin/sh" $(OPV)

## run container in background
docker-run-bg:
	sudo docker run -d -p $(WEBPORT) --rm --name $(PROJECT) $(OPV)

## get into console of container running in background
docker-cli-bg:
	sudo docker exec -it $(PROJECT) /bin/sh

## tails docker logs
docker-logs:
	sudo docker logs -f $(PROJECT)

## stops container running in background
docker-stop:
	sudo docker stop $(PROJECT)


## pushes to docker hub
docker-push:
	sudo docker push $(OPV)

