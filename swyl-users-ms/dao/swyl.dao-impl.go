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
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
// @return *models.User
// 
// @return error
// 
// @TODO should return JWT token
func (ui *UserDaoImpl) Connect(walletAddress *string) (*models.User, error) {
	// declare user placeholder
	var user *models.User = &models.User{}

	// set up find query
	query := bson.D{{Key: "wallet_address", Value: walletAddress}}
	
	// find the user in database using user.wallet_address
	dbResult := ui.mongoCollection.FindOne(ui.ctx, query).Decode(user)

	// logic: if dbResult error == nil => user with `walletAddress` has already connected before
	// logic: if dbResult error != nil => user with `walletAddress` has never connected before
	if (dbResult == nil) {
		return user, nil
	} else if dbResult.Error() == "mongo: no documents in result" {
		// prepare user
		user = &models.User{
			Wallet_address: walletAddress,
			Username: walletAddress,
			Joined_at: time.Now().Unix(),
		}

		// insert new user to internal database
		_, err := ui.mongoCollection.InsertOne(ui.ctx, user)

		// return
		return user, err
	} else {
		return nil, dbResult
	}
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