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
	"Syns/servers/syns-users-ms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in controller
type DemoRequestController struct {
	DemoRequestDao dao.DemoRequestDao
}

// @dev Constructor
func DemoRequestControllerConstructor(demoRequestDao dao.DemoRequestDao) *DemoRequestController{
	return &DemoRequestController {
		DemoRequestDao: demoRequestDao,
	}
}

// @route `POST/submit-demo-request`
//
// @dev handle demo request submission from users
//
// @param gc *gin.Context
func (fc *DemoRequestController) SubmitDemoRequest(gc *gin.Context) {

	// declare param
	var param *models.DemoRequest

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": nil, "error":  err.Error()}); return;
	}

	// struct validation
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": nil, "error": err.Error()}); return;}

	// invoke DemoRequestDao.SubmitDemoRequest() api
	if err := fc.DemoRequestDao.SubmitDemoRequest(&param.Email, &param.Name, &param.Question); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": nil, "error": err.Error()}); return;}

	// alert new demo request submitted
	utils.EmailNotification("DEMO_REQUEST", param);

	// http response
	gc.JSON(200,  gin.H{"msg": "Demo request successfully submitted", "error": nil})
}