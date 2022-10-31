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
	// @param clubId *uint64
	// 
	// @return *models.Club
	// 
	// @return error
	GetClubAt(clubId *uint64) (*models.Club, error)


	// @notice Gets a club at clubId
	// 
	// @param clubOwner *string
	// 
	// @return *models.Club
	// 
	// @return error
	GetClubOwnedBy(clubOwner *string) (*models.Club, error)


	// @notice Updates a Club's Total_member
	// 
	// @param clubOwner *string
	// 
	// @param totalMember *uint64
	UpdateClubTotalMember(clubOwner *string, totalMember *uint64) error
}