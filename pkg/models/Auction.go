package models

type Auction struct {
	name          string
	Bids          []AdObject
	AdPlacementId string
	Winner        AdObject
}

type Bidder struct {
	name string
}
