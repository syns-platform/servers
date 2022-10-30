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
		log.Fatal(e);
	}
}