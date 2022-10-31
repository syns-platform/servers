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
	controllers "Swyl/servers/swyl-club-ms/controllers/club"
	dao "Swyl/servers/swyl-club-ms/dao/club"
	"Swyl/servers/swyl-club-ms/db"
	routers "Swyl/servers/swyl-club-ms/routers/club"
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
	cr				*routers.ClubRouter
)


// @dev Runs before main()
func init() {
   	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}

	// init context
	ctx = context.TODO()

	// init mongo client
	mongoClient = db.EstablishMongoClient(ctx)

	// get clubCollection
	clubCollection = db.GetMongoCollection(mongoClient, "clubs")

	// init ClubDao interface
	ci := dao.ClubDaoConstructor(ctx, clubCollection)

	// init ClubController
	cc := controllers.ClubControllerConstructor(ci)

	// init ClubRouter
	cr = routers.ClubRouterConstructor(cc)
   
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
	clubBasePath := server.Group("/v1/swyl/club")

	// init Handler
	cr.ClubRoutes(clubBasePath)

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}