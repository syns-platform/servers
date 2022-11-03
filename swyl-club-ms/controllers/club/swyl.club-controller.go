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

//@notice global vars
var validate = validator.New()

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
	var param *models.Club

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}

	// validate struct models.Club
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, err.Error()); return;}

	// extra validatation on param.Club_owner to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(param.Club_owner)
	if err != nil{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - club_owner is not an ETH crypto wallet address"}); return;}

	// invoke ClubDao.CreateClub() api
	if err := cc.ClubDao.CreateClub(param.Club_owner, param.Created_at); err != nil {
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
	// decalre param 
	param := gc.Param("club-owner")

	// test params.club_owner to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(&param)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - Wallet_address is not an ETH crypto wallet address"}); return;}

	// invoke ClubDao.GetClubOwnedBy()
	club, err := cc.ClubDao.GetClubOwnedBy(&param)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, &club)
}

// @notice Method of ClubController struct
// 
// @route `GET/get-all-club`
// 
// @dev Gets all club
// 
// @param gc *gin.Context
func (cc *ClubController) GetAllClubs(gc *gin.Context) {
	// invoke ClubDao.GetAllUsers
	clubs, err := cc.ClubDao.GetAllClubs()
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, &clubs)
}



// @notice Method of ClubController struct
// 
// @route `PATCH/update-club`
// 
// @dev Updates a Club
// 
// @param gc *gin.Context
func (cc *ClubController) UpdateClub(gc *gin.Context) {
	// Declare param struct
	type Param struct {
		Club_owner *string `json:"club_owner" validate:"required,len=42,alphanum"`
		Context	*uint16 `json:"context" validate:"eq=1|eq=0"`
	}
	
	// Declare param placeholder
	param := Param{}

	// bind JSON post data to params
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"erorr": err.Error()}); return;
	}

	// extra validation on struct Param
	if err1 := validate.Struct(param); err1 != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err1.Error()}); return;}

	// test params.wallet_address to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(param.Club_owner)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - wallet_address is not an ETH crypto wallet address"}); return;}

	// invoke ClubDao.UpdateClub()
	if err := cc.ClubDao.UpdateClub(param.Club_owner, param.Context); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "Club sucessfully updated"})
}