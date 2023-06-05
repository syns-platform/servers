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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// @notice global var
var validate = validator.New()

// @notice Root struct for other methods in controller
type SynsTokenController struct {
	Syns721TokenDao dao.Syns721TokenDao
}

// @dev Constructor
func Syns721TokenControllerConstructor(Syns721TokenDao dao.Syns721TokenDao) *SynsTokenController{
	return &SynsTokenController {
		Syns721TokenDao: Syns721TokenDao,
	}
}

// @route `POST/mint-new-syns-token`
// 
// @dev handle injecting new syns token to database
// 
// @param gc *gin.Context
func (stc *SynsTokenController) MintNewSyns721Token(gc *gin.Context) {
	// declare param
	var param *models.Syns721SuperNFT

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(400, gin.H{"error": err.Error()}); return;
	}

	// make sure tokenHash is an empty string for new SynsToken
	if strings.Compare(param.TokenHash, "") != 0 {gc.AbortWithStatusJSON(400, gin.H{"error": "Invalid new Syns Token"}); return}

	// check if assetContract is valid
	if !strings.EqualFold(param.AssetContract, OFFICUAL_SYNS_721_SC_ADDR) && !strings.EqualFold(param.AssetContract, OFFICUAL_SYNS_1155_SC_ADDR) {
		gc.AbortWithStatusJSON(400, gin.H{"error": "Invalid asset contract input"}); return
	}

	// struct validation
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(400, gin.H{"error": err.Error()}); return}

	// invoke TokenDao.MintNewSyns721Token() api
	if databaseErr := stc.Syns721TokenDao.MintNewSyns721Token(param); databaseErr != nil {
		gc.AbortWithStatusJSON(500, gin.H{"error": databaseErr.Error()}); return;
	}

	// response 200
	gc.JSON(200, gin.H{"error": nil})
}

// @route `GET/fetch-all-syns-721-super-tokens`
// 
// @dev handle fetching all tokens from backend
// 
// @param gc *gin.Context
func (stc *SynsTokenController) GetAllSyns721SuperTokens(gc *gin.Context) {
	// invokle GetAllSyns721SuperTokens dao method
	syns721SuperTokens, err := stc.Syns721TokenDao.GetAllSyns721SuperTokens()
	if err != nil {gc.AbortWithStatusJSON(500, gin.H{"syns721SuperTokens": nil, "error": err.Error()}); return}

	// httpo response
	gc.JSON(200, gin.H{"syns721SuperTokens": syns721SuperTokens, "error": nil})
}

// @route `GET/fetch-syns-721-super-tokens-owned-by/:token-owner`
// 
// @dev handle fetching all syns 721 super tokens owned by `token_owner` from backend
// 
// @param gc *gin.Context
func (stc *SynsTokenController) GetAllSyns721SuperTokensOwnedBy(gc *gin.Context) {
	// prepare param
	tokenOwner := gc.Param("token-owner")

	// invokle GetAllSyns721SuperTokensOwnedBy dao method
	syns721SuperTokens, err := stc.Syns721TokenDao.GetAllSyns721SuperTokensOwnedBy(tokenOwner)
	if err != nil {gc.AbortWithStatusJSON(500, gin.H{"syns721SuperTokens": nil, "error": err.Error()}); return}

	// httpo response
	gc.JSON(200, gin.H{"syns721SuperTokens": syns721SuperTokens, "error": nil})
}

// @route `GET//fetch-syns-721-super-token/:asset-contract/:token-id`
// 
// @dev handle fetching specific syns 721 super token by assetContract and tokenId
// 
// @param gc *gin.Context
func (stc *SynsTokenController) GetSyns721SuperTokenMetadata(gc *gin.Context) {
	// prepare param
	assetContract := gc.Param("asset-contract")
	tokenId := gc.Param("token-id")

	// invokle GetSyns721SuperToken dao method
	syns721SuperToken, err := stc.Syns721TokenDao.GetSyns721SuperTokenMetadata(assetContract, tokenId)
	if err != nil {gc.AbortWithStatusJSON(500, gin.H{"syns721SuperTokenMetadata": nil, "error": err.Error()}); return}

	// httpo response
	gc.JSON(200, gin.H{"syns721SuperTokenMetadata": syns721SuperToken, "error": nil})
}