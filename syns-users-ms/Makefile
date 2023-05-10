include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
SYNS_USER_MS_SERVER=syns-users-ms
SYNS_USERS_PORT=41105
SYNS_USER_MS_TAG=3.1
SYNS_USER_MS_IMAGE=$(DOCKER_USERNAME)/$(SYNS_USER_MS_SERVER):$(SYNS_USER_MS_TAG)
DOCKER_USERNAME=synsplatform
DOCKER_RMI=docker rmi -f
DOCKER_RM=docker rm -f
DOCKER_PULL=docker pull
DOCKER_PUSH=docker push
DOCKER_IMAGE_LIST_ID=docker images -q
DOCKER_CONTAINER_LIST_ID=docker ps -aq
DOCKER_BUILD_SCRIPT=docker build --no-cache -t $(SYNS_USER_MS_IMAGE) .
DOCKER_RUN_SCRIPT=docker run -d --rm $\
					--name $(SYNS_USER_MS_SERVER) $\
					--env-file .env $\
					-e GIN_MODE=release $\
					-p $(SYNS_USERS_PORT):$(PRODUCTION_PORT) $\
					$(SYNS_USER_MS_IMAGE)

ENV_VARS=$(shell grep -v '^#' .env | xargs)
GOOGLE_CLOUD_PROJECT_ID=syns-platform
GOOGLE_CLOUD_REGION=us-central1


GOOGLE_CLOUD_BUILD_SCRIPT=gcloud builds submit --tag gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(SYNS_USER_MS_SERVER)

GOOGLE_CLOUD_DEPLOY_SCRIPT=gcloud run deploy $(SYNS_USER_MS_SERVER) --source . $\ 
						--image gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(SYNS_USER_MS_SERVER) $\
						--platform managed $\
						--region $(GOOGLE_CLOUD_REGION) $\
						--allow-unauthenticated $\
						--set-env-vars "GIN_MODE=release,$(ENV_VARS)"


#############################################
############### LOCAL BUILD #################
#############################################

# dev-mode
.phony: dev
dev: 
	@$(COMPILEDAEMON_PATH) -command="./$(SYNS_USER_MS_SERVER)"

# local run
.phony: go-run
go-run:
	@go run .


#############################################
############### DOCKER BUILD ################
#############################################
docker-remove-syns-users-img:
	$(DOCKER_RMI) $(SYNS_USER_MS_IMAGE)

docker-build-syns-users: docker-remove-syns-users-img
	$(DOCKER_BUILD_SCRIPT)

docker-run-syns-users:
	$(DOCKER_RUN_SCRIPT)

docker-update-remote-image: docker-build-syns-users
	$(DOCKER_PUSH) $(SYNS_USER_MS_IMAGE)

docker-pull-syns-users:
	$(DOCKER_PULL) $(SYNS_USER_MS_IMAGE)

docker-dev-syns-users: docker-pull-syns-users docker-run-syns-users


.PHONY: docker-clean
docker-clean:
	$(DOCKER_RM) $(SYNS_USER_MS_SERVER) $&
	$(DOCKER_RMI) $(SYNS_USER_MS_IMAGE)



#############################################
######### Google Cloud BUILD ################
#############################################
gcloud-build:
	$(GOOGLE_CLOUD_BUILD_SCRIPT)

gcloud-deploy: gcloud-build
	$(GOOGLE_CLOUD_DEPLOY_SCRIPT)


update-remote-image-docker-gcloud: docker-update-remote-image gcloud-deploy