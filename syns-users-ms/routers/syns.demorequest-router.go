/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

// @imports
import (
	"Syns/servers/syns-users-ms/controllers"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in router
type DemoRequestRouter struct {
   DemoRequestController *controllers.DemoRequestController
}

// @dev Constructor
func DemoRequestRouterConstructor(demoRequestRouter *controllers.DemoRequestController) *DemoRequestRouter {
   return &DemoRequestRouter{
      DemoRequestController: demoRequestRouter,
   }
}


// @notice Method of DemoRequestRouter struct
// 
// @dev Declares a list of endpoints
func (fr * DemoRequestRouter) DemoRequestRoutes(rg *gin.RouterGroup) {
   rg.POST("/submit-demo-request", fr.DemoRequestController.SubmitDemoRequest)
}