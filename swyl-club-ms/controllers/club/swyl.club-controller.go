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
	dao "Swyl/servers/swyl-club-ms/dao/club"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in club-constrollers
type ClubController struct {
	ClubDao dao.ClubDao
}


// @dev Constructor
func ClubControllerConstructor(clubDao dao.ClubDao) *ClubController {
	return &ClubController{
		ClubDao: clubDao,
	}
}


// @notice Method of ClubController struct
// 
// @route `POST/create-club`
// 
// @dev Creates a club
// 
// @param gc *gin.Context
func (cc *ClubController) CreateClub(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"mes": "swylv1.0"})
}


// @notice Method of ClubController struct
// 
// @route `GET/get-club-owned-by/:club-owner`
// 
// @dev Gets a club owned by clubOwner
// 
// @param gc *gin.Context
func (cc *ClubController) GetClubOwnedBy(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"mes": "swylv1.0"})
}


// @notice Method of ClubController struct
// 
// @route `PATCH/update-club`
// 
// @dev Updates a Club
// 
// @param gc *gin.Context
func (cc *ClubController) UpdateClub(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"mes": "swylv1.0"})
}