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
	"errors"
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
		// return OK
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

		// return user and err
		return user, err
	} else {
		// return nil, and other error result from mongoDB
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
	// declare user placeholder
	var user *models.User = &models.User{}

	// set up find query
	query := bson.D{{Key: "wallet_address", Value: walletAddress}}

	// find the user in database using user.wallet_address
	if dbResult := ui.mongoCollection.FindOne(ui.ctx, query).Decode(user); dbResult != nil {return nil, dbResult}

	// return OK
	return user, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets all user.
// 
// @NOTE might not be necessary
// 
// @return []*models.User
func (ui *UserDaoImpl) GetAllUsers() (*[]models.User, error) {
	// Declare a slice of models.User
	var users []models.User

	// find users in database
	cursor, err := ui.mongoCollection.Find(ui.ctx, bson.D{})
	if err != nil {return nil, err}

	// decode cursor into a list of results
	if err = cursor.All(ui.ctx, &users); err != nil {return nil, err}

	// return OK
	return &users, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Updates a user.
// 
// @param *models.User
// 
// @return error
func (ui *UserDaoImpl) UpdateUser(user *models.User) error {
	// set up filter query
	filter := bson.M{"wallet_address": user.Wallet_address}

	// set up update query
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "username", Value: user.Username}}},
		{Key: "$set", Value: bson.D{{Key: "avatar", Value: user.Avatar}}},
		{Key: "$set", Value: bson.D{{Key: "display_name", Value: user.Display_name}}},
		{Key: "$set", Value: bson.D{{Key: "email", Value: user.Email}}},
		{Key: "$set", Value: bson.D{{Key: "bio", Value: user.Bio}}},
		{Key: "$set", Value: bson.D{{Key: "website", Value: user.Website}}},
		{Key: "$set", Value: bson.D{{Key: "social_media", Value: user.Social_media}}},
	}

	// update user 
	result, err := ui.mongoCollection.UpdateOne(ui.ctx, filter, update)
	if err != nil {return err}

	// return error if no document found with declared filter
	if result.MatchedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}

	// return OK
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
	// set up find query
	filter := bson.D{{Key: "wallet_address", Value: walletAddress}}

	// delete user from internal database
	result, err := ui.mongoCollection.DeleteOne(ui.ctx, filter)
	if (err != nil) {return err}

	// return error if no document found with declared filter
	if (result.DeletedCount == 0) {return errors.New("!MONGO - No matched document found with filter")}

	// response OK
	return nil
}