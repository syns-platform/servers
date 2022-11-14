/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @dev Examine access token in headers
//
// @return gin.HandlerFunc
func Authenticate() gin.HandlerFunc {
	return func (gc *gin.Context) {
            // prepare custome claim for decoding
            type SwylClaims struct {
                  Signer                  string `json:"signer"`
                  Signature               string `json:"signature"`
                  LoginMessage            string `json:"loginMessage"`
                  jwt.StandardClaims
            }

            // Get bearer token from Authorization headers
            bearerToken := gc.GetHeader("Authorization")
            if bearerToken == "" {gc.AbortWithStatusJSON(401, gin.H{"error": "!BEARER_TOKEN - no authorization header found"});return;}

            // extra the jwt from bearer token
            accessToken := strings.Split(bearerToken, " ")[1]

            // check if accessToken is not empty
            if accessToken == "" {gc.AbortWithStatusJSON(401, gin.H{"error": "!ACCESS_TOKEN - empty authorization token"})}

            // Decode/validate accessToken
            token, err := jwt.ParseWithClaims(accessToken, &SwylClaims{}, func(token *jwt.Token) (interface{}, error) {
               return []byte(os.Getenv("JWT_SECRET_KEY")), nil
            })
            
            // Implement authenticating logic
            if claims, ok := token.Claims.(*SwylClaims); ok && token.Valid {
                  // prepare signature & loginMessage as bytes
                  byteSignature := hexutil.MustDecode(claims.Signature) // decode signature from ahex string with 0x prefix to []byte
                  byteLoginMessage := []byte(claims.LoginMessage)

                  // calculate a hash for byteLoginMessage which can be safely used to calculate a signature from
                  hashLoginMessage := accounts.TextHash(byteLoginMessage)

                  if byteSignature[crypto.RecoveryIDOffset] == 27 || byteSignature[crypto.RecoveryIDOffset] == 28 {
                        byteSignature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
                  }
                  
                  // verify the signature to find the pubKey created the signature
                  ecdsaPubKey, err := crypto.SigToPub(hashLoginMessage, byteSignature)
                  if err != nil {gc.AbortWithStatusJSON(401, gin.H{"error": err.Error()}); return;}

                  // convert edcsa pubKey to common eth pubKey
                  pubKeyAddress := crypto.PubkeyToAddress(*ecdsaPubKey)
                  
                  // compare pubKeyAddress to claims.Signer, if matched => pass authentication and vice verca
                  if matched := strings.EqualFold(claims.Signer, pubKeyAddress.Hex()); !matched {
                        gc.AbortWithStatusJSON(401, gin.H{"error": "!PUBLIC_KEY - jwt.payload.Signer do not match the public key address recovered from verifying jwt.payload.Signature & jwt.payload.LoginMessage "}); 
                        return;
                  }

                  // pass the signer to the controller handler function
                  gc.Set("Signer", claims.Signer)

                  // move on
                  gc.Next()
            } else {
               gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()}); return;
            }
		
	}
}