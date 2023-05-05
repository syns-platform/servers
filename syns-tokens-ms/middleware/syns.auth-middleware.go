/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// @dev Examne Syns API Key in headers
//
// @return gin.HandleFunc
func AuthenticateSynsAPIKey() gin.HandlerFunc {
      return func (gc *gin.Context) {
            // Get bearer token from Authorization headers
            bearerToken := gc.GetHeader("Authorization")
            if bearerToken == "" {gc.AbortWithStatusJSON(401, gin.H{"error": "!BEARER_TOKEN - no authorization header found"});return;}

            // extract the API_KEY from bearer token
            clientAPIKey := strings.Split(bearerToken, " ")[1]

            // check if accessToken is not empty
            if clientAPIKey == "" {gc.AbortWithStatusJSON(401, gin.H{"error": "!ACCESS_TOKEN - empty authorization token"})}

            // compare clientAPIKey with the API Key from server
            // if don't match, abort request with 401
            if strings.Compare(clientAPIKey, os.Getenv("SYNS_PLATFORM_API_KEY")) != 0 {
                  gc.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
            } 
            
            // move on
            gc.Next()
      }
}