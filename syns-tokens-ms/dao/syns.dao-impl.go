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
	"Syns/servers/syns-tokens-ms/models"
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constants
var (
	MAX_SAFE_VALUE = uint64(9007199254740991); // 9007199254740991 = 2^53 - 1 which is the safe value for Typescript in the client app
)

// @notice Root struct for other methods in dao-impl
type Syns721TokenDaoImpl struct {
	ctx 			context.Context
	mongoCollection		*mongo.Collection
}



// @dev Syns Token Constructor
func Syns721TokenDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) Syns721TokenDao {
	return &Syns721TokenDaoImpl{
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}

// @notice Method of Syns721TokenDaoImpl struct
// 
// @dev Add Syns 721 super token to database
// 
// @param synsNFT *models.Syns721SuperNFT
// 
// @return error
func (sti *Syns721TokenDaoImpl) MintNewSyns721Token(synsNFT *models.Syns721SuperNFT) (error) {
	// marshal synsNFT to byte slice
	synsNFTBytes, _ := json.Marshal(synsNFT)

	// calculate new token hash string
	tokenHash := crypto.Keccak256Hash(synsNFTBytes).Hex()

	// update tokenHash field in synsNFT
	synsNFT.TokenHash = tokenHash


	// make sure there is no duplicated hash in the database
	query := bson.D{{Key: "token_hash", Value: tokenHash}}
	dbRes := sti.mongoCollection.FindOne(sti.ctx, query).Decode(synsNFT)

	// @logic: if dbRes.Error() contains "no documents..." => no token with the same hash has been added to the database
	// @logic: if dbRes.Error() do not contains "no documents..." => a token with the same hash has been added to the database => abort!
	if (dbRes != nil && strings.Contains(dbRes.Error(), "no documents in result")) {
		// inject new Syns  Token to internal database
		_, err := sti.mongoCollection.InsertOne(sti.ctx, synsNFT)
		return err
	} else {
		return errors.New("duplicated token")
	}
}

// @notice Method of Syns721TokenDaoImpl struct
// 
// @dev Update Syns 721 Super Token using Syns Listing from Syns Marketplace
// 
// @param synsListing *models.SynsMarketplaceListing
// 
// @param eventName string
// 
// @return error
func (sti *Syns721TokenDaoImpl) UpdatedSyns721SuperTokenBySynsListing(synsListing *models.SynsMarketplaceListing, eventName string) (error) {
	// prepare Syns super token
	syns721SuperToken := &models.Syns721SuperNFT{}

	// prepare filter query
	filter := bson.D{}

	// look up token in database
	if (strings.Compare(eventName, "ListingAdded") == 0) {
		filter = bson.D{{Key: "token_id", Value: synsListing.TokenId.Uint64()}}
	} else if (strings.Compare(eventName, "ListingRemoved") == 0) {
		filter = bson.D{{Key: "listing_id", Value: synsListing.ListingId.Uint64()}}
	}

	// prepare buyoutPice 
	ether := new(big.Float)
	ether.SetString(synsListing.BuyoutPricePerToken.String())
	ethValue := ether.Quo(ether, big.NewFloat(1e18))

	// Set the options to return the updated document
	mongoOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// prepare update query based on eventName
	if (strings.Compare(eventName, "ListingAdded") == 0) {
		update := bson.D{
			{Key: "$set", Value: bson.D{{Key: "is_listing", Value: true}}},
			{Key: "$set", Value: bson.D{{Key: "listing_id", Value: synsListing.ListingId.Uint64()}}},
			{Key: "$set", Value: bson.D{{Key: "lister", Value: synsListing.TokenOwner.Hex()}}},
			{Key: "$set", Value: bson.D{{Key: "start_sale", Value: synsListing.StartSale.Uint64()}}},
			{Key: "$set", Value: bson.D{{Key: "currency", Value: synsListing.Currency.Hex()}}},
			{Key: "$set", Value: bson.D{{Key: "buyout_price_per_token", Value: ethValue.String()}}},
		}
	
		// update updated syns721SuperToken in internal database
		if dbRes := sti.mongoCollection.FindOneAndUpdate(sti.ctx, filter, update, mongoOptions).Decode(syns721SuperToken); dbRes != nil {
			return dbRes
		}
	} else if (strings.Compare(eventName, "ListingRemoved") == 0) {
		update := bson.D{
			{Key: "$set", Value: bson.D{{Key: "is_listing", Value: false}}},
			{Key: "$set", Value: bson.D{{Key: "listing_id", Value: MAX_SAFE_VALUE}}},
			{Key: "$set", Value: bson.D{{Key: "lister", Value: synsListing.TokenOwner.Hex()}}},
			{Key: "$set", Value: bson.D{{Key: "start_sale", Value: synsListing.StartSale.Uint64()}}},
			{Key: "$set", Value: bson.D{{Key: "currency", Value: synsListing.Currency.Hex()}}},
			{Key: "$set", Value: bson.D{{Key: "buyout_price_per_token", Value: ethValue.String()}}},
		}
	
		// update updated syns721SuperToken in internal database
		if dbRes := sti.mongoCollection.FindOneAndUpdate(sti.ctx, filter, update, mongoOptions).Decode(syns721SuperToken); dbRes != nil {
			return dbRes
		}
	}

	// record old token hash
	oldTokenHash := syns721SuperToken.TokenHash

	// marshal updated syns721SuperToken to byte slice
	synsNFTBytes, _ := json.Marshal(syns721SuperToken)

	// calculate new token hash string
	newTokenHash := crypto.Keccak256Hash(synsNFTBytes).Hex()

	// update filter query based on old token hash
	filter = bson.D{{Key: "token_hash", Value: oldTokenHash}}

	// update update query to update token_hash
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token_hash", Value: newTokenHash}}}}

	// update record with new token_hash
	if dbRes := sti.mongoCollection.FindOneAndUpdate(sti.ctx, filter, update); dbRes.Err() != nil {
		return dbRes.Err()
	}

	return nil
}

// @notice Transfer token from lister to buyer
// 
// @param listingId uint64
// 
// @param buyerAddr string
// 
// @return error
func (sti *Syns721TokenDaoImpl) TransferSyns721SuperToken(listingId uint64, buyerAddr string) (error) {
	// prepare Syns721SuperToken
	syns721SuperToken := &models.Syns721SuperNFT{}

	// prepare filter query
	filter := bson.D{{Key: "listing_id", Value: listingId}}

	// prepare update query
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "token_owner", Value: buyerAddr}}},
		{Key: "$set", Value: bson.D{{Key: "is_listing", Value: false}}},
		{Key: "$set", Value: bson.D{{Key: "listing_id", Value: MAX_SAFE_VALUE}}},
		{Key: "$set", Value: bson.D{{Key: "lister", Value: buyerAddr}}},
		{Key: "$set", Value: bson.D{{Key: "start_sale", Value: 0}}},
		{Key: "$set", Value: bson.D{{Key: "currency", Value: common.Address{}.Hex()}}},
		{Key: "$set", Value: bson.D{{Key: "buyout_price_per_token", Value: "0.0"}}},
	}

	// Set the options to return the updated document
	mongoOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// update record and decode the record to syns721SuperToken
	if dbRes := sti.mongoCollection.FindOneAndUpdate(sti.ctx, filter, update, mongoOptions).Decode(syns721SuperToken); dbRes != nil {
		return dbRes
	}

	// record old token hash
	oldTokenHash := syns721SuperToken.TokenHash

	// marshal updated syns721SuperToken to byte slice
	synsNFTBytes, _ := json.Marshal(syns721SuperToken)

	// calculate new token hash string
	newTokenHash := crypto.Keccak256Hash(synsNFTBytes).Hex()

	// update filter query based on old token hash
	filter = bson.D{{Key: "token_hash", Value: oldTokenHash}}

	// update update query to update token_hash
	update = bson.D{{Key: "$set", Value: bson.D{{Key: "token_hash", Value: newTokenHash}}}}

	// update record with new token_hash
	if dbRes := sti.mongoCollection.FindOneAndUpdate(sti.ctx, filter, update); dbRes.Err() != nil {
		return dbRes.Err()
	}

	return nil
}

