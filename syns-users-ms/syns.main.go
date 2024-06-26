/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package main

import (
	"Syns/servers/syns-users-ms/controllers"
	"Syns/servers/syns-users-ms/dao"
	"Syns/servers/syns-users-ms/db"
	"Syns/servers/syns-users-ms/routers"
	"Syns/servers/syns-users-ms/utils"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice global variables
var (
	server			*gin.Engine
	ctx			context.Context
	mongoClient		*mongo.Client
	ur 			*routers.UserRouter
	fr 			*routers.FeedbackRouter
	drr 			*routers.DemoRequestRouter
)

// @dev Runs before main()
func init() {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}

	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

	// init context
	ctx = context.TODO()

	// init mongo client
	mongoClient = db.EstablishMongoClient(ctx)

	// get userCollection
	userCollection := db.GetMongoCollection(mongoClient, "users")
	// get feedbackCollection
	feedbackCollection := db.GetMongoCollection(mongoClient, "feedback")
	// get feedbackCollection
	demoRequestCollection := db.GetMongoCollection(mongoClient, "demo-requests")
	
	// init UserDao interface
	ui := dao.UserDaoConstructor(ctx, userCollection)
	// init FeedbackDao interface
	fi := dao.FeedbackDaoConstructor(ctx, feedbackCollection)
	// init DemoRequestDao interface
	dri := dao.DemoRequestConstructor(ctx, demoRequestCollection)

	// init UserController
	uc := controllers.UserControllerConstructor(ui)
	// init FeedbackController
	fc := controllers.FeedbackControllerConstructor(fi)
	// init DemoRequestController
	drc := controllers.DemoRequestControllerConstructor(dri)

	// init UserRouter
	ur = routers.UserRouterConstructor(uc)
	// init FeedbackRouter
	fr = routers.FeedbackRouterConstructor(fc)
	// init DemoRequestRouter
	drr = routers.DemoRequestRouterConstructor(drc)
}

// @dev Root function
func main() {
	// setup CORS
	server.Use(utils.SetupCorsConfig())

	// defer a call to `Disconnect()` after instantiating client
	defer func() {if err := mongoClient.Disconnect(ctx); err != nil {panic(err)}}()
	
	// Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// init basePath
	userBasePath := server.Group("/v2/syns/user")
	tokenBasePath := server.Group("/v2/syns/token")
	feedbackBasePath := server.Group("/v2/syns/feedback/")
	demoRequestBasePath := server.Group("/v2/syns/demo-request/")

	// init Handler
	ur.UserRoutes(userBasePath)
	fr.FeedbackRoutes(feedbackBasePath)
	routers.TokenRoutes(tokenBasePath)
	drr.DemoRequestRoutes(demoRequestBasePath)

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}