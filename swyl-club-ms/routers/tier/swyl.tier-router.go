/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

import (
	controllers "Swyl/servers/swyl-club-ms/controllers/tier"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in router
type TierRouter struct {
   TierController *controllers.TierController
}

// @Constructor
func TierRouterConstructor(tierController *controllers.TierController) *TierRouter {
   return &TierRouter{
      TierController: tierController,
   }
}


// @notice method of TierRouter struct
// 
// @dev Declare a list of endpoints 
func (tr *TierRouter) TierRoutes(rg *gin.RouterGroup) {
   rg.POST("/create-tier", tr.TierController.CreateTier)
   rg.GET("/get-tier-at/:tier_id", tr.TierController.GetTierAt)
   rg.GET("/get-all-tiers-owned-by/:club_owner", tr.TierController.GetAllTiersOwnedBy)
   rg.PATCH("/update-tier", tr.TierController.UpdateTier)
   rg.DELETE("/delete-tier-at/:tier_id", tr.TierController.DeleteTier)
}