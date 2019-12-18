package betfair

import "context"

type ListCompetitionsRequest struct {
	Context context.Context
	Filter MarketFilter `json:"filter"`
	Locale string `json:"locale, omitempty"`
}

type ListEventTypesRequest struct {
	Context context.Context
	Filter MarketFilter `json:"filter"`
	Locale string `json:"locale, omitempty"`
}