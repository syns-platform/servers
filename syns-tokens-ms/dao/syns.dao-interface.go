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

	// @notice Add a new Syns Token to database
	// 
	// @param synsNFT *models.Syns721SuperNFT
	// 
	// @return error
	MintNewSyns721Token(synsNFT *models.Syns721SuperNFT) (error)

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
}
