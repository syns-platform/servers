/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

import (
	"Syns/servers/syns-users-ms/controllers"

	"github.com/gin-gonic/gin"
)

// @dev Declares list of endpoints
func UtilsRouter (rg *gin.RouterGroup) {
	rg.GET("/get-all-syns-tokens/:asset-contract", controllers.GetAllSynsTokens)
	rg.GET("/get-nfts-owned-by/:owner-addr/:asset-contract", controllers.GetSynsTokensOwnedBy)
}