/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

// @notice struct for Syns Token
type SynsNFT struct {
	TokenID          int       `json:"tokenId"`
	AssetContract    string    `json:"assetContract"`
	TokenOwner       string    `json:"tokenOwner"`
	OriginalOwner    string    `json:"originalOwner"`
	TokenURI         string    `json:"tokenURI"`
	Image            string    `json:"image"`
	Audio            string    `json:"audio"`
	ERCType          string    `json:"ercType"`
	Quantity         int       `json:"quantity"`
	OriginalQuantity int       `json:"originalQuantity"`
	IsListing        bool      `json:"isListing"`
	ListingID        int       `json:"listingId"`
	RoyaltyBps       int       `json:"royaltyBps"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Age              int       `json:"age"`
	SharableLink     string    `json:"sharableLink"`
}