// @notice Get all Syns 721 Super Token
// 
// @return *[]models.Syns721SuperNFT
// 
// @return error
func (sti *Syns721TokenDaoImpl) GetAllSyns721SuperTokens() (*[]models.Syns721SuperNFT, error) {
	// prepare tokens placeholder
	var syns721SuperTokens []models.Syns721SuperNFT

	// fetch all tokens in database
	cursor, dbRes := sti.mongoCollection.Find(sti.ctx, bson.D{})
	if dbRes != nil {return nil, dbRes}

	// decode cursor into declared token placeholder
	if decodedErr := cursor.All(sti.ctx, &syns721SuperTokens); decodedErr != nil {return nil, decodedErr}

	// return OK
	return &syns721SuperTokens, nil
}

// @notice Get all Syns 721 Super Token owned by an address
// 
// @param tokenOwner string
// 
// @return *[]models.Syns721SuperNFT
// 
// @return error
func (sti *Syns721TokenDaoImpl) GetAllSyns721SuperTokensOwnedBy(tokenOwner string) (*[]models.Syns721SuperNFT, error) {
	// prepare tokens placeholder
	var syns721SuperTokens []models.Syns721SuperNFT

	// prepare find query
	query := bson.D{{Key: "token_owner", Value: primitive.Regex{Pattern: "^" + tokenOwner + "$", Options: "i"}}}
	

	// find tokens owned my tokenOwner
	cursor, dbRes := sti.mongoCollection.Find(sti.ctx, query); 
	if dbRes != nil {return nil, dbRes}

	// decode cursor to the placeholder
	if decodedErr := cursor.All(sti.ctx, &syns721SuperTokens); decodedErr != nil {return nil, decodedErr}

	// return OK
	return &syns721SuperTokens, nil
}