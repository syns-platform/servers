/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package onchain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// @dev transfer sign-up reward to `toAddress`
//
// @param toAddress string
func TransferEth(toAddress string) (error) {
	// prepare ethereum client
	client, _ := ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/" +os.Getenv("ALCHEMY_API_KEY"))

	// prepare context
	ctx := context.Background()

	// convert env.privateKey to ecdsa.PrivateKey
	synsServicePrivateKey, err := crypto.HexToECDSA(os.Getenv("SYNS_SERVICE_ACCOUNT_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
		return err
	}

	// get ecdsa.PublicKey from privateKey
    synsServicePublicKeyECDSA, ok := synsServicePrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// convert ecdsa.PublicKey to hex address
	synsServiceHexAddress := crypto.PubkeyToAddress(*synsServicePublicKeyECDSA)

	// prepare nonce in the background
	// @notice `PendingNonceAt` returns the account nonce of the given account in the pending state. This is the nonce that should be used for the next transaction.
	// @notice Using `PendingNonceAt` avoids nonce collisions and ensure that your transactions are executed in the correct order on the blockchain.
	nonce, err := client.PendingNonceAt(ctx, synsServiceHexAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	// prepare gasLimit & gas price for the transaction
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err !=nil {
		log.Fatal(err)
		return err
	}

	// 50000000000000000 wei ~ 0.05 ETH
	value := new(big.Int).Add(big.NewInt(50000000000000000), gasPrice) // Syns Service Account will take care of the most of the gasPrice deducted from `value`

	// prepare recipient address
	recipientAddress := common.HexToAddress(toAddress)
	var data []byte

	// prepare transacton type
	tx := types.NewTransaction(nonce, recipientAddress, value, gasLimit, gasPrice, data)

	// get chainId
	chainId, _ := client.ChainID(ctx)

	// signing transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), synsServicePrivateKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// send transaction to blockchain
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return nil
}