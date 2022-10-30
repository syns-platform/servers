include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
SERVER=server


############### LOCAL BUILD ###############

# dev-mode
.phony: dev-mode
dev-mode: 
	@$(COMPILEDAEMON_PATH) -command="./$(SERVER)"

# local run
.phony: go-run
go-run:
	@go run .