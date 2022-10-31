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
	"Swyl/servers/swyl-users-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in dao-impl
type UserDaoImpl struct {
	ctx 				context.Context
	mongoCollection		*mongo.Collection
}


// @dev Constructor
func UserDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) UserDao {
	return &UserDaoImpl{
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}


// @notice Method of UserDaoImpl struct
// 
// @dev Connects to an account stored in the internal database using wallet address.
// 		Create a new account on first connect.
// 
// @param user	*models.User
// 
// @return error
func (ui *UserDaoImpl) Connect(user *models.User) error {
	return nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a user at wallet address.
// 
// @param walletAddress *string
// 
// @return *models.User
// 
// @return error
func (ui *UserDaoImpl) GetUserAt(walletAddress *string) (*models.User, error) {
	return nil, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets all user.
// 
// @NOTE might not be necessary
// 
// @return []*models.User
func (ui *UserDaoImpl) GetAllUsers() []*models.User {
	return nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Updates a user.
// 
// @param *models.User
// 
// @return error
func (ui *UserDaoImpl) UpdateUser(*models.User) error {
	return nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Deletes a user at wallet address.
// 
// @param walletAddress *string
// 
// @return error
func (ui *UserDaoImpl) DeactivateUserAt(walletAddress *string) error {
	return nil
}