/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package db

// @import
import (
	"Swyl/server/utils"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @dev Creates a MongoDB instance
//
// @return *mongo.Client
func EstablishMongoClient(ctx context.Context) *mongo.Client {
	// get the mongoDB uri
	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {log.Fatal("!MONGODB_URI - uri is not defined.")}
	
	// Establish the connection
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	utils.HandleException(err)

	// defer a call to `Disconnect()` after instantiating client
	defer func() {if err := mongoClient.Disconnect(ctx); err != nil {panic(err)}}()

	// return mongo client
	log.Println("MongoDB connected...")
	return mongoClient
}

// @dev Gets a mongodb collection
// 
// @param mongoClient *mongo.Client
//  
// @param collectionName string
// 
// @return *mongo.Collection
func GetMongoCollection(mongoClient *mongo.Client, collectionName string) *mongo.Collection {
	// get the collection
	collection := mongoClient.Database(os.Getenv("MONGO_DB")).Collection(collectionName)

	// return the collection
	return collection
}
