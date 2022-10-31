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
	"Swyl/servers/swyl-users-ms/dao"
	"Swyl/servers/swyl-users-ms/models"
	"Swyl/servers/swyl-users-ms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
// @route `POST/connect`
// 
// @dev Connects to an account stored in the internal database using wallet address. Create a new account on first connect.
// 
// @param gc *gin.Context
func (uc *UserController) Connect(gc *gin.Context){
	// declare user
	var params models.User

	// bind json post data to user
	if err := gc.ShouldBindJSON(&params); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}

	// test params.wallet_address to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(params.Wallet_address)
	if err != nil{gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test wallet_address against regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - wallet_address is not an ETH crypto wallet address"}); return;}

	// validate params
	validate := validator.New()
	if err := validate.Struct(params); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, err.Error()); return;}

	// invoke UserDaoImpl.Connect() api
	foundUser, err := uc.UserDao.Connect(params.Wallet_address)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(200, gin.H{"user": foundUser})
}


// @notice Method of UserController struct
// 
// @route `GET/get-user-at`
// 
// @dev Respond with a user at param `wallet-address`
// 
// @param gc *gin.Context
func (uc *UserController) GetUserAt(gc *gin.Context) {}


// @notice Method of UserController struct
// 
// @route `GET/get-all-users`
// @NOTE might not be necessary
// 
// @dev Reponsd with a slice of models.User
// 
// @param gc *gin.Context
func (uc *UserController) GetAllUsers(gc *gin.Context) {}


// @notice Method of UserController struct
// 
// @route `PATCH/update-user`
// 
// @dev Allows a user to update their account
// 
// @param gc *gin.Context
func (uc *UserController) UpdateUser(gc *gin.Context) {}


// @notice Method of UserController struct
// 
// @route `DELETE/deactivate-user-at/:wallet-address`
// 
// @param gc *gin.Context
// 
// @return error
func (uc *UserController) DeactivateUserAt(gc *gin.Context) {}