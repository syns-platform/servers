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
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in sub-dao-impl
type SubDaoImpl struct {
	ctx			context.Context
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
func (si *SubDaoImpl) GetAllSubsAt(tierId *string, clubOwner *string) (*[]models.Subscription, error) {
	// Declare subs holder
	subs := &[]models.Subscription{}

	// set up tierObjectId 
	tierObjectId, err := primitive.ObjectIDFromHex(*tierId)
	if err != nil {return nil, err}

	// prepare filter
	filter := bson.D{{Key: "tier_id", Value: tierObjectId}, {Key: "club_owner", Value: clubOwner}}

	// get subs from internal database
	cursor, err := si.mongoCollection.Find(si.ctx, filter)
	if err != nil {return nil, err}

	// decode cursor into declared subs
	if err := cursor.All(si.ctx, subs); err != nil {return nil ,err}

	// return OK
	return subs, nil
}


// @notice Method of SubDaoImpl struct
//
// @dev Toggles a subscription status at subId
// 
// @param subId *string
// 
// @return error
func (si *SubDaoImpl) ToggleSubStatusAt(subId *string) error {
	// set up objectId
	objectId, err := primitive.ObjectIDFromHex(*subId)
	if err != nil {return err}

	// prepare filter query
	filter := bson.M{"_id": objectId}

	// find sub at subId
	sub := &models.Subscription{}
	if err := si.mongoCollection.FindOne(si.ctx, filter).Decode(sub); err != nil {return err}

	// prepare update query
	update := bson.D{{Key: "$set", Value: bson.M{"status": !sub.Status}}}

	// update the sub
	if _, err := si.mongoCollection.UpdateOne(si.ctx, filter, update); err != nil {return err}
	
	// response OK
	return nil
}


// @notice Method of SubDaoImpl struct
//
// @dev Lets a subscriber unsubscribe a tier 
// 
// @param subId *string
// 
// @return error
func (si *SubDaoImpl) Unsubscribe(subId *string) error {
	// set up objectId
	objectId, err := primitive.ObjectIDFromHex(*subId)
	if err != nil {return err}

	// prepare filter 
	filter := bson.M{"_id": objectId}

	// delete subscription in the database
	dbRes, err := si.mongoCollection.DeleteOne(si.ctx, filter)
	if err != nil {return err}
	if dbRes.DeletedCount == 0 {return errors.New("!MONGO - No matched document found")}

	// return OK
	return nil
}