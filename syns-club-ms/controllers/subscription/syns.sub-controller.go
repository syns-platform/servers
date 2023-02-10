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
	dao "Syns/servers/syns-club-ms/dao/subscription"
	"Syns/servers/syns-club-ms/models"
	"Syns/servers/syns-club-ms/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// @notice global var
var validate = validator.New()

// @notice Root struct for other methods in sub-controller
type SubController struct {
	SubDao dao.SubDao
}

// @dev Constructor
func SubControllerConstructor(subDao dao.SubDao) *SubController {
	return &SubController{
		SubDao: subDao,
	}
}


// @notice Method of SubController struct
// 
// @route `POST/subscribe`
// 
// @dev Lets a user subscribe to a tier
func (si *SubController) Subscribe(gc *gin.Context) {
	// Declare param *model
	param := &models.Subscription{}

	// bind json post data to param
	if err := gc.ShouldBindJSON(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

	// validate struct param
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

	// test param.Club_owner to match ETH Crypto wallet address convention
	ownerMatched, ownerErr := utils.TestEthAddress(param.Club_owner)
	if ownerErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test clubOwner against regex"}); return;}
	if !ownerMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - clubOwner is not an ETH crypto wallet address"}); return;}

	// test param.Subscriber to match ETH Crypto wallet address convention
	subsMatched, subsErr := utils.TestEthAddress(param.Subscriber)
	if subsErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test subscriber against regex"}); return;}
	if !subsMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - subscriber is not an ETH crypto wallet address"}); return;}

	// invoke SubDao.Subscribe
	if err := si.SubDao.Subscribe(param); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "New subscription has sucessfully been made"})
}


// @notice Method of SubController struct
//
// @route `GET/get-sub-at/:sub_id`
//
// @dev Gets a subscription using subscription_id
func (si *SubController) GetSubscriptionAt(gc *gin.Context) {
	// Handle param
	subId := gc.Param("sub_id")

	// sanitize subId
	matched, err := regexp.MatchString(`^[a-fA-f0-9]{24}$`, subId)
	if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test subId using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - subId is not valid"}); return;}
 
	// invoke SubDao.GetSubscriptionAt
	subs, err := si.SubDao.GetSubscriptionAt(&subId)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, subs)
}


// @notice Method of SubController struct
//
// @route `GET/get-all-subs-at?tier_id=&club_owner=`
//
// @dev Gets all subscriptions onwed at tier_id and by club_owner
func (si *SubController) GetAllSubsAt(gc *gin.Context) {
	// Handle Param
	tierId := gc.Query("tier_id")
	clubOwner := gc.Query("club_owner")

	// sanitize tierId
	matched, err := regexp.MatchString(`^[a-fA-f0-9]{24}$`, tierId)
	if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test tierId using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - tierId is not valid"}); return;}
 
	// test clubOwner to match ETH Crypto wallet address convention
	subsMatched, subsErr := utils.TestEthAddress(&clubOwner)
	if subsErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test clubOwner against regex"}); return;}
	if !subsMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - clubOwner is not an ETH crypto wallet address"}); return;}

	// invoke SubDao.GetAllSubsAt
	subs, err := si.SubDao.GetAllSubsAt(&tierId, &clubOwner)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, subs)
}


// @notice Method of SubController struct
//
// @route `PATCH/update-sub-status/:sub_id`
//
// @dev Updates a subscription status
func (si *SubController) ToggleSubStatusAt(gc *gin.Context) {
	// handle subId param
	subId := gc.Param("sub_id")

	// sanitize subId
	matched, err := regexp.MatchString(`^[a-fA-f0-9]{24}$`, subId)
	if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test subId using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - subId is not valid"}); return;}
 
	// invokde SubDao.ToggleSubStatusAt
	if err := si.SubDao.ToggleSubStatusAt(&subId); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}
	
	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "Subscription's status sucessfully updated"})
}


// @notice Method of SubController struct
//
// @route `DELETE/unsubscribe`
//
// @dev Lets a subscriber unsubscribe a tier 
func (si *SubController) Unsubscribe(gc *gin.Context) {
	// handle subId param
	subId := gc.Param("sub_id")

	// sanitize subId
	matched, err := regexp.MatchString(`^[a-fA-f0-9]{24}$`, subId)
	if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test subId using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - subId is not valid"}); return;}
 
	// invoke SubDao.Unsubscribe
	if err := si.SubDao.Unsubscribe(&subId); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

	// http response
	gc.JSON(http.StatusOK, gin.H{"msg": "Sucessfully unsubscribe!"})
}