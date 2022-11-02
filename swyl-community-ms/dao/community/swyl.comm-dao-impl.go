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
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
// @param comm *models.Community
// 
// @return error
func (ci *CommDaoImpl) CreateComm(comm *models.Community) error {
	// find a community in the database using comm.comm_owner
	filter := bson.M{"community_owner": comm.Community_owner}
	dbRes := ci.mongoCollection.FindOne(ci.ctx, filter);

	// @logic: if dbRes == nil => a club with comm.community_owner has been created
	// @logic: if dbRes != nil => a club with comm.community_owner has NOT been created
	// @logic: else return dbRes
	if (dbRes.Err() == nil) {
		return errors.New("!MONGO - A community has already been created by this commOwner")
	} else if (dbRes.Err().Error() == "mongo: no documents in result") {
		_, err := ci.mongoCollection.InsertOne(ci.ctx, comm)
		return err
	} else {
		return dbRes.Err()
	}
}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets a Comm owned by commOwner
// 
// @param commOwner *string
// 
// @return *models.Community
// 
// @return error
func (ci *CommDaoImpl) GetCommOwnedBy(commOwner *string) (*models.Community, error) {
	// declare comm holder
	comm := &models.Community{}

	// prepare filter
	filter := bson.M{"community_owner": commOwner}

	// find comm in database
	if err := ci.mongoCollection.FindOne(ci.ctx, filter).Decode(comm); err != nil {return nil, err}
	
	// return OK
	return comm, nil
}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets all Comm has ever created
// 
// @NOTE Might not be necessary
// 
// return *[]models.Community
func (ci *CommDaoImpl) GetAllComms() (*[]models.Community, error) {
	// declare comms holder
	comms := &[]models.Community{}

	// get communities from internal database
	cursor, err := ci.mongoCollection.Find(ci.ctx, bson.M{})
	if err != nil {return nil, err}

	// decode cursor into declared comms
	if err := cursor.All(ci.ctx, comms); err != nil {return nil, err}
	
	// return OK
	return comms, nil
}


// @notice Method of CommDaoImpl struct
// 
// @dev Updates Comm's bio
// 
// @param commOwner *string
// 
// @param commBio *string
// 
// @return error
func (ci *CommDaoImpl) UpdateCommBioOwnedBy(commOwner *string, commBio *string) error {
	// prepare filter query
	filter := bson.M{"community_owner": commOwner}

	// prepare update query
	query := bson.D{{Key: "$set", Value: bson.M{"bio": commBio}}} 

	// update community in database
	dbRes, err := ci.mongoCollection.UpdateOne(ci.ctx, filter, query)
	if err != nil {return err}
	if dbRes.MatchedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}
	
	// return ok
	return nil
}

// @notice Method of CommDaoImpl struct
// 
// @notice Updates Comm's total_followers || Comm's total_posts
// 
// @param commOnwer *string
// 
// @param followerContext int16
// 
// @param postContext int16
// 
// @return error
func (ci *CommDaoImpl) UpdateCommTotalOwnedBy(commOwner *string, followerContext int16, postContext int16) error {
	// prepare filter query
	filter := bson.M{"community_owner": commOwner}

	// find the community with commOwner
	comm := &models.Community{}
	if err := ci.mongoCollection.FindOne(ci.ctx, filter).Decode(comm); err != nil {return err}

	// @logic if comm.total_followers == 0 && followerContext == -1, block
	// @logic if comm.total_posts == 0 && postContext == -1, block
	if ((comm.Total_followers == 0 && followerContext == -1) || (comm.Total_posts == 0 && postContext == -1)) {
		return errors.New("!OVERFLOW - total_followers *uint64 & total_posts *uint64 cannot be negative")
	}
	
	// prepare update query
	update := bson.D{
		{Key: "$set", Value: bson.M{"total_followers": comm.Total_followers + uint64(followerContext)}},
		{Key: "$set", Value: bson.M{"total_posts": comm.Total_posts + uint64(postContext)}},
	}

	// update community
	if _, err := ci.mongoCollection.UpdateOne(ci.ctx, filter, update); err != nil {return err}
	
	// return OK
	return nil
}


// @notice Method of CommDaoImpl struct
// 
// @dev Lets a Swyl user start following a community
// 
// @param follower *string
// 
// @param commOwner *string 
func (ci *CommDaoImpl) Follow(commOnwer *string, follower *string) error {return nil}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets a Swyl follower at followerId
// 
// @param followerId *primitive.ObjectID
// 
// @return *models.Follower
func (ci *CommDaoImpl) GetFollowerAt(followerId *primitive.ObjectID) (*models.Follower, error) {return nil, nil}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets all Swyl followers in a community own by commOwner
// 
// @param commOwner *string
// 
// @return *[]models.Follower
// 
// @return error
func (ci *CommDaoImpl) GetAllFollowersInCommOwnedBy(commOwner *string) (*[]models.Follower, error) {return nil, nil}


// @notice Method of CommDaoImpl struct
// 
// @dev Lets a Swyl user at followerId unfollows a community
// 
// @param followerId *primitive.ObjectID
// 
// @return error
func (ci *CommDaoImpl) Unfollow(followerId *primitive.ObjectID) error {return nil}
