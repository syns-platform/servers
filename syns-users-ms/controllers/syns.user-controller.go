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
	"Syns/servers/syns-users-ms/dao"
	"Syns/servers/syns-users-ms/models"
	"Syns/servers/syns-users-ms/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @notice global var
var validate = validator.New()

// @notice Root struct for other methods in controller
type UserController struct {
	UserDao dao.UserDao
}

// @dev Constructor
func UserControllerConstructor(userDao dao.UserDao) *UserController{
	return &UserController {
		UserDao: userDao,
	}
}

// @notice Method of UserController struct
// 
// @route `GET/get-server-status`
// 
// @dev Health Check Path - monitors the server and for zero downtime deploys.
// 
// @param gc *gin.Context
func (uc *UserController) SynsServerHealthCheck(gc *gin.Context) {
	utils.ReportVisitor(gc.ClientIP())
	gc.JSON(200, gin.H{"syns-user": "Status OK"})
}


// @notice Method of UserController struct
// 
// @route `POST/connect`
// 
// @dev Connects to an account stored in the internal database using wallet address. Create a new account on first connect.
// 
// @param gc *gin.Context
func (uc *UserController) Connect(gc *gin.Context){
	// get verifiedUserWalletAddress from auth middleware
	verifiedUserWalletAddress := gc.GetString("verifiedUserWalletAddress")

	// invoke UserDao.Connect() api
	foundUser, err := uc.UserDao.Connect(&verifiedUserWalletAddress)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(200, &foundUser)
}


// @notice Method of UserController struct
// 
// @route `POST/claim-page`
// 
// @dev Lets a wallet owner claim a Syns page with passed-in username
// 
// @param gc *gin.Context
func (uc *UserController) ClaimPage(gc *gin.Context){
	// get verifiedUserWalletAddress from auth middleware
	verifiedUserWalletAddress := gc.GetString("verifiedUserWalletAddress")

	// declare param
	var param *models.User

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}

	// test param.wallet_address to match ETH Crypto wallet address convention
	walletMatched, err := utils.TestEthAddress(param.Wallet_address)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !walletMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - wallet_address is not an ETH crypto wallet address"}); return;}

	// test if param.wallet_address matched verifiedUserWalletAddress
	if signerMatched := strings.EqualFold(verifiedUserWalletAddress, *param.Wallet_address); !signerMatched {
		gc.AbortWithStatusJSON(401, gin.H{"error": "!SIGNER - request.body.wallet_address do not match verified signer"}); return;
	}

	// strip all special characters off from param.username
	validUsername, err := utils.SanitizeUsername(param.Username)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX"}); return;}

	// update param.Username to new validUsername
	param.Username = validUsername;

	// extra vaidatation on struct param
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}


	// invoke UserDao.ClaimPage() api
	_, claimErr := uc.UserDao.ClaimPage(param)
	if claimErr != nil && claimErr.Error() == "!USERNAME_TAKEN" {gc.AbortWithStatusJSON(400, gin.H{"error": "!USERNAME - username has already been taken"}); return;}
	if claimErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(200, gin.H{"user": "Successfully claim a page"})
}


// @notice Method of UserController struct
// 
// @route `GET/check-username-availability?username=`
// 
// @dev Checks if a username has been taken
// 
// @param gc *gin.Context
func (uc *UserController) CheckUsernameAvailability (gc *gin.Context) {
	// get username from param query
	username := gc.Query("username")

	// strip all special characters off from param.username
	validUsername, err := utils.SanitizeUsername(&username)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX"}); return;}

	// invoke UserDao.CheckUsernameAvailability()
	bool, err := uc.UserDao.CheckUsernameAvailability(validUsername)
	if err != nil {gc.AbortWithStatusJSON(500, gin.H{"error": err.Error()}); return;}

	gc.JSON(200, bool)
}


// @notice Method of UserController struct
// 
// @route `GET/get-user-at`
// 
// @dev Respond with a user at param `wallet-address`
// 
// @param gc *gin.Context
func (uc *UserController) GetUserAt(gc *gin.Context) {
	// declare param
	param := gc.Param("wallet-address")

	// test param.wallet_address to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(&param)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - Wallet_address is not an ETH crypto wallet address"}); return;}

	// invoke UserDao.GetUserAt
	foundUser, err := uc.UserDao.GetUserAt(&param)
	if err != nil {
		// return 200 with error message says no documents found
		if (err.Error() == "mongo: no documents in result") {gc.AbortWithStatusJSON(200, gin.H{"error": err.Error()}); return;}

		// return 500 if other errors
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;
	}

	// http response
	gc.JSON(http.StatusOK, &foundUser)
}

// @notice Method of UserController struct
// 
// @route `GET/get-user-by?username=`
// 
// @dev Respond with a user by param `username`
// 
// @param gc *gin.Context
func (uc *UserController) GetUserBy(gc *gin.Context) {
	// declare param
	username := gc.Query("username")

	// strip all special characters off from param.username
	validUsername, err := utils.SanitizeUsername(&username)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX"}); return;}

	// invoke UserDao.CheckUsernameAvailability()
	foundUser, err := uc.UserDao.GetUserBy(validUsername)
	if err != nil {
		// return 200 with error message says no documents found
		if (err.Error() == "mongo: no documents in result") {gc.AbortWithStatusJSON(200, gin.H{"error": err.Error()}); return;}

		// return 500 if other errors
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;
	}

	// http response
	gc.JSON(http.StatusOK, &foundUser)
}


// @notice Method of UserController struct
// 
// @route `GET/get-all-users`
// @NOTE might not be necessary
// 
// @dev Reponsd with a slice of models.User
// 
// @param gc *gin.Context
func (uc *UserController) GetAllUsers(gc *gin.Context) {
	// invoke UserDao.GetAllUsers
	users, err := uc.UserDao.GetAllUsers()
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, &users)
}


// @notice Method of UserController struct
// 
// @route `PATCH/update-user`
// 
// @dev Allows a user to update their account
// 
// @param gc *gin.Context
func (uc *UserController) UpdateUser(gc *gin.Context) {
	// get verifiedUserWalletAddress from auth middleware
	// verifiedUserWalletAddress := gc.GetString("verifiedUserWalletAddress")
	// declare param as models.User
	var param models.User

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}

	// test param.wallet_address to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(param.Wallet_address)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - wallet_address is not an ETH crypto wallet address"}); return;}

	// extra validation on struct models.User
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, err.Error()); return;}

	// invoke UserDao.UpdateUser()
	if err := uc.UserDao.UpdateUser(&param); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, err.Error()); return;}

	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "User succesfully updated"})
}


// @notice Method of UserController struct
// 
// @route `DELETE/deactivate-user-at/:wallet-address`
// 
// @param gc *gin.Context
// 
// @return error
func (uc *UserController) DeactivateUserAt(gc *gin.Context) {
	// declare param
	param := gc.Param("wallet-address")

	// test param.wallet_address to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(&param)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - Wallet_address is not an ETH crypto wallet address"}); return;}

	// invokde UserDao.DeactivateUserAt
	if err := uc.UserDao.DeactivateUserAt(&param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, err.Error()); return;}

	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "User succesfully deactivated"})
}