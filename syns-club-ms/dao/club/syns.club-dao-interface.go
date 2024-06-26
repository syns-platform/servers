/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import "Syns/servers/syns-club-ms/models"

// @notice ClubDao interface
type ClubDao interface {

	// @notice Creates a club
	// 
	// @param clubOwner *string
	// 
	// @param createdAt uint64
	// 
	// @return error
	CreateClub(clubOwner *string, createdAt uint64) error


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


	// @notice Increase/decrease a Club's total members
	// 
	// @param clubOwner *string
	// 
	// @param context *uint16 0=decrease || 1=increase
	UpdateClub(clubOwner *string, context *uint16) error
}