package models

type AdObject struct {
	AdId            string
	BidPrice        uint64
	IsOpen          bool
	BiddingDuration uint64
	MetaData        map[string]string
}
