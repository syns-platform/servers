/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import (
	"Syns/servers/syns-tokens-ms/models"
)

// @notice Syns721TokenDao interface
type Syns721TokenDao interface {

	// @notice Add a new Syns 721 Super Token to database
	// 
	// @param synsNFT *models.Syns721SuperNFT
	// 
	// @return error
	MintNewSyns721Token(synsNFT *models.Syns721SuperNFT) (error)

	// @notice Update Syns 721 Super Token using Syns Listing from Syns Marketplace
	// 
	// @param synsListing models.SynsMarketplaceListing
	// 
	// @param eventName string
	// 
	// @return error
	UpdatedSyns721SuperTokenBySynsListing(synsListing *models.SynsMarketplaceListing, eventName string) (error)

	// @notice Transfer token from lister to buyer
	// 
	// @param listingId uint64
	// 
	// @param buyerAddr string
	// 
	// @return error
	TransferSyns721SuperToken(listingId uint64, buyerAddr string) (error)

	// @notice Get all Syns 721 Super Token
	// 
	// @return *[]models.Syns721SuperNFT
	// 
	// @return error
	GetAllSyns721SuperTokens() (*[]models.Syns721SuperNFT, error)

	// @notice Get all Syns 721 Super Token owned by an address
	// 
	// @param tokenOwner string
	// 
	// @return *[]models.Syns721SuperNFT
	// 
	// @return error
	GetAllSyns721SuperTokensOwnedBy(tokenOwner string) (*[]models.Syns721SuperNFT, error)

	// @notice Get single Syns 721 Super Token
	// 
	// @param assetContract string
	// 
	// @param tokenId string
	// 
	// @return *models.Syns721SuperNFT
	// 
	// @return error
	GetSyns721SuperTokenMetadata(assetContract, tokenId string) (*models.Syns721SuperNFT, error)
}
