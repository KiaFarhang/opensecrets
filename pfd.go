package opensecrets

type MemberProfile struct {
	Name              string `json:"name"`
	DataYear          int    `json:"data_year,string"`
	MemberId          string `json:"member_id"`
	NetLow            int    `json:"net_low,string"`
	NetHigh           int    `json:"net_high,string"`
	PositionHeldCount int    `json:"position_held_count,string"`
	AssetCount        int    `json:"asset_count,string"`
	AssetLow          int    `json:"asset_low,string"`
	AssetHigh         int    `json:"asset_high,string"`
	TransactionCount  int    `json:"transaction_count,string"`
	TransactionLow    int    `json:"tx_low,string"`
	TransactionHigh   int    `json:"tx_high,string"`
	Source            string `json:"source"`
	Origin            string `json:"origin"`
	UpdateTimestamp   string `json:"update_timestamp"`
	Assets            []Asset
	Transactions      []Transaction
	Positions         []Position
}

type Asset struct {
	Name         string `json:"name"`
	HoldingsLow  int    `json:"holdings_low,string"`
	HoldingsHigh int    `json:"holdings_high,string"`
	Industry     string `json:"industry"`
	Sector       string `json:"sector"`
	SubsidiaryOf string `json:"subsidiary_of"`
}

type Transaction struct {
	AssetName         string `json:"asset_name"`
	TransactionDate   string `json:"tx_date"`
	TransactionAction string `json:"tx_action"`
	ValueLow          int    `json:"value_low,string"`
	ValueHigh         int    `json:"value_high,string"`
}

type Position struct {
	Title        string `json:"title"`
	Organization string `json:"organization"`
}

type assetWrapper struct {
	Attributes Asset `json:"@attributes"`
}

type assetResponse struct {
	Wrapper []assetWrapper `json:"asset"`
}

type transactionWrapper struct {
	Attributes Transaction `json:"@attributes"`
}

type transactionResponse struct {
	Wrapper []transactionWrapper `json:"transaction"`
}

type positionWrapper struct {
	Attributes Position `json:"@attributes"`
}

type positionResponse struct {
	Wrapper []positionWrapper `json:"position"`
}

type memberPFDInnerWrapper struct {
	Attributes   MemberProfile       `json:"@attributes"`
	Assets       assetResponse       `json:"assets"`
	Transactions transactionResponse `json:"transactions"`
	Positions    positionResponse    `json:"positions"`
}

type memberPFDOuterWrapper struct {
	Profile memberPFDInnerWrapper `json:"member_profile"`
}

type memberPFDResponseWrapper struct {
	Response memberPFDOuterWrapper `json:"response"`
}
