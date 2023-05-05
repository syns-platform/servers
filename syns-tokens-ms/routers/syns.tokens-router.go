/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

import (
	"Syns/servers/syns-tokens-ms/controllers"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in router
type SynsTokenkRouter struct {
	SynsTokenController *controllers.SynsTokenController
 }
 
 // @dev Constructor
 func SynsTokenRouterConstructor(synsTokenController *controllers.SynsTokenController) *SynsTokenkRouter {
	return &SynsTokenkRouter{
	   SynsTokenController: synsTokenController,
	}
 }

// @dev Declares list of endpoints
func (str *SynsTokenkRouter) TokenRouter (rg *gin.RouterGroup) {
	rg.GET("/get-all-syns-tokens/utils/:asset-contract", controllers.GetAllSynsTokens)
	rg.GET("/get-nfts-owned-by/utils/:owner-addr/:asset-contract", controllers.GetSynsTokensOwnedBy)
	rg.GET("/get-token-metadata/utils/:asset-contract/:token-id/:token-type", controllers.GetTokenMetadata)
	rg.GET("/get-owners-for-token/utils/:asset-contract/:token-id", controllers.GetOwnersForToken)
}