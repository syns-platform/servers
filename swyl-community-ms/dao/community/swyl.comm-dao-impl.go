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
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in comm-dao-impl
type CommDaoImpl struct {
	ctx					context.Context
	commCollection		*mongo.Collection
	followerCollection 	*mongo.Collection
}

// @dev Constructor
func CommDaoConstructor(ctx context.Context, commCollection *mongo.Collection, followerCollection *mongo.Collection) CommDao {
	return &CommDaoImpl {
		ctx: ctx,
		commCollection: commCollection,
		followerCollection: followerCollection,
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
	dbRes := ci.commCollection.FindOne(ci.ctx, filter);

	// @logic: if dbRes == nil => a club with comm.community_owner has been created
	// @logic: if dbRes != nil => a club with comm.community_owner has NOT been created
	// @logic: else return dbRes
	if (dbRes.Err() == nil) {
		return errors.New("!MONGO - A community has already been created by this commOwner")
	} else if (dbRes.Err().Error() == "mongo: no documents in result") {
		_, err := ci.commCollection.InsertOne(ci.ctx, comm)
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
	if err := ci.commCollection.FindOne(ci.ctx, filter).Decode(comm); err != nil {return nil, err}
	
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
	cursor, err := ci.commCollection.Find(ci.ctx, bson.M{})
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
	dbRes, err := ci.commCollection.UpdateOne(ci.ctx, filter, query)
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
	if err := ci.commCollection.FindOne(ci.ctx, filter).Decode(comm); err != nil {return err}

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
	if _, err := ci.commCollection.UpdateOne(ci.ctx, filter, update); err != nil {return err}
	
	// return OK
	return nil
}


// @notice Method of CommDaoImpl struct
// 
// @logic if a user has already followed the community => unfollow
// 
// @logic if a user does NOT followed the community => follow
// 
// @dev Lets a Swyl user start following a community
// 
// @param follower *models.Follower
func (ci *CommDaoImpl) ToggleFollow(follower *models.Follower) (int, error) {
	// Set up filter
	filter := bson.D{
		{Key: "community_owner", Value: follower.Community_owner},
		{Key: "follower", Value: follower.Follower},
	}

	// find follower in database
	dbRes := ci.followerCollection.FindOne(ci.ctx, filter)

	// @logic: if dbRes.Err() == nil => a document with follower.Commynity_owner & follower.Follower is found => followed => unfollow
	// @logic: if dbRes.Err() == nil => a document with follower.Commynity_owner & follower.Follower is NOT found => unfollowed => follow
	if dbRes.Err() == nil {
		// delete follower document
		_, err := ci.followerCollection.DeleteOne(ci.ctx, filter)
		return 0, err
	} else if dbRes.Err().Error() == "mongo: no documents in result" {
		// set up follower_id
		follower.Follower_ID = primitive.NewObjectID()

		// set up follow_at
		follower.Follow_at = uint64(time.Now().Unix())

		// insert new follower to internal database
		if _, err := ci.followerCollection.InsertOne(ci.ctx, follower); err != nil {return -1, err}

		// return OK
		return 1, nil
	} else {
		// return err
		return -1, dbRes.Err()
	}
}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets a Swyl follower at followerId
// 
// @param followerId *string
// 
// @return *models.Follower
func (ci *CommDaoImpl) GetFollowerAt(followerId *string) (*models.Follower, error) {
	// declare follower placeholder
	follower := &models.Follower{}

	// prepare objectId
	objectId, err := primitive.ObjectIDFromHex(*followerId)
	if err != nil {return nil, err}

	// prepare filter query
	qeury := bson.M{"_id": objectId}

	// find follower in database
	if err := ci.followerCollection.FindOne(ci.ctx, qeury).Decode(follower); err != nil {return nil, err}
	
	// return OK
	return follower, nil
}


// @notice Method of CommDaoImpl struct
// 
// @dev Gets all Swyl followers in a community own by commOwner
// 
// @param commOwner *string
// 
// @return *[]models.Follower
// 
// @return error
func (ci *CommDaoImpl) GetAllFollowersInCommOwnedBy(commOwner *string) (*[]models.Follower, error) {
	// declare followers placeholder
	followers := &[]models.Follower{}

	// prepare filter query
	filter := bson.M{"community_owner": commOwner}

	// find followers
	cursor, err := ci.followerCollection.Find(ci.ctx, filter)
	if err != nil {return nil, err}

	// decode cursor to 
	if err := cursor.All(ci.ctx, followers); err != nil {return nil, err}

	// return OK
	return followers, nil
}