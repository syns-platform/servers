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
	// @param synsNFT SynsNFT
	// 
	// @return error
	MintNewSynsToken(synsNFT *models.SuperSyns721NFT) (error)
}
