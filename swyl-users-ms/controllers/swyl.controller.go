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
)

// @notice Root struct for other methods in controllers
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
// @dev
// 
// @param gc *gin.Context
// 
// @return error
func (uc *UserController) Connect(user *models.User ){}


// @notice Method of UserController struct
// 
// @dev
// 
// @param gc *gin.Context
// 
// @return *models.User
// 
// @return 
func (uc *UserController) GetUserAt(walletAddress *string) {}


// @notice Method of UserController struct
// 
// @dev
// 
// @NOTE might not be necessary
// 
// @return []*models.User
func (uc *UserController) GetAllUsers() {}


// @notice Method of UserController struct
// 
// @dev
// 
// @param gc *gin.Context
// 
// @return `r
func (uc *UserController) UpdateUser(*models.User) {}


// @notice Method of UserController struct
// 
// @dev
// 
// @param gc *gin.Context
// 
// @return error
func (uc *UserController) DeleteUserAt(walletAddress *string) {}