/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

import (
	"Swyl/servers/swyl-users-ms/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @notice global var
// var validate = validator.New()

// @route `POST/generate-access-token`
//
// @dev Generate a new JWT from the payload passed in from http.body
//
// @param gc *gin.Context
func GenerateAccessToken(gc *gin.Context) {
	// prepare payloadParamHolder
	type payloadParamHolder struct {
		PubKey string `json:"pubKey" validate:"required,len=42,alphanum"`
		Signature string `json:"signature" validate:"required,len=132,alphanum"`
		MessageToSign string `json:"messageToSign" validate:"required"`
	}

	// declare param
	var param *payloadParamHolder

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}

	// test param.pubKey to match ETH Crypto wallet address convention
	pubKeyMatched, err := utils.TestEthAddress(&param.PubKey)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test param.pubKey against regex"}); return;}
	if !pubKeyMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - param.pubKey is not an ETH crypto wallet address"}); return;}

	// test param.Signature to see if it's a valid ethereum signed signature
	signedMsgMatched, err := utils.TestSignature(&param.Signature)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test param.signature against regex"}); return;}
	if !signedMsgMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!SIGNATURE - param.signature is not an ETH cyphertext"}); return;}

	// extra validatation on struct param
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}


	// Create a new token object
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pubKey": param.PubKey,
		"signature": param.Signature,
		"messageToSign": param.MessageToSign,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		gc.AbortWithStatusJSON(500, gin.H{"error": err.Error()}); return;
	}

	// send the JWT back in form of httpOnly cookie
	// @TODO: Figure out how to access httpOnly from other microservices
	gc.SetSameSite(http.SameSiteNoneMode)
	gc.SetCookie("Authorization", accessToken, 3600*24, "", "", false, true)

	// gc.JSON(200, gin.H{"msg": "JWT successfully generated"})
	gc.JSON(200, gin.H{"token": accessToken})
}