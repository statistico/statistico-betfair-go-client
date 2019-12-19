package betfair

type ListCompetitionsRequest struct {
	Filter MarketFilter `json:"filter"`
	Locale string       `json:"locale, omitempty"`
}

type ListEventTypesRequest struct {
	Filter MarketFilter `json:"filter"`
	Locale string       `json:"locale, omitempty"`
}

type ListEventsRequest struct {
	Filter MarketFilter `json:"filter"`
	Locale string       `json:"locale, omitempty"`
}

type ListMarketCatalogueRequest struct {
	Filter MarketFilter `json:"filter"`
	//MarketProjection []string     `json:"marketProjection, omitempty"`
	//Sort             string       `json:"sort, omitempty"`
	MaxResults int    `json:"maxResults,omitempty"`
	Locale     string `json:"locale, omitempty"`
}
