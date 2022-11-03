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
	controllers "Swyl/servers/swyl-community-ms/controllers/post"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other method in post-router
type PostRouter struct {
	PostController controllers.PostController
}

// @dev constructor
func PostRouterConstructor(postController *controllers.PostController) *PostRouter {
	return &PostRouter{
		PostController: *postController,
	}
}

// @notice Method of PostRouter struct
// 
// @dev Declares a list of endpoints
func (pr *PostRouter) PostRoutes(rg *gin.RouterGroup) {
	// ############ post routes ############
	rg.POST("/create-post", pr.PostController.CreatePost)
	rg.GET("/get-post-at/:post_id", pr.PostController.GetPostAt)
	rg.GET("/get-all-posts-by/:community_owner", pr.PostController.GetPostsBy)
	rg.PATCH("/update-post-content", pr.PostController.UpdatePostContent)
	rg.PATCH("/react-post", pr.PostController.ReactPost)
	rg.DELETE("/delete-post-at/:post_id", pr.PostController.DeletePostAt)

	// ############ comment routes ############
	rg.POST("/comment", pr.PostController.Comment)
	rg.GET("/get-comment-at/:comment_id", pr.PostController.GetCommentAt)
	rg.GET("/get-all-comments-at/:post_id", pr.PostController.GetAllCommentsAt)
	rg.PATCH("/update-comment-content", pr.PostController.UpdateCommentContent)
	rg.PATCH("/react-comment", pr.PostController.ReactComment)
	rg.DELETE("/delete-comment-at/:comment_id", pr.PostController.DeleteCommentAt)

	// ############ comment routes ############
	rg.POST("/reply", pr.PostController.Reply)
	rg.GET("/get-reply-at/:reply_id", pr.PostController.GetReplyAt)
	rg.GET("/get-all-replies-at/:comment_id", pr.PostController.GetAllRepliesAt)
	rg.PATCH("/update-reply-content", pr.PostController.UpdateReplyContent)
	rg.PATCH("/react-reply", pr.PostController.ReactReply)
	rg.DELETE("/delete-reply-at/:reply_id", pr.PostController.ReactReply)

}