/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

// @import
import (
	"Swyl/servers/swyl-users-ms/controllers"
	"Swyl/servers/swyl-users-ms/middleware"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in router
type UserRouter struct {
   UserController *controllers.UserController
}

// @dev Constructor
func UserRouterConstructor(userController *controllers.UserController) *UserRouter {
   return &UserRouter{
      UserController: userController,
   }
}


// @notice Method of UserController struct
// 
// @dev Declares a list of endpoints
func (ur * UserRouter) UserRoutes(rg *gin.RouterGroup) {
   rg.GET("/get-user-by", ur.UserController.GetUserBy)
   rg.GET("/get-all-user", ur.UserController.GetAllUsers)
   rg.GET("/get-user-at/:wallet-address", ur.UserController.GetUserAt)
   rg.GET("/get-server-status", ur.UserController.SwylServerHealthCheck)
   rg.POST("/connect", middleware.Authenticate(), ur.UserController.Connect)
   rg.POST("/claim-page", middleware.Authenticate(), ur.UserController.ClaimPage)
   rg.PATCH("/update-user", middleware.Authenticate(), ur.UserController.UpdateUser)
   rg.GET("/check-username-availability", ur.UserController.CheckUsernameAvailability)
   rg.DELETE("/deactivate-user-at/:wallet-address", middleware.Authenticate(), ur.UserController.DeactivateUserAt)
}