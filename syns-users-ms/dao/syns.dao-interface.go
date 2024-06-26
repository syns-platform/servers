/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import "Syns/servers/syns-users-ms/models"

// @notice UserDao interface
type UserDao interface {

	// @notice Connects to an account stored in the internal database using wallet address.
	// 		   Create a new account on first connect.
	// 
	// @param walletAddress string
	// 
	// @return *models.User
	// 
	// @return bool - first connect check
	// 
	// @return error
	Connect(walletAddress string) (*models.User, bool, error)

	// @notice Lets a wallet owner claim a Syns page with passed-in username
	// 
	// @param userParam	*models.User
	// 
	// @return *mdoels.User
	// 
	// @return error
	ClaimPage(userParam *models.User) (*models.User, error)

	// @notice Checks if a username has been taken
	// 
	// @param username string
	// 
	// @return bool
	CheckUsernameAvailability(username string) (bool, error)

	// @notice Gets a user at wallet address.
	// 
	// @param walletAddress string
	// 
	// @return *models.User
	// 
	// @return error
	GetUserAt(walletAddress string) (*models.User, error)

	// @notice Gets a user by username.
	// 
	// @param username string
	// 
	// @return *models.User
	// 
	// @return error
	GetUserBy(username string) (*models.User, error)

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
	UpdateUser(user *models.User) error

	// @notice Deletes a user at wallet address.
	// 
	// @param walletAddress string
	// 
	// @return error
	DeactivateUserAt(walletAddress string) error
}

// @notice FeedbackDao interface
type FeedbackDao interface {
	// @notice Handle feedback submission
	// 
	// @param email string
	// 
	// @param feedback string
	// 
	// @return error
	SubmitFeedback(email, feedback string) (error)
}

// @notice DemoRequestDao interface
type DemoRequestDao interface {
	// @notice Handle demo request submission
	// 
	// @param email string
	// 
	// @param name string
	// 
	// @param question string
	// 
	// @return error
	SubmitDemoRequest(email, name, question string) (error)
}