/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

import (
	"Syns/servers/syns-users-ms/dao"
	"Syns/servers/syns-users-ms/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in controller
type FeedbackController struct {
	FeedbackDao dao.FeedbackDao
}

// @dev Constructor
func FeedbackControllerConstructor(feedbackDao dao.FeedbackDao) *FeedbackController{
	return &FeedbackController {
		FeedbackDao: feedbackDao,
	}
}

// @route `POST/submit-feedback`
//
// @dev handle feedback submission from users
//
// @param gc *gin.Context
func (fc *FeedbackController) SubmitFeedback(gc *gin.Context) {

	// declare param
	var param *models.Feedback

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":  err.Error()}); return;
	}

	// struct validation
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

	// invoke FeedbackDao.SubmitFeedback() api
	if err := fc.FeedbackDao.SubmitFeedback(&param.Email, &param.Feedback); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(200,  gin.H{"msg": "Feedback successfully submitted"})
}