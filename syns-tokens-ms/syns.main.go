/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package main

import (
	"Syns/servers/syns-tokens-ms/controllers"
	"Syns/servers/syns-tokens-ms/dao"
	"Syns/servers/syns-tokens-ms/db"
	"Syns/servers/syns-tokens-ms/onchain"
	"Syns/servers/syns-tokens-ms/routers"
	"Syns/servers/syns-tokens-ms/utils"
	"context"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice global variables
var (
	server			*gin.Engine
	ctx			context.Context
	mongoClient		*mongo.Client
	syns721TokenCollection		*mongo.Collection
	tr 			*routers.SynsTokenkRouter
	to			*onchain.SynsTokenOnchain
	client		*ethclient.Client
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

	// get synsTokenCollection
	syns721TokenCollection = db.GetMongoCollection(mongoClient, "syns-721-tokens")

	// init Syns721TokenDao interface
	ti := dao.Syns721TokenDaoConstructor(ctx, syns721TokenCollection)
	
	// init SynsTokenController
	tc := controllers.Syns721TokenControllerConstructor(ti)

	// init Syns721TokenOnchainConstructor
	to = onchain.Syns721TokenOnchainConstructor(ti)

	// init SynsTokenRouter
	tr = routers.SynsTokenRouterConstructor(tc)

	// Connect to the Ethereum client
	client, _ = ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/"+os.Getenv("ALCHEMY_API_KEY"))
	
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
	tokenBasePath := server.Group("/v2/syns/nfts/721/")

	// init Handler
	tr.Syns721TokenRouter(tokenBasePath)

	// @notice listen to newTokenMintedTo on-chain event to add new token to database
	to.HandleNewSyns721TokenMinted(client, "contract-artifacts/SynsERC721.json")

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}