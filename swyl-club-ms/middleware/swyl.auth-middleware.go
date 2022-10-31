/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package middleware

import "github.com/gin-gonic/gin"

// @dev Examine access token in headers
//
// @return gin.HandlerFunc
func Authenticate() gin.HandlerFunc {
	return func (gc *gin.Context) {
		// @TODO examine jwt token
	}
}