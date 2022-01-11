package models

// Personal finance information for a member of Congress, or someone in the judicial or executive branches.
type MemberProfile struct {
	Name              string `json:"name"`
	DataYear          int    `json:"data_year,string"`
	MemberId          string `json:"member_id"`       // CRP ID
	NetLow            int    `json:"net_low,string"`  // Calculated low range of the person's net worth
	NetHigh           int    `json:"net_high,string"` // Calculated high range of the person's net worth
	PositionHeldCount int    `json:"position_held_count,string"`
	AssetCount        int    `json:"asset_count,string"`
	AssetLow          int    `json:"asset_low,string"`  // Calculated low range value of the person's assets
	AssetHigh         int    `json:"asset_high,string"` // Calculated high range value of the person's assets
	TransactionCount  int    `json:"transaction_count,string"`
	TransactionLow    int    `json:"tx_low,string"`    // Lowest value of transaction range
	TransactionHigh   int    `json:"tx_high,string"`   // Highest value of transaction range
	Source            string `json:"source"`           // Link to this data on OpenSecrets.org
	Origin            string `json:"origin"`           // Attribute to display
	UpdateTimestamp   string `json:"update_timestamp"` // Date this data was last updated (M/DD/YY)
	Assets            []Asset
	Transactions      []Transaction
	Positions         []Position
}

// An asset reported by a member.
type Asset struct {
	Name         string `json:"name"`
	HoldingsLow  int    `json:"holdings_low,string"`  // Least the asset is worth
	HoldingsHigh int    `json:"holdings_high,string"` // Most the asset is worth
	Industry     string `json:"industry"`
	Sector       string `json:"sector"`        // Sector ID
	SubsidiaryOf string `json:"subsidiary_of"` // Parent organization
}

// Financial transaction done during period.
type Transaction struct {
	AssetName         string `json:"asset_name"`
	TransactionDate   string `json:"tx_date"`           // Mon DD YYYY
	TransactionAction string `json:"tx_action"`         // Buy, Sold, Exchanged
	ValueLow          int    `json:"value_low,string"`  // Minimum value of transaction
	ValueHigh         int    `json:"value_high,string"` // Maximum value of transaction
}

// Position held by a member.
type Position struct {
	Title        string `json:"title"`        // Position title
	Organization string `json:"organization"` // Organization with which position is held
}
