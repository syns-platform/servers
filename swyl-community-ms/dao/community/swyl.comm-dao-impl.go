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
// @dev Updates Comm's total_followers || Comm's total_posts
// 
// @param commOwner *string
// 
// @return error
func (ci *CommDaoImpl) UpdateCommOwnedBy(commOwner *string) error {return nil}