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

// @notice SynsTokenDao interface
type SynsTokenDao interface {

	// @notice Add a new Syns Token to database
	// 
	// @param synsNFT SynsNFT
	// 
	// @return error
	MintNewSynsToken(synsNFT *models.SynsNFT) (error)
}
