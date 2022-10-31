/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package main

import (
	"Swyl/servers/swyl-users-ms/controllers"
	"Swyl/servers/swyl-users-ms/dao"
	"Swyl/servers/swyl-users-ms/db"
	"Swyl/servers/swyl-users-ms/routers"
	"Swyl/servers/swyl-users-ms/utils"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice global variables
var (
	server			*gin.Engine
	ctx				context.Context
	mongoClient		*mongo.Client
	userCollection	*mongo.Collection
	ur 				*routers.UserRouter
)

// @dev Runs before main()
func init() {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}

	// init context
	ctx = context.TODO()

	// init mongo client
	mongoClient = db.EstablishMongoClient(ctx)

	// get userCollection
	userCollection = db.GetMongoCollection(mongoClient, "users")

	// init UserDao interface
	ui := dao.UserDaoConstructor(ctx, userCollection)

	// init UserController
	uc := controllers.UserControllerConstructor(ui)

	// init UserRouter
	ur = routers.UserRouterConstructor(uc)

	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})
}

// @dev Root function
func main() {
	// defer a call to `Disconnect()` after instantiating client
	defer func() {if err := mongoClient.Disconnect(ctx); err != nil {panic(err)}}()
	
	// Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// init basePath
	basePath := server.Group("/v1/swyl")

	// init Handler
	ur.UserRoutes(basePath)

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}