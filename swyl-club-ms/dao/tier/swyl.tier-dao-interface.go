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

// @notice TierDao interface
type TierDao interface {

   	// @notice Lets a club owner create a tier to internal database
	// 
	// @param tier *models.Tier
	// 
	// @return error
   	CreateTier(tier *models.Tier) error

   	// @notice Gets a Tier at tierId and clubOwner
	// 
	// @param tierId *string
	// 
	// @return *models.Tier
	// 
	// @return error
   	GetTierAt(tierId *string) (*models.Tier, error)


	// @notice Gets all tiers owned by clubOwner
	// 
	// @param clubOwner *string
	// 
	// @return *[]models.Tier
	// 
	// @return error
	GetAllTiersOwnedBy(clubOwner *string) (*[]models.Tier, error)


	// @notice Lets a clubOwner update a tier
	// 
	// @param tier *models.Tier
	// 
	// @return error
	UpdateTier(tier *models.Tier) error

	// @notice Lets a clubOwner delete a tier
	// 
	// @param tierId *string
	// 
	// @return error
	DeleteTier(clubId *string) error
}