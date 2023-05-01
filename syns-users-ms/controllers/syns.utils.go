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
	"os"
	"strconv"
	"strings"
	"time"

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
// @honor Moralis API
//
// @param gc *gin.Context
func GetAllSynsTokens(gc *gin.Context) {

	assetContract := gc.Param("asset-contract")
  
	url := MORALIS_BASE_URL+assetContract+"?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"
  
	// prepare response object
	var resObject map[string]interface{}

	// process http request
	resObject = utils.DoHttp(url, "X-API-Key", os.Getenv("MORALIS_API_KEY"), &resObject)

	// return to client
	gc.JSON(200, gin.H{"nfts": resObject["result"]})
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

	// prepare response object
	var resObject map[string]interface{}

	// process http request
	resObject = utils.DoHttp(url, "", "", &resObject)

	// return to client
	gc.JSON(200, gin.H{"nfts": resObject["ownedNfts"]})
}


// @route `GET/get-token-metadata/:asset-contract/:token-id/:token-type`
// 
// @dev Get token metadata based on asset contract and tokenId
//
// @honor Alchemy API
//
// @param gc *gin.Context
func GetTokenMetadata(gc *gin.Context) {
	// prepare params
	assetContract := gc.Param("asset-contract")
	tokenId := gc.Param("token-id")
	tokenType := gc.Param("token-type")

	// prepare urls
	alchemyUrl := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTMetadata?contractAddress="+assetContract+"&tokenId="+tokenId+"&tokenType="+tokenType+"&refreshCache=false"
	moralisUrl := MORALIS_BASE_URL+assetContract+"/"+tokenId+"/?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"


	// prepare response objects
	var alchemyResObject map[string]interface{}
	var moralisResObject map[string]interface{}

	// process http requests
	alchemyResObject = utils.DoHttp(alchemyUrl, "", "", &alchemyResObject)
	moralisResObject = utils.DoHttp(moralisUrl,"X-API-Key", os.Getenv("MORALIS_API_KEY"), &moralisResObject)

	// prepare integer fields
	tokenIdInt, _ := strconv.Atoi(tokenId)
	quantityInt, _ := strconv.Atoi(moralisResObject["amount"].(string))
	unixTime, _ := time.Parse("2006-01-02T15:04:05.000Z", moralisResObject["last_token_uri_sync"].(string))
	ercType := "ERC-721"
	if strings.Compare(tokenType, "ERC1155") == 0 {
		ercType = "ERC-1155"
	} 
	
	// prepare SynsNFT struct
	SynsNFT := models.SynsNFT{
		TokenID: tokenIdInt,
		AssetContract: assetContract,
		TokenOwner: moralisResObject["owner_of"].(string),
		OriginalOwner: moralisResObject["minter_address"].(string),
		TokenURI: alchemyResObject["tokenUri"].(map[string]interface{})["raw"].(string),
		Image: strings.Replace(alchemyResObject["metadata"].(map[string]interface{})["image"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
		Audio: strings.Replace(alchemyResObject["metadata"].(map[string]interface{})["audio"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
		ERCType: ercType,
		Quantity: quantityInt,
		OriginalQuantity: quantityInt,
		IsListing: false,
		ListingID: -1,
		RoyaltyBps: -1,
		Name: alchemyResObject["metadata"].(map[string]interface{})["name"].(string),
		Description: alchemyResObject["description"].(string),
		Age: int(unixTime.Unix()),
		SharableLink: os.Getenv("OFFICIAL_PLATOFORM_URL")+"/syns-token/"+assetContract+"/"+tokenId,
	}


	// return to client
	gc.JSON(200, SynsNFT)
}


// @route `GET/get-owners-for-token/:asset-contract/:token-id`
// 
// @dev Get an object of current owner and original creator of a token
//
// @honor Moralis API
//
// @param gc *gin.Context
func GetOwnersForToken(gc *gin.Context) {
	// prepare params
	assetContract := gc.Param("asset-contract")
	tokenId := gc.Param("token-id")

	// prepare url
	url := MORALIS_BASE_URL+assetContract+"/"+tokenId+"/owners?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"

	// prepare response object
	var resObject map[string]interface{}

	// process http request
	resObject = utils.DoHttp(url, "X-API-Key", os.Getenv("MORALIS_API_KEY"), &resObject)

	// extract the result field
	result := resObject["result"].([]interface{})[0].(map[string]interface{})

	// return to client
	gc.JSON(200, gin.H{"originalCreator": result["minter_address"], "currentOwner": result["owner_of"]})
}
