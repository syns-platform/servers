include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
SERVER=syns-community-ms


############### LOCAL BUILD ###############

# dev-mode
.phony: dev-mode
dev: 
	@$(COMPILEDAEMON_PATH) -command="./$(SERVER)"

# local run
.phony: go-run
go-run:
	@go run .