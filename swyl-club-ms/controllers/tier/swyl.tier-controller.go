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

	"github.com/gin-gonic/gin"
)

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
func (tc *TierController) CreateTier(gc *gin.Context) {}


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