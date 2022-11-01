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
func HandleException(e error) {
	if (e != nil) {
		log.Panic(e);
	}
}

// @dev Handle testing `wallet_address` to match ETH crypto wallet address
// 
// @param wallet_address *string
// 
// @return bool
func TestEthAddress(wallet_address *string) (bool, error) {
	pattern := os.Getenv("ETH_ADDRESS_REGEX")
	return regexp.MatchString(pattern, *wallet_address)
}
