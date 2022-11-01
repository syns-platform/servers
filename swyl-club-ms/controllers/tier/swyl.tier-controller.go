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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
      gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"erorr": err.Error()}); return;
   }

   // http reponse
   gc.JSON(200, gin.H{"msg": "Tier is sucessfully created"})
}


// @notice Method of TierController struct
// 
// @route `@GET/get-tier-at/:tier_id`
// 
// @dev Gets a Tier at tierId and clubOwner
// 
// @param gc *gin.Context
func (tc *TierController) GetTierAt(gc *gin.Context) {
   // get tierId from param
   tierId := gc.Param("tier_id")

   // sanitize tierId
   matched, err := regexp.MatchString(`^[a-zA-Z0-9]{24}$`, tierId)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - clubId is not valid"}); return;}

   
   // invoke TierDao.GetTierAt
   tier, err := tc.TierDao.GetTierAt(&tierId)
   if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}
   
   // http reponse
   gc.JSON(200, tier)
}


// @notice Method of TierController struct
// 
// @route `GET//get-all-tiers-owned-by/:club_owner`
// 
// @dev Gets all tiers owned by clubOwner
// 
// @param gc *gin.Context
func (tc *TierController) GetAllTiersOwnedBy(gc *gin.Context) {
   // Handle param
   clubOwner := gc.Param("club_owner")

   // validate clubOwner to match ETH Crypto Wallet address convention
   matched, err := utils.TestEthAddress(&clubOwner)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - club_owner is not an ETH crypto wallet address"}); return;}

   // invode TierDao.GetAllTiersOwnedBy
   tiers, err := tc.TierDao.GetAllTiersOwnedBy(&clubOwner)
   if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"erorr": err.Error()}); return;}

   gc.JSON(200, tiers)
}


// @notice Method of TierController struct
// 
// @route `PATCH/update-tier`
// 
// @dev Lets a clubOwner update a tier
// 
// @param gc *gin.Context
func (tc *TierController) UpdateTier(gc *gin.Context) {
   // prepare param holder
   param := &models.Tier{}

   // bind json post data to param holder
   if err := gc.ShouldBindJSON(param); err != nil {
      gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
   }


   // sanitizing param.Tier_id
   tierId := primitive.ObjectID.Hex(param.Tier_ID)
   matched, err := regexp.MatchString(`^[a-zA-Z0-9]{24}$`, tierId)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - clubId is not valid"}); return;}

   // validate struct
   if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // invoke TierDao.UpdateTier
   if err := tc.TierDao.UpdateTier(param); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

   gc.JSON(200, gin.H{"msg": "Club sucessfully updated"})
}


// @notice Method of TierController struct
// 
// @route `DELETE/delete-tier-at/:tier_id`
// 
// @param gc *gin.Context
func (tc *TierController) DeleteTier(gc *gin.Context){
   // Handle param
   param := gc.Param("tier_id")

   // sanitizing param
   matched, err := regexp.MatchString(`^[a-zA-Z0-9]{24}$`, param)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - clubId is not valid"}); return;}

   // invoke TierDao.DeleteTier
   if err := tc.TierDao.DeleteTier(&param); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

   gc.JSON(200, gin.H{"msg": "Tier sucessfully deleted"})
}