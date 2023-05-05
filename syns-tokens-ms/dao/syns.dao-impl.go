/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

import (
	"Syns/servers/syns-tokens-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @import

// @notice Root struct for other methods in dao-impl
type SynsTokenDaoImpl struct {
	ctx 			context.Context
	mongoCollection		*mongo.Collection
}



// @dev Syns Token Constructor
func SynsTokenDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) SynsTokenDao {
	return &SynsTokenDaoImpl{
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}

// @notice Method of SynsTokenDaoImpl struct
// 
// @dev Add token
// 
// @param walletAddress *string
// 
// @return *models.SynsNFT
// 
// @return error
func (st *SynsTokenDaoImpl) MintNewSynsToken(synsNFT *models.SynsNFT) (error) {
	return nil
}
