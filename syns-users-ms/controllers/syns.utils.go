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
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Constants
var (
	ALCHEMY_BASE_URL = "https://polygon-mumbai.g.alchemy.com/nft/v2/"
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2/"
	OFFICUAL_SYNS_721_SC_ADDR = "0xfDe11549f6133020721975BAc8A054EF6FCb4C0f"
	OFFICUAL_SYNS_1155_SC_ADDR = "0x8aa884a1297f10C5B9Daa48Cd8e85Acb4C713933"
)

// @route `GET/get-all-syns-tokens/:asset-contract`
//
// @dev Get all Syns tokens from blockchain
//
// @honor Moralis API
//
// @param gc *gin.Context
func GetAllSynsTokens(gc *gin.Context) {
	// prepare asset contract
	assetContract := gc.Param("asset-contract")

	// sanity check param
	if !strings.EqualFold(assetContract, OFFICUAL_SYNS_721_SC_ADDR) && !strings.EqualFold(assetContract, OFFICUAL_SYNS_1155_SC_ADDR) {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"nfts": nil, "error": "Bad request - Invalid Syns Asset Contract Address"}); return;}
	}

	// prepare url
	moralisUrl := MORALIS_BASE_URL+"nft/"+assetContract+"?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"
	alchemyUrl := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTsForCollection?contractAddress="+assetContract+"&withMetadata=true"
  
	// prepare response object
	var moralisResObject map[string]interface{}
	var alchemyResObject map[string]interface{}

	// process http request
	moralisResObject = utils.DoHttp(moralisUrl, "X-API-Key", os.Getenv("MORALIS_API_KEY"), &moralisResObject)
	alchemyResObject = utils.DoHttp(alchemyUrl, "", "", &alchemyResObject)

	// prepare list of nfts from both APIs
	moralisNFTs := moralisResObject["result"].([]interface{})
	alchemyNFTs := alchemyResObject["nfts"].([]interface{})

	// sort the list of Nfts by token_id from largest to smallest
	sort.Slice(moralisNFTs, func(i, j int) bool {
		// access token_ids from elements
		tokenId_i := moralisNFTs[i].(map[string]interface{})["token_id"].(string)
		tokenId_j := moralisNFTs[j].(map[string]interface{})["token_id"].(string)

		// convert token_ids to integer
		tokenIdInt_i, _ := strconv.Atoi(tokenId_i)
		tokenIdInt_j, _ := strconv.Atoi(tokenId_j)

		// return sorting logic
		return tokenIdInt_i < tokenIdInt_j
	})

	// prepare an array of Syns Tokens
	var SynsNFTs []models.SynsNFT

	// implement Syns Token logics
	for i := 0; i < len(moralisNFTs); i ++ {
		// prepare fields
		tokenIdInt, _ := strconv.Atoi(moralisNFTs[i].(map[string]interface{})["token_id"].(string))
		quantityInt, _ := strconv.Atoi(moralisNFTs[i].(map[string]interface{})["amount"].(string))
		ercType := "ERC-721"
		if strings.Compare(moralisNFTs[i].(map[string]interface{})["contract_type"].(string), "ERC1155") == 0 {
			ercType = "ERC-1155"
		}
		
		// prepare SynsNFT struct 
		SynsNFT := models.SynsNFT{
			TokenID: tokenIdInt,
			AssetContract: assetContract,
			TokenOwner: "0x0000000000000000000000000000000000000000",
			OriginalOwner: moralisNFTs[i].(map[string]interface{})["minter_address"].(string),
			TokenURI: alchemyNFTs[i].(map[string]interface{})["tokenUri"].(map[string]interface{})["raw"].(string),
			Image: strings.Replace(alchemyNFTs[i].(map[string]interface{})["metadata"].(map[string]interface{})["image"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
			Audio: strings.Replace(alchemyNFTs[i].(map[string]interface{})["metadata"].(map[string]interface{})["audio"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
			ERCType: ercType,
			Quantity: quantityInt,
			OriginalQuantity: quantityInt,
			IsListing: false,
			ListingID: -1,
			RoyaltyBps: -1,
			Name: alchemyNFTs[i].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string),
			Description: alchemyNFTs[i].(map[string]interface{})["description"].(string),
			Age: 0, // for future, use Moralis Track NFT transfers API to calculate token block_timestamp if neccessary
			SharableLink: os.Getenv("OFFICIAL_PLATOFORM_URL")+"/syns-token/"+assetContract+"/"+moralisNFTs[i].(map[string]interface{})["token_id"].(string),
		}

		// append new SynsNFT to SynsNFTs
		SynsNFTs = append(SynsNFTs, SynsNFT)
	}

	// return to client
	gc.JSON(200, gin.H{"nfts": SynsNFTs, "error": nil})
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

	// sanity check param
	if !strings.EqualFold(assetContract, OFFICUAL_SYNS_721_SC_ADDR) && !strings.EqualFold(assetContract, OFFICUAL_SYNS_1155_SC_ADDR) {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"nfts": nil, "error": "Bad request - Invalid Syns Asset Contract Address"}); return;}
	}

	// test addresses to match ETH Crypto wallet address convention
	ownerAddrMatched, ownerErr := utils.TestEthAddress(&ownerAddr)
	if ownerErr != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"nfits":nil, "error": "!REGEX - cannot test ownerAddr against regex"}); return;}
	if !ownerAddrMatched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"nfits":nil, "error": "!ETH_ADDRESS - ownerAddr is not an ETH address"}); return;}

	// prepare url
	moralisUrl := MORALIS_BASE_URL+ownerAddr+"/nft?chain=mumbai&format=decimal&token_addresses%5B0%5D="+assetContract+"&normalizeMetadata=true&media_items=false"
	alchemyUrl := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTs?owner="+ownerAddr+"&contractAddresses[]="+assetContract+"&withMetadata=true&pageSize=100"
  
	// prepare response object
	var moralisResObject map[string]interface{}
	var alchemyResObject map[string]interface{}

	// process http request
	moralisResObject = utils.DoHttp(moralisUrl, "X-API-Key", os.Getenv("MORALIS_API_KEY"), &moralisResObject)
	alchemyResObject = utils.DoHttp(alchemyUrl, "", "", &alchemyResObject)

	
	// prepare list of nfts from both APIs
	moralisNFTs := moralisResObject["result"].([]interface{})
	alchemyNFTs := alchemyResObject["ownedNfts"].([]interface{})

	// sort the list of Nfts by token_id from largest to smallest
	sort.Slice(moralisNFTs, func(i, j int) bool {
		// access token_ids from elements
		tokenId_i := moralisNFTs[i].(map[string]interface{})["token_id"].(string)
		tokenId_j := moralisNFTs[j].(map[string]interface{})["token_id"].(string)

		// convert token_ids to integer
		tokenIdInt_i, _ := strconv.Atoi(tokenId_i)
		tokenIdInt_j, _ := strconv.Atoi(tokenId_j)

		// return sorting logic
		return tokenIdInt_i < tokenIdInt_j
	})

	// prepare an array of Syns Tokens
	var SynsNFTs []models.SynsNFT

	// implement Syns Token logics
	for i := 0; i < len(moralisNFTs); i ++ {
		// prepare fields
		tokenIdInt, _ := strconv.Atoi(moralisNFTs[i].(map[string]interface{})["token_id"].(string))
		quantityInt, _ := strconv.Atoi(moralisNFTs[i].(map[string]interface{})["amount"].(string))
		ercType := "ERC-721"
		if strings.Compare(moralisNFTs[i].(map[string]interface{})["contract_type"].(string), "ERC1155") == 0 {
			ercType = "ERC-1155"
		}
		
		// prepare SynsNFT struct 
		SynsNFT := models.SynsNFT{
			TokenID: tokenIdInt,
			AssetContract: assetContract,
			TokenOwner: ownerAddr,
			OriginalOwner: moralisNFTs[i].(map[string]interface{})["minter_address"].(string),
			TokenURI: alchemyNFTs[i].(map[string]interface{})["tokenUri"].(map[string]interface{})["raw"].(string),
			Image: strings.Replace(alchemyNFTs[i].(map[string]interface{})["metadata"].(map[string]interface{})["image"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
			Audio: strings.Replace(alchemyNFTs[i].(map[string]interface{})["metadata"].(map[string]interface{})["audio"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
			ERCType: ercType,
			Quantity: quantityInt,
			OriginalQuantity: quantityInt,
			IsListing: false,
			ListingID: -1,
			RoyaltyBps: -1,
			Name: alchemyNFTs[i].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string),
			Description: alchemyNFTs[i].(map[string]interface{})["description"].(string),
			Age: 0, // for future, use Moralis Track NFT transfers API to calculate token block_timestamp if neccessary
			SharableLink: os.Getenv("OFFICIAL_PLATOFORM_URL")+"/syns-token/"+assetContract+"/"+moralisNFTs[i].(map[string]interface{})["token_id"].(string),
		}

		// append new SynsNFT to SynsNFTs
		SynsNFTs = append(SynsNFTs, SynsNFT)
	}
	// return to client
	gc.JSON(200, gin.H{"nfts": SynsNFTs, "error": nil})
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

	// sanity check params
	if !strings.EqualFold(assetContract, OFFICUAL_SYNS_721_SC_ADDR) && !strings.EqualFold(assetContract, OFFICUAL_SYNS_1155_SC_ADDR) {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"SynsTokenMetadata": nil, "error": "Bad request - Invalid Syns Asset Contract Address"}); return;}
	}
	if _, err := strconv.Atoi(tokenId); err != nil {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"SynsTokenMetadata": nil, "error": "Bad request - Invalid token ID"}); return;}
	}
	if strings.Compare(tokenType, "ERC721") != 0 && strings.Compare(tokenType, "ERC1155") != 0 {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"SynsTokenMetadata": nil, "error": "Bad request - Invalid token type"}); return;}
	}

	// prepare urls
	alchemyUrl := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTMetadata?contractAddress="+assetContract+"&tokenId="+tokenId+"&tokenType="+tokenType+"&refreshCache=false"
	moralisUrl := MORALIS_BASE_URL+"nft/"+assetContract+"/"+tokenId+"/?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"
	moraliTransfersUrl := MORALIS_BASE_URL+"nft/"+assetContract+"/"+tokenId+"/transfers?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"

	// prepare response objects
	var alchemyResObject map[string]interface{}
	var moralisResObject map[string]interface{}
	var moralisTransferResObject map[string]interface{}

	// process http requests
	alchemyResObject = utils.DoHttp(alchemyUrl, "", "", &alchemyResObject)
	moralisResObject = utils.DoHttp(moralisUrl,"X-API-Key", os.Getenv("MORALIS_API_KEY"), &moralisResObject)
	moralisTransferResObject = utils.DoHttp(moraliTransfersUrl,"X-API-Key", os.Getenv("MORALIS_API_KEY"), &moralisTransferResObject)

	// make sure tokenID is a valid tokenID within the smart contract (i.e. check if alchemyResObject and moralisResObject returns non-empty metadata)
	if moralisResObject["message"] != nil && strings.Contains(moralisResObject["message"].(string), "No metadata found!") {
		{gc.AbortWithStatusJSON(http.StatusNotFound, gin.H{"SynsTokenMetadata": nil, "error": "No metadata found!"}); return;}
	}

	// prepare fields
	tokenIdInt, _ := strconv.Atoi(tokenId)
	quantityInt, _ := strconv.Atoi(moralisResObject["amount"].(string))
	ercType := "ERC-721"
	if strings.Compare(tokenType, "ERC1155") == 0 {
		ercType = "ERC-1155"
	} 

	// prepare token age
	transferResultArray := moralisTransferResObject["result"].([]interface{})
 	tokenAge, _ := time.Parse("2006-01-02T15:04:05.000Z", transferResultArray[len(transferResultArray) - 1].(map[string]interface{})["block_timestamp"].(string))
	
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
		Age: int(tokenAge.Unix()),
		SharableLink: os.Getenv("OFFICIAL_PLATOFORM_URL")+"/syns-token/"+assetContract+"/"+tokenId,
	}

	// return to client
	gc.JSON(200, gin.H{"SynsTokenMetadata": SynsNFT, "error": nil})
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

	// sanity check params
	if !strings.EqualFold(assetContract, OFFICUAL_SYNS_721_SC_ADDR) && !strings.EqualFold(assetContract, OFFICUAL_SYNS_1155_SC_ADDR) {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"owners": nil, "error": "Bad request - Invalid Syns Asset Contract Address"}); return;}
	}
	if _, err := strconv.Atoi(tokenId); err != nil {
		{gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"owners": nil, "error": "Bad request - Invalid token ID"}); return;}
	}

	// prepare url
	url := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getOwnersForToken?contractAddress="+assetContract+"&tokenId="+tokenId

	// prepare response object
	var resObject map[string]interface{}

	// process http request
	resObject = utils.DoHttp(url, "X-API-Key", os.Getenv("MORALIS_API_KEY"), &resObject)

	
	// make sure tokenID is a valid tokenID within the smart contract (i.e. check if alchemyResObject and moralisResObject returns non-empty metadata)
	if resObject == nil {
		{gc.AbortWithStatusJSON(http.StatusNotFound, gin.H{"owners": nil, "error": "No metadata found!"}); return;}
	}
	
	// return to client
	gc.JSON(200, gin.H{"owners": resObject["owners"].([]interface{}), "error": nil})
}
