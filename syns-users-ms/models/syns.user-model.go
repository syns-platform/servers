/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

// @notice The information related to a Syns user
type User struct {
	Wallet_address 			*string					`json:"wallet_address" bson:"wallet_address" validate:"required,len=42,alphanum"`
	Username 			*string					`json:"username" bson:"username" validate:"omitempty,min=2,max=30"`
	Avatar				*string					`json:"avatar" bson:"avatar"`	
	Display_name 			*string					`json:"display_name" bson:"display_name" validate:"omitempty,min=3,max=40"`
	Email 				*string					`json:"email" bson:"email" validate:"omitempty,min=3,max=40,email"`
	Bio 				*string					`json:"bio" bson:"bio" validate:"omitempty,max=200"`
	Website 			*string					`json:"website" bson:"website" validate:"omitempty,url"`
	Social_media 			[]SocialMedia				`json:"social_media" bson:"social_media"`
	Joined_at			int64					`json:"joined_at" bson:"joined_at"`
}


// @notice The information related to the links to a user's social media
type SocialMedia struct {
	Media 				string 					`json:"media" bson:"media"`
	Url				string					`json:"url" bson:"url" validate:"omitempty,url"`
}