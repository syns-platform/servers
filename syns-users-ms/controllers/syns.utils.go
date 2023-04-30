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
	"os"

	"github.com/gin-gonic/gin"
)

// Constants
var (
	ALCHEMY_BASE_URL = "https://polygon-mumbai.g.alchemy.com/nft/v2/"
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2/nft/"
)

// @route `GET/get-all-syns-tokens/:asset-contract`
//
// @dev Get all Syns tokens from blockchain
//
// @honor Moralis
//
// @param gc *gin.Context
func GetAllSynsTokens(gc *gin.Context) {

	assetContract := gc.Param("asset-contract")
  
	url := MORALIS_BASE_URL+assetContract+"?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"
  
	body := utils.DoHttp(url, "X-API-Key", os.Getenv("MORALIS_API_KEY"))

	// prepare response object
	NFTRes := &models.MoralisNFTResponse{}
  
	// parse json from []byte to JSON
	json.Unmarshal(body, NFTRes)
  
	  // return this to client
	gc.JSON(200, gin.H{"nfts": NFTRes})
  }

// @route `GET/get-nfts-owned-by/:owner-addr/:asset-contract`
// 
// @dev Get all Syns tokens owned by an owner address
//
// @honor Alchemy API
//
// @param gc *gin.Context
func GetSynsTokensOwnedBy(gc *gin.Context) {
	// prepare params
	ownerAddr := gc.Param("owner-addr")
	assetContract := gc.Param("asset-contract")

	// prepare url
	url := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTs?owner="+ownerAddr+"&contractAddresses[]="+assetContract+"&withMetadata=true&pageSize=100"

	// process http request
	body := utils.DoHttp(url, "", "")

	// prepare response object
	NFTRes := &models.AlchemyNFTResponse{}

	// parse json from []byte to JSON
	json.Unmarshal(body, NFTRes)

	// return this to client
	gc.JSON(200, gin.H{"nfts": NFTRes.OwnedNfts})
}