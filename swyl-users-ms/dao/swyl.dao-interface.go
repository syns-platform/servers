/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import "Swyl/servers/swyl-users-ms/models"

// @notice Dao interface
type UserDao interface {

	// @notice Connects to an account stored in the internal database using wallet address.
	// 		   Create a new account on first connect.
	// 
	// @param user	*models.User
	// 
	// @return error
	Connect(walletAddress *string) (*models.User ,error);

	// @notice Gets a user at wallet address.
	// 
	// @param walletAddress *string
	// 
	// @return *models.User
	// 
	// @return error
	GetUserAt(walletAddress *string) (*models.User, error)

	// @notice Gets all user.
	// 
	// @NOTE might not be necessary
	// 
	// @return []*models.User
	GetAllUsers() (*[]models.User, error)

	// @notice Updates a user.
	// 
	// @param *models.User
	// 
	// @return error
	UpdateUser(*models.User) error

	// @notice Deletes a user at wallet address.
	// 
	// @param walletAddress *string
	// 
	// @return error
	DeactivateUserAt(walletAddress *string) error
}