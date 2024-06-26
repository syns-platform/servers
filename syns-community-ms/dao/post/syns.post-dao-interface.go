/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

import "Syns/servers/syns-community-ms/models"

// @notice Dao interface
type PostDao interface {

	// #############################################
	// 			 		Post APIs 
	// #############################################

	// @notice Lets a commOwner create a post
	// 
	// @param post *models.Post
	// 
	// @return error
	CreatePost(post *models.Post) error

	// @notice Gets a post at postId
	// 
	// @param postId *string
	// 
	// @return *models.Post
	GetPostAt(postId *string) (*models.Post, error)
	
	// @notice Gets all posts created by commOwner
	// 
	// @param commOwner *string
	// 
	// @return *[]models.Post
	GetPostsBy(commOnwer *string) (*[]models.Post, error)

	// @notice Lets a commOwner update a post - only post.Title and post.Content are allowed
	// 
	// @param *model.Post
	// 
	// @return error
	UpdatePostContent(post *models.Post) error

	// @notice Lets a user react to a post
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
	ReactPost(reaction *models.Reaction) error

	// @notice Lets a commOwner delete own post at postId
	// 
	// @param postId *string
	// 
	// @return error
	DeletePostAt(postId *string) error


	// #############################################
	// 			 		Comment APIs 
	// #############################################

	// @notice Lets a user comment on a post
	// 
	// @param comment *models.Comment
	// 
	// @return error
	Comment(comment *models.Comment) error

	// @notice Gets a comment at commentId
	// 
	// @param commentId *string
	// 
	// @return *models.Comment
	// 
	// @return error
	GetCommentAt(commentId *string) (*models.Comment, error)

	// @notice Gets all comments at postId
	// 
	// @param postId *string
	// 
	// @return *[]models.Comment
	// 
	// @return error
	GetAllCommentsAt(postId *string) (*[]models.Comment, error)

	// @notice Lets a user to update own comment - only comment.Content is allowed
	// 
	// @param comment *models.Comment
	// 
	// @return error
	UpdateCommentContent(comment *models.Comment) (error)

	// @notice Lets a user react to a comment
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
	ReactComment(reaction *models.Reaction) error

	// @notice Lets a user to delete own comment at comment_id
	// 
	// @param commentId *string
	// 
	// @return error
	DeleteCommentAt(commentId *string) error


	// #############################################
	// 			 		Reply APIs 
	// #############################################

	// @notice Lets a user reply to a comment
	// 
	// @param reply *models.Reply
	// 
	// @return error
	Reply(reply *models.Reply) error

	// @notice Gets a reply at replyId
	// 
	// @param replyId *string
	// 
	// @return *model.Reply
	// 
	// @return error
	GetReplyAt(replyId *string) (*models.Reply, error)

	// @notice Gets all replies at commentId
	// 
	// @param commentId *string
	// 
	// @return *[]models.Reply
	// 
	// @return error
	GetAllRepliesAt(commentId *string) (*[]models.Reply, error)

	// @notice Lets a user update own reply - only reply.Content is allowed
	// 
	// @param reply *models.Reply
	// 
	// @return error
	UpdateReplyContent(reply *models.Reply) error

	// @notice Lets a user react a reply
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
	ReactReply(reaction *models.Reaction) error

	// @notice Lets a user delete own reply at replyId
	// 
	// @param replyId *string
	// 
	// @return error
	DeleteReplyAt(replyId *string) error
}