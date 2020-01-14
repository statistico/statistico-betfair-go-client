package betfair

type ListCompetitionsRequest struct {
	Filter MarketFilter `json:"filter"`
	Locale string       `json:"locale,omitempty"`
}

type ListEventTypesRequest struct {
	Filter MarketFilter `json:"filter"`
	Locale string       `json:"locale,omitempty"`
}

type ListEventsRequest struct {
	Filter MarketFilter `json:"filter"`
	Locale string       `json:"locale,omitempty"`
}

type ListMarketBookRequest struct {
	MarketIDs                     []string        `json:"marketIds"`
	PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
	OrderProjection               string          `json:"orderProjection,omitempty"`
	MatchProjection               string          `json:"matchProjection,omitempty"`
	IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
	PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
	CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
	CurrencyCode                  string          `json:"currencyCode,omitempty"`
	Locale                        string          `json:"locale,omitempty"`
	MatchedSince                  string          `json:"matchedSince,omitempty"`
	BetIDs                        []string        `json:"betids,omitempty"`
}

type ListMarketCatalogueRequest struct {
	Filter           MarketFilter `json:"filter"`
	MarketProjection []string     `json:"marketProjection,omitempty"`
	Sort             string       `json:"sort"`
	MaxResults       int          `json:"maxResults"`
	Locale           string       `json:"locale,omitempty"`
}

type ListRunnerBookRequest struct {
	MarketID                      string          `json:"marketId"`
	SelectionID                   uint64          `json:"selectionId"`
	Handicap                      float32         `json:"handicap,omitempty"`
	PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
	OrderProjection               string          `json:"orderProjection,omitempty"`
	MatchProjection               string          `json:"matchProjection,omitempty"`
	IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
	PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
	CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
	CurrencyCode                  string          `json:"currencyCode,omitempty"`
	Locale                        string          `json:"locale,omitempty"`
	MatchedSince                  string          `json:"matchedSince,omitempty"`
	BetIDs                        []string        `json:"betids,omitempty"`
}
