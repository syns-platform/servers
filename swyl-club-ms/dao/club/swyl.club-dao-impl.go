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
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in club-dao-impl
type ClubDaoImpl struct {
	ctx 				context.Context
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
// @return error
func (ci *ClubDaoImpl) CreateClub(clubOwner *string) error {
	// prepare club placeholder
	// club := &models.Club{}

	// set up find query
	query := bson.M{"club_owner": clubOwner}

	// check if club has already been created with the `clubOwner`
	dbRes := ci.mongCollection.FindOne(ci.ctx, query)

	// logic: if dbRes error != nil => club with `clubOwner` has never been created
	// logic: if dbRes error == nil => club with `clubOwner` has already been created
	if dbRes.Err() == nil {
		return errors.New("!MONGO - A club has already been created by this clubOwner")
	} else if dbRes.Err().Error() == "mongo: no documents in result" {
		// prepare newClub
		newClub := &models.Club{
				Club_owner: clubOwner,
				Created_at: uint64(time.Now().Unix()),
				Total_member: uint64(0),
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
	return nil, nil
}


// @notice Method of UserDaoImpl struct
// 
// @dev Updates a Club
// 
// @param clubOwner *string
// 
// @param totalMember *uint64
func (ci *ClubDaoImpl) UpdateClub(clubOwner *string, totalMember *uint64) error {
	return nil
}