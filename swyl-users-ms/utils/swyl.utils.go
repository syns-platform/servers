/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package utils

// @import
import (
	"log"
	"os"
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @dev Loads environment variables
func LoadEnvVars() {
	err := godotenv.Load();
	HandleException(err);
}

// @dev Handdle error exception
//
// @param e error - the passed in error
func HandleException(e error) {if (e != nil) {log.Panic(e)}}

// @dev Handle testing `wallet_address` to match ETH crypto wallet address
// 
// @param wallet_address *string
// 
// @return bool
// 
// @return error
func TestEthAddress(wallet_address *string) (bool, error) {
	pattern := os.Getenv("ETH_ADDRESS_REGEX")
	return regexp.MatchString(pattern, *wallet_address)
}

// @dev Handle testing `signature` to match ETH signed signature
// 
// @param signature *string
// 
// @return bool
// 
// @return error
func TestSignature(signature *string) (bool, error) {
	pattern := os.Getenv("SIGNATURE_REGEX")
	return regexp.MatchString(pattern, *signature)
}

// @dev Sets up config for cors
// 
// @return gin.HandlerFunc
func SetupCorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: 		[]string{os.Getenv("CORS_ALLOW_ORIGIN")},
		AllowMethods:		[]string{"POST", "PATCH", "PUT", "DELETE", "GET"},
		AllowHeaders: 		[]string{"Origin"},	
		AllowCredentials: 	true,
		MaxAge: 			12*time.Hour,
	})
}
