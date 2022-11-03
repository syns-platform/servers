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
	dao "Swyl/servers/swyl-community-ms/dao/post"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in post-controller
type PostController struct {
	PostDao *dao.PostDao
}

// @dev Constructor
func PostControllerConstructor(postDao *dao.PostDao) *PostController {
	return &PostController {
		PostDao: postDao,
	}
}

// @notice Method of PostController struct
// 
// @route `POST/create-post`
// 
// @dev Lets a commOwner create a post
// 
// @param gc *gin.Context
func (pi *PostController) CreatePost(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `GET/get-post-at/:post_id`
// 
// @dev Gets a post at postId
// 
// @param gc *gin.Context
func (pi *PostController) GetPostAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `GET/get-all-posts-by/:community_owner`
// 
// @dev Gets all posts created by commOwner
// 
// @param gc *gin.Context
func (pi *PostController) GetPostsBy(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/update-post-content`
// 
// @dev Lets a commOwner update a post - only post.Title and post.Content are allowed
// 
// @param gc *gin.Context
func (pi *PostController) UpdatePost(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/react-post`
// 
// @dev Lets a user react to a post
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to post.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from post.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in post.Reaction
// 
// @param gc *gin.Context
func (pi *PostController) ReactPost(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `POST/comment-post`
// 
// @dev Lets a user comment on a post
// 
// @param gc *gin.Context
func (pi *PostController) Comment(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/update-comment`
// 
// @dev Lets a user to update own comment - only comment.Content is allowed
// 
// @param gc *gin.Context
func (pi *PostController) UpdateComment(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/react-comment`
// 
// @dev Lets a user react to a comment
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to comment.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from comment.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in comment.Reaction
// 
// @param gc *gin.Context
func (pi *PostController) ReactComment(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `POST/create-reply`
// 
// @dev Lets a user reply to a comment
// 
// @param gc *gin.Context
func (pi *PostController) Reply(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/update-reply`
// 
// @dev Lets a user update own reply - only reply.Content is allowed
// 
// @param gc *gin.Context
func (pi *PostController) UpdateReply(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `POST/create-tier`
// 
// @dev Lets a user react a reply
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to comment.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from comment.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in comment.Reaction
// 
// @param gc *gin.Context
func (pi *PostController) ReactReply(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}