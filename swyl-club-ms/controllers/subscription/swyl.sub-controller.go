/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

// @import
import (
	dao "Swyl/servers/swyl-club-ms/dao/subscription"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in sub-controller
type SubController struct {
	SubDao dao.SubDao
}

// @dev Constructor
func SubControllerConstructor(subDao dao.SubDao) *SubController {
	return &SubController{
		SubDao: subDao,
	}
}


// @notice Method of SubController struct
// 
// @route `POST/subscribe`
// 
// @dev Lets a user subscribe to a tier
func (si *SubController) Subscribe(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg": "swyl-v1"})
}


// @notice Method of SubController struct
//
// @route `GET/get-sub-at/:sub_id`
//
// @dev Gets a subscription using subscription_id
func (si *SubController) GetSubscriptionAt(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg": "swyl-v1"})
}


// @notice Method of SubController struct
//
// @route `GET/get-all-subs-at?tier_id=&club_owner=`
//
// @dev Gets all subscriptions onwed at tier_id and by club_owner
func (si *SubController) GetAllSubsAt(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg": "swyl-v1"})
}


// @notice Method of SubController struct
//
// @route `PATCH/update-sub-status`
//
// @dev Updates a subscription status
func (si *SubController) UpdateSubStatus(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg": "swyl-v1"})
}


// @notice Method of SubController struct
//
// @route `DELETE/unsubscribe`
//
// @dev Lets a subscriber unsubscribe a tier 
func (si *SubController) Unsubscribe(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg": "swyl-v1"})
}