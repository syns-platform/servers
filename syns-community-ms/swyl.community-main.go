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
	commControllers "Syns/servers/syns-community-ms/controllers/community"
	postControllers "Syns/servers/syns-community-ms/controllers/post"
	commDao "Syns/servers/syns-community-ms/dao/community"
	postDao "Syns/servers/syns-community-ms/dao/post"
	"Syns/servers/syns-community-ms/db"
	commRouters "Syns/servers/syns-community-ms/routers/community"
	postRouters "Syns/servers/syns-community-ms/routers/post"
	"Syns/servers/syns-community-ms/utils"
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

// @notice global variables
var (
	server			*gin.Engine
	ctx 			context.Context
	cr 			*commRouters.CommRouter
	pr 			*postRouters.PostRouter
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
	mongoClient := db.EstablishMongoClient(ctx)

	// ############ init community router ############
	commCollection := db.GetMongoCollection(mongoClient, "communities") //init communities collection
	followerCollection := db.GetMongoCollection(mongoClient, "followers") //init follower collection
	cd := commDao.CommDaoConstructor(ctx, commCollection, followerCollection) // init CommDao
	cc := commControllers.CommControllerConstructor(cd) // init CommController
	cr = commRouters.CommRouterConstructor(cc) // init CommRouter


	// ############ init post router ############
	postCollection := db.GetMongoCollection(mongoClient, "posts") //init posts collection
	commentCollection := db.GetMongoCollection(mongoClient, "comments") //init comments collection
	replyCollection := db.GetMongoCollection(mongoClient, "replies") //init replies collection
	pd := postDao.PostDaoConstructor(ctx, postCollection, commentCollection, replyCollection) // init PostDao
	pc := postControllers.PostControllerConstructor(pd) // init PostController
	pr = postRouters.PostRouterConstructor(pc) //init PostRouter
}


// @dev Root function
func main() {
	// catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// init basePath
	commBasePath := server.Group("v2/syns/community") // community base path
	postBasePath := server.Group("v2/syns/community/post") // post base path

	// init handlers
	cr.CommRoutes(commBasePath) // community router
	pr.PostRoutes(postBasePath) // post routers

	// run server 
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
	
}