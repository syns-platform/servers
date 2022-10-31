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

	"github.com/gin-gonic/gin"
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
func (uc *UserController) Connect(gc *gin.Context){}


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