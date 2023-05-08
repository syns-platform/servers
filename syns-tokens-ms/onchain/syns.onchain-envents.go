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
	onchain "Syns/servers/syns-tokens-ms/onchain/utils"
	"Syns/servers/syns-tokens-ms/utils"
	"context"
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
	OFFICUAL_SYNS_MARKETPLACE_SC_ADDR = utils.HandleReadFile("contract-artifacts/SynsMarketplaceAddress.json")["address"].(string)
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

// @dev Handle adding new minted token to database based on `newTokenMintedTo` event onchain
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
			case log := <-synsErc721EventLogs:
				// prepare lister
				minterAddr := common.HexToAddress(log.Topics[1].Hex())

				// prepare tokenId
				tokenId, _ := strconv.ParseInt(log.Topics[2].Hex(), 0, 64)

				// prepare royaltyBps
				royaltyBps, _ := strconv.ParseInt(log.Topics[3].Hex(), 0, 64)

				// prepare tokenAge
				blockerHeader, _ := client.BlockByNumber(context.Background(), big.NewInt(0).SetUint64(log.BlockNumber))
				tokenAge := blockerHeader.Time

				// prepare tokenUri
				event := abi.Events["newTokenMintedTo"]
				decoded, _ := event.Inputs.UnpackValues(log.Data)
				tokenUri := decoded[0].(string)

				// prepare new Syns 721 super token
				synsSuperToken := onchain.PrepareNewMintedSyns721SuperNFT(minterAddr.Hex(), tokenUri, int(tokenId), int(royaltyBps), tokenAge())

				// add new token to database
				sto.Syns721TokenDao.MintNewSyns721Token(synsSuperToken)
			}
		}
	}()
}