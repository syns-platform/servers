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
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in post-dao-impl
type PostDaoImpl struct {
	ctx 				context.Context
	postCollection 		*mongo.Collection
	commentCollection	*mongo.Collection
	replyCollection		*mongo.Collection
}

// @dev Constructor
func PostDaoConstructor(
	ctx context.Context, 
	postCollection *mongo.Collection,
	commentCollection *mongo.Collection,
	replyCollection *mongo.Collection,
) PostDao {
	return &PostDaoImpl {
		ctx: ctx,
		postCollection: postCollection,
		commentCollection: commentCollection,
		replyCollection: replyCollection,
	}
}


// #############################################
// 			 		Post APIs 
// #############################################

// @notice Method of UserDaoImpl struct
// 
// @dev Lets a commOwner create a post
// 
// @param post *models.Post
// 
// @return error
func (pi *PostDaoImpl) CreatePost(post *models.Post) error {
	// set up post's objectId
	post.Post_ID = primitive.NewObjectID()

	// set up post.Created_at
	post.Created_at = uint64(time.Now().Unix())

	// insert post to internal database
	_, err := pi.postCollection.InsertOne(pi.ctx, post)
	return err
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a post at postId
// 
// @param postId *string
// 
// @return *models.Post
func (pi *PostDaoImpl) GetPostAt(postId *string) (*models.Post, error) {
	// prepare post placeholder
	post := &models.Post{}

	// prepare objectId
	objectId, idErr := primitive.ObjectIDFromHex(*postId)
	if idErr != nil {return nil, idErr}

	// prepare filter query
	filter := bson.M{"_id": objectId}

	// find post in database and decode to post placeholder
	err := pi.postCollection.FindOne(pi.ctx, filter).Decode(post)
	
	// return OK
	return post, err
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets all posts created by commOwner
// 
// @param commOwner *string
// 
// @return *[]models.Post
func (pi *PostDaoImpl) GetPostsBy(commOnwer *string) (*[]models.Post, error) {
	// prepare posts placeholder
	posts := &[]models.Post{}

	// prepare filter query
	filter := bson.M{"community_owner": commOnwer}
	
	// find all posts by commOnwer
	cursor, err := pi.postCollection.Find(pi.ctx, filter); if err != nil {return nil, err}

	// decode cursor into posts placeholder
	decodeErr := cursor.All(pi.ctx, posts)

	// return OK
	return posts, decodeErr
}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a commOwner update a post - only post.Title and post.Content are allowed
// 
// @param *model.Post
// 
// @return error
func (pi *PostDaoImpl) UpdatePostContent(post *models.Post) error {
	// prepare filter query
	filter := bson.M{"_id": post.Post_ID}

	// find the post in database
	if dbRes := pi.postCollection.FindOne(pi.ctx, filter); dbRes.Err() != nil {return dbRes.Err()}
	
	// prepare update query
	update := bson.D{
		{Key: "$set", Value: bson.M{"title": post.Title}},
		{Key: "$set", Value: bson.M{"content": post.Content}},
	}

	// update the post in database
	_, err := pi.postCollection.UpdateOne(pi.ctx, filter, update)

	// return OK
	return err
}


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
func (pi *PostDaoImpl) ReactPost(reaction *models.Reaction) error {
	// set up reaction.React_at
	reaction.React_at = uint64(time.Now().Unix())

	// prepare filter
	filter := bson.M{"_id": reaction.Post_ID}

	// find the post in postCollection
	targetPost := &models.Post{}
	if err := pi.postCollection.FindOne(pi.ctx, filter).Decode(targetPost); err != nil {return err}

	// get Reaction slice from targetPost
	reactions := targetPost.Reaction

	// loop through reactions slice to examine each reactionItem in the slice
	// @logic-a if len(reactions) == 0 => no reaction yet => simply add the reaction to reactions slice
	if len(reactions) == 0 {
		reactions = append(reactions, *reaction)
	} else {
		// @logic-b if reaction.Reacter is not found in reactions slice, simply add the reaction to reactions slice
		// @logic-c if a reactionItem.Reacter == reaction.Reacter => this request is to either update the react_type, or unreact the post
		// @logic-c.1 if the reactionItem.React_type != reaction.React_type => update the new react_type to reactionItem
		// @logic-c.2 if the reactionItem.React_type == reaction.React_type => remove the reactionItem off of the reactions slice
		shouldUpdateReact := false
		shouldUnreact := false
		unreactIndex := 0

		for index, reactionItem := range reactions {
			if (*reactionItem.Reacter == *reaction.Reacter && *reactionItem.React_type != *reaction.React_type) { //logic-c.1
				// update reactionItem.React_type
				reactions[index].React_type = reaction.React_type
				shouldUpdateReact = true
				break
			} else if (*reactionItem.Reacter == *reaction.Reacter && *reactionItem.React_type == *reaction.React_type) { //logic-c.2
				// update shouldUnreact & unreactIndex
				unreactIndex = index
				shouldUnreact = true
				break
			}
		}

		if shouldUnreact { // logic-c.2
			reactions[unreactIndex] = reactions[len(reactions) - 1]
			reactions = reactions[:len(reactions)-1]
		} else if !shouldUpdateReact && !shouldUnreact { // logic-b
			reactions = append(reactions, *reaction)
		}
	}

	// prepare update query
	update := bson.D{
		{Key: "$set", Value: bson.M{"reaction": reactions}},
	}

	// update post
	_, err := pi.postCollection.UpdateOne(pi.ctx, filter, update)

	// return OK
	return err
}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a commOwner delete own post at
// 
// @param postId *string
// 
// @return error
func (pi *PostDaoImpl) DeletePostAt(postId *string) error {
	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {return err}

	// prepare filter query
	filter := bson.M{"_id": objectId}

	// delete post
	dbRes, err := pi.postCollection.DeleteOne(pi.ctx, filter) 
	if err != nil {return nil}
	if dbRes.DeletedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}
	
	// return ok
	return nil
}


// #############################################
// 			 		Comment APIs 
// #############################################

// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user comment on a post
// 
// @param comment *models.Comment
// 
// @return error
func (pi *PostDaoImpl) Comment(comment *models.Comment) error {
	// make sure comment.Post_id is a valid post
	filter := bson.M{"_id": comment.Post_ID}
	if dbRes := pi.postCollection.FindOne(pi.ctx, filter); dbRes.Err() != nil {return errors.New("!MONGO - no post document found with post_id")}

	// prepare comment
	comment.Comment_ID = primitive.NewObjectID()
	comment.Comment_at = uint64(time.Now().Unix())

	// insert comment to commentCollection
	_, err := pi.commentCollection.InsertOne(pi.ctx, comment)
	
	// return OK
	return err
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a comment at commentId
// 
// @param commentId *string
// 
// @return *models.Comment
// 
// @return error
func (pi *PostDaoImpl) GetCommentAt(commentId *string) (*models.Comment, error) {
	// decalre post placeholder
	post := &models.Comment{}

	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*commentId); if err != nil {return nil, err}

	// prepare filter query
	filter := bson.M{"_id": objectId}

	// find the comment in database
	if err := pi.commentCollection.FindOne(pi.ctx, filter).Decode(post); err != nil {return nil, err}
	
	// return OK
	return post, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets all comments at postId
// 
// @param postId *string
// 
// @return *[]models.Comment
// 
// @return error
func (pi *PostDaoImpl) GetAllCommentsAt(postId *string) (*[]models.Comment, error) {
	// declare posts placeholder
	posts := &[]models.Comment{}

	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*postId); if err != nil {return nil, err}

	// prepare filter query
	filter := bson.M{"post_id": objectId}

	// find posts in database
	cursor, err := pi.commentCollection.Find(pi.ctx, filter); if err != nil {return nil, err}
	
	// decode cursor to posts placeholder
	if err := cursor.All(pi.ctx, posts); err != nil {return nil, err}

	// return OK
	return posts, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user to update own comment - only comment.Content is allowed
// 
// @param comment *models.Comment
// 
// @return error
func (pi *PostDaoImpl) UpdateCommentContent(comment *models.Comment) error {
	// prepare filter query
	filter := bson.M{
		"_id": comment.Comment_ID,
		"post_id": comment.Post_ID,
		"commenter": comment.Commenter,
	}

	// prepare update query
	update := bson.D{{Key: "$set", Value: bson.M{"content": comment.Content}}}

	// update comment
	dbRes, err := pi.commentCollection.UpdateOne(pi.ctx, filter, update)
	if err != nil {return err}
	if dbRes.MatchedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}
	
	// return OK
	return nil
}


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
func (pi *PostDaoImpl) ReactComment(reaction *models.Reaction) error {
	// @notice reaction.Post_ID = _id (a.k.a Comment_ID in this case)

	// prepare filter query
	filter := bson.M{"_id": reaction.Post_ID}

	// declare comment holder
	targetComment := &models.Comment{}

	// find targetComment
	if err := pi.commentCollection.FindOne(pi.ctx, filter).Decode(targetComment); err != nil {return err}

	// get the Reaction field in targetComment
	reactions := targetComment.Reaction

	// logic-a: if reactions is an empty slice, simply add the new passed in `reaction` to reactions slice
	if len(reactions) == 0 {
		reactions = append(reactions, *reaction)
	} else {
		// logic-b: if reaction.Reacter is NOT found in reactions slice => simply add the new passed in `reaction` to reaction slice
		// logic-c: if reaction.Reacter is found in reactions slice => the request is either to update the reaction or unreact the reaction
		// logic-c.1 if reaction.React_type != reactionItem.React_type => update reaction
		// logic-c.2 if reaction.Reaction_type == reactionItem.React_type => unreact reaction

		shouldUpdateReact := false
		shouldUnreact := false
		targetIndex := 0

		for index, reactionItem := range reactions {
			if *reaction.Reacter == *reactionItem.Reacter && *reaction.React_type != *reactionItem.React_type { //logic-c.1
				reactions[index].React_type = reaction.React_type
				shouldUpdateReact = true
			} else if *reaction.Reacter == *reactionItem.Reacter && *reaction.React_type == *reactionItem.React_type { //logic-c.2
				targetIndex = index
				shouldUnreact = true
			}
		}

		if shouldUnreact { // logic-c.2
			// swap the targetItem with the lastItem
			reactions[targetIndex] = reactions[len(reactions) - 1] 
			
			// exclude last item
			reactions = reactions[:len(reactions) - 1] 
		} else if !shouldUnreact && !shouldUpdateReact { // logic-b
			reactions = append(reactions, *reaction)
		}
	}

	// prepare update query
	update := bson.D{
		{Key: "$set", Value: bson.M{"reaction": reactions}},
	}

	// update the comment
	dbRes, err := pi.commentCollection.UpdateOne(pi.ctx, filter, update)
	if err != nil {return err}
	if dbRes.MatchedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}
	
	
	// return OK
	return nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user to delete own comment at commentId
// 
// @param commentId *string
// 
// @return error
func (pi *PostDaoImpl) DeleteCommentAt(commentId *string) error {
	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*commentId); if err != nil {return nil}
	
	// prepare filter query
	filter := bson.M{"_id": objectId}

	// delete the comment
	dbRes, err := pi.commentCollection.DeleteOne(pi.ctx, filter);
	if err != nil {return nil}
	if dbRes.DeletedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}
	
	// return OK
	return nil
}


