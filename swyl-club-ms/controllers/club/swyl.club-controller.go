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
	"Swyl/servers/swyl-club-ms/models"
	"Swyl/servers/swyl-club-ms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
	// declare params 
	var params *models.Club

	// bind json post data to params
	if err := gc.ShouldBindJSON(&params); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}

	// extra validation on struct models.Club
	validate := validator.New()
	if err := validate.Struct(params); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, err.Error()); return;}

	// test params.Club_owner to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(params.Club_owner)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - wallet_address is not an ETH crypto wallet address"}); return;}

	// invoke ClubDaoImpl.CreateClub() api
	if err := cc.ClubDao.CreateClub(params.Club_owner); err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;
	}

	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "Club sucessfully created"})
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