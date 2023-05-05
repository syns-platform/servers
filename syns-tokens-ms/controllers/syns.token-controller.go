/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

import (
	"Syns/servers/syns-tokens-ms/dao"
	"Syns/servers/syns-tokens-ms/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// @notice global var
var validate = validator.New()

// @notice Root struct for other methods in controller
type SynsTokenController struct {
	SynsTokenDao dao.SynsTokenDao
}

// @dev Constructor
func SynsTokenControllerConstructor(synsTokenDao dao.SynsTokenDao) *SynsTokenController{
	return &SynsTokenController {
		SynsTokenDao: synsTokenDao,
	}
}

// @route `POST/mint-new-syns-token`
// 
// @dev handle injecting new syns token to database
// 
// @param gc *gin.Context
func (stc *SynsTokenController) MintNewSynsToken(gc *gin.Context) {
	// declare param
	var param *models.SynsNFT

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{"error": err.Error()}); return;
	}

	// struct validation
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(400, gin.H{"error": err.Error()}); return}

	// invoke TokenDao.MintNewSynsToken() api
	if databaseErr := stc.SynsTokenDao.MintNewSynsToken(param); databaseErr != nil {
		gc.AbortWithStatusJSON(500, gin.H{"error": databaseErr.Error()})
	}

	// response 200
	gc.JSON(200, gin.H{"error": nil})
}