include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
SWYL_SERVER=swyl-users-ms
SWYL_USER_MS_IMAGE=logann131/swyl-users-ms:$(SWYL_USER_MS_TAG)
SWYL_USERS_PORT=41105
SWYL_USER_MS_TAG=1.0
DOCKER_RMI=docker rmi -f
DOCKER_RM=docker rm -f
DOCKER_PULL=docker pull
DOCKER_PUSH=docker push
DOCKER_IMAGE_LIST_ID=docker images -q
DOCKER_CONTAINER_LIST_ID=docker ps -aq
DOCKER_BUILD_SCRIPT = docker build --no-cache -t $(SWYL_USER_MS_IMAGE) .
DOCKER_RUN_SCRIPT = docker run -d --rm $\
					--name $(SWYL_SERVER) $\
					--env-file .env $\
					-e GIN_MODE=release $\
					-p $(SWYL_USERS_PORT):$(PRODUCTION_PORT) $\
					$(SWYL_USER_MS_IMAGE)


#############################################
############### LOCAL BUILD #################
#############################################

# dev-mode
.phony: dev
dev: 
	@$(COMPILEDAEMON_PATH) -command="./$(SWYL_SERVER)"

# local run
.phony: go-run
go-run:
	@go run .


#############################################
############### DOCKER BUILD ################
#############################################
docker-remove-swyl-users-img:
	$(DOCKER_RMI) $(SWYL_USER_MS_IMAGE)

docker-build-swyl-users: docker-remove-swyl-users-img
	$(DOCKER_BUILD_SCRIPT)

docker-run-swyl-users:
	$(DOCKER_RUN_SCRIPT)

docker-update-remote-image: docker-build-swyl-users
	$(DOCKER_PUSH) $(SWYL_USER_MS_IMAGE)

docker-pull-swyl-users:
	$(DOCKER_PULL) $(SWYL_USER_MS_IMAGE)

docker-dev-swyl-users: docker-pull-swyl-users docker-run-swyl-users


.PHONY: docker-clean
docker-clean:
	$(DOCKER_RM) $$($(DOCKER_CONTAINER_LIST_ID)) $&
	$(DOCKER_RMI) $$($(DOCKER_IMAGE_LIST_ID))