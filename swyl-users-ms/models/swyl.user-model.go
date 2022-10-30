/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @notice The information related to a Swyl user
type User struct {
	ID				primitive.ObjectID			`bson:"_id"`
	Wallet_address 			string					`json:"wallet_address" bson:"wallet_address" validate:"required, min:42, max:42"`
	Username 			string					`json:"username" bson:"username" validate:"required, min:3, max:20"`
	Avatar				string					`json:"avatar" bson:"avatar"`	
	Display_name 			string					`json:"display_name" bson:"display_name" validate:"min:3, max:40"`
	Email 				string					`json:"email" bson:"email" validate:"min:3, max:40"`
	Bio 				string					`json:"bio" bson:"bio" validate:"max: 200"`
	Website 			string					`json:"website" bson:"website"`
	Joined_at			time.Time				`json:"joined_at" bson:"joined_at"`
	Social_media 			[]SocialMedia				`json:"social_media" bson:"social_media"`
}


// @notice The information related to the links to a user's social media
type SocialMedia struct {
	Media 				string 					`json:"media" bson:"media"`
	Url				string					`json:"url" bson:"url"`
}