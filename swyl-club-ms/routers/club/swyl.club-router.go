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
	controllers "Swyl/servers/swyl-club-ms/controllers/club"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in router
type ClubRouter struct {
	ClubController *controllers.ClubController
}


// @dev Constructor
func ClubRouterConstructor(clubController *controllers.ClubController) *ClubRouter {
	return &ClubRouter{
		ClubController: clubController,
	}
}

// @notice Method of ClubRouter struct
// 
// @dev Decalres a list of endpoints
func (cr *ClubRouter) ClubRoutes(rg *gin.RouterGroup) {
	rg.POST("/create-club", cr.ClubController.CreateClub)
	rg.GET("/get-club-owned-by/:club-owner", cr.ClubController.GetClubOwnedBy)
	rg.GET("/get-all-clubs", cr.ClubController.GetAllClubs)
	rg.PATCH("/update-club", cr.ClubController.UpdateClub)
}