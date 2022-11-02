/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

import (
	controllers "Swyl/servers/swyl-club-ms/controllers/subscription"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in sub-router
type SubRouter struct {
   SubController *controllers.SubController
}

// @dev Constructor
func SubRouterConstructor(subController *controllers.SubController) *SubRouter {
   return &SubRouter{
      SubController: subController,
   }
}

// @notice Method of SubRouter struct
// 
// @dev Declares a list of endpoints
func (sr *SubRouter) SubRoutes(rg *gin.RouterGroup) {
   rg.POST("/subscribe", sr.SubController.Subscribe)
   rg.GET("/get-sub-at/:sub_id", sr.SubController.GetSubscriptionAt)
   rg.GET("/get-all-subs-at", sr.SubController.GetAllSubsAt)
   rg.PATCH("/update-sub-status", sr.SubController.UpdateSubStatus)
   rg.DELETE("/unsubscribe", sr.SubController.Unsubscribe)
}