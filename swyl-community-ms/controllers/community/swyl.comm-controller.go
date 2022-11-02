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
   gc.JSON(200, gin.H{"msg": "Commnity successfully created"})
}


// @notice Method of CommController struct
// 
// @route `GET/get-community-owned-by`
// 
// @dev Gets a Comm owned by commOwner
func (cc *CommController) GetCommOwnedBy(gc *gin.Context) {
   // Handle commOwner param
   commOwner := gc.Param("comm_owner")

   // test commOwner to match ETH Crypto wallet address convention
	ownerMatched, ownerErr := utils.TestEthAddress(&commOwner)
	if ownerErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test commOwner against regex"}); return;}
	if !ownerMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - commOwner is not an ETH crypto wallet address"}); return;}

   // invoke CommDao.GetCommOwnedBy
   comm, err := cc.CommDao.GetCommOwnedBy(&commOwner)
   if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

   // http response
   gc.JSON(200, gin.H{"msg": comm})
}


// @notice Method of CommController struct
// 
// @route `GET/get-all-comms`
// 
// @dev Gets all Comm has ever created
func (cc *CommController) GetAllComms(gc *gin.Context) {
   // invode CommDao.GetAllClubs
   comms, err := cc.CommDao.GetAllComms()
   if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;}

   // http response
   gc.JSON(200, gin.H{"msg": comms})
}


// @notice Method of CommController struct
// 
// @route `PATCH/update-comm-owned-by/bio`
// 
// @dev Updates Comm's bio
func (cc *CommController) UpdateCommBioOwnedBy(gc *gin.Context){
   // declare param holder
   param := &models.Community{}

   // bind json post data to param holder
   if err := gc.ShouldBindJSON(&param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // validate param struct
   if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // test param.Community_owner to match ETH Crypto wallet address convention
	ownerMatched, ownerErr := utils.TestEthAddress(param.Community_owner)
	if ownerErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test community_owner against regex"}); return;}
	if !ownerMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - community_owner is not an ETH crypto wallet address"}); return;}

   // invoke CommDao.UpdateCommBioOwnedBy
   if err := cc.CommDao.UpdateCommBioOwnedBy(param.Community_owner, param.Bio); err != nil {
      gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;
   }

   // http response
   gc.JSON(200, gin.H{"msg": "Community Bio successfully updated"})
}


// @notice Method of CommController struct 
// 
// @route `PATCH/update-comm-owned-by/total`
// 
// @dev Updates Comm's total_followers || Comm's total_posts
func (cc *CommController) UpdateCommTotalOwnedBy(gc *gin.Context) {
   // set up param struct
   type Param struct {
      Community_owner 		*string 		`json:"community_owner" bson:"community_owner" validate:"required,len=42,alphanum"`
      Follower_context	 	int16				`json:"follower_context" bson:"follower_context" validate:"omitempty,oneof=-1 0 1"`
      Post_context	 	int16				`json:"post_context" bson:"post_context" validate:"omitempty,oneof=-1 0 1"`
   }

   // delcare param holder
   param := &Param{}

   // bind json post data to param holder
   if err := gc.ShouldBindJSON(&param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // validate param
   if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}

   // test param.Community_owner to match ETH Crypto wallet address convention
	ownerMatched, ownerErr := utils.TestEthAddress(param.Community_owner)
	if ownerErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test community_owner against regex"}); return;}
	if !ownerMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - community_owner is not an ETH crypto wallet address"}); return;}   

   // invokde CommDao.UpdateCommTotalOwnedBy
   if err := cc.CommDao.UpdateCommTotalOwnedBy(param.Community_owner, param.Follower_context, param.Post_context); err != nil {
      gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return;
   }

   // http response
   gc.JSON(200, gin.H{"msg": "Community successfully updated"})
}


// @notice Method of CommController struct 
// 
// @route `POST/follow`
// 
// @dev Lets a Swyl user start following a community
func (cc *CommController) Follow(gc *gin.Context) {gc.JSON(200, "swyl-v1")}


// @notice Method of CommController struct 
// 
// @route `GET/get-follower-at/:follower_id`
// 
// @dev Gets a Swyl follower at followerId
func (cc *CommController) GetFollowerAt(gc *gin.Context) {gc.JSON(200, "swyl-v1")}


// @notice Method of CommController struct 
// 
// @route `GET/get-all-followers/:community_owner`
// 
// @dev Gets all Swyl followers in a community own by commOwner
func (cc *CommController) GetAllFollowersInCommOwnedBy(gc *gin.Context) {gc.JSON(200, "swyl-v1")}


// @notice Method of CommController struct 
// 
// @route `DELETE/unfollow/:follower_id`
// 
// @dev Lets a Swyl user at followerId unfollows a community
func (cc *CommController) Unfollow(gc *gin.Context) {gc.JSON(200, "swyl-v1")}
