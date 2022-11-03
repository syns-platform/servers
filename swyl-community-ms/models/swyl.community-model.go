/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// @notice The information related to a Swyl Community
type Community struct {
	Community_owner 					*string 			`json:"community_owner" bson:"community_owner" validate:"required,len=42,alphanum"`
	Bio							*string				`json:"bio" bson:"bio" validate:"omitempty"`
	Total_followers	 					uint64				`json:"total_followers" bson:"total_followers" validate:"omitempty"`
	Total_posts						uint64				`json:"total_posts" bson:"total_posts" validate:"omitempty"`
}

// @notice The information related to a Swyl Follower
type Follower struct {
	Follower_ID 						primitive.ObjectID 		`json:"follower_Id" bson:"_id"`
	Community_owner 					*string 			`json:"community_owner" bson:"community_owner" validate:"required,len=42,alphanum"`
	Follower						*string 			`json:"follower" bson:"follower" validate:"required,len=42,alphanum"`
	Follow_at						uint64				`json:"follow_at" bson:"follow_at"`
}

// @notice The information related to a Swyl POST
type Post struct {
	Post_ID 						primitive.ObjectID 		`json:"post_id" bson:"_id"`
	Community_owner 					*string 			`json:"community_owner" bson:"community_owner" validate:"required,len=42,alphanum"`
	Title 							*string				`json:"title" bson:"title" validate:"required,min=2,max=100"`
	Content							*string				`json:"content" bson:"content" validate:"required,max=10000"`
	Reaction						[]Reaction			`json:"reaction" bson:"reaction"`
	Created_at						uint64				`json:"created_at" bson:"created_at"`
}

// @notice The information related to a Swyl Reaction
type Reaction struct {
	Reacter 						*string				`json:"reacter" bson:"reacter" validate:"required,len=42,alphanum"`
	Post_ID							primitive.ObjectID 		`json:"post_id" bson:"post_id" validate:"required"`
	React_type 						*string 			`json:"react_type" bson:"react_type" validate:"required,oneof=SUPPORT BRAVO LAUGH FIRE"`
	React_at						uint64				`json:"react_at" bson:"react_at"`
}

// @notice The information related to a Swyl Comment
type Comment struct {
	Comment_ID  						primitive.ObjectID 		`json:"comment_id" bson:"_id"`
	Post_ID							primitive.ObjectID 		`json:"post_id" bson:"post_id" validate:"required"`
	Commenter 						*string 			`json:"commenter" bson:"commenter" validate:"required,len=42,alphanum"`
	Content							*string				`json:"content" bson:"content" validate:"required,max=700"`
	Reaction						[]Reaction			`json:"reaction" bson:"reaction"`
	Comment_at						uint64				`json:"commente_at" bson:"commente_at"`
}

// @notice The information related to a Swyl Reply
type Reply struct {
	Reply_id						primitive.ObjectID 		`json:"reply_id" bson:"_id"`
	Comment_ID  						primitive.ObjectID 		`json:"comment_id" bson:"comment_id"`
	Reply_to						*string				`json:"reply_to" bson:"reply_to" validate:"required,len=42,alphanum"`
	Content							*string				`json:"content" bson:"content" validate:"required,max=1000"`
	Reaction						[]Reaction			`json:"reaction" bson:"reaction"`
	Reply_at						uint64				`json:"reply_at" bson:"reply_at"`
}