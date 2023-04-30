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

type MoralisNFTResponse struct {
	Total     *int    `json:"total"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
	Cursor    string  `json:"cursor"`
	Result    []moralisNFT   `json:"result"`
}

type moralisNFTMetadata struct {
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	AnimationURL *string    `json:"animation_url"`
	ExternalLink *string    `json:"external_link"`
	Image        string     `json:"image"`
	Attributes   []struct{} `json:"attributes"`
}

type moralisNFT struct {
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
	NormalizedMetadata moralisNFTMetadata `json:"normalized_metadata"`
	PossibleSpam      bool        `json:"possible_spam"`
}

///////////////////////////////////////////////////////////////////
/////////////////// 							/////////////////// 
///////////////////   Alchemy NFTs By Owner		/////////////////// 
/////////////////// 							/////////////////// 
///////////////////////////////////////////////////////////////////
type AlchemyNFTsByOwnerResponse struct {
	OwnedNfts 		[]alchemyNFT	`json:"ownedNfts"`
	TotalCount		int				`json:"totalCount"`
	BlockHash		string			`json:"blockHash"`
}

type alchemyNFT struct {
	Contract         alContract         `json:"contract"`
	ID               alID               `json:"id"`
	Balance          string           `json:"balance"`
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	TokenURI         alTokenURI         `json:"tokenUri"`
	Media            []alMedia          `json:"media"`
	Metadata         alMetadata         `json:"metadata"`
	TimeLastUpdated  time.Time        `json:"timeLastUpdated"`
	ContractMetadata alContractMetadata `json:"contractMetadata"`
}

type alContract struct {
	Address string `json:"address"`
}

type alID struct {
	TokenID        string         	`json:"tokenId"`
	TokenMetadata  alTokenMetadata  `json:"tokenMetadata"`
}

type alTokenMetadata struct {
	TokenType string `json:"tokenType"`
}

type alTokenURI struct {
	Gateway string `json:"gateway"`
	Raw     string `json:"raw"`
}

type alMedia struct {
	Gateway   string `json:"gateway"`
	Thumbnail string `json:"thumbnail"`
	Raw       string `json:"raw"`
	Format    string `json:"format"`
	Bytes     int64  `json:"bytes"`
}

type alMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Audio       string `json:"audio"`
}

type alContractMetadata struct {
	Name              string    	`json:"name"`
	Symbol            string    	`json:"symbol"`
	TokenType         string    	`json:"tokenType"`
	ContractDeployer  string    	`json:"contractDeployer"`
	DeployedBlockNumber int64   	`json:"deployedBlockNumber"`
	OpenSea           alOpenSea   	`json:"openSea"`
}

type alOpenSea struct {
	LastIngestedAt time.Time `json:"lastIngestedAt"`
}


///////////////////////////////////////////////////////////////////
/////////////////// 							/////////////////// 
///////////////////   Alchemy NFTs Metadata		/////////////////// 
/////////////////// 							/////////////////// 
///////////////////////////////////////////////////////////////////

type AlchemyNftMetadataResponse struct {
	Contract         alNftMetadataContract     `json:"contract"`
	ID               alNftMetadataID           `json:"id"`
	Title            string                     `json:"title"`
	Description      string                     `json:"description"`
	TokenURI         alNftMetadataTokenURI      `json:"tokenUri"`
	Media            []alNftMetadataMedia       `json:"media"`
	Metadata         alNftMetadataMetadata      `json:"metadata"`
	TimeLastUpdated  time.Time                  `json:"timeLastUpdated"`
	ContractMetadata alNftMetadataContractMeta  `json:"contractMetadata"`
}

type alNftMetadataContract struct {
	Address string `json:"address"`
}

type alNftMetadataID struct {
	TokenID       string                 `json:"tokenId"`
	TokenMetadata alNftMetadataTokenMeta `json:"tokenMetadata"`
}

type alNftMetadataTokenMeta struct {
	TokenType string `json:"tokenType"`
}

type alNftMetadataTokenURI struct {
	Gateway string `json:"gateway"`
	Raw     string `json:"raw"`
}

type alNftMetadataMedia struct {
	Gateway   string `json:"gateway"`
	Thumbnail string `json:"thumbnail"`
	Raw       string `json:"raw"`
	Format    string `json:"format"`
	Bytes     int64  `json:"bytes"`
}

type alNftMetadataMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Audio       string `json:"audio"`
}

type alNftMetadataContractMeta struct {
	Name              string                  `json:"name"`
	Symbol            string                  `json:"symbol"`
	TokenType         string                  `json:"tokenType"`
	ContractDeployer  string                  `json:"contractDeployer"`
	DeployedBlockNumber int64                 `json:"deployedBlockNumber"`
	OpenSea           alNftMetadataOpenSea    `json:"openSea"`
}

type alNftMetadataOpenSea struct {
	LastIngestedAt time.Time `json:"lastIngestedAt"`
}