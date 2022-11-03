/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import (
	"Swyl/servers/swyl-community-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in post-dao-impl
type PostDaoImpl struct {
	ctx 				context.Context
	mongoCollection 	*mongo.Collection
}

// @dev Constructor
func PostDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) PostDao {
	return &PostDaoImpl {
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}



// @notice Method of UserDaoImpl struct
// 
// @dev Lets a commOwner create a post
// 
// @param post *models.Post
// 
// @return error
func (pi *PostDaoImpl) CreatePost(post *models.Post) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a post at postId
// 
// @param postId *string
// 
// @return *models.Post
func (pi *PostDaoImpl) GetPostAt(postId *string) (*models.Post, error) {return nil, nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets all posts created by commOwner
// 
// @param commOwner *string
// 
// @return *[]models.Post
func (pi *PostDaoImpl) GetPostsBy(commOnwer *string) (*[]models.Post, error) {return nil, nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a commOwner update a post - only post.Title and post.Content are allowed
// 
// @param *model.Post
// 
// @return error
func (pi *PostDaoImpl) UpdatePostContent(post *models.Post) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user react to a post
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to post.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from post.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in post.Reaction
// 
// @param reaction *models.Reaction
// 
// @return error
func (pi *PostDaoImpl) ReactPost(reaction *models.Reaction) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user comment on a post
// 
// @param comment *models.Comment
// 
// @return error
func (pi *PostDaoImpl) Comment(comment *models.Comment) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user to update own comment - only comment.Content is allowed
// 
// @param comment *models.Comment
// 
// @return error
func (pi *PostDaoImpl) UpdateComment(comment *models.Comment) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user react to a comment
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to comment.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from comment.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in comment.Reaction
// 
// @param reaction *models.Reaction
// 
// @return error
func (pi *PostDaoImpl) ReactComment(reaction *models.Reaction) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user to delete own comment at commentId
// 
// @param commentId *string
// 
// @return error
func (pi *PostDaoImpl) DeleteCommentAt(commentId *string) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user reply to a comment
// 
// @param reply *models.Reply
// 
// @return error
func (pi *PostDaoImpl) Reply(reply *models.Reply) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user update own reply - only reply.Content is allowed
// 
// @param reply *models.Reply
// 
// @return error
func (pi *PostDaoImpl) UpdateReply(reply *models.Reply) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user react a reply
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to comment.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from comment.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in comment.Reaction
// 
// @param reaction *models.Reaction
// 
// @return error
func (pi *PostDaoImpl) ReactReply(reaction *models.Reply) error {return nil}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user delete own reply at replyId
// 
// @param replyId *string
// 
// @return error
func (pi *PostDaoImpl) DeleteReplyAt(replyId *string) error {return nil}