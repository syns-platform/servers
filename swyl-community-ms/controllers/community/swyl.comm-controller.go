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
	dao "Swyl/servers/swyl-community-ms/dao/community"

	"github.com/gin-gonic/gin"
)

// @notice global var
// var validate = validator.New()

// @notice Root struct for other methods in comm-controller
type CommController struct {
   CommDao dao.CommDao
}

// @dev Constructor 
func CommControllerConstructor(commDao dao.CommDao) *CommController {
   return &CommController{
      CommDao: commDao,
   }
}


// @notice Method of CommController struct
// 
// @route `POST/create-community`
// 
// @dev Lets a Swyl user create a community
// 
// @NOTE Should be fired off when #user/connect api is called
func (cc *CommController) CreateComm(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of CommController struct
// 
// @route `GET/get-community-owned-by`
// 
// @dev Gets a Comm owned by commOwner
func (cc *CommController) GetCommOwnedBy(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of CommController struct
// 
// @route `GET/get-all-comms`
// 
// @dev Gets all Comm has ever created
func (cc *CommController) GetAllComms(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of CommController struct
// 
// @route `PATCH/update-comm-owned-by`
// 
// @dev Updates Comm's total_followers || Comm's total_posts
func (cc *CommController) UpdateCommOwnedBy(gc *gin.Context){
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}