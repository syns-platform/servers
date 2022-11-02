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
	dao "Swyl/servers/swyl-community-ms/dao/community"
	"Swyl/servers/swyl-community-ms/models"
	"Swyl/servers/swyl-community-ms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// @notice global var
var validate = validator.New()

// @notice Root struct for other methods in comm-controller
type CommController struct {
   CommDao dao.CommDao
}

// @dev Constructor 
func CommControllerConstructor(commDao dao.CommDao) *CommController {
   return &CommController{
      CommDao: commDao,
   }
}


// @notice Method of CommController struct
// 
// @route `POST/create-community`
// 
// @dev Lets a Swyl user create a community
// 
// @NOTE Should be fired off when #user/connect api is called
func (cc *CommController) CreateComm(gc *gin.Context) {
   // declare param 
   param := &models.Community{}

   // bind json post data to param
   if err := gc.ShouldBindJSON(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // validate struct param
   if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // test param.Community_owner to match ETH Crypto wallet address convention
	ownerMatched, ownerErr := utils.TestEthAddress(param.Community_owner)
	if ownerErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test community_owner against regex"}); return;}
	if !ownerMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - community_owner is not an ETH crypto wallet address"}); return;}

   // invokde CommDao.create
   if err := cc.CommDao.CreateComm(param); err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

   // http response
   gc.JSON(200, gin.H{"msg": "Commnity successfull created"})
}


// @notice Method of CommController struct
// 
// @route `GET/get-community-owned-by`
// 
// @dev Gets a Comm owned by commOwner
func (cc *CommController) GetCommOwnedBy(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of CommController struct
// 
// @route `GET/get-all-comms`
// 
// @dev Gets all Comm has ever created
func (cc *CommController) GetAllComms(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of CommController struct
// 
// @route `PATCH/update-comm-owned-by`
// 
// @dev Updates Comm's total_followers || Comm's total_posts
func (cc *CommController) UpdateCommOwnedBy(gc *gin.Context){
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}