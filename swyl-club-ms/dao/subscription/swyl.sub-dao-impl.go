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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in sub-dao-impl
type SubDaoImpl struct {
	ctx			 		context.Context
	mongoCollection 	*mongo.Collection
}

// @dev Constructor
func SubDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) SubDao {
	return &SubDaoImpl{
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}

// @notice Method of SubDaoImpl struct
// 
// @dev Lets a user subscribe to a tier
// 
// @param sub *models.Subscription
// 
// @return error
func (si *SubDaoImpl) Subscribe(sub *models.Subscription) error{
	// updated sub.Subscription_ID
	sub.Subscription_ID = primitive.NewObjectID()

	// insert subscription to internal database
	if _, err := si.mongoCollection.InsertOne(si.ctx, sub); err != nil {return err}

	// return OK
	return nil
}


// @notice Method of SubDaoImpl struct
//
// @dev Gets a subscription using subscription_id
// 
// @param subId *string
// 
// @return *models.Subscription
func (si *SubDaoImpl) GetSubscriptionAt(subId *string) (*models.Subscription, error){
	// prepare sub placeholder
	sub := &models.Subscription{}

	// set up ObjectId
	objectId, err := primitive.ObjectIDFromHex(*subId)
	if err != nil {return nil, err}

	// prepare filter
	filter := bson.M{"_id": objectId}

	// find the sub in database
	if err := si.mongoCollection.FindOne(si.ctx, filter).Decode(sub); err != nil {return nil, err}

	// return OK
	return sub, nil
}


// @notice Method of SubDaoImpl struct
//
// @dev Gets all subscriptions onwed at tier_id and by club_owner
// 
// @param tierId *string
// 
// @param clubOwner *string
// 
// @return *[]models.Subscription
// 
// @return error
func (si *SubDaoImpl) GetAllSubsAt(tierId *string, clubOwner *string) (*[]models.Subscription, error) {return nil, nil}


// @notice Method of SubDaoImpl struct
//
// @dev Updates a subscription status
// 
// @param subId *string
// 
// @return error
func (si *SubDaoImpl) UpdateSubStatus(subIb *string) error {return nil}


// @notice Method of SubDaoImpl struct
//
// @dev Lets a subscriber unsubscribe a tier 
// 
// @param tierId *string
// 
// @param subId *string
// 
// @return error
func (si *SubDaoImpl) Unsubscribe(tierId *string, subId *string) error {return nil}