/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @dev Examine access token in headers
//
// @return gin.HandlerFunc
func Authenticate() gin.HandlerFunc {
	return func (gc *gin.Context) {

            // Get bearer token from Authorization headers
            bearerToken := gc.GetHeader("Authorization")

            accessToken := strings.Split(bearerToken, " ")[1]

            // Decode/validate accessToken
            token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
               if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                  gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "!ACCESS_TOKEN - cannot parse token"}); return nil, nil;
               }

               return []byte(os.Getenv("JWT_SECRET_KEY")), nil
            })
            
            // Implement authenticating logic
            if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
               log.Println(claims["signature"])
               log.Println(claims["loginMessage"])
               log.Println(claims["signer"])
            } else {
               gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()}); return;
            }
		
	}
}