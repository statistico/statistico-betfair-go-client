package betfair

type ExBestOffersOverrides struct {
	BestPricesDepth          uint32  `json:"bestPricesDepth"`
	RollupModel              string  `json:"rollupModel,omitempty"`
	RollupLimit              uint32  `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float32 `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    uint32  `json:"rollupLiabilityFactor,omitempty"`
}

type MarketFilter struct {
	TextQuery          string    `json:"textQuery,omitempty"`
	ExchangeIDs        []string  `json:"exchangeIds,omitempty"`
	EventTypeIDs       []string  `json:"eventTypeIds,omitempty"`
	EventIDs           []string  `json:"eventIds,omitempty"`
	CompetitionIDs     []string  `json:"competitionIds,omitempty"`
	MarketIDs          []string  `json:"marketIds,omitempty"`
	Venues             []string  `json:"venues,omitempty"`
	BSPOnly            bool      `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  bool      `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         bool      `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []string  `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string  `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string  `json:"marketTypeCodes,omitempty"`
	MarketStartTime    TimeRange `json:"marketStartTime,omitempty"`
	WithOrders         []string  `json:"withOrders,omitempty"`
	RaceTypes          []string  `json:"raceTypes,omitempty"`
}

type PriceProjection struct {
	PriceData             []string
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise,omitempty"`
	RolloverStakes        bool                  `json:"rolloverStakes,omitempty"`
}

type TimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}
