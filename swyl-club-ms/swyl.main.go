/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package main

// @import
import (
	clubControllers "Swyl/servers/swyl-club-ms/controllers/club"
	subControllers "Swyl/servers/swyl-club-ms/controllers/subscription"
	tierControllers "Swyl/servers/swyl-club-ms/controllers/tier"
	clubDao "Swyl/servers/swyl-club-ms/dao/club"
	subDao "Swyl/servers/swyl-club-ms/dao/subscription"
	tierDao "Swyl/servers/swyl-club-ms/dao/tier"
	"Swyl/servers/swyl-club-ms/db"
	clubRouters "Swyl/servers/swyl-club-ms/routers/club"
	subRouters "Swyl/servers/swyl-club-ms/routers/subscription"
	tierRouters "Swyl/servers/swyl-club-ms/routers/tier"
	"Swyl/servers/swyl-club-ms/utils"
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
	clubCollection	*mongo.Collection
	tierCollection	*mongo.Collection
	subCollection	*mongo.Collection
	cr				*clubRouters.ClubRouter
	tr				*tierRouters.TierRouter
	sr				*subRouters.SubRouter
)


// @dev Runs before main()
func init() {
   	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}

	// init context
	ctx = context.TODO()

	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

	// init mongo client
	mongoClient = db.EstablishMongoClient(ctx)

	// ############ init club router ############
	clubCollection = db.GetMongoCollection(mongoClient, "clubs") // get clubCollection
	ci := clubDao.ClubDaoConstructor(ctx, clubCollection) // init ClubDao interface
	cc := clubControllers.ClubControllerConstructor(ci) // init ClubController
	cr = clubRouters.ClubRouterConstructor(cc) // init ClubRouter
   

	// ############ init tier router ############
	tierCollection = db.GetMongoCollection(mongoClient, "tiers") // get tier Collections
	ti := tierDao.TierDaoConstructor(ctx, tierCollection) // init TierDao interface
	tc := tierControllers.TierControllerConstructor(ti) // init TierController
	tr = tierRouters.TierRouterConstructor(tc)
   	

	// ############ init sub router ############
	subCollection = db.GetMongoCollection(mongoClient, "subs")
	si := subDao.SubDaoConstructor(ctx, subCollection)
	sc := subControllers.SubControllerConstructor(si)
	sr = subRouters.SubRouterConstructor(sc)
}


// @dev Root function
func main() {
	// defer a call to `Disconnect()` after instantiating client
	defer func() {if err := mongoClient.Disconnect(ctx); err != nil {panic(err)}}()
	
   	// Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// init basePath
	clubBasePath := server.Group("/v1/swyl/club") // club bash path
	tierBashPath := server.Group("/v1/swyl/tier") // tier base path
	subBashPath := server.Group("/v1/swyl/sub") // subs base path


	// init Handler
	cr.ClubRoutes(clubBasePath) // club routes
	tr.TierRoutes(tierBashPath) // tier routes
	sr.SubRoutes(subBashPath) // sub routes

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}