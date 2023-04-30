/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

import (
	"Syns/servers/syns-users-ms/models"
	"Syns/servers/syns-users-ms/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// @route `POST/generate-access-token`
//
// @dev Generate a new JWT from the payload passed in from http.body
//
// @param gc *gin.Context
func GenerateAccessToken(gc *gin.Context) {
	// prepare payloadParamHolder
	type payloadParamHolder struct {
		UserWalletAddress string `json:"userWalletAddress" validate:"required,len=42,alphanum"`
		Signature string `json:"signature" validate:"required,len=132,alphanum"`
		LoginMessage string `json:"loginMessage" validate:"required"`
	}

	// declare param
	var param *payloadParamHolder

	// bind json post data to param
	if err := gc.ShouldBindJSON(&param); err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
	}
	
	// test param.UserWalletAddress to match ETH Crypto wallet address convention
	UserWalletAddressMatched, err := utils.TestEthAddress(&param.UserWalletAddress)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test param.UserWalletAddress against regex"}); return;}
	if !UserWalletAddressMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - param.UserWalletAddress is not an ETH crypto wallet address"}); return;}

	// test param.Signature to see if it's a valid ethereum signed signature
	signedMsgMatched, err := utils.TestSignature(&param.Signature)
	if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "!REGEX - cannot test param.signature against regex"}); return;}
	if !signedMsgMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!SIGNATURE - param.signature is not an ETH cyphertext"}); return;}

	// extra validatation on struct param
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;}


	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserWalletAddress": param.UserWalletAddress,
		"signature": param.Signature,
		"loginMessage": param.LoginMessage,
		"exp": time.Now().Add(24*time.Hour).Unix(),
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
	gc.JSON(200, gin.H{"accessToken": accessToken})
}

// @route `GET/get-all-syns-tokens/:asset-contract`
//
// @dev Get all Syns 721 tokens from blockchain
//
// @param gc *gin.Context
func GetAllSyns721Tokens(gc *gin.Context) {

  assetContract := gc.Param("asset-contract")

  url := "https://deep-index.moralis.io/api/v2/nft/"+assetContract+"?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"

  req, _ := http.NewRequest("GET", url, nil)

  req.Header.Add("Accept", "application/json")
  req.Header.Add("X-API-Key", os.Getenv("MORALIS_API_KEY"))

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

	NFTRes := &models.NFTResponse{}

  // parse json from []byte to JSON
  json.Unmarshal(body, NFTRes)

	// return this to client
  gc.JSON(200, gin.H{"nfts": NFTRes})
}