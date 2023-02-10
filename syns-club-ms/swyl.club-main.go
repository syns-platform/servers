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
	clubControllers "Syns/servers/syns-club-ms/controllers/club"
	subControllers "Syns/servers/syns-club-ms/controllers/subscription"
	tierControllers "Syns/servers/syns-club-ms/controllers/tier"
	clubDao "Syns/servers/syns-club-ms/dao/club"
	subDao "Syns/servers/syns-club-ms/dao/subscription"
	tierDao "Syns/servers/syns-club-ms/dao/tier"
	"Syns/servers/syns-club-ms/db"
	clubRouters "Syns/servers/syns-club-ms/routers/club"
	subRouters "Syns/servers/syns-club-ms/routers/subscription"
	tierRouters "Syns/servers/syns-club-ms/routers/tier"
	"Syns/servers/syns-club-ms/utils"
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
	clubCollection		*mongo.Collection
	tierCollection		*mongo.Collection
	subCollection		*mongo.Collection
	cr			*clubRouters.ClubRouter
	tr			*tierRouters.TierRouter
	sr			*subRouters.SubRouter
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
	clubCollection = db.GetMongoCollection(mongoClient, "clubs") // get clubs Collections
	ci := clubDao.ClubDaoConstructor(ctx, clubCollection) // init ClubDao interface
	cc := clubControllers.ClubControllerConstructor(ci) // init ClubController
	cr = clubRouters.ClubRouterConstructor(cc) // init ClubRouter
   

	// ############ init tier router ############
	tierCollection = db.GetMongoCollection(mongoClient, "tiers") // get tiers Collections
	ti := tierDao.TierDaoConstructor(ctx, tierCollection) // init TierDao interface
	tc := tierControllers.TierControllerConstructor(ti) // init TierController
	tr = tierRouters.TierRouterConstructor(tc) // init TierRouter
   	

	// ############ init sub router ############
	subCollection = db.GetMongoCollection(mongoClient, "subs") // get subs Collections
	si := subDao.SubDaoConstructor(ctx, subCollection) // init SubDao interface
	sc := subControllers.SubControllerConstructor(si) // init SubController
	sr = subRouters.SubRouterConstructor(sc) // init SubRouter
}


// @dev Root function
func main() {
	// defer a call to `Disconnect()` after instantiating client
	defer func() {if err := mongoClient.Disconnect(ctx); err != nil {panic(err)}}()
	
   	// Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// init basePath
	clubBasePath := server.Group("/v2/syns/club") // club bash path
	tierBashPath := server.Group("/v2/syns/tier") // tier base path
	subBashPath := server.Group("/v2/syns/sub") // subs base path


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