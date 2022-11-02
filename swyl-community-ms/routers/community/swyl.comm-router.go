/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

// @import
import (
	controllers "Swyl/servers/swyl-community-ms/controllers/community"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in comm- router
type CommRouter struct {
	CommController *controllers.CommController
}

// @dev Constructor
func CommRouterConstructor(commController *controllers.CommController) *CommRouter{
	return &CommRouter{
		CommController: commController,
	}
}

// @notice Method of CommRouter struct
// 
// @dev Declares a list of routes
func (cr *CommRouter) CommRoutes(rg *gin.RouterGroup) {
	rg.POST("create-community", cr.CommController.CreateComm)
	rg.GET("get-comminity-owned-by/:comm_owner", cr.CommController.GetCommOwnedBy)
	rg.GET("get-all-comms", cr.CommController.GetAllComms)
	rg.PATCH("update-comm-owned-by/bio", cr.CommController.UpdateCommBioOwnedBy)
	rg.PATCH("update-comm-owned-by/total", cr.CommController.UpdateCommTotalOwnedBy)
	rg.POST("follow", cr.CommController.Follow)
	rg.GET("get-follower-at/:follower_id", cr.CommController.GetFollowerAt)
	rg.GET("get-all-followers/:community_owner", cr.CommController.GetAllFollowersInCommOwnedBy)
	rg.DELETE("unfollow/:follower_id", cr.CommController.Unfollow)
}