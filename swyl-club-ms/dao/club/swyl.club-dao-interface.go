/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import "Swyl/servers/swyl-club-ms/models"

// @notice Dao interface
type ClubDao interface {

	// @notice Creates a club
	// 
	// @param clubOwner *string
	// 
	// @return error
	CreateClub(clubOwner *string) error


	// @notice Gets a club at clubId
	// 
	// @param clubOwner *string
	// 
	// @return *models.Club
	// 
	// @return error
	GetClubOwnedBy(clubOwner *string) (*models.Club, error)


	// @notice Gets a slice of all clubs
	// 
	// @return *[]models.Club
	// 
	// @return error
	GetAllClubs() (*[]models.Club, error)


	// @notice Updates a Club
	// 
	// @param clubOwner *string
	// 
	// @param totalMember *uint64
	UpdateClub(clubOwner *string, totalMember *uint64) error
}