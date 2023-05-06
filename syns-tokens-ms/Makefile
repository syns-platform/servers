include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
SYNS_TOKEN_MS_SERVER=syns-tokens-ms
SYNS_TOKEN_PORT=41106
SYNS_TOKEN_MS_TAG=3.0
SYNS_TOKEN_MS_IMAGE=$(DOCKER_USERNAME)/$(SYNS_TOKEN_MS_SERVER):$(SYNS_TOKEN_MS_TAG)
DOCKER_USERNAME=synsplatform
DOCKER_RMI=docker rmi -f
DOCKER_RM=docker rm -f
DOCKER_PULL=docker pull
DOCKER_PUSH=docker push
DOCKER_IMAGE_LIST_ID=docker images -q
DOCKER_CONTAINER_LIST_ID=docker ps -aq
DOCKER_BUILD_SCRIPT=docker build --no-cache -t $(SYNS_TOKEN_MS_IMAGE) .
DOCKER_RUN_SCRIPT=docker run -d --rm $\
					--name $(SYNS_TOKEN_MS_SERVER) $\
					--env-file .env $\
					-e GIN_MODE=release $\
					-p $(SYNS_TOKEN_PORT):$(PRODUCTION_PORT) $\
					$(SYNS_TOKEN_MS_IMAGE)

ENV_VARS=$(shell grep -v '^#' .env | xargs)
GOOGLE_CLOUD_PROJECT_ID=syns-platform
GOOGLE_CLOUD_REGION=us-central1


GOOGLE_CLOUD_BUILD_SCRIPT=gcloud builds submit --tag gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(SYNS_TOKEN_MS_SERVER)

GOOGLE_CLOUD_DEPLOY_SCRIPT=gcloud run deploy $(SYNS_TOKEN_MS_SERVER) --source . $\ 
						--image gcr.io/$(GOOGLE_CLOUD_PROJECT_ID)/$(SYNS_TOKEN_MS_SERVER) $\
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
	$(COMPILEDAEMON_PATH) -command="./$(SYNS_TOKEN_MS_SERVER)"

# local run
.phony: go-run
go-run:
	@go run .


#############################################
############### DOCKER BUILD ################
#############################################
docker-remove-syns-token-img:
	$(DOCKER_RMI) $(SYNS_TOKEN_MS_IMAGE)

docker-build-syns-token: docker-remove-syns-token-img
	$(DOCKER_BUILD_SCRIPT)

docker-run-syns-token:
	$(DOCKER_RUN_SCRIPT)

docker-update-remote-image: docker-build-syns-token
	$(DOCKER_PUSH) $(SYNS_TOKEN_MS_IMAGE)

docker-pull-syns-token:
	$(DOCKER_PULL) $(SYNS_TOKEN_MS_IMAGE)

docker-dev-syns-token: docker-pull-syns-token docker-run-syns-token


.PHONY: docker-clean
docker-clean:
	$(DOCKER_RM) $(SYNS_TOKEN_MS_SERVER) $&
	$(DOCKER_RMI) $(SYNS_TOKEN_MS_IMAGE)



#############################################
######### Google Cloud BUILD ################
#############################################
gcloud-build:
	$(GOOGLE_CLOUD_BUILD_SCRIPT)

gcloud-deploy: gcloud-build
	$(GOOGLE_CLOUD_DEPLOY_SCRIPT)

update-remote-image-docker-gcloud: docker-update-remote-image gcloud-deploy