// #############################################
// 			 		Reply APIs 
// #############################################

// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user reply to a comment
// 
// @param reply *models.Reply
// 
// @return error
func (pi *PostDaoImpl) Reply(reply *models.Reply) error {
	// make sure reply is for a valid comment
	filter := bson.M{"_id": reply.Comment_ID,"commenter": reply.Reply_to,}
	if dbRes := pi.commentCollection.FindOne(pi.ctx, filter); dbRes.Err() != nil {return errors.New("!MONGO - No comment is found with comment_id and reply_to")}

	// prepare reply
	reply.Reply_ID = primitive.NewObjectID()
	reply.Reply_at = uint64(time.Now().Unix())

	// insert new reply to database
	_, err := pi.replyCollection.InsertOne(pi.ctx, reply)

	// return OK
	return err
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a reply at replyId
// 
// @param replyId *string
// 
// @return *model.Reply
// 
// @return error
func (pi *PostDaoImpl) GetReplyAt(replyId *string) (*models.Reply, error) {
	// declare reply holder
	reply := &models.Reply{}

	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*replyId); if err != nil {return nil, err}

	// prepare filter query
	filter := bson.M{"_id": objectId}

	// find the reply in database
	if err := pi.replyCollection.FindOne(pi.ctx, filter).Decode(reply); err != nil {return nil, err}
	
	// return OK
	return reply, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets all replies at commentId
// 
// @param commentId *string
// 
// @return *[]models.Reply
// 
// @return error
func (pi *PostDaoImpl) GetAllRepliesAt(commentId *string) (*[]models.Reply, error) {
	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*commentId); if err != nil {return nil, err}

	// prepare filter
	commentFilter := bson.M{"_id": objectId}
	repliesFilter := bson.M{"comment_id": objectId}

	// make sure commentId points at a valid comment
	if dbRes := pi.commentCollection.FindOne(pi.ctx, commentFilter); dbRes.Err() != nil {return nil, errors.New("!MONGO - No comment is found with comment_id")}

	// find the replies
	replies := &[]models.Reply{}
	cursor, err := pi.replyCollection.Find(pi.ctx, repliesFilter); if err != nil {return nil, err}

	// decode cursor into repleis holder
	if err := cursor.All(pi.ctx, replies); err != nil {return nil, err}
	
	// return OK
	return replies, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Lets a user update own reply - only reply.Content is allowed
// 
// @param reply *models.Reply
// 
// @return error
func (pi *PostDaoImpl) UpdateReplyContent(reply *models.Reply) error {
	// prepare filter query
	commentFilter := bson.M{"_id": reply.Comment_ID}
	replyFilter := bson.M{
		"_id": reply.Reply_ID, 
		"comment_id": reply.Comment_ID,
		"reply_to": reply.Reply_to,
	}

	// check if reply.Comment_ID points at a valid comment
	if dbRes := pi.commentCollection.FindOne(pi.ctx, commentFilter); dbRes.Err() != nil {return errors.New("!MONGO - No comment is found with comment_id") }

	// prepare update query
	update := bson.D{{Key: "$set", Value: bson.M{"content": reply.Content}}}
	
	// update reply
	dbRes, err := pi.replyCollection.UpdateOne(pi.ctx, replyFilter, update)
	if err != nil {return nil}
	if dbRes.MatchedCount == 0 {return errors.New("!MONGO - No matched document found")}

	// return OK
	return nil
}


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