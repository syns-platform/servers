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

	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	// look up token in database
	query := bson.D{{Key: "token_id", Value: int(synsListing.TokenId.Int64())}}
	if dbRes := sti.mongoCollection.FindOne(sti.ctx, query).Decode(syns721SuperToken); dbRes != nil {
		return dbRes
	}

	// prepare currentTokenHash
	currentTokenHash := syns721SuperToken.TokenHash

	// @logic: if eventName == "ListingAdded", update new listing
	// @logic: if eventName == "listingRemoved", remove listing
	if (strings.Compare(eventName, "ListingAdded") == 0) {
		// prepare buyoutPice 
		ether := new(big.Float)
		ether.SetString(synsListing.BuyoutPricePerToken.String())
		ethValue := ether.Quo(ether, big.NewFloat(1e18))


		// update listing information in syns721SuperToken
		syns721SuperToken.IsListing = true
		syns721SuperToken.ListingID = int(synsListing.ListingId.Int64())
		syns721SuperToken.Lister = synsListing.TokenOwner.Hex()
		syns721SuperToken.StartSale = synsListing.StartSale.Uint64()
		syns721SuperToken.Currency = synsListing.Currency.Hex()
		syns721SuperToken.BuyouPricePerToken = ethValue.String()


		// marshal updated syns721SuperToken to byte slice
		synsNFTBytes, _ := json.Marshal(syns721SuperToken)

		// calculate new token hash string
		newTokenHash := crypto.Keccak256Hash(synsNFTBytes).Hex()

		// update tokenHash
		syns721SuperToken.TokenHash = newTokenHash

		// prepare filter query
		filter := bson.D{{Key: "token_hash", Value: currentTokenHash}}

		// prepare update query
		update := bson.D{
			{Key: "$set", Value: bson.D{{Key: "token_hash", Value: newTokenHash}}},
			{Key: "$set", Value: bson.D{{Key: "is_listing", Value: true}}},
			{Key: "$set", Value: bson.D{{Key: "listing_id", Value: int(synsListing.ListingId.Int64())}}},
			{Key: "$set", Value: bson.D{{Key: "lister", Value: synsListing.TokenOwner.Hex()}}},
			{Key: "$set", Value: bson.D{{Key: "start_sale", Value: synsListing.StartSale.Uint64()}}},
			{Key: "$set", Value: bson.D{{Key: "currency", Value: synsListing.Currency.Hex()}}},
			{Key: "$set", Value: bson.D{{Key: "buyout_price_per_token", Value: ethValue.String()}}},
		}

		// update updated syns721SuperToken in internal database
		if dbRes := sti.mongoCollection.FindOneAndUpdate(sti.ctx, filter, update); dbRes.Err() != nil {
			return dbRes.Err()
		}
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
	query := bson.D{{Key: "token_owner", Value: strings.ToLower(tokenOwner)}}

	// find tokens owned my tokenOwner
	cursor, dbRes := sti.mongoCollection.Find(sti.ctx, query); 
	if dbRes != nil {return nil, dbRes}

	// decode cursor to the placeholder
	if decodedErr := cursor.All(sti.ctx, &syns721SuperTokens); decodedErr != nil {return nil, decodedErr}

	// return OK
	return &syns721SuperTokens, nil
}