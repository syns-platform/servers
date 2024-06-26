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

// @notice SubDao interface
type SubDao interface {

	// @notice Lets a user subscribe to a tier
	// 
	// @param sub *models.Subscription
	// 
	// @return error
	Subscribe(sub *models.Subscription) error


	// @notice Gets a subscription using subscription_id
	// 
	// @param subId *string
	// 
	// @return *models.Subscription
	GetSubscriptionAt(subId *string) (*models.Subscription, error)


	// @notice Gets all subscriptions onwed at tier_id and by club_owner
	// 
	// @param tierId *string
	// 
	// @param clubOwner *string
	// 
	// @return *[]models.Subscription
	// 
	// @return error
	GetAllSubsAt(tierId *string, clubOwner *string) (*[]models.Subscription, error)


	// @notice Toggles a subscription status at subId
	// 
	// @param subId *string
	// 
	// @return error
	ToggleSubStatusAt(subId *string) error


	// @notice Lets a subscriber unsubscribe a tier 
	// 
	// @param tierId *string
	// 
	// @return error
	Unsubscribe(subId *string) error
}