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
	TokenID          int       `json:"tokenId" bson:"token_id" validate:"required,number"`
	AssetContract    string    `json:"assetContract" bson:"asset_contract" validate:"required,eth_addr"`
	TokenOwner       string    `json:"tokenOwner" bson:"token_owner" validate:"required,eth_addr"`
	OriginalOwner    string    `json:"originalOwner" bson:"original_owner" validate:"required,eth_addr"`
	TokenURI         string    `json:"tokenURI" bson:"token_uri" validate:"required"`
	Image            string    `json:"image" bson:"image" validate:"required"`
	Audio            string    `json:"audio" bson:"audio" validate:"required"`
	ERCType          string    `json:"ercType" bson:"erc_type" validate:"required"`
	Quantity         int       `json:"quantity" bson:"quantity" validate:"required,number"`
	OriginalQuantity int       `json:"originalQuantity" bson:"original_quantity" validate:"required,number"`
	IsListing        bool      `json:"isListing" bson:"is_listing"`
	ListingID        int       `json:"listingId" bson:"listing_id" validate:"required,number"`
	RoyaltyBps       int       `json:"royaltyBps" bson:"royalty_bps" validate:"required,number"`
	Name             string    `json:"name" bson:"name" validate:"required"`
	Description      string    `json:"description" bson:"description" validate:"required"`
	Age              int       `json:"age" bson:"age" validate:"required,number"`
	SharableLink     string    `json:"sharableLink" bson:"sharable_link" validate:"required"`
}

// @notice struct for Syns Token
type Syns721SuperNFT struct {
	TokenHash		 string    `json:"token_hash" bson:"token_hash" validate:"omitempty"`
	TokenID          int       `json:"tokenId" bson:"token_id" validate:"omitempty,number"`
	AssetContract    string    `json:"assetContract" bson:"asset_contract" validate:"required,eth_addr"`
	TokenOwner       string    `json:"tokenOwner" bson:"token_owner" validate:"required,eth_addr"`
	OriginalOwner    string    `json:"originalOwner" bson:"original_owner" validate:"required,eth_addr"`
	TokenURI         string    `json:"tokenURI" bson:"token_uri" validate:"required"`
	Image            string    `json:"image" bson:"image" validate:"required"`
	Audio            string    `json:"audio" bson:"audio" validate:"required"`
	ERCType          string    `json:"ercType" bson:"erc_type" validate:"required"`
	Quantity         int       `json:"quantity" bson:"quantity" validate:"required,number"`
	IsListing        bool      `json:"isListing" bson:"is_listing"`
	ListingID        int       `json:"listingId" bson:"listing_id" validate:"omitempty,number"`
	RoyaltyBps       int       `json:"royaltyBps" bson:"royalty_bps" validate:"omitempty,number"`
	Name             string    `json:"name" bson:"name" validate:"required"`
	Description      string    `json:"description" bson:"description" validate:"omitempty"`
	Age              int       `json:"age" bson:"age" validate:"required,number"`
	SharableLink     string    `json:"sharableLink" bson:"sharable_link" validate:"required"`
	Lister			 string	   `json:"lister" bson:"lister" validate:"required,eth_addr"`
	StartSale		 int	   `json:"startSale" bson:"start_sale" validate:"omitempty,number"`
	EndSale			 int	   `json:"endSale" bson:"end_sale" validate:"required,number"`
	Currency		 string	   `json:"currency" bson:"currency" validate:"required,eth_addr"`
	BuyouPricePerToken	string	   `json:"buyoutPricePerToken" bson:"buyout_price_per_token" validate:"omitempty"`
}
