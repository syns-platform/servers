/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

// @notice The information related to a Syns user
type Feedback struct {
	Email string `json:"email"`
	Feedback string `json:"feedback" validate:"required"`
	Submitted_at int64	`json:"submitted_at" bson:"submitted_at"`
}