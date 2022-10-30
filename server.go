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
	"Swyl/server/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @notice global variables
var (
	server *gin.Engine
)

// @dev Runs before main()
func init() {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}

	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})
}

// @Root function
func main() {
	// Catch all unallowed HTTP methods sent to the server
	server.HandleMethodNotAllowed = true

	// base path
	basePath := server.Group("/v1/swyl")

	// basic first GET request
	basePath.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "pang",
		})
	})

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}