/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

import "Swyl/servers/swyl-community-ms/models"

// @notice Dao interface
type CommDao interface {

	// @notice Lets a Swyl user create a community
	// 
	// @NOTE Should be fired off when #user/connect api is called
	// 
	// @param comm *models.Community
	// 
	// @return error
	CreateComm(comm *models.Community) error

	
	// @notice Gets a Comm owned by commOwner
	// 
	// @param commOwner *string
	// 
	// @return *models.Community
	// 
	// @return error
	GetCommOwnedBy(commOwner *string) (*models.Community, error)


	// @notice Gets all Comm has ever created
	// 
	// @NOTE Might not be necessary
	// 
	// return *[]models.Community
	GetAllComms() (*[]models.Community, error)


	// @notice Updates Comm's total_followers || Comm's total_posts
	// 
	// @param commOwner *string
	// 
	// @return error
	UpdateCommOwnedBy(commOwner *string) error
	
}