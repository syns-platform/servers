/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package main

// @improt
import (
	commControllers "Swyl/servers/swyl-community-ms/controllers/community"
	commDao "Swyl/servers/swyl-community-ms/dao/community"
	"Swyl/servers/swyl-community-ms/db"
	commRouters "Swyl/servers/swyl-community-ms/routers/community"
	"Swyl/servers/swyl-community-ms/utils"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice global variables
var (
	server			*gin.Engine
	ctx 			context.Context
	mongoClient		*mongo.Client
	commCollection	*mongo.Collection
	cr 				*commRouters.CommRouter
)

// @dev Runs before main()
func init() {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}

	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

	// init mongo client
	mongoClient = db.EstablishMongoClient(ctx)

	// ############ init club router ############
	commCollection := db.GetMongoCollection(mongoClient, "communities") //init communities collection
	followerCollection := db.GetMongoCollection(mongoClient, "followers") //init follower collection
	cd := commDao.CommDaoConstructor(ctx, commCollection, followerCollection) // init CommDao
	cc := commControllers.CommControllerConstructor(cd) // init CommController
	cr = commRouters.CommRouterConstructor(cc) // init CommRouter
}


// @dev Root function
func main() {
	// catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// init basePath
	commBasePath := server.Group("v1/swyl/community") // community base path

	// init handlers
	cr.CommRoutes(commBasePath)

	// run server 
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
	
}