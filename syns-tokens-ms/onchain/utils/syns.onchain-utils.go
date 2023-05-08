/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package onchain

import (
	"Syns/servers/syns-tokens-ms/models"
	"Syns/servers/syns-tokens-ms/utils"
	"context"
	"encoding/json"
	"log"
	"math"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// constants
var (
	OFFICIAL_SYNS_721_SC_ADDR = utils.HandleReadFile("contract-artifacts/SynsERC721Address.json")["address"].(string)
	ALCHEMY_BASE_URL = "https://polygon-mumbai.g.alchemy.com/nft/v2/"
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2/"
)

// @dev access contract ABI
//
// @param relativePath string - path to JSON file
//
// @return ABI - string
func StringifyContractABI(relativePath string) string {
	// handle read file from disk
	contractABI := utils.HandleReadFile(relativePath)

	// Extract the "abi" field from the contract ABI
	abiBytes, err := json.Marshal(contractABI["abi"])
	if err != nil {
		log.Fatal(err)
	}

	// Convert the ABI bytes into a string and return it
	return string(abiBytes)
}

// @dev handle listening to `eventName` event on chain
// 
// @param eventName string
//
// @return ethereum.Subscription
// 
// @return chan types.Log
func ListenToOnChainEvent (client *ethclient.Client, abi abi.ABI, eventName, contractAddr string) (ethereum.Subscription, chan types.Log) {
	// Create event filter to listen for `eventName` event
    eventFilter := ethereum.FilterQuery{
        Addresses: []common.Address{common.HexToAddress(contractAddr)},
        Topics: [][]common.Hash{{abi.Events[eventName].ID}},
    }

	 // Create channel to receive event logs
	 eventLogs := make(chan types.Log)

	 // Start event subscription
	 subscription, err := client.SubscribeFilterLogs(context.Background(), eventFilter, eventLogs)
	 if err != nil {
		 log.Fatal(err)
	 }

	 return subscription, eventLogs
}

// @dev prepare Syns721SuperNFT token
// 
// @param minterAddress string
// 
// @param tokenId string
// 
// @param royaltyBps string
func PrepareNewMintedSyns721SuperNFT(minterAddress, tokenURI string, tokenId common.Hash, royaltyBps uint8, tokenAge uint64) (synsSuperToken *models.Syns721SuperNFT) {
	// prepare urls
	tokenUriCloudfareUrl := strings.Replace(tokenURI, "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1)

	// prepare response objects
	var tokenUriObj map[string]interface{}

	// process http requests
	tokenUriObj = utils.DoHttp(tokenUriCloudfareUrl, "", "", &tokenUriObj)

	// prepare SynsNFT struct
	synsSuperToken = &models.Syns721SuperNFT{
		TokenID: tokenId.Big().Uint64(),
		AssetContract: OFFICIAL_SYNS_721_SC_ADDR,
		TokenOwner: minterAddress,
		OriginalOwner: minterAddress,
		TokenURI: tokenURI,
		Image: strings.Replace(tokenUriObj["image"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
		Audio: strings.Replace(tokenUriObj["audio"].(string), "ipfs://", "https://cloudflare-ipfs.com/ipfs/",1),
		ERCType: "ERC-721",
		Quantity: 1,
		IsListing: false,
		ListingID: math.MaxUint64,
		RoyaltyBps: royaltyBps,
		Name: tokenUriObj["name"].(string),
		Description: tokenUriObj["description"].(string),
		Age: tokenAge,
		SharableLink: os.Getenv("OFFICIAL_PLATOFORM_URL")+"/syns-token/"+OFFICIAL_SYNS_721_SC_ADDR+"/"+tokenId.Big().String(),
		Lister: minterAddress,
		StartSale: 0,
		EndSale: 0,
		Currency: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		BuyouPricePerToken: "0",
	}

	return synsSuperToken
}
