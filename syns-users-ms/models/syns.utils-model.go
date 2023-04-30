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

type NFTMetadata struct {
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	AnimationURL *string    `json:"animation_url"`
	ExternalLink *string    `json:"external_link"`
	Image        string     `json:"image"`
	Attributes   []struct{} `json:"attributes"`
}

type NFT struct {
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
	NormalizedMetadata NFTMetadata `json:"normalized_metadata"`
	PossibleSpam      bool        `json:"possible_spam"`
}

type NFTResponse struct {
	Total     *int    `json:"total"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
	Cursor    string  `json:"cursor"`
	Result    []NFT   `json:"result"`
}
