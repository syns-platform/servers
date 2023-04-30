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
type FeedbackRouter struct {
   FeedbackController *controllers.FeedbackController
}

// @dev Constructor
func FeedbackRouterConstructor(feedbackController *controllers.FeedbackController) *FeedbackRouter {
   return &FeedbackRouter{
      FeedbackController: feedbackController,
   }
}


// @notice Method of FeedbackRouter struct
// 
// @dev Declares a list of endpoints
func (fr * FeedbackRouter) FeedbackRoutes(rg *gin.RouterGroup) {
   rg.POST("/submit-feedback", fr.FeedbackController.SubmitFeedback)
}