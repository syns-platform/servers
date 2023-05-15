/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

// @notice The information related to a Demo Request
type DemoRequest struct {
	Email string `json:"email" validate:"email"`
   Name string `json:"name" validate:"omitempty"`
	Question string `json:"feedback" validate:"omitempty"`
	Submitted_at int64	`json:"submitted_at" bson:"submitted_at"`
}