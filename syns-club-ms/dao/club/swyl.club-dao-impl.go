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

// @notice Root struct for other methods in club-dao-impl
type ClubDaoImpl struct {
	ctx 			context.Context
	mongCollection		*mongo.Collection
}

// @dev Constructor
func ClubDaoConstructor (ctx context.Context, mongCollection *mongo.Collection) ClubDao {
	return &ClubDaoImpl{
		ctx: ctx,
		mongCollection: mongCollection,
	}
}


// @notice Method of UserDaoImpl struct
// 
// @dev Creates a club
// 
// @param clubOwner *string
// 
// @param createdAt uint64
// 
// @return error
func (ci *ClubDaoImpl) CreateClub(clubOwner *string, createdAt uint64) error {
	// set up find query
	query := bson.M{"club_owner": clubOwner}

	// check if club has already been created with the `clubOwner`
	dbRes := ci.mongCollection.FindOne(ci.ctx, query)

	// @logic: if dbRes error != nil => club with `clubOwner` has never been created
	// @logic: if dbRes error == nil => club with `clubOwner` has already been created
	// @logic: else return dbRes
	if dbRes.Err() == nil {
		return errors.New("!MONGO - A club has already been created by this clubOwner")
	} else if dbRes.Err().Error() == "mongo: no documents in result" {
		// prepare newClub
		newClub := &models.Club{
				Club_owner: clubOwner,
				Created_at: createdAt,
				Total_members: uint64(0),
			}

		// insert newClub to internal database
		_, err := ci.mongCollection.InsertOne(ci.ctx, newClub)
	
		// return
		return err 
	} else {
		return dbRes.Err()
	}
}


// @notice Method of UserDaoImpl struct
// 
// @dev Gets a club owned by clubOwner
// 
// @param clubOwner *string
// 
// @return *models.Club
// 
// @return error
func (ci *ClubDaoImpl) GetClubOwnedBy(clubOwner *string) (*models.Club, error) {
	// declare club placeholder
	club := &models.Club{}

	// set up find query
	query := bson.M{"club_owner": clubOwner}

	// find the club in database using clubOwner
	if dbRes := ci.mongCollection.FindOne(ci.ctx, query).Decode(club); dbRes != nil {return nil, dbRes}

	// return OK
	return club, nil
}

// @notice Method of UserDaoImpl struct
// 
// @dev Gets a slice of all the club
// 
// @return *[]models.Club
// 
// @return error
func (ci *ClubDaoImpl) GetAllClubs() (*[]models.Club, error) {
	// Declare a slice of placeholder models.User 
	clubs := &[]models.Club{}

	// Find clubs in database
	cursor, err := ci.mongCollection.Find(ci.ctx, bson.D{})
	if err != nil {return nil, err}

	// decode cursor into the declared slice
	if err := cursor.All(ci.ctx, clubs); err != nil {return nil, err}

	// return OK
	return clubs, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Increase/decrease a Club's total members
// 
// @param clubOwner *string
// 
// @param context *uint16 0=decrease || 1=increase
func (ci *ClubDaoImpl) UpdateClub(clubOwner *string, context *uint16) error {
	// set up filter query
	filter := bson.M{"club_owner": clubOwner}

	// get club from database
	club := &models.Club{}
	if err := ci.mongCollection.FindOne(ci.ctx, filter).Decode(club); err != nil {return err}

	// examine club.Total_members
	if (club.Total_members == 0 && *context == 0) {return errors.New("!TOTAL_MEMBERS - Cannot decrease total members")}

	// declare update query
	var update primitive.D

	// set up update query
	if *context == 0 {
		update = bson.D {
			{Key: "$inc", Value: bson.D{{Key: "total_members", Value: -1}}},
		}
	} else {
		update = bson.D {
			{Key: "$inc", Value: bson.D{{Key: "total_members", Value: 1}}},
		}
	}

	// update Club
	updateRes, err := ci.mongCollection.UpdateOne(ci.ctx, filter, update)
	if err != nil {return err}

	// return error if no document found with declared filter
	if updateRes.MatchedCount == 0 {return errors.New("!MONGO - No matched document found with filter")}

	// return OK
	return nil
}