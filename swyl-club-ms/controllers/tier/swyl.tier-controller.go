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
	dao "Swyl/servers/swyl-club-ms/dao/tier"
	"Swyl/servers/swyl-club-ms/models"
	"Swyl/servers/swyl-club-ms/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// global vars
var validate = validator.New()

// @notice Root struct for other methods in controller
type TierController struct {
   TierDao dao.TierDao
}

// @dev Constructor
func TierControllerConstructor(tierDao dao.TierDao) *TierController {
   return &TierController{
      TierDao: tierDao,
   }
}


// @notice Method of TierController struct
// 
// @route `POST/create-tier`
// 
// @dev Lets a club owner create a tier to internal database
// 
// @param gc *gin.Context
func (tc *TierController) CreateTier(gc *gin.Context) {
   // declare params
   param := &models.Tier{}

   // bind json post data to param
   if err := gc.ShouldBindJSON(param); err != nil {
      gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
   }

   // validate struct models.Tier
   if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"erorr": err.Error()}); return;}

   // extra validation on param.club_owner to match ETH Crypto Wallet address convention
   matched, err := utils.TestEthAddress(param.Club_owner)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - club_owner is not an ETH crypto wallet address"}); return;}

   // invoke TierDao.CreateTier() api
   if err := tc.TierDao.CreateTier(param); err !=nil {
      gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"erorr": err.Error()})
   }

   // http reponse
   gc.JSON(200, gin.H{"msg": "Tier is sucessfully created"})
}


// @notice Method of TierController struct
// 
// @route `POST/connect`
// 
// @dev Gets a Tier at tierId and clubOwner
// 
// @param gc *gin.Context
func (tc *TierController) GetTierAt(gc *gin.Context) {}


// @notice Method of TierController struct
// 
// @route `POST/connect`
// 
// @dev Gets all tiers owned by clubOwner
// 
// @param gc *gin.Context
func (tc *TierController) GetAllTiersOwnedBy(gc *gin.Context) {}


// @notice Method of TierController struct
// 
// @route `POST/connect`
// 
// @dev Lets a clubOwner update a tier
// 
// @param gc *gin.Context
func (tc *TierController) UpdateTier(gc *gin.Context) {}


// @notice Method of TierController struct
// 
// @route `POST/connect`
// 
// @param gc *gin.Context
func (tc *TierController) DeleteTier(gc *gin.Context){}