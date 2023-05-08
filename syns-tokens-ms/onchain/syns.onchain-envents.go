/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package onchain

import (
	"Syns/servers/syns-tokens-ms/dao"
	"Syns/servers/syns-tokens-ms/models"
	onchain "Syns/servers/syns-tokens-ms/onchain/utils"
	"Syns/servers/syns-tokens-ms/utils"
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// constants
var (
	OFFICIAL_SYNS_721_SC_ADDR = utils.HandleReadFile("contract-artifacts/SynsERC721Address.json")["address"].(string)
	OFFICIAL_SYNS_MARKETPLACE_SC_ADDR = utils.HandleReadFile("contract-artifacts/SynsMarketplaceAddress.json")["address"].(string)
)


// @notice Root struct for other methods in controller
type SynsTokenOnchain struct {
	Syns721TokenDao dao.Syns721TokenDao
}

// @dev Constructor
func Syns721TokenOnchainConstructor(Syns721TokenDao dao.Syns721TokenDao) *SynsTokenOnchain{
	return &SynsTokenOnchain {
		Syns721TokenDao: Syns721TokenDao,
	}
}

// @dev Handle adding new minted token to database based on `newTokenMintedTo` event from SynsERC721 Smart Contract
// 
// @param client *ethclient.Client
// 
// @param pathToContract string
// 
// @param pathToContract string
func (sto *SynsTokenOnchain) HandleNewSyns721TokenMinted(client *ethclient.Client, pathToContract string) {
	// Extract the contract ABI from the JSON file
	contractABI := onchain.StringifyContractABI(pathToContract)

	// Parse the ABI into a Go type for ERC721 token contract
    abi, err := abi.JSON(strings.NewReader(contractABI))
    if err != nil {
        log.Fatal(err)
    }

	// prepare synsErc721Subscription and synsErc721EventLogs from onchain `newTokenMintedTo` event
	synsErc721Subscription, synsErc721EventLogs := onchain.ListenToOnChainEvent(client, abi, "newTokenMintedTo", OFFICIAL_SYNS_721_SC_ADDR)


	// Start event loop in background to do database logics
	go func() {
		for {
			select {
			case err := <-synsErc721Subscription.Err():
				log.Fatal(err)
			case eventLog := <-synsErc721EventLogs:
				// prepare lister
				minterAddr := common.HexToAddress(eventLog.Topics[1].Hex())

				// prepare tokenId
				tokenId, _ := strconv.ParseInt(eventLog.Topics[2].Hex(), 0, 64)

				// prepare royaltyBps
				royaltyBps, _ := strconv.ParseInt(eventLog.Topics[3].Hex(), 0, 64)

				// prepare tokenAge
				blockerHeader, _ := client.BlockByNumber(context.Background(), big.NewInt(0).SetUint64(eventLog.BlockNumber))
				tokenAge := blockerHeader.Time

				// prepare tokenUri
				event := abi.Events["newTokenMintedTo"]
				decoded, _ := event.Inputs.UnpackValues(eventLog.Data)
				tokenUri := decoded[0].(string)

				// prepare new Syns 721 super token
				synsSuperToken := onchain.PrepareNewMintedSyns721SuperNFT(minterAddr.Hex(), tokenUri, int(tokenId), int(royaltyBps), tokenAge())

				// add new token to database
				sto.Syns721TokenDao.MintNewSyns721Token(synsSuperToken)
			}
		}
	}()
}


// @dev Handle update listing information by listening to `ListingAdded` event from SynsMarketplace Smart Contract
// 
// @param client *ethclient.Client
// 
// @param pathToContract string
// 
// @param pathToContract string
func (sto *SynsTokenOnchain) HandleNewSyns721ListingAdded(client *ethclient.Client, pathToContract string) {
	// prepare eventName
	eventName := "ListingAdded"

	// Extract the contract ABI from the JSON file
	stringifiedContractABI := onchain.StringifyContractABI(pathToContract)

	// Parse the ABI into a Go type for ERC721 token contract
    contractABI, err := abi.JSON(strings.NewReader(stringifiedContractABI))
    if err != nil {
        log.Fatal(err)
    }

	// prepare synsErc721Subscription and synsErc721EventLogs from onchain `newTokenMintedTo` event
	synsListingSubscription, synsListingEventLogs := onchain.ListenToOnChainEvent(client, contractABI, eventName, OFFICIAL_SYNS_MARKETPLACE_SC_ADDR)

	// Start event loop in background to do database logics
	go func() {
		for {
			select {
			case err := <-synsListingSubscription.Err(): 
				log.Fatal(err)
			case eventLog := <-synsListingEventLogs:
				// prepare listingAddedEvent struct
				var listingAddedEvent struct {
					ListingId            *big.Int
					AssetContract        common.Address
					Lister               common.Address
					Listing              models.SynsMarketplaceListing
				}

				// Unpackge into listingAddedEvent
				err = contractABI.UnpackIntoInterface(&listingAddedEvent, eventName, eventLog.Data)
				if err != nil {
					fmt.Println("Failed to decode event log data:", err)
					return
				}

				// @logic if listing is created by tokenType = 1 (i.e. ERC721) => update syns721SuperToken in database
				if (listingAddedEvent.Listing.TokenType == 1) {
					if dbRes := sto.Syns721TokenDao.UpdatedSyns721SuperTokenBySynsListing(&listingAddedEvent.Listing, eventName); dbRes != nil {
						log.Fatal(dbRes)
					} else {
						log.Println("New Event Alert - ListingAdded: Successfully updated Syns 721 super token based on Syns Listing!")
					}
				}
			}
		}
	}()
}



// @dev Handle adding new minted token to database based on `newTokenMintedTo` event onchain
// 
// @param client *ethclient.Client
// 
// @param pathToContract string
// 
// @param pathToContract string
func (sto *SynsTokenOnchain) HandleSyns721RemoveSale(client *ethclient.Client, pathToContract string) {
	// prepare eventName
	eventName := "ListingRemoved"

	// Extract the contract ABI from the JSON file
	stringifiedContractABI := onchain.StringifyContractABI(pathToContract)

	// Parse the ABI into a Go type for ERC721 token contract
    contractABI, err := abi.JSON(strings.NewReader(stringifiedContractABI))
    if err != nil {
        log.Fatal(err)
    }

	// prepare synsErc721Subscription and synsErc721EventLogs from onchain `newTokenMintedTo` event
	synsListingSubscription, synsListingEventLogs := onchain.ListenToOnChainEvent(client, contractABI, eventName, OFFICIAL_SYNS_MARKETPLACE_SC_ADDR)

	// Start event loop in background to do database logics
	go func() {
		for {
			select {
			case err := <-synsListingSubscription.Err(): 
				log.Fatal(err)
			case eventLog := <-synsListingEventLogs:
				// prepare synsListing
				synsListing := models.SynsMarketplaceListing{
					ListingId: eventLog.Topics[1].Big(),
					TokenOwner: common.HexToAddress(eventLog.Topics[2].Hex()),
					StartSale: big.NewInt(0),
					BuyoutPricePerToken: big.NewInt(0),
				}

				// update super token logics
				if dbRes := sto.Syns721TokenDao.UpdatedSyns721SuperTokenBySynsListing(&synsListing, eventName); dbRes != nil {
					log.Fatal(dbRes)
				} else {
					log.Println("New Event Alert - ListingRemoved: Successfully updated Syns 721 super token based on Syns Listing!")
				}
			}
		}
	}()
}