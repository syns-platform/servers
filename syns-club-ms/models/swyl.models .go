/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// @notice The information related to a Swyl club
type Club struct {
	Club_owner					*string		`json:"club_owner" bson:"club_owner" validate:"required,len=42,alphanum"`
	Created_at					uint64		`json:"created_at" bson:"created_at" validate:"required"`
	Total_members					uint64		`json:"total_members" bson:"total_members"`
}


// @notice The information related to a Swyl Tier
type Tier struct {
	Tier_ID						primitive.ObjectID		`json:"tier_id" bson:"_id, omitempty"`
	Club_owner					*string 			`json:"club_owner" bson:"club_owner" validate:"required,len=42,alphanum"`
	Tier_name					*string				`json:"tier_name" bson:"tier_name" validate:"required,min=2,max=20"`
	Tier_img					*string				`json:"tier_img" bson:"tier_img"`
	Tier_bio					*string				`json:"tier_bio" bson:"tier_bio" validate:"omitempty,min=2,max=200"`
	Tier_fee					*uint64				`json:"tier_fee" bson:"tier_fee" validate:"required"`
	Tier_limit					*uint64				`json:"tier_limit" bson:"tier_limit" validate:"required,ne=0"`
	Tier_welcome_msg				*string				`json:"tier_welcome_msg" bson:"tier_welcome_msg" validate:"omitempty,min=2,max=100"`
	Created_at					uint64				`json:"created_at" bson:"created_at" validate:"required"`
}


// @notice The information related to a Swyl Subscription
// @TODO: add username, email, displayname
type Subscription struct {
	Subscription_ID					primitive.ObjectID 		`json:"subscription_id" bason:"_id"`
	Club_owner					*string 			`json:"club_owner" bson:"club_owner" validate:"required,len=42,alphanum"`
	Tier_ID						primitive.ObjectID		`json:"tier_id" bson:"tier_id, omitempty"`
	Subscriber					*string 			`json:"subscriber" bson:"subscriber" validate:"required,len=42,alphanum"`
	Status						bool				`json:"status" bson:"status"`
	Joined_at					uint64				`json:"joined_at" bson:"joined_at" validate:"required"`
}

