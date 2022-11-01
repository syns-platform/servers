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
	"Swyl/servers/swyl-club-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in dao-impl
type TierDaoImpl struct {
   ctx               context.Context
   mongoCollection   *mongo.Collection
} 


// @dev Constructor
func TierDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) TierDao {
   return &TierDaoImpl {
      ctx: ctx,
      mongoCollection: mongoCollection,
   }
}

// @notice Method of TierDaoImpl struct
// 
// @dev Lets a club owner create a tier to internal database
// 
// @param tier *models.Tier
// 
// @return error
func (ti *TierDaoImpl) CreateTier(tier *models.Tier) error {
   return nil
}


// @notice Method of TierDaoImpl struct
// 
// @dev Gets a Tier at tierId and clubOwner
// 
// @param clubId *uint64
// 
// @param clubOwner *string
// 
// @return *models.Tier
// 
// @return error
func (ti *TierDaoImpl) GetTierAt(clubId *uint64, clubOwner *string) (*models.Tier, error) {
   return nil, nil
}


// @notice Method of TierDaoImpl struct
// 
// @dev Gets all tiers owned by clubOwner
// 
// @param clubOwner *string
// 
// @return *[]models.Tier
// 
// @return error
func (ti *TierDaoImpl) GetAllTiersOwnedBy(clubOwner *string) (*[]models.Tier, error) {
   return nil, nil
}


// @notice Method of TierDaoImpl struct
// 
// @dev Lets a clubOwner update a tier
// 
// @param tier *models.Tier
// 
// @return error
func (ti *TierDaoImpl) UpdateTier(tier *models.Tier) error {
   return nil
}


// @notice Lets a clubOwner delete a tier
// 
// @param tierId *uint64
// 
// @param clubOwner *string
// 
// @return error
func (ti *TierDaoImpl) DeleteTier(tierId *uint64, clubOwner *string) error {
   return nil
}