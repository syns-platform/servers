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
	"Swyl/servers/swyl-community-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in comm-dao-impl
type CommDaoImpl struct {
	ctx					context.Context
	mongoCollection		*mongo.Collection
}

// @dev Constructor
func CommDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) CommDao {
	return &CommDaoImpl {
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}


// @notice Method of CommDaoImpl struct
// 
// @dev Lets a Swyl user create a community
// 
// @NOTE Should be fired off when #user/connect api is called
// 
// @param commOwner *string
// 
// @return error
func (ci *CommDaoImpl) CreateComm(commOwner *string) error {return nil}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets a Comm owned by commOwner
// 
// @param commOwner *string
// 
// @return *models.Community
// 
// @return error
func (ci *CommDaoImpl) GetCommOwnedBy(commOwner *string) (*models.Community, error) {return nil, nil}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets all Comm has ever created
// 
// @NOTE Might not be necessary
// 
// return *[]models.Community
func (ci *CommDaoImpl) GetAllComms() (*models.Community, error) {return nil, nil}


// @notice Method of CommDaoImpl struct
// 
// @dev Updates Comm's total_followers || Comm's total_posts
// 
// @param commOwner *string
// 
// @return error
func (ci *CommDaoImpl) UpdateCommOwnedBy(commOwner *string) error {return nil}