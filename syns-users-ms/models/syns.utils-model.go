/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

import (
	"time"
)

///////////////////////////////////////////////////////////
/////////////////// 					///////////////////
/////////////////// 	   Moralis		///////////////////
/////////////////// 					///////////////////
///////////////////////////////////////////////////////////

type MoralisNFTMetadata struct {
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	AnimationURL *string    `json:"animation_url"`
	ExternalLink *string    `json:"external_link"`
	Image        string     `json:"image"`
	Attributes   []struct{} `json:"attributes"`
}

type MoralisNFT struct {
	TokenAddress      string      `json:"token_address"`
	TokenID           string      `json:"token_id"`
	Amount            string      `json:"amount"`
	TokenHash         string      `json:"token_hash"`
	BlockNumberMinted string      `json:"block_number_minted"`
	UpdatedAt         *time.Time  `json:"updated_at"`
	ContractType      string      `json:"contract_type"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	TokenURI          string      `json:"token_uri"`
	Metadata          string      `json:"metadata"`
	LastTokenURISync  time.Time   `json:"last_token_uri_sync"`
	LastMetadataSync  time.Time   `json:"last_metadata_sync"`
	MinterAddress     string      `json:"minter_address"`
	NormalizedMetadata MoralisNFTMetadata `json:"normalized_metadata"`
	PossibleSpam      bool        `json:"possible_spam"`
}

type MoralisNFTResponse struct {
	Total     *int    `json:"total"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
	Cursor    string  `json:"cursor"`
	Result    []MoralisNFT   `json:"result"`
}

///////////////////////////////////////////////////////////
/////////////////// 					/////////////////// 
/////////////////// 	   Alchemy		/////////////////// 
/////////////////// 					/////////////////// 
/////////////////////////////////////////////////////////// 
type AlchemyNFTResponse struct {
	OwnedNfts 		[]AlchemyNFT	`json:"ownedNfts"`
	TotalCount		int				`json:"totalCount"`
	BlockHash		string			`json:"blockHash"`
}

type AlchemyNFT struct {
	Contract         AlContract         `json:"contract"`
	ID               AlID               `json:"id"`
	Balance          string           `json:"balance"`
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	TokenURI         AlTokenURI         `json:"tokenUri"`
	Media            []AlMedia          `json:"media"`
	Metadata         AlMetadata         `json:"metadata"`
	TimeLastUpdated  time.Time        `json:"timeLastUpdated"`
	ContractMetadata AlContractMetadata `json:"contractMetadata"`
}

type AlContract struct {
	Address string `json:"address"`
}

type AlID struct {
	TokenID        string         	`json:"tokenId"`
	TokenMetadata  AlTokenMetadata  `json:"tokenMetadata"`
}

type AlTokenMetadata struct {
	TokenType string `json:"tokenType"`
}

type AlTokenURI struct {
	Gateway string `json:"gateway"`
	Raw     string `json:"raw"`
}

type AlMedia struct {
	Gateway   string `json:"gateway"`
	Thumbnail string `json:"thumbnail"`
	Raw       string `json:"raw"`
	Format    string `json:"format"`
	Bytes     int64  `json:"bytes"`
}

type AlMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Audio       string `json:"audio"`
}

type AlContractMetadata struct {
	Name              string    	`json:"name"`
	Symbol            string    	`json:"symbol"`
	TokenType         string    	`json:"tokenType"`
	ContractDeployer  string    	`json:"contractDeployer"`
	DeployedBlockNumber int64   	`json:"deployedBlockNumber"`
	OpenSea           AlOpenSea   	`json:"openSea"`
}

type AlOpenSea struct {
	LastIngestedAt time.Time `json:"lastIngestedAt"`
}